package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/tj/go-spin"
)

var (
	outputPath = flag.String("out", "", "Specify a `dir` for the output.")
	noCGO      = flag.Bool("nocgo", false, "Do not include a cgo-specific header in resulting files.")
	ccDefs     = flag.Bool("ccdefs", false, "Use built-in defines from a hosted C-compiler.")
	ccIncl     = flag.Bool("ccincl", false, "Use built-in sys include paths from a hosted C-compiler.")
	maxMem     = flag.String("maxmem", "0x7fffffff", "Specifies platform's memory cap the generated code.")
	fancy      = flag.Bool("fancy", true, "Enable fancy output in the term.")
	nostamp    = flag.Bool("nostamp", false, "Disable printing timestamps in the output files.")
	debug      = flag.Bool("debug", false, "Enable some debug info.")
)

const logo = `Copyright (c) 2015-2017 Maxim Kupriianov <max@kc.vc>
Based on a C99 compiler front end by Jan Mercl <0xjnml@gmail.com>
`

func init() {
	if *debug {
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetFlags(0)
	}
	flag.Usage = func() {
		fmt.Println(logo)
		fmt.Printf("Usage: %s package1.yml [package2.yml] ...\n", os.Args[0])
		fmt.Printf("See https://github.com/xlab/c-for-go for examples and documentation.\n\n")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		fmt.Println()
		log.Fatalln("[ERR] no package configuration files have been provided.")
	}
}

func main() {
	s := spin.New()

	var wg sync.WaitGroup
	doneChan := make(chan struct{})
	for _, cfgPath := range getConfigPaths() {
		if *fancy {
			wg.Add(1)
			go func() {
				for {
					select {
					case <-doneChan:
						doneChan = make(chan struct{})
						fmt.Printf("\r  \033[36mprocessing %s\033[m done.\n", cfgPath)
						wg.Done()
						return
					default:
						fmt.Printf("\r  \033[36mprocessing %s\033[m %s", cfgPath, s.Next())
						time.Sleep(100 * time.Millisecond)
					}
				}
			}()
		}

		var t0 time.Time
		if *debug {
			t0 = time.Now()
		}
		process, err := NewProcess(cfgPath, *outputPath)
		if err != nil {
			log.Fatalln("[ERR]", err)
		}
		process.Generate(*noCGO)
		if err := process.Flush(*noCGO); err != nil {
			log.Fatalln("[ERR]", err)
		}
		if *debug {
			fmt.Printf("done in %v\n", time.Now().Sub(t0))
		}
		if *fancy {
			close(doneChan)
			wg.Wait()
		}
	}
}

func getConfigPaths() (paths []string) {
	for _, path := range flag.Args() {
		if info, err := os.Stat(path); err != nil {
			log.Fatalln("[ERR] cannot locate the specified path:", path)
		} else if info.IsDir() {
			if path, ok := configFromDir(path); ok {
				paths = append(paths, path)
				continue
			}
			log.Fatalln("[ERR] cannot find any config file in:", path)
		}
		paths = append(paths, path)
	}
	return
}

func configFromDir(path string) (string, bool) {
	possibleNames := []string{"c-for-go.yaml", "c-for-go.yml"}
	if base := filepath.Base(path); len(base) > 0 {
		possibleNames = append(possibleNames,
			fmt.Sprintf("%s.yaml", base), fmt.Sprintf("%s.yml", base))
	}
	for _, name := range possibleNames {
		path := filepath.Join(path, name)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, true
		}
	}
	return "", false
}
