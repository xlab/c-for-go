// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"github.com/cznic/c/internal/xc"
	"go/token"
)

// AbstractDeclarator represents data reduced by productions:
//
//	AbstractDeclarator:
//	        Pointer
//	|       PointerOpt DirectAbstractDeclarator  // Case 1
type AbstractDeclarator struct {
	indirection              int  // 'int **i': 2.
	specifier                Type // 'int i': specifier is 'int'
	Case                     int
	DirectAbstractDeclarator *DirectAbstractDeclarator
	Pointer                  *Pointer
	PointerOpt               *PointerOpt
}

func (a *AbstractDeclarator) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AbstractDeclarator) String() string {
	return PrettyString(a)
}

// AbstractDeclaratorOpt represents data reduced by productions:
//
//	AbstractDeclaratorOpt:
//	        /* empty */
//	|       AbstractDeclarator  // Case 1
type AbstractDeclaratorOpt struct {
	AbstractDeclarator *AbstractDeclarator
}

func (a *AbstractDeclaratorOpt) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AbstractDeclaratorOpt) String() string {
	return PrettyString(a)
}

// AdditiveExpression represents data reduced by productions:
//
//	AdditiveExpression:
//	        MultiplicativeExpression
//	|       AdditiveExpression '+' MultiplicativeExpression  // Case 1
//	|       AdditiveExpression '-' MultiplicativeExpression  // Case 2
type AdditiveExpression struct {
	AdditiveExpression       *AdditiveExpression
	Case                     int
	MultiplicativeExpression *MultiplicativeExpression
	Token                    xc.Token
}

func (a *AdditiveExpression) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AdditiveExpression) String() string {
	return PrettyString(a)
}

// AndExpression represents data reduced by productions:
//
//	AndExpression:
//	        EqualityExpression
//	|       AndExpression '&' EqualityExpression  // Case 1
type AndExpression struct {
	AndExpression      *AndExpression
	Case               int
	EqualityExpression *EqualityExpression
	Token              xc.Token
}

func (a *AndExpression) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AndExpression) String() string {
	return PrettyString(a)
}

// ArgumentExpressionList represents data reduced by productions:
//
//	ArgumentExpressionList:
//	        AssignmentExpression
//	|       ArgumentExpressionList ',' AssignmentExpression  // Case 1
type ArgumentExpressionList struct {
	ArgumentExpressionList *ArgumentExpressionList
	AssignmentExpression   *AssignmentExpression
	Case                   int
	Token                  xc.Token
}

func (a *ArgumentExpressionList) reverse() *ArgumentExpressionList {
	if a == nil {
		return nil
	}

	na := a
	nb := na.ArgumentExpressionList
	for nb != nil {
		nc := nb.ArgumentExpressionList
		nb.ArgumentExpressionList = na
		na = nb
		nb = nc
	}
	a.ArgumentExpressionList = nil
	return na
}

func (a *ArgumentExpressionList) fragment() interface{} { return a.reverse() }

// String implements fmt.Stringer.
func (a *ArgumentExpressionList) String() string {
	return PrettyString(a)
}

// ArgumentExpressionListOpt represents data reduced by productions:
//
//	ArgumentExpressionListOpt:
//	        /* empty */
//	|       ArgumentExpressionList  // Case 1
type ArgumentExpressionListOpt struct {
	ArgumentExpressionList *ArgumentExpressionList
}

func (a *ArgumentExpressionListOpt) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *ArgumentExpressionListOpt) String() string {
	return PrettyString(a)
}

// AssignmentExpression represents data reduced by productions:
//
//	AssignmentExpression:
//	        ConditionalExpression
//	|       UnaryExpression AssignmentOperator AssignmentExpression  // Case 1
type AssignmentExpression struct {
	AssignmentExpression  *AssignmentExpression
	AssignmentOperator    *AssignmentOperator
	Case                  int
	ConditionalExpression *ConditionalExpression
	UnaryExpression       *UnaryExpression
}

func (a *AssignmentExpression) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AssignmentExpression) String() string {
	return PrettyString(a)
}

// AssignmentExpressionOpt represents data reduced by productions:
//
//	AssignmentExpressionOpt:
//	        /* empty */
//	|       AssignmentExpression  // Case 1
type AssignmentExpressionOpt struct {
	AssignmentExpression *AssignmentExpression
}

func (a *AssignmentExpressionOpt) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AssignmentExpressionOpt) String() string {
	return PrettyString(a)
}

// AssignmentOperator represents data reduced by productions:
//
//	AssignmentOperator:
//	        '='
//	|       "*="   // Case 1
//	|       "/="   // Case 2
//	|       "%="   // Case 3
//	|       "+="   // Case 4
//	|       "-="   // Case 5
//	|       "<<="  // Case 6
//	|       ">>="  // Case 7
//	|       "&="   // Case 8
//	|       "^="   // Case 9
//	|       "|="   // Case 10
type AssignmentOperator struct {
	Case  int
	Token xc.Token
}

func (a *AssignmentOperator) fragment() interface{} { return a }

// String implements fmt.Stringer.
func (a *AssignmentOperator) String() string {
	return PrettyString(a)
}

// BlockItem represents data reduced by productions:
//
//	BlockItem:
//	        Declaration
//	|       Statement    // Case 1
type BlockItem struct {
	Case        int
	Declaration *Declaration
	Statement   *Statement
}

func (b *BlockItem) fragment() interface{} { return b }

// String implements fmt.Stringer.
func (b *BlockItem) String() string {
	return PrettyString(b)
}

// BlockItemList represents data reduced by productions:
//
//	BlockItemList:
//	        BlockItem
//	|       BlockItemList BlockItem  // Case 1
type BlockItemList struct {
	BlockItem     *BlockItem
	BlockItemList *BlockItemList
	Case          int
}

func (b *BlockItemList) reverse() *BlockItemList {
	if b == nil {
		return nil
	}

	na := b
	nb := na.BlockItemList
	for nb != nil {
		nc := nb.BlockItemList
		nb.BlockItemList = na
		na = nb
		nb = nc
	}
	b.BlockItemList = nil
	return na
}

func (b *BlockItemList) fragment() interface{} { return b.reverse() }

// String implements fmt.Stringer.
func (b *BlockItemList) String() string {
	return PrettyString(b)
}

// BlockItemListOpt represents data reduced by productions:
//
//	BlockItemListOpt:
//	        /* empty */
//	|       BlockItemList  // Case 1
type BlockItemListOpt struct {
	BlockItemList *BlockItemList
}

func (b *BlockItemListOpt) fragment() interface{} { return b }

// String implements fmt.Stringer.
func (b *BlockItemListOpt) String() string {
	return PrettyString(b)
}

// CastExpression represents data reduced by productions:
//
//	CastExpression:
//	        UnaryExpression
//	|       '(' TypeName ')' CastExpression  // Case 1
type CastExpression struct {
	Case            int
	CastExpression  *CastExpression
	Token           xc.Token
	Token2          xc.Token
	TypeName        *TypeName
	UnaryExpression *UnaryExpression
}

func (c *CastExpression) fragment() interface{} { return c }

// String implements fmt.Stringer.
func (c *CastExpression) String() string {
	return PrettyString(c)
}

// CompoundStatement represents data reduced by production:
//
//	CompoundStatement:
//	        '{' BlockItemListOpt '}'
type CompoundStatement struct {
	Declarations     *Bindings
	BlockItemListOpt *BlockItemListOpt
	Token            xc.Token
	Token2           xc.Token
}

func (c *CompoundStatement) fragment() interface{} { return c }

// String implements fmt.Stringer.
func (c *CompoundStatement) String() string {
	return PrettyString(c)
}

