package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/mflag.v1"
)

var (
	configPath = mflag.String([]string{"c", "-config", "-cfg"}, "", "Specify the config file path, defaults to <package>.yml or cgogen.yml.")
	outputPath = mflag.String([]string{"o", "-out", "-out-dir"}, "", "Specify the output directory")
)

var packageName string

func init() {
	log.SetFlags(0)
	mflag.Usage = func() {
		log.Printf("Usage: %s <package>\n", os.Args[0])
		mflag.PrintDefaults()
	}
	mflag.Parse()
	if packageName = mflag.Arg(0); len(packageName) == 0 {
		mflag.Usage()
		Errorf("no package name specified")
	}
}

func main() {
	cgogen, err := NewCGOGen(packageName, getConfigPath(), *outputPath)
	if err != nil {
		Errorf(err.Error())
	}
	cgogen.Generate()
}

func getConfigPath() (str string) {
	if path := *configPath; len(path) > 0 {
		if info, err := os.Stat(path); err != nil || info.IsDir() {
			Errorf("cannot locate the specified config file: %s", path)
		}
		return path
	}
	paths := []string{
		fmt.Sprintf("%s.yaml", filepath.Base(packageName)),
		fmt.Sprintf("%s.yml", filepath.Base(packageName)),
		"cgogen.yaml",
		"cgogen.yml",
	}
	for _, path := range paths {
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path
		}
	}
	Errorf("config path was not specified, also could not locate neither of: %v", strings.Join(paths, ", "))
	return
}

func Errorf(format string, a ...interface{}) {
	log.Fatalf("[ERROR]: "+format+"\n", a...)
}
