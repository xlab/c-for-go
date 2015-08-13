package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/xlab/cgogen/generator"
	"github.com/xlab/cgogen/parser"
	"github.com/xlab/cgogen/translator"
)

type WriterOpt int

const (
	WriterDoc WriterOpt = iota
	WriterInludes
	WriterConst
	WriterTypes
	WriterUnions
	WriterMain
)

var writerNames = map[WriterOpt]string{
	WriterDoc:     "doc",
	WriterInludes: "includes",
	WriterConst:   "const",
	WriterTypes:   "types",
	WriterUnions:  "unions",
}

type CGOGen struct {
	gen     *generator.Generator
	writers map[WriterOpt]*os.File
}

type CGOGenConfig struct {
	Generator  *generator.Config
	Translator *translator.Config
	Parser     *parser.Config
}

func NewCGOGen(packageName, configPath, outputPath string) (*CGOGen, error) {
	cfgData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg CGOGenConfig
	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		return nil, err
	}
	p, err := parser.New(cfg.Parser)
	if err != nil {
		return nil, err
	}
	unit, err := p.Parse()
	if err != nil {
		return nil, err
	}
	tl, err := translator.New(cfg.Translator)
	if err != nil {
		return nil, err
	}
	if err := tl.Learn(unit); err != nil {
		return nil, err
	}
	pkg := filepath.Base(packageName)
	gen, err := generator.New(pkg, cfg.Generator, tl)
	if err != nil {
		return nil, err
	}
	cgogen := &CGOGen{
		gen:     gen,
		writers: make(map[WriterOpt]*os.File),
	}
	filePrefix := filepath.Join(outputPath, packageName)
	if err := os.MkdirAll(filePrefix, 0655); err != nil {
		return nil, err
	}
	if f, err := os.Create(filepath.Join(filePrefix, fmt.Sprintf("%s.go", pkg))); err != nil {
		return nil, err
	} else {
		cgogen.writers[WriterMain] = f
	}
	for opt, name := range writerNames {
		if f, err := os.Create(filepath.Join(filePrefix, fmt.Sprintf("%s.go", name))); err == nil {
			cgogen.writers[opt] = f
		}
	}
	return cgogen, nil
}

func (c *CGOGen) Generate() {
	main := c.writers[WriterMain]
	if wr, ok := c.writers[WriterDoc]; ok {
		c.gen.WriteDoc(wr)
		c.gen.WritePackageHeader(main)
	} else {
		c.gen.WriteDoc(main)
	}
	if wr, ok := c.writers[WriterInludes]; ok {
		c.gen.WriteIncludes(wr)
	} else {
		c.gen.WriteIncludes(main)
	}
	if wr, ok := c.writers[WriterConst]; ok {
		c.gen.WriteConst(wr)
	} else {
		c.gen.WriteConst(main)
	}
	if wr, ok := c.writers[WriterTypes]; ok {
		c.gen.WriteTypedefs(wr)
	} else {
		c.gen.WriteTypedefs(main)
	}
	if wr, ok := c.writers[WriterUnions]; ok {
		c.gen.WriteUnions(wr)
	} else {
		c.gen.WriteUnions(main)
	}
	c.gen.WriteDeclares(main)
}

func (c *CGOGen) Close() {
	for _, w := range c.writers {
		w.Close()
	}
}
