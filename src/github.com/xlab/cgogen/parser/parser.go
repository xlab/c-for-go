package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cznic/c/internal/cc"
	"github.com/cznic/c/internal/xc"
)

type Config struct {
	Arch            TargetArch
	DefinesPath     string
	IncludePaths    []string
	SysIncludePaths []string
	TargetPaths     []string
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
	if cfg.Arch != Arch32 && cfg.Arch != Arch64 {
		// default to 64-bit arch
		cfg.Arch = Arch64
	}
	return cfg
}

type Parser struct {
	cfg        *Config
	predefined string
	model      cc.Model
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
	}
	if len(cfg.TargetPaths) == 0 {
		return nil, errors.New("parser: no target paths specified")
	}
	if model, ok := models[cfg.Arch]; !ok {
		return nil, errors.New(fmt.Sprintf("parser: the target model not defined: %v", cfg.Arch))
	} else {
		p.model = model
	}
	if def, ok := predefines[cfg.Arch]; !ok {
		p.predefined = predefinedBase
	} else {
		p.predefined = def
	}
	if len(cfg.IncludePaths) == 0 {
		cfg.IncludePaths = []string{"."}
	}
	cfg.SysIncludePaths = []string{"/usr/include", "/usr/local/Cellar/libvpx/1.4.0/include/vpx"}
	if len(cfg.DefinesPath) > 0 {
		if buf, err := ioutil.ReadFile(cfg.DefinesPath); err != nil {
			return nil, errors.New("parser: custom defines file provided but can't be read")
		} else if len(buf) > 0 {
			p.predefined = fmt.Sprintf("%s\n// custom defines below\n%s", p.predefined, buf)
		}
	}
	return p, nil
}

func (p *Parser) Parse() (unit *cc.TranslationUnit, macros []string, err error) {
	unit, err = cc.Parse(p.predefined, p.cfg.TargetPaths, p.model,
		cc.IncludePaths(p.cfg.IncludePaths),
		cc.SysIncludePaths(p.cfg.SysIncludePaths))
	if err != nil {
		unit = nil
		return
	}
	for id := range cc.Macros {
		if macro := xc.Dict.S(id); len(macro) > 0 {
			macros = append(macros, string(macro))
		}
	}
	return
}