// ConditionalExpression represents data reduced by productions:
//
//	ConditionalExpression:
//	        LogicalOrExpression
//	|       LogicalOrExpression '?' ExpressionList ':' ConditionalExpression  // Case 1
type ConditionalExpression struct {
	Case                  int
	ConditionalExpression *ConditionalExpression
	ExpressionList        *ExpressionList
	LogicalOrExpression   *LogicalOrExpression
	Token                 xc.Token
	Token2                xc.Token
}

func (c *ConditionalExpression) fragment() interface{} { return c }

// String implements fmt.Stringer.
func (c *ConditionalExpression) String() string {
	return PrettyString(c)
}

// Constant represents data reduced by productions:
//
//	Constant:
//	        CHARCONST
//	|       FLOATCONST         // Case 1
//	|       INTCONST           // Case 2
//	|       LONGCHARCONST      // Case 3
//	|       LONGSTRINGLITERAL  // Case 4
//	|       STRINGLITERAL      // Case 5
type Constant struct {
	Case  int
	Token xc.Token
}

func (c *Constant) fragment() interface{} { return c }

// String implements fmt.Stringer.
func (c *Constant) String() string {
	return PrettyString(c)
}

// ConstantExpression represents data reduced by production:
//
//	ConstantExpression:
//	        ConditionalExpression
type ConstantExpression struct {
	ConditionalExpression *ConditionalExpression
}

func (c *ConstantExpression) fragment() interface{} { return c }

// String implements fmt.Stringer.
func (c *ConstantExpression) String() string {
	return PrettyString(c)
}

// ControlLine represents data reduced by productions:
//
//	ControlLine:
//	        PPDEFINE IDENTIFIER ReplacementList
//	|       PPDEFINE IDENTIFIER_LPAREN "..." ')' ReplacementList                                // Case 1
//	|       PPDEFINE IDENTIFIER_LPAREN IdentifierList ',' "..." ')' ReplacementList             // Case 2
//	|       PPDEFINE IDENTIFIER_LPAREN IdentifierListOpt ')' ReplacementList                    // Case 3
//	|       PPERROR PpTokenListOpt                                                              // Case 4
//	|       PPHASH_NL                                                                           // Case 5
//	|       PPINCLUDE PpTokenList                                                               // Case 6
//	|       PPLINE PpTokenList                                                                  // Case 7
//	|       PPPRAGMA PpTokenListOpt                                                             // Case 8
//	|       PPUNDEF IDENTIFIER '\n'                                                             // Case 9
//	|       PPASSERT PpTokenList                                                                // Case 10
//	|       PPDEFINE IDENTIFIER_LPAREN IDENTIFIER "..." ')' ReplacementList                     // Case 11
//	|       PPDEFINE IDENTIFIER_LPAREN IdentifierList ',' IDENTIFIER "..." ')' ReplacementList  // Case 12
//	|       PPIDENT PpTokenList                                                                 // Case 13
//	|       PPIMPORT PpTokenList                                                                // Case 14
//	|       PPINCLUDE_NEXT PpTokenList                                                          // Case 15
//	|       PPUNASSERT PpTokenList                                                              // Case 16
//	|       PPWARNING PpTokenList                                                               // Case 17
type ControlLine struct {
	Case              int
	IdentifierList    *IdentifierList
	IdentifierListOpt *IdentifierListOpt
	PpTokenList       PpTokenList
	PpTokenListOpt    PpTokenList
	ReplacementList   PpTokenList
	Token             xc.Token
	Token2            xc.Token
	Token3            xc.Token
	Token4            xc.Token
	Token5            xc.Token
	Token6            xc.Token
}

func (c *ControlLine) fragment() interface{} { return c }

// String implements fmt.Stringer.
func (c *ControlLine) String() string {
	return PrettyString(c)
}

// Declaration represents data reduced by production:
//
//	Declaration:
//	        DeclarationSpecifiers InitDeclaratorListOpt ';'
type Declaration struct {
	IsFileScope           bool
	IsTypedef             bool
	DeclarationSpecifiers *DeclarationSpecifiers
	InitDeclaratorListOpt *InitDeclaratorListOpt
	Token                 xc.Token
}

func (d *Declaration) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *Declaration) String() string {
	return PrettyString(d)
}

// DeclarationList represents data reduced by productions:
//
//	DeclarationList:
//	        Declaration
//	|       DeclarationList Declaration  // Case 1
type DeclarationList struct {
	Case            int
	Declaration     *Declaration
	DeclarationList *DeclarationList
}

func (d *DeclarationList) reverse() *DeclarationList {
	if d == nil {
		return nil
	}

	na := d
	nb := na.DeclarationList
	for nb != nil {
		nc := nb.DeclarationList
		nb.DeclarationList = na
		na = nb
		nb = nc
	}
	d.DeclarationList = nil
	return na
}

func (d *DeclarationList) fragment() interface{} { return d.reverse() }

// String implements fmt.Stringer.
func (d *DeclarationList) String() string {
	return PrettyString(d)
}

// DeclarationListOpt represents data reduced by productions:
//
//	DeclarationListOpt:
//	        /* empty */
//	|       DeclarationList  // Case 1
type DeclarationListOpt struct {
	DeclarationList *DeclarationList
}

func (d *DeclarationListOpt) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DeclarationListOpt) String() string {
	return PrettyString(d)
}

// DeclarationSpecifiers represents data reduced by productions:
//
//	DeclarationSpecifiers:
//	        StorageClassSpecifier DeclarationSpecifiersOpt
//	|       TypeSpecifier DeclarationSpecifiersOpt          // Case 1
//	|       TypeQualifier DeclarationSpecifiersOpt          // Case 2
//	|       FunctionSpecifier DeclarationSpecifiersOpt      // Case 3
type DeclarationSpecifiers struct {
	IsAuto                   bool // StorageClassSpecifier "auto" is present.
	IsConst                  bool // TypeQualifier "const" is present.
	IsExtern                 bool // StorageClassSpecifier "extern" is present.
	IsInline                 bool // FunctionSpecifier "inline" is present.
	IsRegister               bool // StorageClassSpecifier "register" is present.
	IsRestrict               bool // TypeQualifier "restrict" is present.
	IsStatic                 bool // StorageClassSpecifier "static" is present.
	IsTypedef                bool // StorageClassSpecifier "typedef" is present.
	IsVolatile               bool // TypeQualifier "volatile" is present.
	typ                      int  // One of tsVoid, tsChar, tsUChar, ...
	typeSpecifiers           int  //
	Case                     int
	DeclarationSpecifiersOpt *DeclarationSpecifiersOpt
	FunctionSpecifier        *FunctionSpecifier
	StorageClassSpecifier    *StorageClassSpecifier
	TypeQualifier            *TypeQualifier
	TypeSpecifier            *TypeSpecifier
}

func (d *DeclarationSpecifiers) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DeclarationSpecifiers) String() string {
	return PrettyString(d)
}

// DeclarationSpecifiersOpt represents data reduced by productions:
//
//	DeclarationSpecifiersOpt:
//	        /* empty */
//	|       DeclarationSpecifiers  // Case 1
type DeclarationSpecifiersOpt struct {
	DeclarationSpecifiers *DeclarationSpecifiers
}

func (d *DeclarationSpecifiersOpt) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DeclarationSpecifiersOpt) String() string {
	return PrettyString(d)
}

// Declarator represents data reduced by production:
//
//	Declarator:
//	        PointerOpt DirectDeclarator
type Declarator struct {
	Initializer      *Initializer // Non nil if Declarator is part of InitDeclarator with Initializer.
	IsDefinition     bool         // Whether Declarator is part of an InitDeclarator with Initializer or part of a FunctionDefinition.
	IsTypedef        bool
	SUSpecifier0     *StructOrUnionSpecifier0 // Non nil if Declarator declares a field.
	DirectDeclarator *DirectDeclarator
	PointerOpt       *PointerOpt
}

