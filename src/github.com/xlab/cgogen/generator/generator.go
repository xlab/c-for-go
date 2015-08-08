package generator

import (
	"errors"

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
