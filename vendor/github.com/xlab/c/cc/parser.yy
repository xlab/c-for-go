%{
// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.10. Substantial portions of expression AST size
// optimizations are from [2], license of which follows.

// ----------------------------------------------------------------------------

// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This grammar is derived from the C grammar in the 'ansitize'
// program, which carried this notice:
// 
// Copyright (c) 2006 Russ Cox, 
// 	Massachusetts Institute of Technology
// 
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the
// Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute,
// sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
// 
// The above copyright notice and this permission notice shall
// be included in all copies or substantial portions of the
// Software.
// 
// The software is provided "as is", without warranty of any
// kind, express or implied, including but not limited to the
// warranties of merchantability, fitness for a particular
// purpose and noninfringement.  In no event shall the authors
// or copyright holders be liable for any claim, damages or
// other liability, whether in an action of contract, tort or
// otherwise, arising from, out of or in connection with the
// software or the use or other dealings in the software.

package cc

import (
	"github.com/xlab/c/xc"
	"github.com/cznic/golex/lex"
)
%}

%union {
	Token			xc.Token
	groupPart		node
	node			node
	toks			PPTokenList
}

%token
	/*yy:token "'%c'"            */ CHARCONST		"character constant"
	/*yy:token "1.%d"            */ FLOATCONST		"floating-point constant"
	/*yy:token "%c"              */ IDENTIFIER		"identifier"
	/*yy:token "%c("             */ IDENTIFIER_LPAREN	"identifier immediatelly followed by '('"
	/*yy:token "%d"              */ INTCONST		"integer constant"
	/*yy:token "L'%c'"           */ LONGCHARCONST		"long character constant"
	/*yy:token "L\"%c\""         */ LONGSTRINGLITERAL	"long string constant"
	/*yy:token "<%c.h>"          */ PPHEADER_NAME		"header name"
	/*yy:token "%d"              */ PPNUMBER		"preprocessing number"
	/*yy:token "\"%c\""          */ STRINGLITERAL		"string literal"

	/*yy:token "\U00100000"      */	PREPROCESSING_FILE	1048576	"preprocessing file prefix"	// 0x100000 = 1048576
	/*yy:token "\U00100001"      */	CONSTANT_EXPRESSION	1048577	"constant expression prefix"
	/*yy:token "\U00100002"      */	TRANSLATION_UNIT	1048578	"translation unit prefix"

	/*yy:token "\n#define"       */	PPDEFINE		"#define"
	/*yy:token "\n#elif"         */	PPELIF			"#elif"
	/*yy:token "\n#else"         */	PPELSE			"#else"
	/*yy:token "\n#endif"        */	PPENDIF			"#endif"
	/*yy:token "\n#error"        */	PPERROR			"#error"
	/*yy:token "\n#"             */	PPHASH_NL		"#"
	/*yy:token "\n#if"           */	PPIF			"#if"
	/*yy:token "\n#ifdef"        */	PPIFDEF			"#ifdef"
	/*yy:token "\n#ifndef"       */	PPIFNDEF		"#ifndef"
	/*yy:token "\n#include"      */	PPINCLUDE		"#include"
	/*yy:token "\n#line"         */	PPLINE			"#line"
	/*yy:token "\n#foo"          */	PPNONDIRECTIVE		"#foo"
	/*yy:token "other_%c"        */ PPOTHER			"ppother"
	/*yy:token "\n##"            */	PPPASTE			"##"
	/*yy:token "\n#pragma"       */	PPPRAGMA		"#pragma"
	/*yy:token "\n#undef"        */	PPUNDEF			"#undef"

	ADDASSIGN			"+="
	ANDAND				"&&"
	ANDASSIGN			"&="
	ARROW				"->"
	AUTO				"auto"
	BOOL				"_Bool"
	BREAK				"break"
	CASE				"case"
	CHAR				"char"
	COMPLEX				"_Complex"
	CONST				"const"
	CONTINUE			"continue"
	DDD				"..."
	DEC				"--"
	DEFAULT				"default"
	DIVASSIGN			"/="
	DO				"do"
	DOUBLE				"double"
	ELSE				"else"
	ENUM				"enum"
	EQ				"=="
	EXTERN				"extern"
	FLOAT				"float"
	FOR				"for"
	GEQ				">="
	GOTO				"goto"
	IF				"if"
	INC				"++"
	INLINE				"inline"
	INT				"int"
	LEQ				"<="
	LONG				"long"
	LSH				"<<"
	LSHASSIGN			"<<="
	MODASSIGN			"%="
	MULASSIGN			"*="
	NEQ				"!="
	ORASSIGN			"|="
	OROR				"||"
	REGISTER			"register"
	RESTRICT			"restrict"
	RETURN				"return"
	RSH				">>"
	RSHASSIGN			">>="
	SHORT				"short"
	SIGNED				"signed"
	SIZEOF				"sizeof"
	STATIC				"static"
	STRUCT				"struct"
	SUBASSIGN			"-="
	SWITCH				"switch"
	TYPEDEF				"typedef"
	TYPEDEFNAME			"typedefname"
	UNION				"union"
	UNSIGNED			"unsigned"
	VOID				"void"
	VOLATILE			"volatile"
	WHILE				"while"
	XORASSIGN			"^="

%type	<toks>
	PPTokenList			"token list"
	PPTokenListOpt			"optional token list"
	ReplacementList			"replacement list"
	TextLine			"text line"

%type	<groupPart>
	GroupPart			"group part"