func (d *Declarator) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *Declarator) String() string {
	return PrettyString(d)
}

// DeclaratorOpt represents data reduced by productions:
//
//	DeclaratorOpt:
//	        /* empty */
//	|       Declarator   // Case 1
type DeclaratorOpt struct {
	Declarator *Declarator
}

func (d *DeclaratorOpt) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DeclaratorOpt) String() string {
	return PrettyString(d)
}

// Designation represents data reduced by production:
//
//	Designation:
//	        DesignatorList '='
type Designation struct {
	DesignatorList *DesignatorList
	Token          xc.Token
}

func (d *Designation) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *Designation) String() string {
	return PrettyString(d)
}

// DesignationOpt represents data reduced by productions:
//
//	DesignationOpt:
//	        /* empty */
//	|       Designation  // Case 1
type DesignationOpt struct {
	Designation *Designation
}

func (d *DesignationOpt) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DesignationOpt) String() string {
	return PrettyString(d)
}

// Designator represents data reduced by productions:
//
//	Designator:
//	        '[' ConstantExpression ']'
//	|       '.' IDENTIFIER              // Case 1
type Designator struct {
	Case               int
	ConstantExpression *ConstantExpression
	Token              xc.Token
	Token2             xc.Token
}

func (d *Designator) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *Designator) String() string {
	return PrettyString(d)
}

// DesignatorList represents data reduced by productions:
//
//	DesignatorList:
//	        Designator
//	|       DesignatorList Designator  // Case 1
type DesignatorList struct {
	Case           int
	Designator     *Designator
	DesignatorList *DesignatorList
}

func (d *DesignatorList) reverse() *DesignatorList {
	if d == nil {
		return nil
	}

	na := d
	nb := na.DesignatorList
	for nb != nil {
		nc := nb.DesignatorList
		nb.DesignatorList = na
		na = nb
		nb = nc
	}
	d.DesignatorList = nil
	return na
}

func (d *DesignatorList) fragment() interface{} { return d.reverse() }

// String implements fmt.Stringer.
func (d *DesignatorList) String() string {
	return PrettyString(d)
}

// DirectAbstractDeclarator represents data reduced by productions:
//
//	DirectAbstractDeclarator:
//	        '(' AbstractDeclarator ')'
//	|       DirectAbstractDeclaratorOpt '[' AssignmentExpressionOpt ']'                             // Case 1
//	|       DirectAbstractDeclaratorOpt '[' TypeQualifierList AssignmentExpressionOpt ']'           // Case 2
//	|       DirectAbstractDeclaratorOpt '[' "static" TypeQualifierListOpt AssignmentExpression ']'  // Case 3
//	|       DirectAbstractDeclaratorOpt '[' TypeQualifierList "static" AssignmentExpression ']'     // Case 4
//	|       DirectAbstractDeclaratorOpt '[' '*' ']'                                                 // Case 5
//	|       '(' ParameterTypeListOpt ')'                                                            // Case 6
//	|       DirectAbstractDeclarator '(' ParameterTypeListOpt ')'                                   // Case 7
type DirectAbstractDeclarator struct {
	indirection                 int  // 'int **i': 2.
	specifier                   Type // 'int i': specifier is 'int'
	AbstractDeclarator          *AbstractDeclarator
	AssignmentExpression        *AssignmentExpression
	AssignmentExpressionOpt     *AssignmentExpressionOpt
	Case                        int
	DirectAbstractDeclarator    *DirectAbstractDeclarator
	DirectAbstractDeclaratorOpt *DirectAbstractDeclaratorOpt
	ParameterTypeListOpt        *ParameterTypeListOpt
	Token                       xc.Token
	Token2                      xc.Token
	Token3                      xc.Token
	TypeQualifierList           *TypeQualifierList
	TypeQualifierListOpt        *TypeQualifierListOpt
}

func (d *DirectAbstractDeclarator) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DirectAbstractDeclarator) String() string {
	return PrettyString(d)
}

// DirectAbstractDeclaratorOpt represents data reduced by productions:
//
//	DirectAbstractDeclaratorOpt:
//	        /* empty */
//	|       DirectAbstractDeclarator  // Case 1
type DirectAbstractDeclaratorOpt struct {
	DirectAbstractDeclarator *DirectAbstractDeclarator
}

func (d *DirectAbstractDeclaratorOpt) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DirectAbstractDeclaratorOpt) String() string {
	return PrettyString(d)
}

// DirectDeclarator represents data reduced by productions:
//
//	DirectDeclarator:
//	        IDENTIFIER
//	|       '(' Declarator ')'                                                           // Case 1
//	|       DirectDeclarator '[' TypeQualifierListOpt AssignmentExpressionOpt ']'        // Case 2
//	|       DirectDeclarator '[' "static" TypeQualifierListOpt AssignmentExpression ']'  // Case 3
//	|       DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'     // Case 4
//	|       DirectDeclarator '[' TypeQualifierListOpt '*' ']'                            // Case 5
//	|       DirectDeclarator '(' DirectDeclarator2                                       // Case 6
type DirectDeclarator struct {
	indirection             int  // 'int **i': 2.
	specifier               Type // 'int i': specifier is 'int'
	AssignmentExpression    *AssignmentExpression
	AssignmentExpressionOpt *AssignmentExpressionOpt
	Case                    int
	Declarator              *Declarator
	DirectDeclarator        *DirectDeclarator
	DirectDeclarator2       *DirectDeclarator2
	Token                   xc.Token
	Token2                  xc.Token
	Token3                  xc.Token
	TypeQualifierList       *TypeQualifierList
	TypeQualifierListOpt    *TypeQualifierListOpt
}

func (d *DirectDeclarator) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DirectDeclarator) String() string {
	return PrettyString(d)
}

// DirectDeclarator2 represents data reduced by productions:
//
//	DirectDeclarator2:
//	        ParameterTypeList ')'
//	|       IdentifierListOpt ')'  // Case 1
type DirectDeclarator2 struct {
	Case              int
	IdentifierListOpt *IdentifierListOpt
	ParameterTypeList *ParameterTypeList
	Token             xc.Token
}

func (d *DirectDeclarator2) fragment() interface{} { return d }

// String implements fmt.Stringer.
func (d *DirectDeclarator2) String() string {
	return PrettyString(d)
}

// ElifGroup represents data reduced by production:
//
//	ElifGroup:
//	        PPELIF PpTokenList GroupListOpt
type ElifGroup struct {
	GroupListOpt *GroupListOpt
	PpTokenList  PpTokenList
	Token        xc.Token
}

func (e *ElifGroup) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ElifGroup) String() string {
	return PrettyString(e)
}

// ElifGroupList represents data reduced by productions:
//
//	ElifGroupList:
//	        ElifGroup
//	|       ElifGroupList ElifGroup  // Case 1
type ElifGroupList struct {
	Case          int
	ElifGroup     *ElifGroup
	ElifGroupList *ElifGroupList
}

func (e *ElifGroupList) reverse() *ElifGroupList {
	if e == nil {
		return nil
	}

	na := e
	nb := na.ElifGroupList
	for nb != nil {
		nc := nb.ElifGroupList
		nb.ElifGroupList = na
		na = nb
		nb = nc
	}
	e.ElifGroupList = nil
	return na
}

func (e *ElifGroupList) fragment() interface{} { return e.reverse() }

