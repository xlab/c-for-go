package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	outputPath = flag.String("out", "", "Specify the `directory` of output files.")
)

var packageName string

func init() {
	log.SetFlags(0)
	flag.Usage = func() {
		log.Printf("Usage: %s <config1> [config2] ...\n", os.Args[0])
		log.Println("Config is either .yml or .yaml file, defining a standalone package.")
		log.Println("See http://xxx for examples an documentation.")
		log.Println("\nOptions:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		Errorf("no configs supplied")
	}
}

func main() {
	for _, cfgPath := range getConfigPaths() {
		log.Printf("processing %s ...", cfgPath)
		ts := time.Now()
		cgogen, err := NewCGOGen(cfgPath, *outputPath)
		if err != nil {
			Errorf(err.Error())
		}
		cgogen.Generate()
		if err := cgogen.Flush(); err != nil {
			Errorf(err.Error())
		}
		log.Printf("done in %v\n", time.Now().Sub(ts))
	}
}

func getConfigPaths() (paths []string) {
	for _, path := range flag.Args() {
		if info, err := os.Stat(path); err != nil {
			Errorf("cannot locate the specified path: %s", path)
		} else if info.IsDir() {
			if path, ok := configFromDir(path); ok {
				paths = append(paths, path)
				continue
			}
			Errorf("cannot find any config file in: %s", path)
		}
		paths = append(paths, path)
	}
	return
}

func configFromDir(path string) (string, bool) {
	possibleNames := []string{"cgogen.yaml", "cgogen.yml"}
	if base := filepath.Base(path); len(base) > 0 {
		possibleNames = append(possibleNames,
			fmt.Sprintf("%s.yaml", base), fmt.Sprintf("%s.yaml", base))
	}
	for _, name := range possibleNames {
		path := filepath.Join(path, name)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, true
		}
	}
	return "", false
}

func Errorf(format string, a ...interface{}) {
	log.Fatalf("[ERROR]: "+format+"\n", a...)
}
