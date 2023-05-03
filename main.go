package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/xlab/c-for-go/generator"
	"github.com/xlab/c-for-go/parser"
	"github.com/xlab/c-for-go/translator"
	"github.com/xlab/pkgconfig/pkg"
	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"modernc.org/cc/v4"
	"os"
	"path/filepath"
	"strings"
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
		fmt.Print(logo)
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

type Buf int

const (
	BufDoc Buf = iota
	BufConst
	BufTypes
	BufUnions
	BufHelpers
	BufMain
)

var goBufferNames = map[Buf]string{
	BufDoc:     "doc",
	BufConst:   "const",
	BufTypes:   "types",
	BufUnions:  "unions",
	BufHelpers: "cgo_helpers",
}

type Process struct {
	cfg          ProcessConfig
	gen          *generator.Generator
	genSync      sync.WaitGroup
	goBuffers    map[Buf]*bytes.Buffer
	chHelpersBuf *bytes.Buffer
	ccHelpersBuf *bytes.Buffer
	outputPath   string
}

type ProcessConfig struct {
	Generator  *generator.Config  `yaml:"GENERATOR"`
	Translator *translator.Config `yaml:"TRANSLATOR"`
	Parser     *parser.Config     `yaml:"PARSER"`
}

func NewProcess(configPath, outputPath string) (*Process, error) {
	cfgData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg ProcessConfig
	if err := yaml.Unmarshal(cfgData, &cfg); err != nil {
		return nil, err
	}
	if cfg.Generator != nil {
		paths := includePathsFromPkgConfig(cfg.Generator.PkgConfigOpts)
		if cfg.Parser == nil {
			cfg.Parser = &parser.Config{}
		}
		cfg.Parser.CCDefs = *ccDefs
		cfg.Parser.CCIncl = *ccIncl
		cfg.Parser.IncludePaths = append(cfg.Parser.IncludePaths, paths...)
		cfg.Parser.IncludePaths = append(cfg.Parser.IncludePaths, filepath.Dir(configPath))
	} else {
		return nil, errors.New("process: generator config was not specified")
	}

	// parse the headers
	unit, err := parser.ParseWith(cfg.Parser)
	if err != nil {
		return nil, err
	}

	if cfg.Translator == nil {
		cfg.Translator = &translator.Config{}
	}
	cfg.Translator.IgnoredFiles = cfg.Parser.IgnoredPaths
	cfg.Translator.LongIs64Bit = unit.ABI.Types[cc.Long].Size == 8
	// learn the model
	tl, err := translator.New(cfg.Translator)
	if err != nil {
		return nil, err
	}
	tl.Learn(unit)

	// begin generation
	pkg := filepath.Base(cfg.Generator.PackageName)
	gen, err := generator.New(pkg, cfg.Generator, tl)
	if err != nil {
		return nil, err
	}
	gen.SetMaxMemory(generator.NewMemSpec(*maxMem))

	if *nostamp {
		gen.DisableTimestamps()
	}
	c := &Process{
		cfg:          cfg,
		gen:          gen,
		goBuffers:    make(map[Buf]*bytes.Buffer),
		chHelpersBuf: new(bytes.Buffer),
		ccHelpersBuf: new(bytes.Buffer),
		outputPath:   outputPath,
	}
	c.goBuffers[BufMain] = new(bytes.Buffer)
	for opt := range goBufferNames {
		c.goBuffers[opt] = new(bytes.Buffer)
	}
	goHelpersBuf := c.goBuffers[BufHelpers]
	go func() {
		c.genSync.Add(1)
		c.gen.MonitorAndWriteHelpers(goHelpersBuf, c.chHelpersBuf, c.ccHelpersBuf)
		c.genSync.Done()
	}()
	return c, nil
}

func (c *Process) Generate(noCGO bool) {
	main := c.goBuffers[BufMain]
	if wr, ok := c.goBuffers[BufDoc]; ok {
		if !c.gen.WriteDoc(wr) {
			c.goBuffers[BufDoc] = nil
		}
		c.gen.WritePackageHeader(main)
	} else {
		c.gen.WriteDoc(main)
	}
	if !noCGO {
		c.gen.WriteIncludes(main)
	}
	if wr, ok := c.goBuffers[BufConst]; ok {
		c.gen.WritePackageHeader(wr)
		if !noCGO {
			c.gen.WriteIncludes(wr)
		}
		if n := c.gen.WriteConst(wr); n == 0 {
			c.goBuffers[BufConst] = nil
		}
	} else {
		c.gen.WriteConst(main)
	}
	if wr, ok := c.goBuffers[BufTypes]; ok {
		c.gen.WritePackageHeader(wr)
		if !noCGO {
			c.gen.WriteIncludes(wr)
		}
		if n := c.gen.WriteTypedefs(wr); n == 0 {
			c.goBuffers[BufTypes] = nil
		}
	} else {
		c.gen.WriteTypedefs(main)
	}
	if !noCGO {
		if wr, ok := c.goBuffers[BufUnions]; ok {
			c.gen.WritePackageHeader(wr)
			c.gen.WriteIncludes(wr)
			if n := c.gen.WriteUnions(wr); n == 0 {
				c.goBuffers[BufUnions] = nil
			}
		} else {
			c.gen.WriteUnions(main)
		}
		c.gen.WriteDeclares(main)
	}
}