// String implements fmt.Stringer.
func (e *ElifGroupList) String() string {
	return PrettyString(e)
}

// ElifGroupListOpt represents data reduced by productions:
//
//	ElifGroupListOpt:
//	        /* empty */
//	|       ElifGroupList  // Case 1
type ElifGroupListOpt struct {
	ElifGroupList *ElifGroupList
}

func (e *ElifGroupListOpt) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ElifGroupListOpt) String() string {
	return PrettyString(e)
}

// ElseGroup represents data reduced by production:
//
//	ElseGroup:
//	        PPELSE '\n' GroupListOpt
type ElseGroup struct {
	GroupListOpt *GroupListOpt
	Token        xc.Token
	Token2       xc.Token
}

func (e *ElseGroup) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ElseGroup) String() string {
	return PrettyString(e)
}

// ElseGroupOpt represents data reduced by productions:
//
//	ElseGroupOpt:
//	        /* empty */
//	|       ElseGroup    // Case 1
type ElseGroupOpt struct {
	ElseGroup *ElseGroup
}

func (e *ElseGroupOpt) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ElseGroupOpt) String() string {
	return PrettyString(e)
}

// EndifLine represents data reduced by production:
//
//	EndifLine:
//	        PPENDIF PpTokenListOpt
type EndifLine struct {
	PpTokenListOpt PpTokenList
	Token          xc.Token
}

func (e *EndifLine) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *EndifLine) String() string {
	return PrettyString(e)
}

// EnumSpecifier represents data reduced by productions:
//
//	EnumSpecifier:
//	        EnumSpecifier0 '{' EnumeratorList '}'
//	|       EnumSpecifier0 '{' EnumeratorList ',' '}'  // Case 1
//	|       "enum" IDENTIFIER                          // Case 2
type EnumSpecifier struct {
	Case           int
	EnumSpecifier0 *EnumSpecifier0
	EnumeratorList *EnumeratorList
	Token          xc.Token
	Token2         xc.Token
	Token3         xc.Token
}

func (e *EnumSpecifier) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *EnumSpecifier) String() string {
	return PrettyString(e)
}

// EnumSpecifier0 represents data reduced by production:
//
//	EnumSpecifier0:
//	        "enum" IdentifierOpt
type EnumSpecifier0 struct {
	IdentifierOpt *IdentifierOpt
	Token         xc.Token
}

func (e *EnumSpecifier0) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *EnumSpecifier0) String() string {
	return PrettyString(e)
}

// EnumerationConstant represents data reduced by production:
//
//	EnumerationConstant:
//	        IDENTIFIER
type EnumerationConstant struct {
	Token xc.Token
}

func (e *EnumerationConstant) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *EnumerationConstant) String() string {
	return PrettyString(e)
}

// Enumerator represents data reduced by productions:
//
//	Enumerator:
//	        EnumerationConstant
//	|       EnumerationConstant '=' ConstantExpression  // Case 1
type Enumerator struct {
	Case                int
	ConstantExpression  *ConstantExpression
	EnumerationConstant *EnumerationConstant
	Token               xc.Token
}

func (e *Enumerator) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *Enumerator) String() string {
	return PrettyString(e)
}

// EnumeratorList represents data reduced by productions:
//
//	EnumeratorList:
//	        Enumerator
//	|       EnumeratorList ',' Enumerator  // Case 1
type EnumeratorList struct {
	Case           int
	Enumerator     *Enumerator
	EnumeratorList *EnumeratorList
	Token          xc.Token
}

func (e *EnumeratorList) reverse() *EnumeratorList {
	if e == nil {
		return nil
	}

	na := e
	nb := na.EnumeratorList
	for nb != nil {
		nc := nb.EnumeratorList
		nb.EnumeratorList = na
		na = nb
		nb = nc
	}
	e.EnumeratorList = nil
	return na
}

func (e *EnumeratorList) fragment() interface{} { return e.reverse() }

// String implements fmt.Stringer.
func (e *EnumeratorList) String() string {
	return PrettyString(e)
}

// EqualityExpression represents data reduced by productions:
//
//	EqualityExpression:
//	        RelationalExpression
//	|       EqualityExpression "==" RelationalExpression  // Case 1
//	|       EqualityExpression "!=" RelationalExpression  // Case 2
type EqualityExpression struct {
	Case                 int
	EqualityExpression   *EqualityExpression
	RelationalExpression *RelationalExpression
	Token                xc.Token
}

func (e *EqualityExpression) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *EqualityExpression) String() string {
	return PrettyString(e)
}

// ExclusiveOrExpression represents data reduced by productions:
//
//	ExclusiveOrExpression:
//	        AndExpression
//	|       ExclusiveOrExpression '^' AndExpression  // Case 1
type ExclusiveOrExpression struct {
	AndExpression         *AndExpression
	Case                  int
	ExclusiveOrExpression *ExclusiveOrExpression
	Token                 xc.Token
}

func (e *ExclusiveOrExpression) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ExclusiveOrExpression) String() string {
	return PrettyString(e)
}

// ExpressionList represents data reduced by productions:
//
//	ExpressionList:
//	        AssignmentExpression
//	|       ExpressionList ',' AssignmentExpression  // Case 1
type ExpressionList struct {
	AssignmentExpression *AssignmentExpression
	Case                 int
	ExpressionList       *ExpressionList
	Token                xc.Token
}

func (e *ExpressionList) reverse() *ExpressionList {
	if e == nil {
		return nil
	}

	na := e
	nb := na.ExpressionList
	for nb != nil {
		nc := nb.ExpressionList
		nb.ExpressionList = na
		na = nb
		nb = nc
	}
	e.ExpressionList = nil
	return na
}

func (e *ExpressionList) fragment() interface{} { return e.reverse() }

// String implements fmt.Stringer.
func (e *ExpressionList) String() string {
	return PrettyString(e)
}

// ExpressionOpt represents data reduced by productions:
//
//	ExpressionOpt:
//	        /* empty */
//	|       ExpressionList  // Case 1
type ExpressionOpt struct {
	ExpressionList *ExpressionList
}

func (e *ExpressionOpt) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ExpressionOpt) String() string {
	return PrettyString(e)
}

// ExpressionStatement represents data reduced by production:
//
//	ExpressionStatement:
//	        ExpressionOpt ';'
type ExpressionStatement struct {
	ExpressionOpt *ExpressionOpt
	Token         xc.Token
}

func (e *ExpressionStatement) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ExpressionStatement) String() string {
	return PrettyString(e)
}

// ExternalDeclaration represents data reduced by productions:
//
//	ExternalDeclaration:
//	        FunctionDefinition
//	|       Declaration         // Case 1
type ExternalDeclaration struct {
	Case               int
	Declaration        *Declaration
	FunctionDefinition *FunctionDefinition
}

func (e *ExternalDeclaration) fragment() interface{} { return e }

// String implements fmt.Stringer.
func (e *ExternalDeclaration) String() string {
	return PrettyString(e)
}

// FunctionDefinition represents data reduced by production:
//
//	FunctionDefinition:
//	        DeclarationSpecifiers Declarator DeclarationListOpt CompoundStatement
type FunctionDefinition struct {
	fnScope               *Bindings
	CompoundStatement     *CompoundStatement
	DeclarationListOpt    *DeclarationListOpt
	DeclarationSpecifiers *DeclarationSpecifiers
	Declarator            *Declarator
}

func (f *FunctionDefinition) fragment() interface{} { return f }

// String implements fmt.Stringer.
func (f *FunctionDefinition) String() string {
	return PrettyString(f)
}

// FunctionSpecifier represents data reduced by production:
//
//	FunctionSpecifier:
//	        "inline"
type FunctionSpecifier struct {
	Token xc.Token
}