%type	<node>
	AbstractDeclarator		"abstract declarator"
	AbstractDeclaratorOpt		"optional abstract declarator"
	ArgumentExpressionList		"argument expression list"
	ArgumentExpressionListOpt	"optional argument expression list"
	BlockItem			"block item"
	BlockItemList			"block item list"
	BlockItemListOpt		"optional block item list"
	CompoundStatement		"compound statement"
	ConstantExpression		"constant expression"
	ControlLine			"control line"
	Declaration			"declaration"
	DeclarationList			"declaration list"
	DeclarationListOpt		"optional declaration list"
	DeclarationSpecifiers		"declaration specifiers"
	DeclarationSpecifiersOpt	"optional declaration specifiers"
	Declarator			"declarator"
	DeclaratorOpt			"optional declarator"
	Designation			"designation"
	DesignationOpt			"optional designation"
	Designator			"designator"
	DesignatorList			"designator list"
	DirectAbstractDeclarator	"direct abstract declarator"
	DirectAbstractDeclaratorOpt	"optional direct abstract declarator"
	DirectDeclarator		"direct declarator"
	ElifGroup			"elif group"
	ElifGroupList			"elif group list"
	ElifGroupListOpt		"optional elif group list"
	ElseGroup			"else group"
	ElseGroupOpt			"optional else group"
	EndifLine			"endif line"
	EnumSpecifier			"enum specifier"
	EnumerationConstant		"enumearation constant"
	Enumerator			"enumerator"
	EnumeratorList			"enumerator list"
	Expression			"expression"
	ExpressionList			"expression list"
	ExpressionListOpt		"optional expression list"
	ExpressionOpt			"optional expression"
	ExpressionStatement		"expression statement"
	ExternalDeclaration		"external declaration"
	FunctionDefinition		"function definition"
	FunctionSpecifier		"function specifier"
	GroupList			"group list"
	GroupListOpt			"optional group list"
	IdentifierList			"identifier list"
	IdentifierListOpt		"optional identifier list"
	IdentifierOpt			"optional identifier"
	IfGroup				"if group"
	IfSection			"if section"
	InitDeclarator			"init declarator"
	InitDeclaratorList		"init declarator list"
	InitDeclaratorListOpt		"optional init declarator list"
	Initializer			"initializer"
	InitializerList			"initializer list"
	IterationStatement		"iteration statement"
	JumpStatement			"jump statement"
	LabeledStatement		"labeled statement"
	ParameterDeclaration		"parameter declaration"
	ParameterList			"parameter list"
	ParameterTypeList		"parameter type list"
	ParameterTypeListOpt		"optional parameter type list"
	Pointer				"pointer"
	PointerOpt			"optional pointer"
	PreprocessingFile		"preprocessing file"
	SelectionStatement		"selection statement"
	SpecifierQualifierList		"specifier qualifier list"
	SpecifierQualifierListOpt	"optional specifier qualifier list"
	Statement			"statement"
	StorageClassSpecifier		"storage class specifier"
	StructDeclaration		"struct declaration"
	StructDeclarationList		"struct declaration list"
	StructDeclarator		"struct declarator"
	StructDeclaratorList		"struct declarator list"
	StructOrUnion			"struct-or-union"
	StructOrUnionSpecifier		"struct-or-union specifier"
	TranslationUnit			"translation unit"
	TypeName			"type name"
	TypeQualifier			"type qualifier"
	TypeQualifierList		"type qualifier list"
	TypeQualifierListOpt		"optional type qualifier list"
	TypeSpecifier			"type specifier"

%precedence	NOELSE
%precedence	ELSE

%right		'=' ADDASSIGN ANDASSIGN DIVASSIGN LSHASSIGN MODASSIGN MULASSIGN
		ORASSIGN RSHASSIGN SUBASSIGN XORASSIGN
		
%left		':' '?'
%left		OROR
%left		ANDAND
%left		'|'
%left		'^'
%left		'&'
%left		EQ NEQ
%left		'<' '>' GEQ LEQ
%left		LSH RSH
%left		'+' '-' 
%left		'%' '*' '/'
%precedence	CAST
%left		'!' '~' SIZEOF UNARY
%right		'(' '.' '[' ARROW DEC INC

%%

//yy:ignore
Start:
	PREPROCESSING_FILE PreprocessingFile
	{
		lx.preprocessingFile = $2.(*PreprocessingFile)
	}
|	CONSTANT_EXPRESSION ConstantExpression
	{
		lx.constantExpression = $2.(*ConstantExpression)
	}
|	TRANSLATION_UNIT TranslationUnit
	{
		if lx.report.Errors(false) == nil && lx.scope.kind != ScopeFile {
			panic("internal error")
		}

		lx.translationUnit = $2.(*TranslationUnit).reverse()
		lx.translationUnit.Declarations = lx.scope
	}

// [0](6.4.4.3)
EnumerationConstant:
	IDENTIFIER

// [0](6.5.2)
ArgumentExpressionList:
	Expression
|	ArgumentExpressionList ',' Expression

ArgumentExpressionListOpt:
	/* empty */ {}
|	ArgumentExpressionList

// [0](6.5.16)
//yy:field	Type	Type		// Type of expression.
//yy:field	Value	interface{}	// Non nil for certain constant expressions.
//yy:field	scope	*Bindings	// Case 0: IDENTIFIER resolution scope.
Expression:
	IDENTIFIER
	{
		lhs.scope = lx.scope
	}
|	CHARCONST
|	FLOATCONST
|	INTCONST
|	LONGCHARCONST
|	LONGSTRINGLITERAL
|	STRINGLITERAL
|	'(' ExpressionList ')'
|	Expression '[' ExpressionList ']'
|	Expression '(' ArgumentExpressionListOpt ')'
	{
		o := lhs.ArgumentExpressionListOpt
		if o == nil {
			break
		}

		m := lx.model
		lhs.Expression.eval(m)
		s := m.adjustFnArgs
		m.adjustFnArgs = true
		for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
			l.Expression.eval(m)
		}
		m.adjustFnArgs = s
	}
|	Expression '.' IDENTIFIER
|	Expression "->" IDENTIFIER
|	Expression "++"
|	Expression "--"
|	'(' TypeName ')' '{' InitializerList CommaOpt '}'
|	"++" Expression
|	"--" Expression
|	'&' Expression %prec UNARY
|	'*' Expression %prec UNARY
|	'+' Expression %prec UNARY
|	'-' Expression %prec UNARY
|	'~' Expression
|	'!' Expression
|	"sizeof" Expression
|	"sizeof" '(' TypeName ')' %prec SIZEOF
|	'(' TypeName ')' Expression %prec CAST
|	Expression '*' Expression
|	Expression '/' Expression
|	Expression '%' Expression
|	Expression '+' Expression
|	Expression '-' Expression
|	Expression "<<" Expression
|	Expression ">>" Expression
|	Expression '<' Expression
|	Expression '>' Expression
|	Expression "<=" Expression
|	Expression ">=" Expression
|	Expression "==" Expression
|	Expression "!=" Expression
|	Expression '&' Expression
|	Expression '^' Expression
|	Expression '|' Expression
|	Expression "&&" Expression
|	Expression "||" Expression
|	Expression '?' ExpressionList ':' Expression
|	Expression '=' Expression
|	Expression "*=" Expression
|	Expression "/=" Expression
|	Expression "%=" Expression
|	Expression "+=" Expression
|	Expression "-=" Expression
|	Expression "<<=" Expression
|	Expression ">>=" Expression
|	Expression "&=" Expression
|	Expression "^=" Expression
|	Expression "|=" Expression


