package parser

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
)

type Config struct {
	Arch         string   `yaml:"Arch"`
	IncludePaths []string `yaml:"IncludePaths"`
	SourcesPaths []string `yaml:"SourcesPaths"`
	IgnoredPaths []string `yaml:"IgnoredPaths"`

	Defines map[string]interface{} `yaml:"Defines"`

	CCDefs bool `yaml:"-"`
	CCIncl bool `yaml:"-"`
}

func ParseWith(cfg *Config) (*cc.AST, error) {
	if len(cfg.SourcesPaths) == 0 {
		return nil, errors.New("parser: no target paths specified")
	}
	cfg, err := checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	var predefined string
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
	if cfg.CCDefs || cfg.CCIncl {
		cppPath := "cpp"
		if v, ok := os.LookupEnv("CPP"); ok {
			cppPath = v
		} else if v, ok = os.LookupEnv("CC"); ok {
			cppPath = v
		}
		// if expanded, err := exec.LookPath(cppPath); err == nil {
		// 	cppPath = expanded
		// }
		ccpredef, _, sysIncludePaths, err := hostCppConfig(cppPath)
		if err != nil {
			log.Println("[WARN] `cpp -dM` failed:", err)
		} else {
			if cfg.CCIncl && len(sysIncludePaths) > 0 {
				// add on top of sysIncludePaths if allowed by config
				cfg.IncludePaths = append(sysIncludePaths, cfg.IncludePaths...)
			}
			ccDefs = ccpredef
			ccDefsOK = true
		}
	}
	if cfg.CCDefs && ccDefsOK {
		predefined += fmt.Sprintf("\n%s", ccDefs)
	}
	ccConfig, _ := cc.NewConfig(runtime.GOOS, cfg.Arch)
	// Let cc provide all predefines and builtins. Only append custom definitions.
	ccConfig.Predefined += predefined
	ccConfig.IncludePaths = append(ccConfig.IncludePaths, cfg.IncludePaths...)
	var sources []cc.Source
	sources = append(sources, cc.Source{Name: "<predefined>", Value: ccConfig.Predefined})
	sources = append(sources, cc.Source{Name: "<builtin>", Value: cc.Builtin})
	for _, sourceEntry := range cfg.SourcesPaths {
		sources = append(sources, cc.Source{
			Name: sourceEntry,
		})
	}
	return cc.Translate(ccConfig, sources)
}

func checkConfig(cfg *Config) (*Config, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	if len(cfg.Arch) == 0 {
		cfg.Arch = runtime.GOARCH
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

// hostCppConfig returns the system C preprocessor configuration, or an error,
// if any.  The configuration is obtained by running the cpp command. For the
// predefined macros list the '-dM' options is added. For the include paths
// lists, the option '-v' is added and the output is parsed to extract the
// "..." include and <...> include paths. To add any other options to cpp, list
// them in opts.
//
// The function relies on a POSIX compatible C preprocessor installed.
// Execution of HostConfig is not free, so caching the results is recommended
// whenever possible.
func hostCppConfig(cpp string, opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) {
	nullPath := "/dev/null"
	newLine := "\n"
	if runtime.GOOS == "windows" {
		nullPath = "nul"
		newLine = "\r\n"
	}
	args := append(append([]string{"-dM"}, opts...), nullPath)
	pre, err := exec.Command(cpp, args...).CombinedOutput()
	if err != nil {
		return "", nil, nil, err
	}

	args = append(append([]string{"-v"}, opts...), nullPath)
	out, err := exec.Command(cpp, args...).CombinedOutput()
	if err != nil {
		return "", nil, nil, err
	}

	a := strings.Split(string(out), newLine)
	for i := 0; i < len(a); {
		switch a[i] {
		case "#include \"...\" search starts here:":
		loop:
			for i = i + 1; i < len(a); {
				switch v := a[i]; {
				case strings.HasPrefix(v, "#") || v == "End of search list.":
					break loop
				default:
					includePaths = append(includePaths, strings.TrimSpace(v))
					i++
				}
			}
		case "#include <...> search starts here:":
			for i = i + 1; i < len(a); {
				switch v := a[i]; {
				case strings.HasPrefix(v, "#") || v == "End of search list.":
					return string(pre), includePaths, sysIncludePaths, nil
				default:
					sysIncludePaths = append(sysIncludePaths, strings.TrimSpace(v))
					i++
				}
			}
		default:
			i++
		}
	}
	return "", nil, nil, fmt.Errorf("failed parsing %s -v output", cpp)
}