func (f *FunctionSpecifier) fragment() interface{} { return f }

// String implements fmt.Stringer.
func (f *FunctionSpecifier) String() string {
	return PrettyString(f)
}

// GroupListOpt represents data reduced by productions:
//
//	GroupListOpt:
//	        /* empty */
//	|       GroupList    // Case 1
type GroupListOpt struct {
	GroupList *GroupList
}

func (g *GroupListOpt) fragment() interface{} { return g }

// String implements fmt.Stringer.
func (g *GroupListOpt) String() string {
	return PrettyString(g)
}

// IdentifierList represents data reduced by productions:
//
//	IdentifierList:
//	        IDENTIFIER
//	|       IdentifierList ',' IDENTIFIER  // Case 1
type IdentifierList struct {
	Case           int
	IdentifierList *IdentifierList
	Token          xc.Token
	Token2         xc.Token
}

func (i *IdentifierList) reverse() *IdentifierList {
	if i == nil {
		return nil
	}

	na := i
	nb := na.IdentifierList
	for nb != nil {
		nc := nb.IdentifierList
		nb.IdentifierList = na
		na = nb
		nb = nc
	}
	i.IdentifierList = nil
	return na
}

func (i *IdentifierList) fragment() interface{} { return i.reverse() }

// String implements fmt.Stringer.
func (i *IdentifierList) String() string {
	return PrettyString(i)
}

// IdentifierListOpt represents data reduced by productions:
//
//	IdentifierListOpt:
//	        /* empty */
//	|       IdentifierList  // Case 1
type IdentifierListOpt struct {
	IdentifierList *IdentifierList
}

func (i *IdentifierListOpt) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *IdentifierListOpt) String() string {
	return PrettyString(i)
}

// IdentifierOpt represents data reduced by productions:
//
//	IdentifierOpt:
//	        /* empty */
//	|       IDENTIFIER   // Case 1
type IdentifierOpt struct {
	Token xc.Token
}

func (i *IdentifierOpt) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *IdentifierOpt) String() string {
	return PrettyString(i)
}

// IfGroup represents data reduced by productions:
//
//	IfGroup:
//	        PPIF PpTokenList GroupListOpt
//	|       PPIFDEF IDENTIFIER '\n' GroupListOpt   // Case 1
//	|       PPIFNDEF IDENTIFIER '\n' GroupListOpt  // Case 2
type IfGroup struct {
	Case         int
	GroupListOpt *GroupListOpt
	PpTokenList  PpTokenList
	Token        xc.Token
	Token2       xc.Token
	Token3       xc.Token
}

func (i *IfGroup) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *IfGroup) String() string {
	return PrettyString(i)
}

// IfSection represents data reduced by production:
//
//	IfSection:
//	        IfGroup ElifGroupListOpt ElseGroupOpt EndifLine
type IfSection struct {
	ElifGroupListOpt *ElifGroupListOpt
	ElseGroupOpt     *ElseGroupOpt
	EndifLine        *EndifLine
	IfGroup          *IfGroup
}

func (i *IfSection) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *IfSection) String() string {
	return PrettyString(i)
}

// InclusiveOrExpression represents data reduced by productions:
//
//	InclusiveOrExpression:
//	        ExclusiveOrExpression
//	|       InclusiveOrExpression '|' ExclusiveOrExpression  // Case 1
type InclusiveOrExpression struct {
	Case                  int
	ExclusiveOrExpression *ExclusiveOrExpression
	InclusiveOrExpression *InclusiveOrExpression
	Token                 xc.Token
}

func (i *InclusiveOrExpression) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *InclusiveOrExpression) String() string {
	return PrettyString(i)
}

// InitDeclarator represents data reduced by productions:
//
//	InitDeclarator:
//	        Declarator
//	|       Declarator '=' Initializer  // Case 1
type InitDeclarator struct {
	Case        int
	Declarator  *Declarator
	Initializer *Initializer
	Token       xc.Token
}

func (i *InitDeclarator) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *InitDeclarator) String() string {
	return PrettyString(i)
}

// InitDeclaratorList represents data reduced by productions:
//
//	InitDeclaratorList:
//	        InitDeclarator
//	|       InitDeclaratorList ',' InitDeclarator  // Case 1
type InitDeclaratorList struct {
	Case               int
	InitDeclarator     *InitDeclarator
	InitDeclaratorList *InitDeclaratorList
	Token              xc.Token
}

func (i *InitDeclaratorList) reverse() *InitDeclaratorList {
	if i == nil {
		return nil
	}

	na := i
	nb := na.InitDeclaratorList
	for nb != nil {
		nc := nb.InitDeclaratorList
		nb.InitDeclaratorList = na
		na = nb
		nb = nc
	}
	i.InitDeclaratorList = nil
	return na
}

func (i *InitDeclaratorList) fragment() interface{} { return i.reverse() }

// String implements fmt.Stringer.
func (i *InitDeclaratorList) String() string {
	return PrettyString(i)
}

// InitDeclaratorListOpt represents data reduced by productions:
//
//	InitDeclaratorListOpt:
//	        /* empty */
//	|       InitDeclaratorList  // Case 1
type InitDeclaratorListOpt struct {
	InitDeclaratorList *InitDeclaratorList
}

func (i *InitDeclaratorListOpt) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *InitDeclaratorListOpt) String() string {
	return PrettyString(i)
}

// Initializer represents data reduced by productions:
//
//	Initializer:
//	        AssignmentExpression
//	|       '{' InitializerList '}'      // Case 1
//	|       '{' InitializerList ',' '}'  // Case 2
type Initializer struct {
	AssignmentExpression *AssignmentExpression
	Case                 int
	InitializerList      *InitializerList
	Token                xc.Token
	Token2               xc.Token
	Token3               xc.Token
}

func (i *Initializer) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *Initializer) String() string {
	return PrettyString(i)
}

// InitializerList represents data reduced by productions:
//
//	InitializerList:
//	        DesignationOpt Initializer
//	|       InitializerList ',' DesignationOpt Initializer  // Case 1
type InitializerList struct {
	Case            int
	DesignationOpt  *DesignationOpt
	Initializer     *Initializer
	InitializerList *InitializerList
	Token           xc.Token
}

func (i *InitializerList) reverse() *InitializerList {
	if i == nil {
		return nil
	}

	na := i
	nb := na.InitializerList
	for nb != nil {
		nc := nb.InitializerList
		nb.InitializerList = na
		na = nb
		nb = nc
	}
	i.InitializerList = nil
	return na
}

func (i *InitializerList) fragment() interface{} { return i.reverse() }

// String implements fmt.Stringer.
func (i *InitializerList) String() string {
	return PrettyString(i)
}

// IterationStatement represents data reduced by productions:
//
//	IterationStatement:
//	        "while" '(' ExpressionList ')' Statement
//	|       "do" Statement "while" '(' ExpressionList ')' ';'                          // Case 1
//	|       "for" '(' ExpressionOpt ';' ExpressionOpt ';' ExpressionOpt ')' Statement  // Case 2
//	|       "for" '(' Declaration ExpressionOpt ';' ExpressionOpt ')' Statement        // Case 3
type IterationStatement struct {
	Case           int
	Declaration    *Declaration
	ExpressionList *ExpressionList
	ExpressionOpt  *ExpressionOpt
	ExpressionOpt2 *ExpressionOpt
	ExpressionOpt3 *ExpressionOpt
	Statement      *Statement
	Token          xc.Token
	Token2         xc.Token
	Token3         xc.Token
	Token4         xc.Token
	Token5         xc.Token
}

func (i *IterationStatement) fragment() interface{} { return i }