ExpressionOpt:
	/* empty */ {}
|	Expression
	{
		lhs.Expression.eval(lx.model)
	}

// [0](6.5.17)
//yy:field	Type	Type		// Type of expression.
//yy:field	Value	interface{}	// Non nil for certain constant expressions.
//yy:list
ExpressionList:
	Expression
|	ExpressionList ',' Expression

ExpressionListOpt:
	/* empty */ {}
|	ExpressionList
	{
		lhs.ExpressionList.eval(lx.model)
	}

// [0](6.6)
//yy:field	Type	Type		// Type of expression.
//yy:field	Value	interface{}	// Non nil for certain constant expressions.
ConstantExpression:
	Expression
	{
		lhs.Value, lhs.Type = lhs.Expression.eval(lx.model)
		if lhs.Value == nil {
			lx.report.Err(lhs.Pos(), "not a constant expression")
		}
	}

// [0](6.7)
//yy:field	declarator	*Declarator	// Synthetic declarator when InitDeclaratorListOpt is nil.
Declaration:
	DeclarationSpecifiers InitDeclaratorListOpt ';'
	{
		o := lhs.InitDeclaratorListOpt
		if o != nil {
			break
		}

		s := lhs.DeclarationSpecifiers
		d := &Declarator{specifier: s}
		dd := &DirectDeclarator{
			Token: xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
			declarator: d,
			idScope: lx.scope,
			specifier: s,
		}
		d.DirectDeclarator = dd
		d.setFull(lx)
		lhs.declarator = d
	}

// [0](6.7)
//yy:field	attr		int	// tsInline, tsTypedefName, ...
//yy:field	typeSpecifier	int	// Encoded combination of tsVoid, tsInt, ...
DeclarationSpecifiers:
	StorageClassSpecifier DeclarationSpecifiersOpt
	{
		lx.scope.specifier = lhs
		a := lhs.StorageClassSpecifier
		b := lhs.DeclarationSpecifiersOpt
		if b == nil {
			lhs.attr = a.attr
			break
		}

		if a.attr&b.attr != 0 {
			lx.report.Err(a.Pos(), "invalid storage class specifier")
			break
		}

		lhs.attr = a.attr|b.attr
		lhs.typeSpecifier = b.typeSpecifier
		if lhs.StorageClassSpecifier.Case != 0 /* "typedef" */ && lhs.IsTypedef() {
			lx.report.Err(a.Pos(), "invalid storage class specifier")
		}
	}
|	TypeSpecifier DeclarationSpecifiersOpt
	{
		lx.scope.specifier = lhs
		a := lhs.TypeSpecifier
		b := lhs.DeclarationSpecifiersOpt
		if b == nil {
			lhs.typeSpecifier = a.typeSpecifier
			break
		}

		lhs.attr = b.attr
		ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
		if _, ok := tsValid[ts]; !ok {
			ts = tsEncode(tsInt)
			lx.report.Err(a.Pos(), "invalid type specifier")
		}
		lhs.typeSpecifier = ts
	}
|	TypeQualifier DeclarationSpecifiersOpt
	{
		lx.scope.specifier = lhs
		a := lhs.TypeQualifier
		b := lhs.DeclarationSpecifiersOpt
		if b == nil {
			lhs.attr = a.attr
			break
		}
	
		if a.attr&b.attr != 0 {
			lx.report.Err(a.Pos(), "invalid type qualifier")
			break
		}

		lhs.attr = a.attr|b.attr
		lhs.typeSpecifier = b.typeSpecifier
		if lhs.IsTypedef() {
			lx.report.Err(a.Pos(), "invalid type qualifier")
		}
	}
|	FunctionSpecifier DeclarationSpecifiersOpt
	{
		lx.scope.specifier = lhs
		a := lhs.FunctionSpecifier
		b := lhs.DeclarationSpecifiersOpt
		if b == nil {
			lhs.attr = a.attr
			break
		}
	
		if a.attr&b.attr != 0 {
			lx.report.Err(a.Pos(), "invalid function specifier")
			break
		}

		lhs.attr = a.attr|b.attr
		lhs.typeSpecifier = b.typeSpecifier
		if lhs.IsTypedef() {
			lx.report.Err(a.Pos(), "invalid function specifier")
		}
	}

//yy:field	attr		int	// tsInline, tsTypedefName, ...
//yy:field	typeSpecifier	int	// Encoded combination of tsVoid, tsInt, ...
DeclarationSpecifiersOpt:
	/* empty */ {}
|	DeclarationSpecifiers
	{
		lhs.attr = lhs.DeclarationSpecifiers.attr
		lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
	}

// [0](6.7)
InitDeclaratorList:
	InitDeclarator
|	InitDeclaratorList ',' InitDeclarator

InitDeclaratorListOpt:
	/* empty */ {}
|	InitDeclaratorList

// [0](6.7)
InitDeclarator:
	Declarator
	{
		lhs.Declarator.setFull(lx)
	}
|	Declarator
	{
		d := $1.(*Declarator)
		d.setFull(lx)
	}
	'=' Initializer
	{
		switch i := lhs.Initializer; i.Case {
		case 0: // Expression
			s := i.Expression.Type
			d := lhs.Declarator.Type
			if !s.CanAssignTo(s) {
				lx.report.Err(i.Pos(), "incompatible types when initializing type '%s' using type ‘%s'", d, s)
			}
		case 1: // '{' InitializerList CommaOpt '}'
			limit := -1
			var checkType Type
			var mb []Member
			d := lhs.Declarator
			k := d.Type.Kind()
			switch k {
			case Array:
				checkType = d.Type.Element()
				limit = d.Type.Elements()
			case Ptr:
				checkType = d.Type.Element()
				d.Type = d.Type.(*ctype).setElements(i.InitializerList.Len())
			case Struct, Union:
				mb, _ = d.Type.Members()
				if mb == nil {
					panic("internal error")
				}

				limit = len(mb)
				if k == Union {
					limit = 1
				}
			default:
				//dbg("", position(d.Pos()), d.Type.Kind())
				panic("TODO")
			}

			values := 0
			for l := i.InitializerList; l != nil; l = l.InitializerList {
				values++
				l.Initializer.typeCheck(checkType, mb, values-1, limit, lx)
			}
		default:
			panic(i.Case)
		}
	}

