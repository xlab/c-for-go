package generator

import (
	"errors"
	"io"
	"math/rand"

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
	rand          *rand.Rand
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
		rand:        rand.New(rand.NewSource(+79269965690)),
	}
	return gen, nil
}

func (gen *Generator) WriteConst(wr io.Writer) int {
	var count int
	gen.writeDefinesGroup(wr, gen.tr.Defines())
	writeSpace(wr, 1)
	tagsSeen := make(map[string]bool)
	namesSeen := make(map[string]bool)

	expandEnum := func(decl tl.CDecl) bool {
		if tag := decl.Spec.GetBase(); len(tag) == 0 {
			gen.expandEnumAnonymous(wr, decl, namesSeen)
			return true
		} else if tagsSeen[tag] {
			return false
		} else {
			gen.expandEnum(wr, decl)
			tagsSeen[tag] = true
			return true
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
		if expandEnum(decl) {
			count++
		}
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
		if expandEnum(decl) {
			count++
		}
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
		count++
	}
	return count
}

func (gen *Generator) WriteTypedefs(wr io.Writer) int {
	var count int
	seenStructTags := make(map[string]bool)
	seenUnionTags := make(map[string]bool)
	for _, decl := range gen.tr.Typedefs() {
		if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
			continue
		}
		switch decl.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			var memTip tl.Tip
			if tag := decl.Spec.GetBase(); len(tag) > 0 {
				seenStructTags[tag] = true
				if memTipRx, ok := gen.tr.MemTipRx(tag); ok {
					memTip = memTipRx.Self()
				}
			}
			if !memTip.IsValid() {
				if memTipRx, ok := gen.tr.MemTipRx(decl.Name); ok {
					memTip = memTipRx.Self()
				}
			}
			gen.writeStructTypedef(wr, decl, memTip == tl.TipMemRaw)
		case tl.UnionKind:
			if tag := decl.Spec.GetBase(); len(tag) > 0 {
				seenUnionTags[tag] = true
			}
			gen.writeUnionTypedef(wr, decl)
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
		count++
	}
	for tag, decl := range gen.tr.TagMap() {
		switch decl.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			if seenStructTags[tag] {
				continue
			}
			if !gen.tr.IsAcceptableName(tl.TargetPublic, tag) {
				continue
			}
			decl.Name = tag
			if memTipRx, ok := gen.tr.MemTipRx(tag); ok {
				gen.writeStructTypedef(wr, decl, memTipRx.Self() == tl.TipMemRaw)
			}
			gen.writeStructTypedef(wr, decl, false)
			writeSpace(wr, 1)
			count++
		case tl.UnionKind:
			if seenUnionTags[tag] {
				continue
			}
			if !gen.tr.IsAcceptableName(tl.TargetPublic, tag) {
				continue
			}
			decl.Name = tag
			gen.writeUnionTypedef(wr, decl)
			writeSpace(wr, 1)
			count++
		}
	}
	return count
}

func (gen *Generator) WriteDeclares(wr io.Writer) int {
	var count int
	for _, decl := range gen.tr.Declares() {
		if decl.IsStatic {
			continue
		}
		const public = true
		switch decl.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			}
			gen.writeStructDeclaration(wr, decl, tl.NoTip, public)
		case tl.UnionKind:
			if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			}
			gen.writeUnionDeclaration(wr, decl, tl.NoTip, public)
		case tl.EnumKind:
			if !decl.IsTemplate() {
				if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
					continue
				}
				gen.writeEnumDeclaration(wr, decl, tl.NoTip, public)
			}
		case tl.FunctionKind:
			if !gen.tr.IsAcceptableName(tl.TargetFunction, decl.Name) {
				continue
			}
			// defaults to ref for the returns
			ptrTip := tl.TipPtrRef
			if ptrTipRx, ok := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name); ok {
				if tip := ptrTipRx.Self(); tip.IsValid() {
					ptrTip = tip
				}
			}
			gen.writeFunctionDeclaration(wr, decl, ptrTip, public)
		}
		writeSpace(wr, 1)
		count++
	}
	return count
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

// randPostfix generates a simply random 4-byte postfix. Doesn't require a crypto package.
func (gen *Generator) randPostfix() int32 {
	return 0x0f000000 + gen.rand.Int31n(0x00ffffff)
}
