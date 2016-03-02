package generator

import (
	"bytes"
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeDefinesGroup(wr io.Writer, defines []*tl.CDecl) (n int) {
	writeStartConst(wr)
	for _, decl := range defines {
		if !decl.IsDefine {
			continue
		}
		if len(decl.Expression) == 0 {
			decl.Expression = skipNameStr
		}
		name := gen.tr.TransformName(tl.TargetConst, decl.Name)
		fmt.Fprintf(wr, "// %s as defined in %s\n", name,
			gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
		if decl.Value != nil {
			fmt.Fprintf(wr, "%s = %v", name, decl.Value)
		} else {
			fmt.Fprintf(wr, "%s = %s", name, decl.Expression)
		}
		writeSpace(wr, 1)
		n++
	}
	writeEndConst(wr)
	return
}

func (gen *Generator) writeConstDeclaration(wr io.Writer, decl *tl.CDecl) {
	declName := gen.tr.TransformName(tl.TargetConst, decl.Name)
	fmt.Fprintf(wr, "// %s as declared in %s\n", declName,
		gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
	if len(decl.Expression) == 0 {
		decl.Expression = skipNameStr
	}
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	if decl.Value != nil {
		fmt.Fprintf(wr, "const %s %s = %v", declName, goSpec, decl.Value)
		return
	}
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
		fmt.Fprintf(wr, "// %s as declared in %s\n", typeName,
			gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
		fmt.Fprintf(wr, "type %s %s\n", typeName, enumType)
		writeSpace(wr, 1)
		fmt.Fprintf(wr, "// %s enumeration from %s\n", typeName,
			gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
	}
	writeStartConst(wr)
	for _, m := range spec.Members {
		if !gen.tr.IsAcceptableName(tl.TargetConst, m.Name) {
			continue
		}
		mName := gen.tr.TransformName(tl.TargetConst, m.Name)
		if len(mName) == 0 {
			continue
		} else if namesSeen[string(mName)] {
			continue
		} else {
			namesSeen[string(mName)] = true
		}
		if !hasType {
			fmt.Fprintf(wr, "// %s as declared in %s\n", mName,
				gen.tr.SrcLocation(tl.TargetConst, m.Name, m.Pos))
		}
		if len(m.Expression) == 0 {
			fmt.Fprintf(wr, "%s\n", mName)
			continue
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
	fmt.Fprintf(wr, "// %s as declared in %s\n", tagName,
		gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
	fmt.Fprintf(wr, "type %s %s\n", tagName, enumType)
	writeSpace(wr, 1)
	if isTypedef {
		if !bytes.Equal(tagName, declName) && len(declName) > 0 {
			// alias type decl name to the tag
			fmt.Fprintf(wr, "// %s as declared in %s\n", declName,
				gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
			fmt.Fprintf(wr, "type %s %s", declName, tagName)
			writeSpace(wr, 1)
		}
	}

	fmt.Fprintf(wr, "// %s enumeration from %s\n", tagName,
		gen.tr.SrcLocation(tl.TargetConst, decl.Name, decl.Pos))
	writeStartConst(wr)
	for _, m := range spec.Members {
		if !gen.tr.IsAcceptableName(tl.TargetConst, m.Name) {
			continue
		}
		mName := gen.tr.TransformName(tl.TargetConst, m.Name)
		if len(mName) == 0 {
			continue
		}
		if len(m.Expression) == 0 {
			fmt.Fprintf(wr, "%s\n", mName)
			continue
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