// [0](6.7.1)
//yy:field	attr	int
StorageClassSpecifier:
	"typedef"
	{
		lhs.attr = saTypedef
	}
|	"extern"
	{
		lhs.attr = saExtern
	}
|	"static"
	{
		lhs.attr = saStatic
	}
|	"auto"
	{
		lhs.attr = saAuto
	}
|	"register"
	{
		lhs.attr = saRegister
	}

// [0](6.7.2)
//yy:field	scope		*Bindings	// If case TYPEDEFNAME.
//yy:field	typeSpecifier	int		// Encoded combination of tsVoid, tsInt, ...
TypeSpecifier:
	"void"
	{
		lhs.typeSpecifier = tsEncode(tsVoid)
	}
|	"char"
	{
		lhs.typeSpecifier = tsEncode(tsChar)
	}
|	"short"
	{
		lhs.typeSpecifier = tsEncode(tsShort)
	}
|	"int"
	{
		lhs.typeSpecifier = tsEncode(tsInt)
	}
|	"long"
	{
		lhs.typeSpecifier = tsEncode(tsLong)
	}
|	"float"
	{
		lhs.typeSpecifier = tsEncode(tsFloat)
	}
|	"double"
	{
		lhs.typeSpecifier = tsEncode(tsDouble)
	}
|	"signed"
	{
		lhs.typeSpecifier = tsEncode(tsSigned)
	}
|	"unsigned"
	{
		lhs.typeSpecifier = tsEncode(tsUnsigned)
	}
|	"_Bool"
	{
		lhs.typeSpecifier = tsEncode(tsBool)
	}
|	"_Complex"
	{
		lhs.typeSpecifier = tsEncode(tsComplex)
	}
|	StructOrUnionSpecifier
	{
		lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
	}
|	EnumSpecifier
	{
		lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
	}
/*yy:example "\U00100002 typedef int i; i j;" */
|	TYPEDEFNAME
	{
		lhs.typeSpecifier = tsEncode(tsTypedefName)
		_, lhs.scope = lx.scope.Lookup2(NSIdentifiers, lhs.Token.Val)
	}

// [0](6.7.2.1)
//yy:field	alignOf	int
//yy:field	scope	*Bindings
//yy:field	sizeOf	int
StructOrUnionSpecifier:
	StructOrUnion IdentifierOpt
	{
		if o := $2.(*IdentifierOpt); o != nil {
			lx.scope.declareStructTag(o.Token, lx.report)
		}
	}
	'{'
	{
		lx.pushScope(ScopeMembers)
		lx.scope.isUnion = $1.(*StructOrUnion).Case == 1 // "union"
		lx.scope.prevStructDeclarator = nil
	}
	StructDeclarationList '}'
	{
		sc := lx.scope
		lhs.scope = sc
		if sc.bitOffset != 0 {
			finishBitField(lx)
		}

		i := 0
		var bt Type
		var d *Declarator
		for l := lhs.StructDeclarationList; l != nil; l = l.StructDeclarationList {
			for l := l.StructDeclaration.StructDeclaratorList; l != nil; l = l.StructDeclaratorList {
				switch sd := l.StructDeclarator; sd.Case {
				case 0: // Declarator
					d = sd.Declarator
				case 1: // DeclaratorOpt ':' ConstantExpression
					if o := sd.DeclaratorOpt; o != nil {
						x := o.Declarator
						if x.bitOffset == 0  {
							d = x
							bt = lx.scope.bitFieldTypes[i]
							i++
						}
						x.bitFieldType = bt
					}
				}
			}
		}
		lx.scope.bitFieldTypes = nil

		lhs.alignOf = sc.maxAlign
		switch {
		case sc.isUnion:
			lhs.sizeOf = align(sc.maxSize, sc.maxAlign)
		default:
			off := sc.offset
			lhs.sizeOf = align(sc.offset, sc.maxAlign)
			if d != nil {
				d.padding = lhs.sizeOf-off
			}
		}
		if lhs.sizeOf == 0 {
			panic("internal error")
		}

		lx.popScope(lhs.Token2)
		if o := lhs.IdentifierOpt; o != nil {
			lx.scope.defineStructTag(o.Token, lhs, lx.report)
		}
	}
|	StructOrUnion IDENTIFIER
	{
		lx.scope.declareStructTag(lhs.Token, lx.report)
		lhs.scope = lx.scope
	}

// [0](6.7.2.1)
StructOrUnion:
	"struct"
|	"union"

// [0](6.7.2.1)
StructDeclarationList:
	StructDeclaration
|	StructDeclarationList StructDeclaration

// [0](6.7.2.1)
StructDeclaration:
	SpecifierQualifierList StructDeclaratorList ';'

// [0](6.7.2.1)
//yy:field	attr		int	// tsInline, tsTypedefName, ...
//yy:field	typeSpecifier	int	// Encoded combination of tsVoid, tsInt, ...
SpecifierQualifierList:
	TypeSpecifier SpecifierQualifierListOpt
	{
		lx.scope.specifier = lhs
		a := lhs.TypeSpecifier
		b := lhs.SpecifierQualifierListOpt
		if b == nil {
			lhs.typeSpecifier = a.typeSpecifier
			break
		}

		lhs.attr = b.attr
		ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
		if _, ok := tsValid[ts]; !ok {
			lx.report.Err(a.Pos(), "invalid type specifier")
			break
		}

		lhs.typeSpecifier = ts
	}
