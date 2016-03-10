package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/xlab/cgogen/generator"
	"github.com/xlab/cgogen/parser"
	"github.com/xlab/cgogen/translator"
	"github.com/xlab/pkgconfig/pkg"
	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v2"
)

type Buf int

const (
	BufDoc Buf = iota
	BufInludes
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

type CGOGen struct {
	cfg          CGOGenConfig
	gen          *generator.Generator
	genSync      sync.WaitGroup
	goBuffers    map[Buf]*bytes.Buffer
	chHelpersBuf *bytes.Buffer
	ccHelpersBuf *bytes.Buffer
	outputPath   string
}

type CGOGenConfig struct {
	Generator  *generator.Config  `yaml:"GENERATOR"`
	Translator *translator.Config `yaml:"TRANSLATOR"`
	Parser     *parser.Config     `yaml:"PARSER"`
}

func NewCGOGen(configPath, outputPath string) (*CGOGen, error) {
	cfgData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg CGOGenConfig
	if err := yaml.Unmarshal(cfgData, &cfg); err != nil {
		return nil, err
	}
	if cfg.Generator != nil {
		paths := includePathsFromPkgConfig(cfg.Generator.PkgConfigOpts)
		if cfg.Parser == nil {
			cfg.Parser = &parser.Config{}
		}
		cfg.Parser.IncludePaths = append(cfg.Parser.IncludePaths, paths...)
		cfg.Parser.IncludePaths = append(cfg.Parser.IncludePaths, filepath.Dir(configPath))
	} else {
		return nil, errors.New("cgogen: generator config was not specified")
	}

	// parse the headers
	unit, err := parser.ParseWith(cfg.Parser)
	if err != nil {
		return nil, err
	}

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
	c := &CGOGen{
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

func (c *CGOGen) Generate(noCGO bool) {
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

func (c *CGOGen) Flush(noCGO bool) error {
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