// String implements fmt.Stringer.
func (i *IterationStatement) String() string {
	return PrettyString(i)
}

// JumpStatement represents data reduced by productions:
//
//	JumpStatement:
//	        "goto" IDENTIFIER ';'
//	|       "continue" ';'              // Case 1
//	|       "break" ';'                 // Case 2
//	|       "return" ExpressionOpt ';'  // Case 3
type JumpStatement struct {
	Case          int
	ExpressionOpt *ExpressionOpt
	Token         xc.Token
	Token2        xc.Token
	Token3        xc.Token
}

func (j *JumpStatement) fragment() interface{} { return j }

// String implements fmt.Stringer.
func (j *JumpStatement) String() string {
	return PrettyString(j)
}

// LabeledStatement represents data reduced by productions:
//
//	LabeledStatement:
//	        IDENTIFIER ':' Statement
//	|       "case" ConstantExpression ':' Statement  // Case 1
//	|       "default" ':' Statement                  // Case 2
type LabeledStatement struct {
	Case               int
	ConstantExpression *ConstantExpression
	Statement          *Statement
	Token              xc.Token
	Token2             xc.Token
}

func (l *LabeledStatement) fragment() interface{} { return l }

// String implements fmt.Stringer.
func (l *LabeledStatement) String() string {
	return PrettyString(l)
}

// LogicalAndExpression represents data reduced by productions:
//
//	LogicalAndExpression:
//	        InclusiveOrExpression
//	|       LogicalAndExpression "&&" InclusiveOrExpression  // Case 1
type LogicalAndExpression struct {
	Case                  int
	InclusiveOrExpression *InclusiveOrExpression
	LogicalAndExpression  *LogicalAndExpression
	Token                 xc.Token
}

func (l *LogicalAndExpression) fragment() interface{} { return l }

// String implements fmt.Stringer.
func (l *LogicalAndExpression) String() string {
	return PrettyString(l)
}

// LogicalOrExpression represents data reduced by productions:
//
//	LogicalOrExpression:
//	        LogicalAndExpression
//	|       LogicalOrExpression "||" LogicalAndExpression  // Case 1
type LogicalOrExpression struct {
	Case                 int
	LogicalAndExpression *LogicalAndExpression
	LogicalOrExpression  *LogicalOrExpression
	Token                xc.Token
}

func (l *LogicalOrExpression) fragment() interface{} { return l }

// String implements fmt.Stringer.
func (l *LogicalOrExpression) String() string {
	return PrettyString(l)
}

// MultiplicativeExpression represents data reduced by productions:
//
//	MultiplicativeExpression:
//	        CastExpression
//	|       MultiplicativeExpression '*' CastExpression  // Case 1
//	|       MultiplicativeExpression '/' CastExpression  // Case 2
//	|       MultiplicativeExpression '%' CastExpression  // Case 3
type MultiplicativeExpression struct {
	Case                     int
	CastExpression           *CastExpression
	MultiplicativeExpression *MultiplicativeExpression
	Token                    xc.Token
}

func (m *MultiplicativeExpression) fragment() interface{} { return m }

// String implements fmt.Stringer.
func (m *MultiplicativeExpression) String() string {
	return PrettyString(m)
}

// ParameterDeclaration represents data reduced by productions:
//
//	ParameterDeclaration:
//	        DeclarationSpecifiers Declarator
//	|       DeclarationSpecifiers AbstractDeclaratorOpt  // Case 1
type ParameterDeclaration struct {
	AbstractDeclaratorOpt *AbstractDeclaratorOpt
	Case                  int
	DeclarationSpecifiers *DeclarationSpecifiers
	Declarator            *Declarator
}

func (p *ParameterDeclaration) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *ParameterDeclaration) String() string {
	return PrettyString(p)
}

// ParameterList represents data reduced by productions:
//
//	ParameterList:
//	        ParameterDeclaration
//	|       ParameterList ',' ParameterDeclaration  // Case 1
type ParameterList struct {
	Case                 int
	ParameterDeclaration *ParameterDeclaration
	ParameterList        *ParameterList
	Token                xc.Token
}

func (p *ParameterList) reverse() *ParameterList {
	if p == nil {
		return nil
	}

	na := p
	nb := na.ParameterList
	for nb != nil {
		nc := nb.ParameterList
		nb.ParameterList = na
		na = nb
		nb = nc
	}
	p.ParameterList = nil
	return na
}

func (p *ParameterList) fragment() interface{} { return p.reverse() }

// String implements fmt.Stringer.
func (p *ParameterList) String() string {
	return PrettyString(p)
}

// ParameterTypeList represents data reduced by productions:
//
//	ParameterTypeList:
//	        ParameterList
//	|       ParameterList ',' "..."  // Case 1
type ParameterTypeList struct {
	Case          int
	ParameterList *ParameterList
	Token         xc.Token
	Token2        xc.Token
}

func (p *ParameterTypeList) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *ParameterTypeList) String() string {
	return PrettyString(p)
}

// ParameterTypeListOpt represents data reduced by productions:
//
//	ParameterTypeListOpt:
//	        /* empty */
//	|       ParameterTypeList  // Case 1
type ParameterTypeListOpt struct {
	ParameterTypeList *ParameterTypeList
}

func (p *ParameterTypeListOpt) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *ParameterTypeListOpt) String() string {
	return PrettyString(p)
}

// Pointer represents data reduced by productions:
//
//	Pointer:
//	        '*' TypeQualifierListOpt
//	|       '*' TypeQualifierListOpt Pointer  // Case 1
type Pointer struct {
	indirection          int
	Case                 int
	Pointer              *Pointer
	Token                xc.Token
	TypeQualifierListOpt *TypeQualifierListOpt
}

func (p *Pointer) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *Pointer) String() string {
	return PrettyString(p)
}

// PointerOpt represents data reduced by productions:
//
//	PointerOpt:
//	        /* empty */
//	|       Pointer      // Case 1
type PointerOpt struct {
	Pointer *Pointer
}

func (p *PointerOpt) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *PointerOpt) String() string {
	return PrettyString(p)
}

// PostfixExpression represents data reduced by productions:
//
//	PostfixExpression:
//	        PrimaryExpression
//	|       PostfixExpression '[' ExpressionList ']'             // Case 1
//	|       PostfixExpression '(' ArgumentExpressionListOpt ')'  // Case 2
//	|       PostfixExpression '.' IDENTIFIER                     // Case 3
//	|       PostfixExpression "->" IDENTIFIER                    // Case 4
//	|       PostfixExpression "++"                               // Case 5
//	|       PostfixExpression "--"                               // Case 6
//	|       '(' TypeName ')' '{' InitializerList '}'             // Case 7
//	|       '(' TypeName ')' '{' InitializerList ',' '}'         // Case 8
type PostfixExpression struct {
	ArgumentExpressionListOpt *ArgumentExpressionListOpt
	Case                      int
	ExpressionList            *ExpressionList
	InitializerList           *InitializerList
	PostfixExpression         *PostfixExpression
	PrimaryExpression         *PrimaryExpression
	Token                     xc.Token
	Token2                    xc.Token
	Token3                    xc.Token
	Token4                    xc.Token
	Token5                    xc.Token
	TypeName                  *TypeName
}

func (p *PostfixExpression) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *PostfixExpression) String() string {
	return PrettyString(p)
}

// PreprocessingFile represents data reduced by production:
//
//	PreprocessingFile:
//	        GroupList
type PreprocessingFile struct {
	file      *token.File
	GroupList *GroupList
}

func (p *PreprocessingFile) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *PreprocessingFile) String() string {
	return PrettyString(p)
}