|	TypeQualifier SpecifierQualifierListOpt
	{
		lx.scope.specifier = lhs
		a := lhs.TypeQualifier
		b := lhs.SpecifierQualifierListOpt
		if b == nil {
			lhs.attr = a.attr
			break
		}
	
		if a.attr&b.attr != 0 {
			lx.report.Err(a.Pos(), "invalid type qualifier")
			break
		}

		lhs.attr = a.attr|b.attr
		lhs.typeSpecifier = b.typeSpecifier
	}

//yy:field	attr		int	// tsInline, tsTypedefName, ...
//yy:field	typeSpecifier	int	// Encoded combination of tsVoid, tsInt, ...
SpecifierQualifierListOpt:
	/* empty */ {}
|	SpecifierQualifierList
	{
		lhs.attr = lhs.SpecifierQualifierList.attr
		lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
	}

// [0](6.7.2.1)
StructDeclaratorList:
	StructDeclarator
|	StructDeclaratorList ',' StructDeclarator

// [0](6.7.2.1)
StructDeclarator:
	Declarator
	{
		lhs.Declarator.setFull(lx)
		lhs.post(lx)
	}
|	DeclaratorOpt ':' ConstantExpression
	{
		m := lx.model
		e := lhs.ConstantExpression
		if e.Value == nil {
			e.Value, e.Type = m.value2(1, m.IntType)
		}
		if e.Type.Kind() != Int {
			lx.report.Err(e.Pos(), "bit field width not an integer")
			e.Value, e.Type = m.value2(1, m.IntType)
		}
		if o := lhs.DeclaratorOpt; o != nil {
			o.Declarator.setFull(lx)
		}
		lhs.post(lx)
	}

CommaOpt:
	/* empty */ {}
|	','

// [0](6.7.2.2)
EnumSpecifier:
	"enum" IdentifierOpt
	{
		if o := $2.(*IdentifierOpt); o != nil {
			lx.scope.declareEnumTag(o.Token, lx.report)
		}
		lx.iota = 0
	}
	'{' EnumeratorList  CommaOpt '}'
	{
		if o := lhs.IdentifierOpt; o != nil {
			lx.scope.defineEnumTag(o.Token, lhs, lx.report)
		}
	}
|	"enum" IDENTIFIER
	{
		lx.scope.declareEnumTag(lhs.Token2, lx.report)
	}

// [0](6.7.2.2)
EnumeratorList:
	Enumerator
|	EnumeratorList ',' Enumerator

// [0](6.7.2.2)
Enumerator:
	EnumerationConstant
	{
		lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, lx.iota)
	}
|	EnumerationConstant '=' ConstantExpression
	{
		m := lx.model
		e := lhs.ConstantExpression
		if e.Type.Kind() != Int {
			lx.report.Err(e.Pos(), "not an integer constant expression (have '%s')", e.Type)
			e.Value, e.Type = m.value2(1, m.IntType)
			break
		}

		var val int
		switch x := e.Value.(type) {
		case int16:
			val = int(x)
		case int32:
			val = int(x)
		case int64:
			val = int(x)
		default:
			panic("internal error")
		}
		lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, val)
	}

// [0](6.7.3)
//yy:field	attr		int	// tsInline, tsTypedefName, ...
TypeQualifier:
	"const"
	{
		lhs.attr = saConst
	}
|	"restrict"
	{
		lhs.attr = saRestrict
	}
|	"volatile"
	{
		lhs.attr = saVolatile
	}

// [0](6.7.4)
//yy:field	attr		int	// tsInline, tsTypedefName, ...
FunctionSpecifier:
	"inline"
	{
		lhs.attr = saInline
	}

// [0](6.7.5)
//yy:field	Linkage		Linkage
//yy:field	Type		Type
//yy:field	bitFieldType	Type
//yy:field	bitOffset	int
//yy:field	bits		int
//yy:field	offsetOf	int
//yy:field	padding		int
//yy:field	specifier	Specifier
Declarator:
	PointerOpt DirectDeclarator
	{
		lhs.specifier = lx.scope.specifier
		lhs.DirectDeclarator.declarator = lhs
	}

DeclaratorOpt:
	/* empty */ {}
|	Declarator

// [0](6.7.5)
//yy:field	declarator	*Declarator
//yy:field	elements	int
//yy:field	enumVal		int32
//yy:field	idScope		*Bindings	// Of case 0: IDENTIFIER.
//yy:field	isEnumConst	bool
//yy:field	paramsScope	*Bindings
//yy:field	parent		*DirectDeclarator
//yy:field	prev		*Binding	// Existing declaration in same scope, if any.
//yy:field	specifier	Specifier
DirectDeclarator:
	IDENTIFIER
	{
		lhs.specifier = lx.scope.specifier
		lx.scope.declareIdentifier(lhs.Token, lhs, lx.report)
		lhs.idScope = lx.scope
	}
|	'(' Declarator ')'
	{
		lhs.Declarator.specifier = nil
		lhs.Declarator.DirectDeclarator.parent = lhs
	}
|	DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
	{
		lhs.elements = -1
		if o := lhs.ExpressionOpt; o != nil {
			o.Expression.eval(lx.model)
			var err error
			if lhs.elements, err = elements(o.Expression.Value); err != nil {
				lx.report.Err(o.Expression.Pos(), "%s", err)
			}
			
		}
		lhs.DirectDeclarator.parent = lhs
	}
|	DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
	{
		lhs.Expression.eval(lx.model)
		var err error
		if lhs.elements, err = elements(lhs.Expression.Value); err != nil {
			lx.report.Err(lhs.Expression.Pos(), "%s", err)
		}
		lhs.DirectDeclarator.parent = lhs
	}
|	DirectDeclarator '[' TypeQualifierList "static" Expression ']'
	{
		lhs.Expression.eval(lx.model)
		var err error
		if lhs.elements, err = elements(lhs.Expression.Value); err != nil {
			lx.report.Err(lhs.Expression.Pos(), "%s", err)
		}
		lhs.DirectDeclarator.parent = lhs
	}
|	DirectDeclarator '[' TypeQualifierListOpt '*' ']'
	{
		lhs.DirectDeclarator.parent = lhs
		lhs.elements = -1
	}
|	DirectDeclarator '('
	{
		lx.pushScope(ScopeParams)
	}
	ParameterTypeList ')'
	{
		lhs.paramsScope, _ = lx.popScope(lhs.Token2)
		lhs.DirectDeclarator.parent = lhs
	}
