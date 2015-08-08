package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cznic/c/internal/cc"
)

type Config struct {
	Arch               string
	archBits           TargetArchBits
	CustomDefinesPath  string
	WebIncludesEnabled bool
	WebIncludePrefix   string
	IncludePaths       []string
	SysIncludePaths    []string
	TargetPaths        []string
}

func NewConfig(paths ...string) *Config {
	return &Config{
		TargetPaths: paths,
	}
}

func checkConfig(cfg *Config) *Config {
	if cfg == nil {
		cfg = &Config{}
	}
	if arch, ok := arches[cfg.Arch]; !ok {
		// default to 64-bit arch
		cfg.archBits = Arch64
	} else if arch != Arch32 && arch != Arch64 {
		// default to 64-bit arch
		cfg.archBits = Arch64
	}
	return cfg
}

type Parser struct {
	cfg        *Config
	ccCfg      *cc.ParseConfig
	predefined string
}

func New(cfg *Config) (*Parser, error) {
	p := &Parser{
		cfg: checkConfig(cfg),
	}
	if len(cfg.TargetPaths) > 0 {
		// workaround for cznic's cc (it panics if supplied path is a dir)
		var saneFiles []string
		for _, path := range cfg.TargetPaths {
			if info, err := os.Stat(path); err != nil || info.IsDir() {
				continue
			}
			if absPath, err := filepath.Abs(path); err != nil {
				path = absPath
			}
			saneFiles = append(saneFiles, path)
		}
		cfg.TargetPaths = saneFiles
	} else {
		return nil, errors.New("parser: no target paths specified")
	}

	if def, ok := predefines[cfg.archBits]; !ok {
		p.predefined = predefinedBase
	} else {
		p.predefined = def
	}
	if len(cfg.CustomDefinesPath) > 0 {
		if buf, err := ioutil.ReadFile(cfg.CustomDefinesPath); err != nil {
			return nil, errors.New("parser: custom defines file provided but can't be read")
		} else if len(buf) > 0 {
			p.predefined = fmt.Sprintf("%s\n// custom defines below\n%s", p.predefined, buf)
		}
	}
	if ccCfg, err := p.ccParserConfig(); err != nil {
		return nil, err
	} else {
		p.ccCfg = ccCfg
	}
	return p, nil
}

func (p *Parser) ccParserConfig() (*cc.ParseConfig, error) {
	ccCfg := &cc.ParseConfig{
		Predefined:         p.predefined,
		Paths:              p.cfg.TargetPaths,
		SysIncludePaths:    p.cfg.SysIncludePaths,
		IncludePaths:       p.cfg.IncludePaths,
		WebIncludesEnabled: p.cfg.WebIncludesEnabled,
		WebIncludePrefix:   p.cfg.WebIncludePrefix,
	}
	if err := cc.CheckParseConfig(ccCfg); err != nil {
		return nil, err
	}
	return ccCfg, nil
}

func (p *Parser) Parse() (unit *cc.TranslationUnit, err error) {
	// this works as easy as this only with patched cc package, beware when using the vanilla cznic/cc.
	return cc.Parse(p.ccCfg)
}
