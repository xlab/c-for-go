package parser

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cznic/cc"
)

type Config struct {
	Arch         string   `yaml:"Arch"`
	IncludePaths []string `yaml:"IncludePaths"`
	SourcesPaths []string `yaml:"SourcesPaths"`

	Defines map[string]interface{} `yaml:"Defines"`

	CCDefs   bool `yaml:"-"`
	archBits TargetArch
}

func ParseWith(cfg *Config) (*cc.TranslationUnit, error) {
	if len(cfg.SourcesPaths) == 0 {
		return nil, errors.New("parser: no target paths specified")
	}
	cfg, err := checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	predefined := builtinBase
	// user-provided defines take precedence
	for name, value := range cfg.Defines {
		switch v := value.(type) {
		case string:
			predefined += fmt.Sprintf("\n#define %s \"%s\"", name, v)
		case int, int16, int32, int64, uint, uint16, uint32, uint64:
			predefined += fmt.Sprintf("\n#define %s %d", name, v)
		case float32, float64:
			predefined += fmt.Sprintf("\n#define %s %ff", name, v)
		case map[interface{}]interface{}:
			if len(v) == 0 { // the special case for override a definition beforehand
				predefined += fmt.Sprintf("\n#define %s", name)
			}
		}
	}
	var (
		ccDefs   string
		ccDefsOK bool
	)
	if cfg.CCDefs {
		CPP := "cpp"
		if v, ok := os.LookupEnv("CPP"); ok {
			CPP = v
		} else if v, ok = os.LookupEnv("CC"); ok {
			CPP = v
		}
		predefined, _, sysIncludePaths, err := cc.HostCppConfig(CPP)
		if err != nil {
			log.Println("[WARN]:", err)
		} else {
			if len(sysIncludePaths) > 0 {
				// add on top of sysIncludePaths
				cfg.IncludePaths = append(sysIncludePaths, cfg.IncludePaths...)
			}
			ccDefs = predefined
			ccDefsOK = true
		}
	}
	if ccDefsOK {
		predefined += fmt.Sprintf("\n%s", ccDefs)
	} else {
		predefined += basePredefines
		if archDefs, ok := archPredefines[cfg.archBits]; ok {
			predefined += fmt.Sprintf("\n%s", archDefs)
		}
	}
	// undefines?
	predefined += fmt.Sprintf("\n%s", builtinBaseUndef)
	for name, value := range cfg.Defines {
		switch value.(type) {
		case string, int, int16, int32, int64, uint, uint16, uint32, uint64, float32, float64:
			continue
		default: // a corner case: undef using an the nil value
			if value == nil {
				predefined += fmt.Sprintf("\n#undef %s", name)
			}
		}
	}
	model := models[cfg.archBits]
	return cc.Parse(predefined, cfg.SourcesPaths, model, cc.SysIncludePaths(cfg.IncludePaths),
		cc.EnableAnonymousStructFields())
}

func checkConfig(cfg *Config) (*Config, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	if arch, ok := arches[cfg.Arch]; !ok {
		// default to 64-bit arch
		cfg.archBits = Arch64
	} else if arch != Arch32 && arch != Arch64 && arch != Arch48 {
		// default to 64-bit arch
		cfg.archBits = Arch64
	}
	// workaround for cznic's cc (it panics if supplied path is a dir)
	var saneFiles []string
	for _, path := range cfg.SourcesPaths {
		if !filepath.IsAbs(path) {
			if hPath, err := findFile(path, cfg.IncludePaths); err != nil {
				err = fmt.Errorf("parser: file specified but not found: %s (include paths: %s)",
					path, strings.Join(cfg.IncludePaths, ", "))
				return nil, err
			} else {
				path = hPath
			}
		}
		if info, err := os.Stat(path); err != nil || info.IsDir() {
			continue
		}
		if absPath, err := filepath.Abs(path); err != nil {
			path = absPath
		}
		saneFiles = append(saneFiles, path)
	}
	cfg.SourcesPaths = saneFiles
	return cfg, nil
}

func findFile(path string, includePaths []string) (string, error) {
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	for _, inc := range includePaths {
		result := filepath.Join(inc, path)
		if _, err := os.Stat(result); err == nil {
			return result, nil
		}
	}
	return "", errors.New("not found")
}
