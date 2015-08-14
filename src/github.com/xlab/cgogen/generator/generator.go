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
}

type ArchFlagSet struct {
	Name  string
	Arch  []string
	Flags []string
}

type Config struct {
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

func New(packageName string, cfg *Config, tr *tl.Translator) (*Generator, error) {
	if len(packageName) == 0 {
		return nil, errors.New("no package name provided")
	}
	if tr == nil {
		return nil, errors.New("no translator provided")
	}
	gen := &Generator{
		pkg: packageName,
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
	for _, decl := range gen.tr.Typedefs() {
		if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
			continue
		}
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
			if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			}
			gen.writeStructDeclaration(wr, decl, true)
		case tl.EnumKind:
			if !decl.IsTemplate() {
				if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
					continue
				}
				gen.writeEnumDeclaration(wr, decl, true)
			}
		case tl.FunctionKind:
			if !gen.tr.IsAcceptableName(tl.TargetFunction, decl.Name) {
				continue
			}
			gen.writeFunctionDeclaration(wr, decl, true)
		}
		writeSpace(wr, 1)
	}
}