// PrimaryExpression represents data reduced by productions:
//
//	PrimaryExpression:
//	        IDENTIFIER
//	|       Constant                // Case 1
//	|       '(' ExpressionList ')'  // Case 2
type PrimaryExpression struct {
	Case           int
	Constant       *Constant
	ExpressionList *ExpressionList
	Token          xc.Token
	Token2         xc.Token
}

func (p *PrimaryExpression) fragment() interface{} { return p }

// String implements fmt.Stringer.
func (p *PrimaryExpression) String() string {
	return PrettyString(p)
}

// RelationalExpression represents data reduced by productions:
//
//	RelationalExpression:
//	        ShiftExpression
//	|       RelationalExpression '<' ShiftExpression   // Case 1
//	|       RelationalExpression '>' ShiftExpression   // Case 2
//	|       RelationalExpression "<=" ShiftExpression  // Case 3
//	|       RelationalExpression ">=" ShiftExpression  // Case 4
type RelationalExpression struct {
	Case                 int
	RelationalExpression *RelationalExpression
	ShiftExpression      *ShiftExpression
	Token                xc.Token
}

func (r *RelationalExpression) fragment() interface{} { return r }

// String implements fmt.Stringer.
func (r *RelationalExpression) String() string {
	return PrettyString(r)
}

// SelectionStatement represents data reduced by productions:
//
//	SelectionStatement:
//	        "if" '(' ExpressionList ')' Statement
//	|       "if" '(' ExpressionList ')' Statement "else" Statement  // Case 1
//	|       "switch" '(' ExpressionList ')' Statement               // Case 2
type SelectionStatement struct {
	Case           int
	ExpressionList *ExpressionList
	Statement      *Statement
	Statement2     *Statement
	Token          xc.Token
	Token2         xc.Token
	Token3         xc.Token
	Token4         xc.Token
}

func (s *SelectionStatement) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *SelectionStatement) String() string {
	return PrettyString(s)
}

// ShiftExpression represents data reduced by productions:
//
//	ShiftExpression:
//	        AdditiveExpression
//	|       ShiftExpression "<<" AdditiveExpression  // Case 1
//	|       ShiftExpression ">>" AdditiveExpression  // Case 2
type ShiftExpression struct {
	AdditiveExpression *AdditiveExpression
	Case               int
	ShiftExpression    *ShiftExpression
	Token              xc.Token
}

func (s *ShiftExpression) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *ShiftExpression) String() string {
	return PrettyString(s)
}

// SpecifierQualifierList represents data reduced by productions:
//
//	SpecifierQualifierList:
//	        TypeSpecifier SpecifierQualifierListOpt
//	|       TypeQualifier SpecifierQualifierListOpt  // Case 1
type SpecifierQualifierList struct {
	IsConst                   bool // TypeQualifier "const" is present.
	IsRestrict                bool // TypeQualifier "restrict" is present.
	IsVolatile                bool // TypeQualifier "volatile" is present.
	typ                       int  // One of tsVoid, tsChar, tsUChar, ...
	typeSpecifiers            int  //
	Case                      int
	SpecifierQualifierListOpt *SpecifierQualifierListOpt
	TypeQualifier             *TypeQualifier
	TypeSpecifier             *TypeSpecifier
}

func (s *SpecifierQualifierList) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *SpecifierQualifierList) String() string {
	return PrettyString(s)
}

// SpecifierQualifierListOpt represents data reduced by productions:
//
//	SpecifierQualifierListOpt:
//	        /* empty */
//	|       SpecifierQualifierList  // Case 1
type SpecifierQualifierListOpt struct {
	SpecifierQualifierList *SpecifierQualifierList
}

func (s *SpecifierQualifierListOpt) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *SpecifierQualifierListOpt) String() string {
	return PrettyString(s)
}

// Statement represents data reduced by productions:
//
//	Statement:
//	        LabeledStatement
//	|       CompoundStatement    // Case 1
//	|       ExpressionStatement  // Case 2
//	|       SelectionStatement   // Case 3
//	|       IterationStatement   // Case 4
//	|       JumpStatement        // Case 5
type Statement struct {
	Case                int
	CompoundStatement   *CompoundStatement
	ExpressionStatement *ExpressionStatement
	IterationStatement  *IterationStatement
	JumpStatement       *JumpStatement
	LabeledStatement    *LabeledStatement
	SelectionStatement  *SelectionStatement
}

func (s *Statement) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *Statement) String() string {
	return PrettyString(s)
}

// StorageClassSpecifier represents data reduced by productions:
//
//	StorageClassSpecifier:
//	        "typedef"
//	|       "extern"    // Case 1
//	|       "static"    // Case 2
//	|       "auto"      // Case 3
//	|       "register"  // Case 4
type StorageClassSpecifier struct {
	Case  int
	Token xc.Token
}

func (s *StorageClassSpecifier) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StorageClassSpecifier) String() string {
	return PrettyString(s)
}

// StructDeclaration represents data reduced by production:
//
//	StructDeclaration:
//	        SpecifierQualifierList StructDeclaratorListOpt ';'
type StructDeclaration struct {
	SpecifierQualifierList  *SpecifierQualifierList
	StructDeclaratorListOpt *StructDeclaratorListOpt
	Token                   xc.Token
}

func (s *StructDeclaration) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StructDeclaration) String() string {
	return PrettyString(s)
}

// StructDeclarationList represents data reduced by productions:
//
//	StructDeclarationList:
//	        StructDeclaration
//	|       StructDeclarationList StructDeclaration  // Case 1
type StructDeclarationList struct {
	Case                  int
	StructDeclaration     *StructDeclaration
	StructDeclarationList *StructDeclarationList
}

func (s *StructDeclarationList) reverse() *StructDeclarationList {
	if s == nil {
		return nil
	}

	na := s
	nb := na.StructDeclarationList
	for nb != nil {
		nc := nb.StructDeclarationList
		nb.StructDeclarationList = na
		na = nb
		nb = nc
	}
	s.StructDeclarationList = nil
	return na
}

func (s *StructDeclarationList) fragment() interface{} { return s.reverse() }

// String implements fmt.Stringer.
func (s *StructDeclarationList) String() string {
	return PrettyString(s)
}

// StructDeclarator represents data reduced by productions:
//
//	StructDeclarator:
//	        Declarator
//	|       DeclaratorOpt ':' ConstantExpression  // Case 1
type StructDeclarator struct {
	align
	offset
	size
	bits               Type
	Case               int
	ConstantExpression *ConstantExpression
	Declarator         *Declarator
	DeclaratorOpt      *DeclaratorOpt
	Token              xc.Token
}

func (s *StructDeclarator) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StructDeclarator) String() string {
	return PrettyString(s)
}

// StructDeclaratorList represents data reduced by productions:
//
//	StructDeclaratorList:
//	        StructDeclarator
//	|       StructDeclaratorList ',' StructDeclarator  // Case 1
type StructDeclaratorList struct {
	Case                 int
	StructDeclarator     *StructDeclarator
	StructDeclaratorList *StructDeclaratorList
	Token                xc.Token
}

func (s *StructDeclaratorList) reverse() *StructDeclaratorList {
	if s == nil {
		return nil
	}

	na := s
	nb := na.StructDeclaratorList
	for nb != nil {
		nc := nb.StructDeclaratorList
		nb.StructDeclaratorList = na
		na = nb
		nb = nc
	}
	s.StructDeclaratorList = nil
	return na
}

func (s *StructDeclaratorList) fragment() interface{} { return s.reverse() }

// String implements fmt.Stringer.
func (s *StructDeclaratorList) String() string {
	return PrettyString(s)
}

