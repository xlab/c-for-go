// Package pkg provides a pkg-config(1) like interface for parsing and fetching info from .pc files.
package pkg

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type config struct {
	pcPaths []string
	cflags  []string
	seenMap map[string]bool
}

type Config interface {
	// Locate returns the path of a pkg-config file for the specified pkg name,
	// search goes across all paths defined in PKG_CONFIG_PATH env variable.
	Locate(pkgName string) (pcPath string, err error)
	// Load tries to open and parse a .pc file located at given path, following
	// all the required packages recursively if the follow option is set.
	Load(pcPath string, follow bool) error
	// LoadedPkgNames returns a sorted list of all packages being processed (via required too).
	LoadedPkgNames() []string
	// CFlags returns a list of CFlags collected from all the loaded .pc files.
	CFlags() []string
}

// NewConfig creates a new pkg-config lookup helper, you may specify lookup paths explicitly,
// otherwise the helper will try to get them from PKG_CONFIG_PATH.
func NewConfig(pcPaths []string) (Config, error) {
	cfg := &config{
		seenMap: make(map[string]bool),
	}
	if len(pcPaths) > 0 {
		cfg.pcPaths = pcPaths
	} else {
		pkgConfigPath := os.Getenv("PKG_CONFIG_PATH")
		for _, path := range strings.Split(pkgConfigPath, ":") {
			path = strings.TrimSpace(path)
			if len(path) > 0 {
				cfg.pcPaths = append(pcPaths, path)
			}
		}
		if len(cfg.pcPaths) == 0 {
			return nil, errors.New("PKG_CONFIG_PATH is not set")
		}
	}
	return Config(cfg), nil
}

func (c config) Locate(pkgName string) (string, error) {
	pcName := fmt.Sprintf("%s.pc", pkgName)
	for _, pcPath := range c.pcPaths {
		path := filepath.Join(pcPath, pcName)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, nil
		}
	}
	return "", fmt.Errorf("%s not found in any of PKG_CONFIG_PATH paths", pcName)
}

func (c config) LoadedPkgNames() []string {
	var names []string
	for path := range c.seenMap {
		name := filepath.Base(path)
		names = append(names, strings.TrimSuffix(name, ".pc"))
	}
	sort.Sort(sort.StringSlice(names))
	return names
}

func (c *config) Load(pcPath string, follow bool) error {
	if cflags, err := c.readCflags(pcPath, follow); err != nil {
		return err
	} else {
		c.cflags = append(c.cflags, cflags...)
		c.cflags = uniqueSorted(c.cflags)
		return nil
	}
}

func (c config) CFlags() []string {
	return c.cflags
}

func (c *config) readCflags(pcPath string, follow bool) ([]string, error) {
	if c.seenMap[pcPath] {
		return nil, nil
	}
	data, err := ioutil.ReadFile(pcPath)
	if err != nil {
		return nil, err
	}
	c.seenMap[pcPath] = true
	expandMap := make(map[string]string)
	readVars(data, expandMap)
	cflags := readArgs(data, expandMap, "Cflags:")
	if follow {
		requires := readArgs(data, expandMap, "Requires:")
		for _, pkgName := range requires {
			pcPath, err := c.Locate(pkgName)
			if err != nil {
				return nil, fmt.Errorf("required %s.pc error: %s", pkgName, err.Error())
			}
			if c.seenMap[pcPath] {
				continue
			}
			next, err := c.readCflags(pcPath, true)
			if err != nil {
				return nil, fmt.Errorf("required %s.pc error: %s", pkgName, err.Error())
			}
			cflags = append(cflags, next...)
		}
	}
	return cflags, nil
}

func uniqueSorted(flags []string) []string {
	seen := make(map[string]struct{}, len(flags))
	uniqueFlags := make([]string, 0, len(flags))
	for _, flag := range flags {
		if _, ok := seen[flag]; !ok {
			uniqueFlags = append(uniqueFlags, flag)
			seen[flag] = struct{}{}
		}
	}
	// I don't know how to sort this yet.
	// sort.Sort(sort.StringSlice(uniqueFlags))
	return uniqueFlags
}

func readArgs(data []byte, expandMap map[string]string, header string) []string {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			line = strings.TrimSpace(line)
			if idx := strings.Index(line, header); idx >= 0 {
				return splitArgs(expand(line[idx+len(header):], expandMap))
			}
		}
	}
	return nil
}

func readVars(data []byte, expandMap map[string]string) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			idx := strings.IndexRune(line, '=')
			if idx < 0 {
				continue
			}
			key := strings.TrimSpace(line[:idx])
			if len(key) == 0 {
				continue
			}
			value := strings.TrimSpace(line[idx+1:])
			value = expand(value, expandMap)
			expandMap[key] = value
		}
	}
}

func expand(line string, expandMap map[string]string) string {
	var pairs []string
	for k, v := range expandMap {
		pairs = append(pairs, fmt.Sprintf("${%s}", k), v)
	}
	r := strings.NewReplacer(pairs...)
	return r.Replace(line)
}

func splitArgs(line string) []string {
	var args []string
	line = strings.Replace(line, ",", " ", -1)
	for _, arg := range strings.Split(line, " ") {
		if len(arg) > 0 {
			args = append(args, arg)
		}
	}
	return args
}
