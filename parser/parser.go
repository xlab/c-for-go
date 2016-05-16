package parser

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		}
	}
	var (
		ccDefs   string
		ccDefsOK bool
	)
	if cfg.CCDefs {
		ccDefs, ccDefsOK = stealDefinesFromCC()
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
	for name, value := range cfg.Defines {
		switch value.(type) {
		case string, int, int16, int32, int64, uint, uint16, uint32, uint64, float32, float64:
			continue
		default: // a corner case: undef using an unknown value type
			predefined += fmt.Sprintf("\n#undef %s", name)
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

func stealDefinesFromCC() (defs string, ok bool) {
	cc, ok := os.LookupEnv("CC")
	if !ok { // second chance for CPP
		cc, ok = os.LookupEnv("CPP")
		if !ok {
			return
		}
	}
	cmd := exec.Command(cc, "-dM", "-E", "-x", "c", "/dev/null")
	buf, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("[WARN]:", err)
		return
	}
	defs = string(buf)
	return defs, true
}
