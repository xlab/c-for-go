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
	//tagsSeen := make(map[string]struct{})
	for _, decl := range gen.tr.Typedefs() {
		if decl.Spec.Kind() != tl.EnumKind {
			continue
		}
		// spec := decl.Spec.(*tl.CEnumSpec)
		// if len(spec.Tag) > 0 {
		// 	if _, ok := tagsSeen[spec.Tag]; ok {
		// 		continue
		// 	}
		// 	tagsSeen[spec.Tag] = struct{}{}
		// }
		gen.writeEnum(wr, decl)
		writeSpace(wr, 1)
	}
	for _, decl := range gen.tr.Declares() {
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
	//tagsSeen := make(map[string]struct{})
	for _, decl := range gen.tr.Typedefs() {
		// if tag := decl.Spec.GetBase(); len(tag) > 0 {
		// 	if _, ok := tagsSeen[tag]; ok {
		// 		continue
		// 	}
		// 	tagsSeen[tag] = struct{}{}
		// }
		switch decl.Kind() {
		case tl.StructKind:
			gen.writeStructDeclaration(wr, decl)
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, decl)
		}
		writeSpace(wr, 1)
	}
}