|	DirectDeclarator '(' IdentifierListOpt ')'
	{
		lhs.DirectDeclarator.parent = lhs
	}

// [0](6.7.5)
Pointer:
	'*' TypeQualifierListOpt
|	'*' TypeQualifierListOpt Pointer

PointerOpt:
	/* empty */ {}
|	Pointer

// [0](6.7.5)
//yy:field	attr		int	// tsInline, tsTypedefName, ...
TypeQualifierList:
	TypeQualifier
	{
		lhs.attr = lhs.TypeQualifier.attr
	}
|	TypeQualifierList TypeQualifier
	{
		a := lhs.TypeQualifierList
		b := lhs.TypeQualifier
		if a.attr&b.attr != 0 {
			lx.report.Err(b.Pos(), "invalid type qualifier")
			break
		}

		lhs.attr = a.attr|b.attr
	}

TypeQualifierListOpt:
	/* empty */ {}
|	TypeQualifierList

// [0](6.7.5)
//yy:field	params	[]Parameter
ParameterTypeList:
	ParameterList
	{
		lhs.post()
	}
|	ParameterList ',' "..."
	{
		lhs.post()
	}

ParameterTypeListOpt:
	/* empty */ {}
|	ParameterTypeList

// [0](6.7.5)
ParameterList:
	ParameterDeclaration
|	ParameterList ',' ParameterDeclaration

// [0](6.7.5)
//yy:field	declarator	*Declarator
/*TODO
A declaration of a parameter as ‘‘function returning type’’ shall be adjusted
to ‘‘pointer to function returning type’’, as in 6.3.2.1.
*/
ParameterDeclaration:
	DeclarationSpecifiers Declarator
	{
		lhs.Declarator.setFull(lx)
		lhs.declarator = lhs.Declarator
	}
|	DeclarationSpecifiers AbstractDeclaratorOpt
	{
		if o := lhs.AbstractDeclaratorOpt; o != nil {
			lhs.declarator = o.AbstractDeclarator.declarator
			lhs.declarator.setFull(lx)
			break
		}

		d := &Declarator{
			specifier: lx.scope.specifier,
			DirectDeclarator: &DirectDeclarator{
				Case: 0, // IDENTIFIER
			},
		}
		d.DirectDeclarator.declarator = d
		lhs.declarator = d
		d.setFull(lx)
	}

// [0](6.7.5)
IdentifierList:
	IDENTIFIER
|	IdentifierList ',' IDENTIFIER

IdentifierListOpt:
	/* empty */ {}
|	IdentifierList

IdentifierOpt:
	/* empty */ {}
|	IDENTIFIER

