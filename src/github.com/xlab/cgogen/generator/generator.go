package generator

import (
	"errors"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

type Generator struct {
	cfg *Config
	tr  *tl.Translator
}

type ArchFlagSet struct {
	Name  string
	Arch  []string
	Flags []string
}

type Config struct {
	PackageName        string
	PackageDescription string
	PackageLicense     string
	PkgConfigOpts      []string
	CFlags             ArchFlagSet
	LDFlags            ArchFlagSet
	CPPFlags           ArchFlagSet
	CXXFlags           ArchFlagSet
	SysIncludes        []string
	Includes           []string
}

func CheckConfig(cfg *Config) error {
	if len(cfg.PackageName) == 0 {
		return errors.New("no package name specified")
	}
	return nil
}

func New(cfg *Config, tr *tl.Translator) (*Generator, error) {
	if tr == nil {
		return nil, errors.New("no translator provided")
	}
	if err := CheckConfig(cfg); err != nil {
		return nil, err
	}
	gen := &Generator{
		cfg: cfg,
		tr:  tr,
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

	for _, decl := range gen.tr.TagMap() {
		if decl.Spec.Kind() != tl.EnumKind {
			continue
		} else if !decl.IsTemplate() {
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
		gen.writeConstDeclaration(wr, decl)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) WriteTypedefs(wr io.Writer) {
	for _, decl := range gen.tr.Typedefs() {
		switch decl.Kind() {
		case tl.StructKind:
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
}

func (gen *Generator) WriteDeclares(wr io.Writer) {
	for _, decl := range gen.tr.Declares() {
		if decl.IsStatic {
			continue
		}
		switch decl.Kind() {
		case tl.StructKind:
			gen.writeStructDeclaration(wr, decl, true)
		case tl.EnumKind:
			if !decl.IsTemplate() {
				gen.writeEnumDeclaration(wr, decl, true)
			}
		case tl.FunctionKind:
			gen.writeFunctionDeclaration(wr, decl, true)
		}
		writeSpace(wr, 1)
	}
}
