package generator

import (
	"bytes"
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeDefinesGroup(wr io.Writer, defines []*tl.CDecl) {
	if len(defines) == 0 {
		return
	}
	writeStartConst(wr)
	for _, decl := range defines {
		if !decl.IsDefine {
			continue
		}
		if len(decl.Expression) == 0 {
			decl.Expression = skipNameStr
		}
		name := gen.tr.TransformName(tl.TargetConst, decl.Name)
		fmt.Fprintf(wr, "// %s as defined in %s\n", name, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "%s = %s", name, decl.Expression)
		writeSpace(wr, 1)
	}
	writeEndConst(wr)
}

func (gen *Generator) writeConstDeclaration(wr io.Writer, decl *tl.CDecl) {
	declName := gen.tr.TransformName(tl.TargetConst, decl.Name)
	fmt.Fprintf(wr, "// %s as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
	if len(decl.Expression) == 0 {
		decl.Expression = skipNameStr
	}
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	fmt.Fprintf(wr, "const %s %s = %s", declName, goSpec, decl.Expression)
}

func (gen *Generator) expandEnumAnonymous(wr io.Writer, decl *tl.CDecl, namesSeen map[string]bool) {
	var typeName []byte
	var hasType bool
	if decl.IsTypedef {
		if typeName = gen.tr.TransformName(tl.TargetType, decl.Name); len(typeName) > 0 {
			hasType = true
		}
	}

	spec := decl.Spec.(*tl.CEnumSpec)
	if hasType {
		enumType := gen.tr.TranslateSpec(&spec.Type)
		fmt.Fprintf(wr, "// %s as declared in %s\n", typeName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "type %s %s\n", typeName, enumType)
		writeSpace(wr, 1)
		fmt.Fprintf(wr, "// %s enumeration from %s\n", typeName, tl.SrcLocation(decl.Pos))
	}
	writeStartConst(wr)
	for _, m := range spec.Members {
		mName := gen.tr.TransformName(tl.TargetConst, m.Name)
		if len(mName) == 0 {
			continue
		} else if namesSeen[string(mName)] {
			continue
		} else {
			namesSeen[string(mName)] = true
		}
		if !hasType {
			fmt.Fprintf(wr, "// %s as declared in %s\n", mName, tl.SrcLocation(m.Pos))
		}
		if len(m.Expression) == 0 {
			m.Expression = skipNameStr
		}
		if hasType {
			fmt.Fprintf(wr, "%s %s = %s\n", mName, typeName, m.Expression)
			continue
		}
		fmt.Fprintf(wr, "%s = %s\n", mName, m.Expression)
	}
	writeEndConst(wr)
	writeSpace(wr, 1)
}

func (gen *Generator) expandEnum(wr io.Writer, decl *tl.CDecl) {
	var declName []byte
	var isTypedef bool
	if decl.IsTypedef {
		if declName = gen.tr.TransformName(tl.TargetType, decl.Name); len(declName) > 0 {
			isTypedef = true
		}
	}

	spec := decl.Spec.(*tl.CEnumSpec)
	tagName := gen.tr.TransformName(tl.TargetType, decl.Spec.GetBase())
	enumType := gen.tr.TranslateSpec(&spec.Type)
	fmt.Fprintf(wr, "// %s as declared in %s\n", tagName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s\n", tagName, enumType)
	writeSpace(wr, 1)
	if isTypedef {
		if !bytes.Equal(tagName, declName) && len(declName) > 0 {
			// alias type decl name to the tag
			fmt.Fprintf(wr, "// %s as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
			fmt.Fprintf(wr, "type %s %s", declName, tagName)
			writeSpace(wr, 1)
		}
	}

	fmt.Fprintf(wr, "// %s enumeration from %s\n", tagName, tl.SrcLocation(decl.Pos))
	writeStartConst(wr)
	for _, m := range spec.Members {
		mName := gen.tr.TransformName(tl.TargetConst, m.Name)
		if len(mName) == 0 {
			continue
		}
		if len(m.Expression) == 0 {
			m.Expression = skipNameStr
		}
		fmt.Fprintf(wr, "%s %s = %s\n", mName, declName, m.Expression)
	}
	writeEndConst(wr)
	writeSpace(wr, 1)
}

func writeStartConst(wr io.Writer) {
	fmt.Fprintln(wr, "const (")
}

func writeEndConst(wr io.Writer) {
	fmt.Fprintln(wr, ")")
}