// [0](6.7.6)
//yy:field	Type		Type
//yy:field	declarator	*Declarator
TypeName:
	{
		lx.pushScope(ScopeBlock)
	}
	SpecifierQualifierList AbstractDeclaratorOpt
	{
		if o := lhs.AbstractDeclaratorOpt; o != nil {
			lhs.declarator = o.AbstractDeclarator.declarator
		} else {
			d := &Declarator{
				specifier: lhs.SpecifierQualifierList,
				DirectDeclarator: &DirectDeclarator{
					Case: 0, // IDENTIFIER
					idScope: lx.scope,
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
		lhs.Type = lhs.declarator.setFull(lx)
		lx.popScope(xc.Token{})
	}

// [0](6.7.6)
//yy:field	declarator	*Declarator
AbstractDeclarator:
	Pointer
	{
		d := &Declarator{
			specifier: lx.scope.specifier,
			PointerOpt: &PointerOpt {
				Pointer: lhs.Pointer,
			},
			DirectDeclarator: &DirectDeclarator{
				Case: 0, // IDENTIFIER
				idScope: lx.scope,
			},
		}
		d.DirectDeclarator.declarator = d
		lhs.declarator = d
	}
|	PointerOpt DirectAbstractDeclarator
	{
		d := &Declarator{
			specifier: lx.scope.specifier,
			PointerOpt: lhs.PointerOpt,
			DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
		}
		d.DirectDeclarator.declarator = d
		lhs.declarator = d
	}

AbstractDeclaratorOpt:
	/* empty */ {}
|	AbstractDeclarator

// [0](6.7.6)
//yy:field	directDeclarator	*DirectDeclarator
//yy:field	paramsScope		*Bindings
DirectAbstractDeclarator:
	'(' AbstractDeclarator ')'
	{
		lhs.AbstractDeclarator.declarator.specifier = nil
		lhs.directDeclarator = &DirectDeclarator{
			Case: 1, // '(' Declarator ')'
			Declarator: lhs.AbstractDeclarator.declarator,
		}
		lhs.AbstractDeclarator.declarator.DirectDeclarator.parent = lhs.directDeclarator
	}
|	DirectAbstractDeclaratorOpt '[' ExpressionOpt ']'
	{
		if o := lhs.ExpressionOpt; o != nil {
			o.Expression.eval(lx.model)
		}
		var dd *DirectDeclarator
		switch o := lhs.DirectAbstractDeclaratorOpt; {
		case o == nil:
			dd = &DirectDeclarator{
				Case: 0, // IDENTIFIER
			}
		default:
			dd = o.DirectAbstractDeclarator.directDeclarator
		}
		lhs.directDeclarator = &DirectDeclarator{
			Case: 2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			DirectDeclarator: dd,
			ExpressionOpt: lhs.ExpressionOpt,
		}
		dd.parent = lhs.directDeclarator
	}
|	DirectAbstractDeclaratorOpt '[' TypeQualifierList ExpressionOpt ']'
	{
		if o := lhs.ExpressionOpt; o != nil {
			o.Expression.eval(lx.model)
		}
		var dd *DirectDeclarator
		switch o := lhs.DirectAbstractDeclaratorOpt; {
		case o == nil:
			dd = &DirectDeclarator{
				Case: 0, // IDENTIFIER
			}
		default:
			dd = o.DirectAbstractDeclarator.directDeclarator
		}
		lhs.directDeclarator = &DirectDeclarator{
			Case: 2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			DirectDeclarator: dd,
			TypeQualifierListOpt: &TypeQualifierListOpt{ lhs.TypeQualifierList },
			ExpressionOpt: lhs.ExpressionOpt,
		}
		dd.parent = lhs.directDeclarator
	}
|	DirectAbstractDeclaratorOpt '[' "static" TypeQualifierListOpt Expression ']'
	{
		lhs.Expression.eval(lx.model)
		var dd *DirectDeclarator
		switch o := lhs.DirectAbstractDeclaratorOpt; {
		case o == nil:
			dd = &DirectDeclarator{
				Case: 0, // IDENTIFIER
			}
		default:
			dd = o.DirectAbstractDeclarator.directDeclarator
		}
		lhs.directDeclarator = &DirectDeclarator{
			Case: 2, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
			DirectDeclarator: dd,
			TypeQualifierListOpt: lhs.TypeQualifierListOpt,
			Expression: lhs.Expression,
		}
		dd.parent = lhs.directDeclarator
	}
|	DirectAbstractDeclaratorOpt '[' TypeQualifierList "static" Expression ']'
	{
		lhs.Expression.eval(lx.model)
		var dd *DirectDeclarator
		switch o := lhs.DirectAbstractDeclaratorOpt; {
		case o == nil:
			dd = &DirectDeclarator{
				Case: 0, // IDENTIFIER
			}
		default:
			dd = o.DirectAbstractDeclarator.directDeclarator
		}
		lhs.directDeclarator = &DirectDeclarator{
			Case: 4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
			DirectDeclarator: dd,
			TypeQualifierList: lhs.TypeQualifierList,
			Expression: lhs.Expression,
		}
		dd.parent = lhs.directDeclarator
	}
|	DirectAbstractDeclaratorOpt '[' '*' ']'
	{
		var dd *DirectDeclarator
		switch o := lhs.DirectAbstractDeclaratorOpt; {
		case o == nil:
			dd = &DirectDeclarator{
				Case: 0, // IDENTIFIER
			}
		default:
			dd = o.DirectAbstractDeclarator.directDeclarator
		}
		lhs.directDeclarator = &DirectDeclarator{
			Case: 5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
			DirectDeclarator: dd,
		}
		dd.parent = lhs.directDeclarator
	}
|	'('
	{
		lx.pushScope(ScopeParams)
	}
	ParameterTypeListOpt ')'
	{
		lhs.paramsScope, _ = lx.popScope(lhs.Token2)
		switch o := lhs.ParameterTypeListOpt; {
		case o != nil:
			lhs.directDeclarator = &DirectDeclarator{
				Case: 6, // DirectDeclarator '(' ParameterTypeList ')'
				DirectDeclarator: &DirectDeclarator{
					Case: 0, // IDENTIFIER
				},
				ParameterTypeList: o.ParameterTypeList,
			}
		default:
			lhs.directDeclarator = &DirectDeclarator{
				Case: 7, // DirectDeclarator '(' IdentifierListOpt ')'
				DirectDeclarator: &DirectDeclarator{
					Case: 0, // IDENTIFIER
				},
			}
		}
		lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
	}
|	DirectAbstractDeclarator '('
	{
		lx.pushScope(ScopeParams)
	}
	ParameterTypeListOpt ')'
	{
		lhs.paramsScope, _ = lx.popScope(lhs.Token2)
		switch o := lhs.ParameterTypeListOpt; {
		case o != nil:
			lhs.directDeclarator = &DirectDeclarator{
				Case: 6, // DirectDeclarator '(' ParameterTypeList ')'
				DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
				ParameterTypeList: o.ParameterTypeList,
			}
		default:
			lhs.directDeclarator = &DirectDeclarator{
				Case: 7, // DirectDeclarator '(' IdentifierListOpt ')'
				DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
			}
		}
		lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
	}

DirectAbstractDeclaratorOpt:
	/* empty */ {}
|	DirectAbstractDeclarator

// [0](6.7.8)
Initializer:
	Expression
	{
		lhs.Expression.eval(lx.model)
	}
|	'{' InitializerList CommaOpt '}'

// [0](6.7.8)
InitializerList:
	DesignationOpt Initializer
|	InitializerList ',' DesignationOpt Initializer

// [0](6.7.8)
Designation:
	DesignatorList '='

DesignationOpt:
	/* empty */ {}
|	Designation

// [0](6.7.8)
DesignatorList:
	Designator
|	DesignatorList Designator

// [0](6.7.8)
Designator:
	'[' ConstantExpression ']'
|	'.' IDENTIFIER

// [0](6.8)
Statement:
	LabeledStatement
|	CompoundStatement
|	ExpressionStatement
|	SelectionStatement
|	IterationStatement
|	JumpStatement

// [0](6.8.1)
LabeledStatement:
	IDENTIFIER ':' Statement
|	"case" ConstantExpression ':' Statement
|	"default" ':' Statement

// [0](6.8.2)
//yy:field	scope	*Bindings	// Scope of the CompoundStatement.
CompoundStatement:
	'{'
	{
		m := lx.scope.mergeScope
		lx.pushScope(ScopeBlock)
		if m != nil {
			lx.scope.merge(m)
		}
		lx.scope.mergeScope = nil
	}
	BlockItemListOpt '}'
	{
		lhs.scope = lx.scope
		lx.popScope(lhs.Token2)
	}

// [0](6.8.2)
BlockItemList:
	BlockItem
|	BlockItemList BlockItem

BlockItemListOpt:
	/* empty */ {}
|	BlockItemList

// [0](6.8.2)
BlockItem:
	Declaration
|	Statement

// [0](6.8.3)
ExpressionStatement:
	ExpressionListOpt ';'

// [0](6.8.4)
SelectionStatement:
	"if" '(' ExpressionList ')' Statement %prec NOELSE
	{
		lhs.ExpressionList.eval(lx.model)
	}
|	"if" '(' ExpressionList ')' Statement "else" Statement
	{
		lhs.ExpressionList.eval(lx.model)
	}
|	"switch" '(' ExpressionList ')' Statement
	{
		lhs.ExpressionList.eval(lx.model)
	}

// [0](6.8.5)
IterationStatement:
	"while" '(' ExpressionList ')' Statement
	{
		lhs.ExpressionList.eval(lx.model)
	}
|	"do" Statement "while" '(' ExpressionList ')' ';'
	{
		lhs.ExpressionList.eval(lx.model)
	}
|	"for" '(' ExpressionListOpt ';' ExpressionListOpt ';' ExpressionListOpt ')' Statement
|	"for" '(' Declaration ExpressionListOpt ';' ExpressionListOpt ')' Statement

// [0](6.8.6)
JumpStatement:
	"goto" IDENTIFIER ';'
|	"continue" ';'
|	"break" ';'
|	"return" ExpressionListOpt ';'

// [0](6.9)
//yy:field	Declarations	*Bindings
//yy:list
TranslationUnit:
	ExternalDeclaration
|	TranslationUnit ExternalDeclaration

// [0](6.9)
ExternalDeclaration:
	FunctionDefinition
|	Declaration

// [0](6.9.1)
FunctionDefinition:
	DeclarationSpecifiers Declarator DeclarationListOpt
	{
		d := $2.(*Declarator)
		d.setFull(lx)
		if k := d.Type.Kind(); k != Function {
			lx.report.Err(d.Pos(), "declarator is not a function (have '%s': %v)", d.Type, k)
		}
		lx.scope.mergeScope = d.DirectDeclarator.paramsScope
	}
	CompoundStatement
	{
		d := lhs.Declarator
		switch dd := d.DirectDeclarator; dd.Case {
		case 6: // DirectDeclarator '(' ParameterTypeList ')'
			if o := lhs.DeclarationListOpt; o != nil {
				lx.report.Err(o.Pos(), "declaration list not allowed in a function definition with parameter type list")
			}
		case 7: // DirectDeclarator '(' IdentifierListOpt ')'
			if o1, o2 := dd.IdentifierListOpt, lhs.DeclarationListOpt; o1 != nil && o2 == nil {
				lx.report.Err(o1.Pos(), "declaration list required in a function definition without a parameter type list")
			}
		default:
			lx.report.Err(lhs.Declarator.Pos(), "invalid function definition declarator")
		}
	}

// [0](6.9.1)
DeclarationList:
	Declaration
|	DeclarationList Declaration

DeclarationListOpt:
	/* empty */ {}
|	DeclarationList

// ========================================================= PREPROCESSING_FILE

// [0](6.10)
//yy:field	path	string
PreprocessingFile:
	GroupList // No more GroupListOpt due to final '\n' injection.
	{
		lhs.path = lx.file.Name()
	}

// [0](6.10)
GroupList:
	GroupPart
//yy:example "\U00100000int\nf() {}"
|	GroupList GroupPart

GroupListOpt:
	/* empty */ {}
//yy:example "\U00100000 \n#ifndef a\nb\n#elif"
|	GroupList

// [0](6.10)
//yy:ignore
GroupPart:
	ControlLine
	{
		$$ = $1.(node)
	}
|	IfSection
	{
		$$ = $1.(node)
	}
|	PPNONDIRECTIVE PPTokenList '\n'
	{
		$$ = $1
	}
|	TextLine
	{
		$$ = $1
	}

//(6.10)
IfSection:
	IfGroup ElifGroupListOpt ElseGroupOpt EndifLine

//(6.10)
IfGroup:
	PPIF PPTokenList '\n' GroupListOpt
|	PPIFDEF IDENTIFIER '\n' GroupListOpt
|	PPIFNDEF IDENTIFIER '\n' GroupListOpt

// [0](6.10)
ElifGroupList:
	ElifGroup
|	ElifGroupList ElifGroup

ElifGroupListOpt:
	/* empty */ {}
|	ElifGroupList

// [0](6.10)
ElifGroup:
	PPELIF PPTokenList '\n' GroupListOpt

// [0](6.10)
ElseGroup:
	PPELSE '\n' GroupListOpt

ElseGroupOpt:
	/* empty */ {}
|	ElseGroup

// [0](6.10)
EndifLine:
	PPENDIF /* PPTokenListOpt */ //TODO Option enabling the non std PPTokenListOpt part.

// [0](6.10)
ControlLine:
	PPDEFINE IDENTIFIER ReplacementList
|	PPDEFINE IDENTIFIER_LPAREN "..." ')' ReplacementList
|	PPDEFINE IDENTIFIER_LPAREN IdentifierList ',' "..." ')' ReplacementList
|	PPDEFINE IDENTIFIER_LPAREN IdentifierListOpt ')' ReplacementList
|	PPERROR PPTokenListOpt
|	PPHASH_NL
|	PPINCLUDE PPTokenList '\n'
|	PPLINE PPTokenList '\n'
|	PPPRAGMA PPTokenListOpt
|	PPUNDEF IDENTIFIER '\n'

	// Non standard stuff.

|	PPDEFINE IDENTIFIER_LPAREN IDENTIFIER "..." ')' ReplacementList
	{
		if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
			lx.report.ErrTok(lhs.Token4, "missing comma before \"...\"")
		}
	}
|	PPDEFINE IDENTIFIER_LPAREN IdentifierList ',' IDENTIFIER "..." ')' ReplacementList
	{
		if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
			lx.report.ErrTok(lhs.Token6, "missing comma before \"...\"")
		}
	}
|	PPDEFINE '\n'
	{
		if !lx.tweaks.enableEmptyDefine {
			lx.report.ErrTok(lhs.Token2, "expected identifier")
		}
	}
//yy:example	"\U00100000 \n#undef foo(bar)"
|	PPUNDEF IDENTIFIER PPTokenList '\n'
	{
		if !lx.tweaks.enableUndefExtraTokens {
			lx.report.ErrTok(decodeTokens(lhs.PPTokenList, nil)[0], "extra tokens after #undef argument")
		}
	}

// [0](6.10)
//yy:ignore
TextLine:
	PPTokenListOpt

// [0](6.10)
//yy:ignore
ReplacementList:
	PPTokenListOpt

// [0](6.10)
//yy:ignore
PPTokenList:
	PPTokens
	{
		$$ = PPTokenList(dict.ID(lx.encBuf))
		lx.encBuf = lx.encBuf[:0]
		lx.encPos = 0
	}

//yy:ignore
PPTokenListOpt:
	'\n'
	{
		$$ = 0
	}
|	PPTokenList '\n'

//yy:ignore
PPTokens:
	PPOTHER
|	PPTokens PPOTHER