func (c *Process) Flush(noCGO bool) error {
	c.gen.Close()
	c.genSync.Wait()
	filePrefix := filepath.Join(c.outputPath, c.cfg.Generator.PackageName)
	if err := os.MkdirAll(filePrefix, 0755); err != nil {
		return err
	}
	createCHFile := func(name string) (*os.File, error) {
		return os.Create(filepath.Join(filePrefix, fmt.Sprintf("%s.h", name)))
	}
	createCCFile := func(name string) (*os.File, error) {
		return os.Create(filepath.Join(filePrefix, fmt.Sprintf("%s.c", name)))
	}
	createGoFile := func(name string) (*os.File, error) {
		return os.Create(filepath.Join(filePrefix, fmt.Sprintf("%s.go", name)))
	}
	writeGoFile := func(opt Buf, name string) error {
		if buf := c.goBuffers[opt]; buf != nil && buf.Len() > 0 {
			if f, err := createGoFile(name); err == nil {
				if err := flushBufferToFile(buf.Bytes(), f, true); err != nil {
					f.Close()
					return err
				}
				f.Close()
			} else {
				return err
			}
		}
		return nil
	}
	writeCHFile := func(buf *bytes.Buffer, name string) error {
		if f, err := createCHFile(name); err != nil {
			return err
		} else if err := flushBufferToFile(buf.Bytes(), f, false); err != nil {
			f.Close()
			return err
		} else {
			return f.Close()
		}
	}
	writeCCFile := func(buf *bytes.Buffer, name string) error {
		if f, err := createCCFile(name); err != nil {
			return err
		} else if err := flushBufferToFile(buf.Bytes(), f, false); err != nil {
			f.Close()
			return err
		} else {
			return f.Close()
		}
	}

	if !noCGO {
		pkg := filepath.Base(c.cfg.Generator.PackageName)
		if err := writeGoFile(BufMain, pkg); err != nil {
			return err
		}
	}
	for opt, name := range goBufferNames {
		if err := writeGoFile(opt, name); err != nil {
			return err
		}
	}
	if noCGO {
		return nil
	}
	if c.chHelpersBuf.Len() > 0 {
		if err := writeCHFile(c.chHelpersBuf, "cgo_helpers"); err != nil {
			return err
		}
	}
	if c.ccHelpersBuf.Len() > 0 {
		if err := writeCCFile(c.ccHelpersBuf, "cgo_helpers"); err != nil {
			return err
		}
	}
	return nil
}

func flushBufferToFile(buf []byte, f *os.File, fmt bool) error {
	if fmt {
		if fmtBuf, err := imports.Process(f.Name(), buf, nil); err == nil {
			_, err = f.Write(fmtBuf)
			return err
		} else {
			log.Printf("[WARN] cannot gofmt %s: %s\n", f.Name(), err.Error())
			f.Write(buf)
			return nil
		}
	}
	_, err := f.Write(buf)
	return err
}

func includePathsFromPkgConfig(opts []string) []string {
	if len(opts) == 0 {
		return nil
	}
	pc, err := pkg.NewConfig(nil)
	if err != nil {
		log.Println("[WARN]", err)
		return nil
	}
	for _, opt := range opts {
		if strings.HasPrefix(opt, "-") || strings.HasPrefix(opt, "--") {
			continue
		}
		if pcPath, err := pc.Locate(opt); err == nil {
			if err := pc.Load(pcPath, true); err != nil {
				log.Println("[WARN] pkg-config:", err)
			}
		} else {
			log.Printf("[WARN] %s.pc referenced in pkg-config options but cannot be found: %s", opt, err.Error())
		}
	}
	flags := pc.CFlags()
	includePaths := make([]string, 0, len(flags))
	for _, flag := range flags {
		if idx := strings.Index(flag, "-I"); idx >= 0 {
			includePaths = append(includePaths, strings.TrimSpace(flag[idx+2:]))
		}
	}
	return includePaths
}