// StructDeclaratorListOpt represents data reduced by productions:
//
//	StructDeclaratorListOpt:
//	        /* empty */
//	|       StructDeclaratorList  // Case 1
type StructDeclaratorListOpt struct {
	StructDeclaratorList *StructDeclaratorList
}

func (s *StructDeclaratorListOpt) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StructDeclaratorListOpt) String() string {
	return PrettyString(s)
}

// StructOrUnion represents data reduced by productions:
//
//	StructOrUnion:
//	        "struct"
//	|       "union"   // Case 1
type StructOrUnion struct {
	Case  int
	Token xc.Token
}

func (s *StructOrUnion) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StructOrUnion) String() string {
	return PrettyString(s)
}

// StructOrUnionSpecifier represents data reduced by productions:
//
//	StructOrUnionSpecifier:
//	        StructOrUnionSpecifier0 '{' StructDeclarationList '}'
//	|       StructOrUnion IDENTIFIER                               // Case 1
type StructOrUnionSpecifier struct {
	Members *Bindings
	align
	bindings *Bindings
	isUnion  bool
	size
	Case                    int
	StructDeclarationList   *StructDeclarationList
	StructOrUnion           *StructOrUnion
	StructOrUnionSpecifier0 *StructOrUnionSpecifier0
	Token                   xc.Token
	Token2                  xc.Token
}

func (s *StructOrUnionSpecifier) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StructOrUnionSpecifier) String() string {
	return PrettyString(s)
}

// StructOrUnionSpecifier0 represents data reduced by production:
//
//	StructOrUnionSpecifier0:
//	        StructOrUnion IdentifierOpt
type StructOrUnionSpecifier0 struct {
	SUSpecifier   *StructOrUnionSpecifier
	IdentifierOpt *IdentifierOpt
	StructOrUnion *StructOrUnion
}

func (s *StructOrUnionSpecifier0) fragment() interface{} { return s }

// String implements fmt.Stringer.
func (s *StructOrUnionSpecifier0) String() string {
	return PrettyString(s)
}

// TranslationUnit represents data reduced by productions:
//
//	TranslationUnit:
//	        ExternalDeclaration
//	|       TranslationUnit ExternalDeclaration  // Case 1
type TranslationUnit struct {
	Declarations        *Bindings
	Case                int
	ExternalDeclaration *ExternalDeclaration
	TranslationUnit     *TranslationUnit
}

func (t *TranslationUnit) reverse() *TranslationUnit {
	if t == nil {
		return nil
	}

	na := t
	nb := na.TranslationUnit
	for nb != nil {
		nc := nb.TranslationUnit
		nb.TranslationUnit = na
		na = nb
		nb = nc
	}
	t.TranslationUnit = nil
	return na
}

func (t *TranslationUnit) fragment() interface{} { return t.reverse() }

// String implements fmt.Stringer.
func (t *TranslationUnit) String() string {
	return PrettyString(t)
}

// TypeName represents data reduced by production:
//
//	TypeName:
//	        SpecifierQualifierList AbstractDeclaratorOpt
type TypeName struct {
	AbstractDeclaratorOpt  *AbstractDeclaratorOpt
	SpecifierQualifierList *SpecifierQualifierList
}

func (t *TypeName) fragment() interface{} { return t }

// String implements fmt.Stringer.
func (t *TypeName) String() string {
	return PrettyString(t)
}

// TypeQualifier represents data reduced by productions:
//
//	TypeQualifier:
//	        "const"
//	|       "restrict"  // Case 1
//	|       "volatile"  // Case 2
type TypeQualifier struct {
	Case  int
	Token xc.Token
}

func (t *TypeQualifier) fragment() interface{} { return t }

// String implements fmt.Stringer.
func (t *TypeQualifier) String() string {
	return PrettyString(t)
}

// TypeQualifierList represents data reduced by productions:
//
//	TypeQualifierList:
//	        TypeQualifier
//	|       TypeQualifierList TypeQualifier  // Case 1
type TypeQualifierList struct {
	Case              int
	TypeQualifier     *TypeQualifier
	TypeQualifierList *TypeQualifierList
}

func (t *TypeQualifierList) reverse() *TypeQualifierList {
	if t == nil {
		return nil
	}

	na := t
	nb := na.TypeQualifierList
	for nb != nil {
		nc := nb.TypeQualifierList
		nb.TypeQualifierList = na
		na = nb
		nb = nc
	}
	t.TypeQualifierList = nil
	return na
}

func (t *TypeQualifierList) fragment() interface{} { return t.reverse() }

// String implements fmt.Stringer.
func (t *TypeQualifierList) String() string {
	return PrettyString(t)
}

// TypeQualifierListOpt represents data reduced by productions:
//
//	TypeQualifierListOpt:
//	        /* empty */
//	|       TypeQualifierList  // Case 1
type TypeQualifierListOpt struct {
	TypeQualifierList *TypeQualifierList
}

func (t *TypeQualifierListOpt) fragment() interface{} { return t }

// String implements fmt.Stringer.
func (t *TypeQualifierListOpt) String() string {
	return PrettyString(t)
}

// TypeSpecifier represents data reduced by productions:
//
//	TypeSpecifier:
//	        "void"
//	|       "char"                  // Case 1
//	|       "short"                 // Case 2
//	|       "int"                   // Case 3
//	|       "long"                  // Case 4
//	|       "float"                 // Case 5
//	|       "double"                // Case 6
//	|       "signed"                // Case 7
//	|       "unsigned"              // Case 8
//	|       "_Bool"                 // Case 9
//	|       "_Complex"              // Case 10
//	|       StructOrUnionSpecifier  // Case 11
//	|       EnumSpecifier           // Case 12
//	|       TYPEDEFNAME             // Case 13
type TypeSpecifier struct {
	bindings               *Bindings
	case2                  int // {tsStruct,tsUnion}
	Case                   int
	EnumSpecifier          *EnumSpecifier
	StructOrUnionSpecifier *StructOrUnionSpecifier
	Token                  xc.Token
}

func (t *TypeSpecifier) fragment() interface{} { return t }

// String implements fmt.Stringer.
func (t *TypeSpecifier) String() string {
	return PrettyString(t)
}

// UnaryExpression represents data reduced by productions:
//
//	UnaryExpression:
//	        PostfixExpression
//	|       "++" UnaryExpression          // Case 1
//	|       "--" UnaryExpression          // Case 2
//	|       UnaryOperator CastExpression  // Case 3
//	|       "sizeof" UnaryExpression      // Case 4
//	|       "sizeof" '(' TypeName ')'     // Case 5
//	|       "defined" IDENTIFIER          // Case 6
//	|       "defined" '(' IDENTIFIER ')'  // Case 7
type UnaryExpression struct {
	Case              int
	CastExpression    *CastExpression
	PostfixExpression *PostfixExpression
	Token             xc.Token
	Token2            xc.Token
	Token3            xc.Token
	Token4            xc.Token
	TypeName          *TypeName
	UnaryExpression   *UnaryExpression
	UnaryOperator     *UnaryOperator
}

func (u *UnaryExpression) fragment() interface{} { return u }

// String implements fmt.Stringer.
func (u *UnaryExpression) String() string {
	return PrettyString(u)
}

// UnaryOperator represents data reduced by productions:
//
//	UnaryOperator:
//	        '&'
//	|       '*'  // Case 1
//	|       '+'  // Case 2
//	|       '-'  // Case 3
//	|       '~'  // Case 4
//	|       '!'  // Case 5
type UnaryOperator struct {
	Case  int
	Token xc.Token
}

func (u *UnaryOperator) fragment() interface{} { return u }

// String implements fmt.Stringer.
func (u *UnaryOperator) String() string {
	return PrettyString(u)
}
