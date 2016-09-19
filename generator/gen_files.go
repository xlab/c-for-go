package generator

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func genLabel() string {
	tpl := "WARNING: This file has automatically been generated on %s.\nBy https://git.io/cgogen. DO NOT EDIT."
	return fmt.Sprintf(tpl, time.Now().Format(time.RFC1123))
}

func (gen *Generator) WriteDoc(wr io.Writer) bool {
	var hasDoc bool
	if len(gen.cfg.PackageLicense) > 0 {
		writeTextBlock(wr, gen.cfg.PackageLicense)
		writeSpace(wr, 1)
		hasDoc = true
	}
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	if len(gen.cfg.PackageDescription) > 0 {
		writeLongTextBlock(wr, gen.cfg.PackageDescription)
		hasDoc = true
	}
	writePackageName(wr, gen.pkg)
	writeSpace(wr, 1)
	return hasDoc
}

func (gen *Generator) WriteIncludes(wr io.Writer) {
	writeStartComment(wr)
	writePkgConfig(wr, gen.cfg.PkgConfigOpts)
	for _, group := range gen.cfg.FlagGroups {
		writeFlagGroup(wr, group)
	}
	for _, path := range gen.cfg.SysIncludes {
		writeSysInclude(wr, path)
	}
	for _, path := range gen.cfg.Includes {
		writeInclude(wr, path)
	}
	writeCStdIncludes(wr, gen.cfg.SysIncludes)
	fmt.Fprintln(wr, `#include "cgo_helpers.h"`)
	writeEndComment(wr)
	fmt.Fprintln(wr, `import "C"`)
	writeSpace(wr, 1)
}

func hasLib(paths []string, lib string) bool {
	for i := range paths {
		if paths[i] == lib {
			return true
		}
	}
	return false
}

func (gen *Generator) writeGoHelpersHeader(wr io.Writer) {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	writePackageName(wr, gen.pkg)
	writeSpace(wr, 1)
	gen.WriteIncludes(wr)
}

func (gen *Generator) writeCHHelpersHeader(wr io.Writer) {
	if len(gen.cfg.PackageLicense) > 0 {
		writeTextBlock(wr, gen.cfg.PackageLicense)
		writeSpace(wr, 1)
	}
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	for _, path := range gen.cfg.SysIncludes {
		writeSysInclude(wr, path)
	}
	for _, path := range gen.cfg.Includes {
		writeInclude(wr, path)
	}
	writeCStdIncludes(wr, gen.cfg.SysIncludes)
	writeCHPragma(wr)
	writeSpace(wr, 1)
}

func (gen *Generator) writeCCHelpersHeader(wr io.Writer) {
	if len(gen.cfg.PackageLicense) > 0 {
		writeTextBlock(wr, gen.cfg.PackageLicense)
		writeSpace(wr, 1)
	}
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	writeCGOIncludes(wr)
	writeSpace(wr, 1)
}

func writeCGOIncludes(wr io.Writer) {
	fmt.Fprintln(wr, `#include "_cgo_export.h"`)
	fmt.Fprintln(wr, `#include "cgo_helpers.h"`)
}

func writeCHPragma(wr io.Writer) {
	fmt.Fprintln(wr, "#pragma once")
}

func writeCStdIncludes(wr io.Writer, sysIncludes []string) {
	if !hasLib(sysIncludes, "stdlib.h") {
		fmt.Fprintln(wr, "#include <stdlib.h>")
	}
	// if !hasLib(sysIncludes, "stdbool.h") {
	// 	fmt.Fprintln(wr, "#include <stdbool.h>")
	// }
}

func (gen *Generator) WritePackageHeader(wr io.Writer) {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	writePackageName(wr, gen.pkg)
	writeSpace(wr, 1)
}

func writeFlagGroup(wr io.Writer, group TraitFlagGroup) {
	if len(group.Name) == 0 {
		return
	}
	if len(group.Flags) == 0 {
		return
	}
	if len(group.Traits) == 0 {
		fmt.Fprintf(wr, "#cgo %s: %s\n", group.Name, strings.Join(group.Flags, " "))
		return
	}
	traits := strings.Join(group.Traits, " ")
	fmt.Fprintf(wr, "#cgo %s %s: %s\n", traits, group.Name, strings.Join(group.Flags, " "))
}

func writeSysInclude(wr io.Writer, path string) {
	fmt.Fprintf(wr, "#include <%s>\n", path)
}

func writeInclude(wr io.Writer, path string) {
	fmt.Fprintf(wr, "#include \"%s\"\n", path)
}

func writePkgConfig(wr io.Writer, opts []string) {
	if len(opts) == 0 {
		return
	}
	fmt.Fprintf(wr, "#cgo pkg-config: %s\n", strings.Join(opts, " "))
}

func writeStartComment(wr io.Writer) {
	fmt.Fprintln(wr, "/*")
}

func writeEndComment(wr io.Writer) {
	fmt.Fprintln(wr, "*/")
}

func writePackageName(wr io.Writer, name string) {
	if len(name) == 0 {
		name = "main"
	}
	fmt.Fprintf(wr, "package %s\n", name)
}

func writeLongTextBlock(wr io.Writer, text string) {
	if len(text) == 0 {
		return
	}
	writeStartComment(wr)
	fmt.Fprint(wr, text)
	writeSpace(wr, 1)
	writeEndComment(wr)
}

func writeTextBlock(wr io.Writer, text string) {
	if len(text) == 0 {
		return
	}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Fprintf(wr, "// %s\n", line)
	}
}

func writeSourceBlock(wr io.Writer, src string) {
	if len(src) == 0 {
		return
	}
	fmt.Fprint(wr, src)
	writeSpace(wr, 1)
}
