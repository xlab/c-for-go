package generator

import (
	"fmt"
	"io"
	"strings"
)

func (gen *Generator) WriteDoc(wr io.Writer) {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writeLongTextBlock(wr, gen.cfg.PackageDescription)
	writePackageName(wr, gen.cfg.PackageName)
}

func (gen *Generator) WriteIncludes(wr io.Writer) error {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writePackageName(wr, gen.cfg.PackageName)
	writeSpace(wr, 1)
	writeStartComment(wr)
	writePkgConfig(wr, gen.cfg.PkgConfigOpts)
	writeFlagSet(wr, gen.cfg.CPPFlags)
	writeFlagSet(wr, gen.cfg.CXXFlags)
	writeFlagSet(wr, gen.cfg.CFlags)
	writeFlagSet(wr, gen.cfg.LDFlags)
	writeSpace(wr, 1)
	for _, path := range gen.cfg.SysIncludes {
		writeSysInclude(wr, path)
	}
	for _, path := range gen.cfg.Includes {
		writeInclude(wr, path)
	}
	writeEndComment(wr)
	fmt.Fprintln(wr, `include "C"`)
	return nil
}

func (gen *Generator) WritePackage(wr io.Writer) {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writePackageName(wr, gen.cfg.PackageName)
}

func writeFlagSet(wr io.Writer, flags ArchFlagSet) {
	if len(flags.Name) == 0 {
		return
	}
	if len(flags.Flags) == 0 {
		return
	}
	fmt.Fprintf(wr, "#cgo %s %s: %s\n",
		strings.Join(flags.Arch, ","),
		flags.Name,
		strings.Join(flags.Flags, " "),
	)
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

func writeSpace(wr io.Writer, n int) {
	fmt.Fprint(wr, strings.Repeat("\n", n))
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

func writeError(wr io.Writer, err error) {
	fmt.Fprintf(wr, "// error: %v\n", err)
}
