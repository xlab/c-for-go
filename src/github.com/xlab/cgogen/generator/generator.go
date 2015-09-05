package generator

import (
	"errors"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

type Generator struct {
	pkg string
	cfg *Config
	tr  *tl.Translator
	//
	closed        bool
	closeC, doneC chan struct{}
	helpersChan   chan *Helper
}

type ArchFlagSet struct {
	Name  string
	Arch  []string
	Flags []string
}

type Config struct {
	PackageName        string      `yaml:"PackageName"`
	PackageDescription string      `yaml:"PackageDescription"`
	PackageLicense     string      `yaml:"PackageLicense"`
	PkgConfigOpts      []string    `yaml:"PkgConfigOpts"`
	CFlags             ArchFlagSet `yaml:"CFlags"`
	LDFlags            ArchFlagSet `yaml:"LDFlags"`
	CPPFlags           ArchFlagSet `yaml:"CPPFlags"`
	CXXFlags           ArchFlagSet `yaml:"CXXFlags"`
	SysIncludes        []string    `yaml:"SysIncludes"`
	Includes           []string    `yaml:"Includes"`
}

func New(pkg string, cfg *Config, tr *tl.Translator) (*Generator, error) {
	if cfg == nil || len(cfg.PackageName) == 0 {
		return nil, errors.New("no package name provided")
	} else if tr == nil {
		return nil, errors.New("no translator provided")
	}
	gen := &Generator{
		pkg: pkg,
		cfg: cfg,
		tr:  tr,
		//
		helpersChan: make(chan *Helper, 1),
		closeC:      make(chan struct{}),
		doneC:       make(chan struct{}),
	}
	return gen, nil
}

func (gen *Generator) WriteConst(wr io.Writer) {
	gen.writeDefinesGroup(wr, gen.tr.Defines())
	writeSpace(wr, 1)
	tagsSeen := make(map[string]bool)
	namesSeen := make(map[string]bool)

	expandEnum := func(decl tl.CDecl) {
		if tag := decl.Spec.GetBase(); len(tag) == 0 {
			gen.expandEnumAnonymous(wr, decl, namesSeen)
		} else if tagsSeen[tag] {
			return
		} else {
			gen.expandEnum(wr, decl)
			tagsSeen[tag] = true
		}
	}

	for tag, decl := range gen.tr.TagMap() {
		if decl.Spec.Kind() != tl.EnumKind {
			continue
		} else if !decl.IsTemplate() {
			continue
		}
		if !gen.tr.IsAcceptableName(tl.TargetType, tag) {
			continue
		}
		expandEnum(decl)
	}
	for _, decl := range gen.tr.Typedefs() {
		if decl.Spec.Kind() != tl.EnumKind {
			continue
		} else if !decl.IsTemplate() {
			continue
		}
		if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
			continue
		}
		expandEnum(decl)
	}
	for _, decl := range gen.tr.Declares() {
		if decl.IsStatic {
			continue
		}
		if len(decl.Name) == 0 {
			continue
		}
		if decl.Spec.Kind() != tl.TypeKind {
			continue
		} else if !decl.IsConst() {
			continue
		}
		if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
			continue
		}
		gen.writeConstDeclaration(wr, decl)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) WriteTypedefs(wr io.Writer) {
	seenTags := make(map[string]bool)
	for _, decl := range gen.tr.Typedefs() {
		if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
			continue
		}
		switch decl.Kind() {
		case tl.StructKind:
			if tag := decl.Spec.GetBase(); len(tag) > 0 {
				seenTags[tag] = true
			}
			gen.writeStructTypedef(wr, decl)
		case tl.EnumKind:
			if !decl.IsTemplate() {
				gen.writeEnumTypedef(wr, decl)
			}
		case tl.TypeKind:
			gen.writeTypeTypedef(wr, decl)
		case tl.FunctionKind:
			gen.writeFunctionTypedef(wr, decl)
		}
		writeSpace(wr, 1)
	}
	for tag, decl := range gen.tr.TagMap() {
		switch decl.Kind() {
		case tl.StructKind:
			if seenTags[tag] {
				continue
			}
			if !gen.tr.IsAcceptableName(tl.TargetPublic, tag) {
				continue
			}
			decl.Name = tag
			gen.writeStructTypedef(wr, decl)
			writeSpace(wr, 1)
		}
	}
}

func (gen *Generator) WriteDeclares(wr io.Writer) {
	for _, decl := range gen.tr.Declares() {
		if decl.IsStatic {
			continue
		}
		switch decl.Kind() {
		case tl.StructKind:
			if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			}
			gen.writeStructDeclaration(wr, decl, tl.TipPtrRef, true)
		case tl.EnumKind:
			if !decl.IsTemplate() {
				if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
					continue
				}
				gen.writeEnumDeclaration(wr, decl, tl.TipPtrRef, true)
			}
		case tl.FunctionKind:
			if !gen.tr.IsAcceptableName(tl.TargetFunction, decl.Name) {
				continue
			}
			gen.writeFunctionDeclaration(wr, decl, tl.TipPtrRef, true)
		}
		writeSpace(wr, 1)
	}
}

func (gen *Generator) Close() {
	if gen.closed {
		return
	}
	gen.closed = true
	gen.closeC <- struct{}{}
	<-gen.doneC
}

func (gen *Generator) MonitorAndWriteHelpers(goWr, cWr io.Writer, initWrFunc ...func() (io.Writer, error)) {
	seenHelperNames := make(map[string]bool)
	var seenGoHelper bool
	var seenCHelper bool
	for {
		select {
		case <-gen.closeC:
			close(gen.helpersChan)
		case helper, ok := <-gen.helpersChan:
			if !ok {
				close(gen.doneC)
				return
			}
			if seenHelperNames[helper.Name] {
				continue
			}
			seenHelperNames[helper.Name] = true

			var wr io.Writer
			switch helper.Side {
			case NoSide, GoSide:
				if goWr != nil {
					wr = goWr
				} else if len(initWrFunc) < 1 {
					continue
				} else if w, err := initWrFunc[0](); err != nil {
					continue
				} else {
					wr = w
				}
				if !seenGoHelper {
					gen.writeGoHelpersHeader(wr)
					seenGoHelper = true
				}
			case CSide:
				if cWr != nil {
					wr = cWr
				} else if len(initWrFunc) < 2 {
					continue
				} else if w, err := initWrFunc[1](); err != nil {
					continue
				} else {
					wr = w
				}
				if !seenCHelper {
					gen.writeCHelpersHeader(wr)
					seenCHelper = true
				}
			default:
				continue
			}

			if len(helper.Description) > 0 {
				writeTextBlock(wr, helper.Description)
			}
			writeSourceBlock(wr, helper.Source)
			writeSpace(wr, 1)
		}
	}
}
