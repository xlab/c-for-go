// CAUTION: Generated file - DO NOT EDIT.

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

import __yyfmt__ "fmt"

import (
	"github.com/cznic/golex/lex"
	"github.com/xlab/c/xc"
)

type yySymType struct {
	yys       int
	Token     xc.Token
	groupPart node
	node      node
	toks      PPTokenList
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault           = 57444
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ANDAND              = 57347
	ANDASSIGN           = 57348
	ARROW               = 57349
	AUTO                = 57350
	BOOL                = 57351
	BREAK               = 57352
	CASE                = 57353
	CAST                = 57354
	CHAR                = 57355
	CHARCONST           = 57356
	COMPLEX             = 57357
	CONST               = 57358
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57359
	DDD                 = 57360
	DEC                 = 57361
	DEFAULT             = 57362
	DIVASSIGN           = 57363
	DO                  = 57364
	DOUBLE              = 57365
	ELSE                = 57366
	ENUM                = 57367
	EQ                  = 57368
	EXTERN              = 57369
	FLOAT               = 57370
	FLOATCONST          = 57371
	FOR                 = 57372
	GEQ                 = 57373
	GOTO                = 57374
	IDENTIFIER          = 57375
	IDENTIFIER_LPAREN   = 57376
	IF                  = 57377
	INC                 = 57378
	INLINE              = 57379
	INT                 = 57380
	INTCONST            = 57381
	LEQ                 = 57382
	LONG                = 57383
	LONGCHARCONST       = 57384
	LONGSTRINGLITERAL   = 57385
	LSH                 = 57386
	LSHASSIGN           = 57387
	MODASSIGN           = 57388
	MULASSIGN           = 57389
	NEQ                 = 57390
	NOELSE              = 57391
	ORASSIGN            = 57392
	OROR                = 57393
	PPDEFINE            = 57394
	PPELIF              = 57395
	PPELSE              = 57396
	PPENDIF             = 57397
	PPERROR             = 57398
	PPHASH_NL           = 57399
	PPHEADER_NAME       = 57400
	PPIF                = 57401
	PPIFDEF             = 57402
	PPIFNDEF            = 57403
	PPINCLUDE           = 57404
	PPLINE              = 57405
	PPNONDIRECTIVE      = 57406
	PPNUMBER            = 57407
	PPOTHER             = 57408
	PPPASTE             = 57409
	PPPRAGMA            = 57410
	PPUNDEF             = 57411
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57412
	RESTRICT            = 57413
	RETURN              = 57414
	RSH                 = 57415
	RSHASSIGN           = 57416
	SHORT               = 57417
	SIGNED              = 57418
	SIZEOF              = 57419
	STATIC              = 57420
	STRINGLITERAL       = 57421
	STRUCT              = 57422
	SUBASSIGN           = 57423
	SWITCH              = 57424
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57425
	TYPEDEFNAME         = 57426
	UNARY               = 57427
	UNION               = 57428
	UNSIGNED            = 57429
	VOID                = 57430
	VOLATILE            = 57431
	WHILE               = 57432
	XORASSIGN           = 57433
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -280
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (276x)
		42:      1,   // '*' (247x)
		57375:   2,   // IDENTIFIER (201x)
		38:      3,   // '&' (197x)
		43:      4,   // '+' (197x)
		45:      5,   // '-' (197x)
		57361:   6,   // DEC (197x)
		57378:   7,   // INC (197x)
		41:      8,   // ')' (176x)
		59:      9,   // ';' (175x)
		44:      10,  // ',' (168x)
		91:      11,  // '[' (149x)
		33:      12,  // '!' (130x)
		126:     13,  // '~' (130x)
		57356:   14,  // CHARCONST (130x)
		57371:   15,  // FLOATCONST (130x)
		57381:   16,  // INTCONST (130x)
		57384:   17,  // LONGCHARCONST (130x)
		57385:   18,  // LONGSTRINGLITERAL (130x)
		57419:   19,  // SIZEOF (130x)
		57421:   20,  // STRINGLITERAL (130x)
		57358:   21,  // CONST (111x)
		57413:   22,  // RESTRICT (111x)
		57431:   23,  // VOLATILE (111x)
		125:     24,  // '}' (109x)
		58:      25,  // ':' (102x)
		57351:   26,  // BOOL (101x)
		57355:   27,  // CHAR (101x)
		57357:   28,  // COMPLEX (101x)
		57365:   29,  // DOUBLE (101x)
		57367:   30,  // ENUM (101x)
		57370:   31,  // FLOAT (101x)
		57380:   32,  // INT (101x)
		57383:   33,  // LONG (101x)
		57417:   34,  // SHORT (101x)
		57418:   35,  // SIGNED (101x)
		57422:   36,  // STRUCT (101x)
		57426:   37,  // TYPEDEFNAME (101x)
		57428:   38,  // UNION (101x)
		57429:   39,  // UNSIGNED (101x)
		57430:   40,  // VOID (101x)
		57420:   41,  // STATIC (96x)
		57344:   42,  // $end (95x)
		57350:   43,  // AUTO (90x)
		57369:   44,  // EXTERN (90x)
		57379:   45,  // INLINE (90x)
		57412:   46,  // REGISTER (90x)
		57425:   47,  // TYPEDEF (90x)
		61:      48,  // '=' (87x)
		57480:   49,  // Expression (87x)
		93:      50,  // ']' (81x)
		46:      51,  // '.' (76x)
		123:     52,  // '{' (75x)
		37:      53,  // '%' (68x)
		47:      54,  // '/' (68x)
		60:      55,  // '<' (68x)
		62:      56,  // '>' (68x)
		63:      57,  // '?' (68x)
		94:      58,  // '^' (68x)
		124:     59,  // '|' (68x)
		57346:   60,  // ADDASSIGN (68x)
		57347:   61,  // ANDAND (68x)
		57348:   62,  // ANDASSIGN (68x)
		57349:   63,  // ARROW (68x)
		57363:   64,  // DIVASSIGN (68x)
		57368:   65,  // EQ (68x)
		57373:   66,  // GEQ (68x)
		57382:   67,  // LEQ (68x)
		57386:   68,  // LSH (68x)
		57387:   69,  // LSHASSIGN (68x)
		57388:   70,  // MODASSIGN (68x)
		57389:   71,  // MULASSIGN (68x)
		57390:   72,  // NEQ (68x)
		57392:   73,  // ORASSIGN (68x)
		57393:   74,  // OROR (68x)
		57415:   75,  // RSH (68x)
		57416:   76,  // RSHASSIGN (68x)
		57423:   77,  // SUBASSIGN (68x)
		57433:   78,  // XORASSIGN (68x)
		10:      79,  // '\n' (57x)
		57408:   80,  // PPOTHER (51x)
		57397:   81,  // PPENDIF (44x)
		57432:   82,  // WHILE (41x)
		57352:   83,  // BREAK (40x)
		57353:   84,  // CASE (40x)
		57359:   85,  // CONTINUE (40x)
		57362:   86,  // DEFAULT (40x)
		57364:   87,  // DO (40x)
		57372:   88,  // FOR (40x)
		57374:   89,  // GOTO (40x)
		57377:   90,  // IF (40x)
		57396:   91,  // PPELSE (40x)
		57414:   92,  // RETURN (40x)
		57424:   93,  // SWITCH (40x)
		57395:   94,  // PPELIF (39x)
		57394:   95,  // PPDEFINE (34x)
		57398:   96,  // PPERROR (34x)
		57399:   97,  // PPHASH_NL (34x)
		57401:   98,  // PPIF (34x)
		57402:   99,  // PPIFDEF (34x)
		57403:   100, // PPIFNDEF (34x)
		57404:   101, // PPINCLUDE (34x)
		57405:   102, // PPLINE (34x)
		57406:   103, // PPNONDIRECTIVE (34x)
		57410:   104, // PPPRAGMA (34x)
		57411:   105, // PPUNDEF (34x)
		57530:   106, // TypeQualifier (28x)
		57481:   107, // ExpressionList (26x)
		57366:   108, // ELSE (22x)
		57504:   109, // PPTokenList (22x)
		57506:   110, // PPTokens (22x)
		57476:   111, // EnumSpecifier (20x)
		57525:   112, // StructOrUnion (20x)
		57526:   113, // StructOrUnionSpecifier (20x)
		57533:   114, // TypeSpecifier (20x)
		57482:   115, // ExpressionListOpt (18x)
		57505:   116, // PPTokenListOpt (16x)
		57459:   117, // DeclarationSpecifiers (15x)
		57487:   118, // FunctionSpecifier (15x)
		57520:   119, // StorageClassSpecifier (15x)
		57453:   120, // CompoundStatement (13x)
		57484:   121, // ExpressionStatement (12x)
		57501:   122, // IterationStatement (12x)
		57502:   123, // JumpStatement (12x)
		57503:   124, // LabeledStatement (12x)
		57515:   125, // SelectionStatement (12x)
		57519:   126, // Statement (12x)
		57511:   127, // Pointer (11x)
		57512:   128, // PointerOpt (10x)
		57455:   129, // ControlLine (8x)
		57461:   130, // Declarator (8x)
		57490:   131, // GroupPart (8x)
		57494:   132, // IfGroup (8x)
		57495:   133, // IfSection (8x)
		57527:   134, // TextLine (8x)
		57456:   135, // Declaration (7x)
		57488:   136, // GroupList (6x)
		57514:   137, // ReplacementList (6x)
		57454:   138, // ConstantExpression (5x)
		57360:   139, // DDD (5x)
		57489:   140, // GroupListOpt (5x)
		57516:   141, // SpecifierQualifierList (5x)
		57531:   142, // TypeQualifierList (5x)
		57445:   143, // AbstractDeclarator (4x)
		57460:   144, // DeclarationSpecifiersOpt (4x)
		57465:   145, // Designator (4x)
		57507:   146, // ParameterDeclaration (4x)
		57532:   147, // TypeQualifierListOpt (4x)
		57452:   148, // CommaOpt (3x)
		57463:   149, // Designation (3x)
		57464:   150, // DesignationOpt (3x)
		57466:   151, // DesignatorList (3x)
		57483:   152, // ExpressionOpt (3x)
		57496:   153, // InitDeclarator (3x)
		57499:   154, // Initializer (3x)
		57508:   155, // ParameterList (3x)
		57509:   156, // ParameterTypeList (3x)
		57434:   157, // $@1 (2x)
		57440:   158, // $@6 (2x)
		57441:   159, // $@7 (2x)
		57446:   160, // AbstractDeclaratorOpt (2x)
		57449:   161, // BlockItem (2x)
		57462:   162, // DeclaratorOpt (2x)
		57467:   163, // DirectAbstractDeclarator (2x)
		57468:   164, // DirectAbstractDeclaratorOpt (2x)
		57469:   165, // DirectDeclarator (2x)
		57470:   166, // ElifGroup (2x)
		57477:   167, // EnumerationConstant (2x)
		57478:   168, // Enumerator (2x)
		57485:   169, // ExternalDeclaration (2x)
		57486:   170, // FunctionDefinition (2x)
		57491:   171, // IdentifierList (2x)
		57492:   172, // IdentifierListOpt (2x)
		57493:   173, // IdentifierOpt (2x)
		57497:   174, // InitDeclaratorList (2x)
		57498:   175, // InitDeclaratorListOpt (2x)
		57500:   176, // InitializerList (2x)
		57510:   177, // ParameterTypeListOpt (2x)
		57517:   178, // SpecifierQualifierListOpt (2x)
		57521:   179, // StructDeclaration (2x)
		57523:   180, // StructDeclarator (2x)
		57529:   181, // TypeName (2x)
		57435:   182, // $@10 (1x)
		57436:   183, // $@2 (1x)
		57437:   184, // $@3 (1x)
		57438:   185, // $@4 (1x)
		57439:   186, // $@5 (1x)
		57442:   187, // $@8 (1x)
		57443:   188, // $@9 (1x)
		57447:   189, // ArgumentExpressionList (1x)
		57448:   190, // ArgumentExpressionListOpt (1x)
		57450:   191, // BlockItemList (1x)
		57451:   192, // BlockItemListOpt (1x)
		1048577: 193, // CONSTANT_EXPRESSION (1x)
		57457:   194, // DeclarationList (1x)
		57458:   195, // DeclarationListOpt (1x)
		57471:   196, // ElifGroupList (1x)
		57472:   197, // ElifGroupListOpt (1x)
		57473:   198, // ElseGroup (1x)
		57474:   199, // ElseGroupOpt (1x)
		57475:   200, // EndifLine (1x)
		57479:   201, // EnumeratorList (1x)
		57376:   202, // IDENTIFIER_LPAREN (1x)
		1048576: 203, // PREPROCESSING_FILE (1x)
		57513:   204, // PreprocessingFile (1x)
		57518:   205, // Start (1x)
		57522:   206, // StructDeclarationList (1x)
		57524:   207, // StructDeclaratorList (1x)
		1048578: 208, // TRANSLATION_UNIT (1x)
		57528:   209, // TranslationUnit (1x)
		57444:   210, // $default (0x)
		57354:   211, // CAST (0x)
		57345:   212, // error (0x)
		57391:   213, // NOELSE (0x)
		57400:   214, // PPHEADER_NAME (0x)
		57407:   215, // PPNUMBER (0x)
		57409:   216, // PPPASTE (0x)
		57427:   217, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'*'",
		"IDENTIFIER",
		"'&'",
		"'+'",
		"'-'",
		"DEC",
		"INC",
		"')'",
		"';'",
		"','",
		"'['",
		"'!'",
		"'~'",
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"STRINGLITERAL",
		"CONST",
		"RESTRICT",
		"VOLATILE",
		"'}'",
		"':'",
		"BOOL",
		"CHAR",
		"COMPLEX",
		"DOUBLE",
		"ENUM",
		"FLOAT",
		"INT",
		"LONG",
		"SHORT",
		"SIGNED",
		"STRUCT",
		"TYPEDEFNAME",
		"UNION",
		"UNSIGNED",
		"VOID",
		"STATIC",
		"$end",
		"AUTO",
		"EXTERN",
		"INLINE",
		"REGISTER",
		"TYPEDEF",
		"'='",
		"Expression",
		"']'",
		"'.'",
		"'{'",
		"'%'",
		"'/'",
		"'<'",
		"'>'",
		"'?'",
		"'^'",
		"'|'",
		"ADDASSIGN",
		"ANDAND",
		"ANDASSIGN",
		"ARROW",
		"DIVASSIGN",
		"EQ",
		"GEQ",
		"LEQ",
		"LSH",
		"LSHASSIGN",
		"MODASSIGN",
		"MULASSIGN",
		"NEQ",
		"ORASSIGN",
		"OROR",
		"RSH",
		"RSHASSIGN",
		"SUBASSIGN",
		"XORASSIGN",
		"'\\n'",
		"PPOTHER",
		"PPENDIF",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"GOTO",
		"IF",
		"PPELSE",
		"RETURN",
		"SWITCH",
		"PPELIF",
		"PPDEFINE",
		"PPERROR",
		"PPHASH_NL",
		"PPIF",
		"PPIFDEF",
		"PPIFNDEF",
		"PPINCLUDE",
		"PPLINE",
		"PPNONDIRECTIVE",
		"PPPRAGMA",
		"PPUNDEF",
		"TypeQualifier",
		"ExpressionList",
		"ELSE",
		"PPTokenList",
		"PPTokens",
		"EnumSpecifier",
		"StructOrUnion",
		"StructOrUnionSpecifier",
		"TypeSpecifier",
		"ExpressionListOpt",
		"PPTokenListOpt",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"StorageClassSpecifier",
		"CompoundStatement",
		"ExpressionStatement",
		"IterationStatement",
		"JumpStatement",
		"LabeledStatement",
		"SelectionStatement",
		"Statement",
		"Pointer",
		"PointerOpt",
		"ControlLine",
		"Declarator",
		"GroupPart",
		"IfGroup",
		"IfSection",
		"TextLine",
		"Declaration",
		"GroupList",
		"ReplacementList",
		"ConstantExpression",
		"DDD",
		"GroupListOpt",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"AbstractDeclarator",
		"DeclarationSpecifiersOpt",
		"Designator",
		"ParameterDeclaration",
		"TypeQualifierListOpt",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"InitDeclarator",
		"Initializer",
		"ParameterList",
		"ParameterTypeList",
		"$@1",
		"$@6",
		"$@7",
		"AbstractDeclaratorOpt",
		"BlockItem",
		"DeclaratorOpt",
		"DirectAbstractDeclarator",
		"DirectAbstractDeclaratorOpt",
		"DirectDeclarator",
		"ElifGroup",
		"EnumerationConstant",
		"Enumerator",
		"ExternalDeclaration",
		"FunctionDefinition",
		"IdentifierList",
		"IdentifierListOpt",
		"IdentifierOpt",
		"InitDeclaratorList",
		"InitDeclaratorListOpt",
		"InitializerList",
		"ParameterTypeListOpt",
		"SpecifierQualifierListOpt",
		"StructDeclaration",
		"StructDeclarator",
		"TypeName",
		"$@10",
		"$@2",
		"$@3",
		"$@4",
		"$@5",
		"$@8",
		"$@9",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"BlockItemList",
		"BlockItemListOpt",
		"CONSTANT_EXPRESSION",
		"DeclarationList",
		"DeclarationListOpt",
		"ElifGroupList",
		"ElifGroupListOpt",
		"ElseGroup",
		"ElseGroupOpt",
		"EndifLine",
		"EnumeratorList",
		"IDENTIFIER_LPAREN",
		"PREPROCESSING_FILE",
		"PreprocessingFile",
		"Start",
		"StructDeclarationList",
		"StructDeclaratorList",
		"TRANSLATION_UNIT",
		"TranslationUnit",
		"$default",
		"CAST",
		"error",
		"NOELSE",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
		"UNARY",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {205, 2},
		2:   {205, 2},
		3:   {205, 2},
		4:   {167, 1},
		5:   {189, 1},
		6:   {189, 3},
		7:   {190, 0},
		8:   {190, 1},
		9:   {49, 1},
		10:  {49, 1},
		11:  {49, 1},
		12:  {49, 1},
		13:  {49, 1},
		14:  {49, 1},
		15:  {49, 1},
		16:  {49, 3},
		17:  {49, 4},
		18:  {49, 4},
		19:  {49, 3},
		20:  {49, 3},
		21:  {49, 2},
		22:  {49, 2},
		23:  {49, 7},
		24:  {49, 2},
		25:  {49, 2},
		26:  {49, 2},
		27:  {49, 2},
		28:  {49, 2},
		29:  {49, 2},
		30:  {49, 2},
		31:  {49, 2},
		32:  {49, 2},
		33:  {49, 4},
		34:  {49, 4},
		35:  {49, 3},
		36:  {49, 3},
		37:  {49, 3},
		38:  {49, 3},
		39:  {49, 3},
		40:  {49, 3},
		41:  {49, 3},
		42:  {49, 3},
		43:  {49, 3},
		44:  {49, 3},
		45:  {49, 3},
		46:  {49, 3},
		47:  {49, 3},
		48:  {49, 3},
		49:  {49, 3},
		50:  {49, 3},
		51:  {49, 3},
		52:  {49, 3},
		53:  {49, 5},
		54:  {49, 3},
		55:  {49, 3},
		56:  {49, 3},
		57:  {49, 3},
		58:  {49, 3},
		59:  {49, 3},
		60:  {49, 3},
		61:  {49, 3},
		62:  {49, 3},
		63:  {49, 3},
		64:  {49, 3},
		65:  {152, 0},
		66:  {152, 1},
		67:  {107, 1},
		68:  {107, 3},
		69:  {115, 0},
		70:  {115, 1},
		71:  {138, 1},
		72:  {135, 3},
		73:  {117, 2},
		74:  {117, 2},
		75:  {117, 2},
		76:  {117, 2},
		77:  {144, 0},
		78:  {144, 1},
		79:  {174, 1},
		80:  {174, 3},
		81:  {175, 0},
		82:  {175, 1},
		83:  {153, 1},
		84:  {157, 0},
		85:  {153, 4},
		86:  {119, 1},
		87:  {119, 1},
		88:  {119, 1},
		89:  {119, 1},
		90:  {119, 1},
		91:  {114, 1},
		92:  {114, 1},
		93:  {114, 1},
		94:  {114, 1},
		95:  {114, 1},
		96:  {114, 1},
		97:  {114, 1},
		98:  {114, 1},
		99:  {114, 1},
		100: {114, 1},
		101: {114, 1},
		102: {114, 1},
		103: {114, 1},
		104: {114, 1},
		105: {183, 0},
		106: {184, 0},
		107: {113, 7},
		108: {113, 2},
		109: {112, 1},
		110: {112, 1},
		111: {206, 1},
		112: {206, 2},
		113: {179, 3},
		114: {141, 2},
		115: {141, 2},
		116: {178, 0},
		117: {178, 1},
		118: {207, 1},
		119: {207, 3},
		120: {180, 1},
		121: {180, 3},
		122: {148, 0},
		123: {148, 1},
		124: {185, 0},
		125: {111, 7},
		126: {111, 2},
		127: {201, 1},
		128: {201, 3},
		129: {168, 1},
		130: {168, 3},
		131: {106, 1},
		132: {106, 1},
		133: {106, 1},
		134: {118, 1},
		135: {130, 2},
		136: {162, 0},
		137: {162, 1},
		138: {165, 1},
		139: {165, 3},
		140: {165, 5},
		141: {165, 6},
		142: {165, 6},
		143: {165, 5},
		144: {186, 0},
		145: {165, 5},
		146: {165, 4},
		147: {127, 2},
		148: {127, 3},
		149: {128, 0},
		150: {128, 1},
		151: {142, 1},
		152: {142, 2},
		153: {147, 0},
		154: {147, 1},
		155: {156, 1},
		156: {156, 3},
		157: {177, 0},
		158: {177, 1},
		159: {155, 1},
		160: {155, 3},
		161: {146, 2},
		162: {146, 2},
		163: {171, 1},
		164: {171, 3},
		165: {172, 0},
		166: {172, 1},
		167: {173, 0},
		168: {173, 1},
		169: {158, 0},
		170: {181, 3},
		171: {143, 1},
		172: {143, 2},
		173: {160, 0},
		174: {160, 1},
		175: {163, 3},
		176: {163, 4},
		177: {163, 5},
		178: {163, 6},
		179: {163, 6},
		180: {163, 4},
		181: {159, 0},
		182: {163, 4},
		183: {187, 0},
		184: {163, 5},
		185: {164, 0},
		186: {164, 1},
		187: {154, 1},
		188: {154, 4},
		189: {176, 2},
		190: {176, 4},
		191: {149, 2},
		192: {150, 0},
		193: {150, 1},
		194: {151, 1},
		195: {151, 2},
		196: {145, 3},
		197: {145, 2},
		198: {126, 1},
		199: {126, 1},
		200: {126, 1},
		201: {126, 1},
		202: {126, 1},
		203: {126, 1},
		204: {124, 3},
		205: {124, 4},
		206: {124, 3},
		207: {188, 0},
		208: {120, 4},
		209: {191, 1},
		210: {191, 2},
		211: {192, 0},
		212: {192, 1},
		213: {161, 1},
		214: {161, 1},
		215: {121, 2},
		216: {125, 5},
		217: {125, 7},
		218: {125, 5},
		219: {122, 5},
		220: {122, 7},
		221: {122, 9},
		222: {122, 8},
		223: {123, 3},
		224: {123, 2},
		225: {123, 2},
		226: {123, 3},
		227: {209, 1},
		228: {209, 2},
		229: {169, 1},
		230: {169, 1},
		231: {182, 0},
		232: {170, 5},
		233: {194, 1},
		234: {194, 2},
		235: {195, 0},
		236: {195, 1},
		237: {204, 1},
		238: {136, 1},
		239: {136, 2},
		240: {140, 0},
		241: {140, 1},
		242: {131, 1},
		243: {131, 1},
		244: {131, 3},
		245: {131, 1},
		246: {133, 4},
		247: {132, 4},
		248: {132, 4},
		249: {132, 4},
		250: {196, 1},
		251: {196, 2},
		252: {197, 0},
		253: {197, 1},
		254: {166, 4},
		255: {198, 3},
		256: {199, 0},
		257: {199, 1},
		258: {200, 1},
		259: {129, 3},
		260: {129, 5},
		261: {129, 7},
		262: {129, 5},
		263: {129, 2},
		264: {129, 1},
		265: {129, 3},
		266: {129, 3},
		267: {129, 2},
		268: {129, 3},
		269: {129, 6},
		270: {129, 8},
		271: {129, 2},
		272: {129, 4},
		273: {134, 1},
		274: {137, 1},
		275: {109, 1},
		276: {116, 1},
		277: {116, 2},
		278: {110, 1},
		279: {110, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 42}:   "invalid empty input",
		yyXError{483, -1}: "expected #endif",
		yyXError{485, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{402, -1}: "expected $end",
		yyXError{403, -1}: "expected $end",
		yyXError{336, -1}: "expected '('",
		yyXError{337, -1}: "expected '('",
		yyXError{338, -1}: "expected '('",
		yyXError{340, -1}: "expected '('",
		yyXError{366, -1}: "expected '('",
		yyXError{146, -1}: "expected ')'",
		yyXError{153, -1}: "expected ')'",
		yyXError{160, -1}: "expected ')'",
		yyXError{186, -1}: "expected ')'",
		yyXError{189, -1}: "expected ')'",
		yyXError{192, -1}: "expected ')'",
		yyXError{200, -1}: "expected ')'",
		yyXError{205, -1}: "expected ')'",
		yyXError{211, -1}: "expected ')'",
		yyXError{227, -1}: "expected ')'",
		yyXError{232, -1}: "expected ')'",
		yyXError{273, -1}: "expected ')'",
		yyXError{356, -1}: "expected ')'",
		yyXError{362, -1}: "expected ')'",
		yyXError{443, -1}: "expected ')'",
		yyXError{444, -1}: "expected ')'",
		yyXError{452, -1}: "expected ')'",
		yyXError{455, -1}: "expected ')'",
		yyXError{458, -1}: "expected ')'",
		yyXError{290, -1}: "expected ':'",
		yyXError{329, -1}: "expected ':'",
		yyXError{390, -1}: "expected ':'",
		yyXError{306, -1}: "expected ';'",
		yyXError{335, -1}: "expected ';'",
		yyXError{342, -1}: "expected ';'",
		yyXError{343, -1}: "expected ';'",
		yyXError{345, -1}: "expected ';'",
		yyXError{349, -1}: "expected ';'",
		yyXError{352, -1}: "expected ';'",
		yyXError{354, -1}: "expected ';'",
		yyXError{360, -1}: "expected ';'",
		yyXError{369, -1}: "expected ';'",
		yyXError{311, -1}: "expected '='",
		yyXError{165, -1}: "expected '['",
		yyXError{424, -1}: "expected '\\n'",
		yyXError{430, -1}: "expected '\\n'",
		yyXError{433, -1}: "expected '\\n'",
		yyXError{435, -1}: "expected '\\n'",
		yyXError{462, -1}: "expected '\\n'",
		yyXError{467, -1}: "expected '\\n'",
		yyXError{470, -1}: "expected '\\n'",
		yyXError{477, -1}: "expected '\\n'",
		yyXError{482, -1}: "expected '\\n'",
		yyXError{488, -1}: "expected '\\n'",
		yyXError{171, -1}: "expected ']'",
		yyXError{179, -1}: "expected ']'",
		yyXError{223, -1}: "expected ']'",
		yyXError{250, -1}: "expected ']'",
		yyXError{41, -1}:  "expected '{'",
		yyXError{43, -1}:  "expected '{'",
		yyXError{279, -1}: "expected '{'",
		yyXError{281, -1}: "expected '{'",
		yyXError{259, -1}: "expected '}'",
		yyXError{263, -1}: "expected '}'",
		yyXError{276, -1}: "expected '}'",
		yyXError{330, -1}: "expected '}'",
		yyXError{46, -1}:  "expected CommaOpt or one of [',', '}']",
		yyXError{242, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{257, -1}: "expected CommaOpt or one of [',', '}']",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{199, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{164, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{332, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{312, -1}: "expected compound statement or '{'",
		yyXError{316, -1}: "expected compound statement or '{'",
		yyXError{309, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{49, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{247, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{294, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{328, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{314, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{351, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{293, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{191, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{6, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{244, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{194, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{161, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{291, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{475, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{481, -1}: "expected endif line or #endif",
		yyXError{410, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{473, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{44, -1}:  "expected enumerator list or identifier",
		yyXError{275, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{71, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{95, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{367, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{371, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{375, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{379, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{57, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{69, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{239, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{168, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{222, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{274, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{59, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{60, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{61, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{62, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{77, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{78, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{79, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{80, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{81, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{82, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{83, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{84, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{85, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{86, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{87, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{88, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{89, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{90, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{91, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{92, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{93, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{94, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{96, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{97, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{98, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{99, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{100, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{101, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{102, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{103, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{104, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{105, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{106, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{120, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{121, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{148, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{174, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{180, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{216, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{219, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{172, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{214, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{5, -1}:   "expected external declaration or one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{464, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{404, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{73, -1}:  "expected identifier",
		yyXError{74, -1}:  "expected identifier",
		yyXError{208, -1}: "expected identifier",
		yyXError{248, -1}: "expected identifier",
		yyXError{341, -1}: "expected identifier",
		yyXError{412, -1}: "expected identifier",
		yyXError{413, -1}: "expected identifier",
		yyXError{420, -1}: "expected identifier",
		yyXError{439, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{398, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{241, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{255, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{243, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{261, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{395, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{258, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{50, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{51, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{52, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{53, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{54, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{55, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{70, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{75, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{76, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{107, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{108, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{109, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{110, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{111, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{112, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{113, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{114, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{115, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{116, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{117, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{123, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{124, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{125, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{126, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{127, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{128, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{129, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{130, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{131, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{132, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{133, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{134, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{135, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{136, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{137, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{138, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{139, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{140, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{141, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{142, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{143, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{147, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{151, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{184, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{240, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{264, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{265, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{266, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{267, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{268, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{269, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{270, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{271, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{272, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{118, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{122, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{144, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{149, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{320, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{254, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{167, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{175, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{181, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{217, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{220, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{405, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{406, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{407, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{409, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{416, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{421, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{423, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{426, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{429, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{431, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{432, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{434, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{436, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{437, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{440, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{446, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{447, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{449, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{454, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{457, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{460, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{461, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{466, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{486, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{487, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{489, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{465, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{469, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{472, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{474, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{479, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{480, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{387, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{400, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{38, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{318, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{401, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{34, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{35, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{169, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{177, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{322, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{323, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{324, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{325, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{326, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{327, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{346, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{347, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{348, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{350, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{358, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{364, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{370, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{374, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{378, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{382, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{384, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{385, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{392, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{394, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{331, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{333, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{334, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{386, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{245, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{252, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{42, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{280, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{277, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{299, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{301, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{302, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{303, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{304, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{305, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{236, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{237, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', '*', ':', '[', identifier]",
		yyXError{197, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{198, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{201, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{210, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{212, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{218, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{221, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{224, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{225, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{159, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{235, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{163, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{176, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{178, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{182, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{183, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{185, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{193, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{229, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{233, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{292, -1}: "expected one of ['(', identifier]",
		yyXError{321, -1}: "expected one of [')', ',', ';']",
		yyXError{441, -1}: "expected one of [')', ',', ...]",
		yyXError{451, -1}: "expected one of [')', ',', ...]",
		yyXError{145, -1}: "expected one of [')', ',']",
		yyXError{152, -1}: "expected one of [')', ',']",
		yyXError{162, -1}: "expected one of [')', ',']",
		yyXError{188, -1}: "expected one of [')', ',']",
		yyXError{190, -1}: "expected one of [')', ',']",
		yyXError{195, -1}: "expected one of [')', ',']",
		yyXError{196, -1}: "expected one of [')', ',']",
		yyXError{206, -1}: "expected one of [')', ',']",
		yyXError{207, -1}: "expected one of [')', ',']",
		yyXError{209, -1}: "expected one of [')', ',']",
		yyXError{228, -1}: "expected one of [')', ',']",
		yyXError{368, -1}: "expected one of [')', ',']",
		yyXError{372, -1}: "expected one of [')', ',']",
		yyXError{376, -1}: "expected one of [')', ',']",
		yyXError{380, -1}: "expected one of [')', ',']",
		yyXError{442, -1}: "expected one of [')', ',']",
		yyXError{289, -1}: "expected one of [',', ':', ';']",
		yyXError{119, -1}: "expected one of [',', ':']",
		yyXError{397, -1}: "expected one of [',', ';', '=']",
		yyXError{260, -1}: "expected one of [',', ';', '}']",
		yyXError{287, -1}: "expected one of [',', ';']",
		yyXError{288, -1}: "expected one of [',', ';']",
		yyXError{295, -1}: "expected one of [',', ';']",
		yyXError{298, -1}: "expected one of [',', ';']",
		yyXError{307, -1}: "expected one of [',', ';']",
		yyXError{308, -1}: "expected one of [',', ';']",
		yyXError{396, -1}: "expected one of [',', ';']",
		yyXError{399, -1}: "expected one of [',', ';']",
		yyXError{45, -1}:  "expected one of [',', '=', '}']",
		yyXError{48, -1}:  "expected one of [',', '=', '}']",
		yyXError{150, -1}: "expected one of [',', ']']",
		yyXError{47, -1}:  "expected one of [',', '}']",
		yyXError{68, -1}:  "expected one of [',', '}']",
		yyXError{256, -1}: "expected one of [',', '}']",
		yyXError{262, -1}: "expected one of [',', '}']",
		yyXError{278, -1}: "expected one of [',', '}']",
		yyXError{246, -1}: "expected one of ['.', '=', '[']",
		yyXError{249, -1}: "expected one of ['.', '=', '[']",
		yyXError{251, -1}: "expected one of ['.', '=', '[']",
		yyXError{253, -1}: "expected one of ['.', '=', '[']",
		yyXError{414, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{422, -1}: "expected one of ['\\n', ppother]",
		yyXError{425, -1}: "expected one of ['\\n', ppother]",
		yyXError{427, -1}: "expected one of ['\\n', ppother]",
		yyXError{313, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{315, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{31, -1}:  "expected one of ['{', identifier]",
		yyXError{32, -1}:  "expected one of ['{', identifier]",
		yyXError{285, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{296, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{300, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{450, -1}: "expected one of [..., identifier]",
		yyXError{157, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{72, -1}:  "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{317, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{319, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{7, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{8, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{355, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{361, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{344, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{353, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{359, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{213, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{202, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{166, -1}: "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{170, -1}: "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{463, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{468, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{471, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{478, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{484, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{203, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{30, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{310, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{187, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{230, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{231, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{155, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{156, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{415, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{419, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{158, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{226, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{204, -1}: "expected parameter type list or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{234, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{438, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{445, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{448, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{453, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{456, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{459, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{154, -1}: "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{339, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{357, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{363, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{373, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{377, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{381, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{383, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{388, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{391, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{393, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{282, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{283, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{284, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{286, -1}: "expected struct declarator list or one of ['(', '*', ':', identifier]",
		yyXError{297, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{428, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{408, -1}: "expected token list or ppother",
		yyXError{411, -1}: "expected token list or ppother",
		yyXError{417, -1}: "expected token list or ppother",
		yyXError{418, -1}: "expected token list or ppother",
		yyXError{476, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{173, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{215, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{365, -1}: "expected while",
		yyXError{3, 42}:   "unexpected EOF",
		yyXError{2, 42}:   "unexpected EOF",
		yyXError{4, 42}:   "unexpected EOF",
	}

	yyParseTab = [490][]uint16{
		// 0
		{193: 283, 203: 282, 205: 281, 208: 284},
		{42: 280},
		{79: 703, 705, 95: 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 685, 690, 687, 689, 136: 684, 204: 683},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 338, 138: 682},
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 286, 290, 287, 135: 320, 169: 318, 319, 209: 285},
		// 5
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 277, 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 286, 290, 287, 135: 320, 169: 681, 319},
		{131, 438, 131, 9: 199, 127: 572, 571, 130: 589, 153: 587, 174: 588, 586},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 582, 290, 287, 144: 585},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 582, 290, 287, 144: 584},
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 582, 290, 287, 144: 583},
		// 10
		{203, 203, 203, 8: 203, 203, 203, 203, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 582, 290, 287, 144: 581},
		{194, 194, 194, 8: 194, 194, 194, 194, 21: 194, 194, 194, 26: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 43: 194, 194, 194, 194, 194},
		{193, 193, 193, 8: 193, 193, 193, 193, 21: 193, 193, 193, 26: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 43: 193, 193, 193, 193, 193},
		{192, 192, 192, 8: 192, 192, 192, 192, 21: 192, 192, 192, 26: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 43: 192, 192, 192, 192, 192},
		{191, 191, 191, 8: 191, 191, 191, 191, 21: 191, 191, 191, 26: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 43: 191, 191, 191, 191, 191},
		// 15
		{190, 190, 190, 8: 190, 190, 190, 190, 21: 190, 190, 190, 26: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 43: 190, 190, 190, 190, 190},
		{189, 189, 189, 8: 189, 189, 189, 189, 21: 189, 189, 189, 25: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 43: 189, 189, 189, 189, 189},
		{188, 188, 188, 8: 188, 188, 188, 188, 21: 188, 188, 188, 25: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 43: 188, 188, 188, 188, 188},
		{187, 187, 187, 8: 187, 187, 187, 187, 21: 187, 187, 187, 25: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 43: 187, 187, 187, 187, 187},
		{186, 186, 186, 8: 186, 186, 186, 186, 21: 186, 186, 186, 25: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 43: 186, 186, 186, 186, 186},
		// 20
		{185, 185, 185, 8: 185, 185, 185, 185, 21: 185, 185, 185, 25: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 43: 185, 185, 185, 185, 185},
		{184, 184, 184, 8: 184, 184, 184, 184, 21: 184, 184, 184, 25: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 43: 184, 184, 184, 184, 184},
		{183, 183, 183, 8: 183, 183, 183, 183, 21: 183, 183, 183, 25: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 43: 183, 183, 183, 183, 183},
		{182, 182, 182, 8: 182, 182, 182, 182, 21: 182, 182, 182, 25: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 43: 182, 182, 182, 182, 182},
		{181, 181, 181, 8: 181, 181, 181, 181, 21: 181, 181, 181, 25: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 43: 181, 181, 181, 181, 181},
		// 25
		{180, 180, 180, 8: 180, 180, 180, 180, 21: 180, 180, 180, 25: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 43: 180, 180, 180, 180, 180},
		{179, 179, 179, 8: 179, 179, 179, 179, 21: 179, 179, 179, 25: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 43: 179, 179, 179, 179, 179},
		{178, 178, 178, 8: 178, 178, 178, 178, 21: 178, 178, 178, 25: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 43: 178, 178, 178, 178, 178},
		{177, 177, 177, 8: 177, 177, 177, 177, 21: 177, 177, 177, 25: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 43: 177, 177, 177, 177, 177},
		{176, 176, 176, 8: 176, 176, 176, 176, 21: 176, 176, 176, 25: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 43: 176, 176, 176, 176, 176},
		// 30
		{2: 560, 52: 113, 173: 559},
		{2: 171, 52: 171},
		{2: 170, 52: 170},
		{2: 322, 52: 113, 173: 321},
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 25: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 43: 149, 149, 149, 149, 149, 50: 149},
		// 35
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 25: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 43: 148, 148, 148, 148, 148, 50: 148},
		{147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 25: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 43: 147, 147, 147, 147, 147, 50: 147},
		{146, 146, 146, 8: 146, 146, 146, 146, 21: 146, 146, 146, 26: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 43: 146, 146, 146, 146, 146},
		{21: 53, 53, 53, 26: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		{21: 51, 51, 51, 26: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
		// 40
		{21: 50, 50, 50, 26: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		{52: 156, 185: 323},
		{154, 154, 154, 8: 154, 154, 154, 154, 21: 154, 154, 154, 25: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 43: 154, 154, 154, 154, 154, 52: 112},
		{52: 324},
		{2: 325, 167: 328, 327, 201: 326},
		// 45
		{10: 276, 24: 276, 48: 276},
		{10: 555, 24: 158, 148: 556},
		{10: 153, 24: 153},
		{10: 151, 24: 151, 48: 329},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 338, 138: 348},
		// 50
		{271, 271, 3: 271, 271, 271, 271, 271, 271, 271, 271, 271, 24: 271, 271, 42: 271, 48: 271, 50: 271, 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{270, 270, 3: 270, 270, 270, 270, 270, 270, 270, 270, 270, 24: 270, 270, 42: 270, 48: 270, 50: 270, 270, 53: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{269, 269, 3: 269, 269, 269, 269, 269, 269, 269, 269, 269, 24: 269, 269, 42: 269, 48: 269, 50: 269, 269, 53: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{268, 268, 3: 268, 268, 268, 268, 268, 268, 268, 268, 268, 24: 268, 268, 42: 268, 48: 268, 50: 268, 268, 53: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		{267, 267, 3: 267, 267, 267, 267, 267, 267, 267, 267, 267, 24: 267, 267, 42: 267, 48: 267, 50: 267, 267, 53: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		// 55
		{266, 266, 3: 266, 266, 266, 266, 266, 266, 266, 266, 266, 24: 266, 266, 42: 266, 48: 266, 50: 266, 266, 53: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{265, 265, 3: 265, 265, 265, 265, 265, 265, 265, 265, 265, 24: 265, 265, 42: 265, 48: 265, 50: 265, 265, 53: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 111, 111, 111, 26: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 398, 107: 432, 158: 434, 181: 553},
		{352, 357, 3: 370, 360, 361, 356, 355, 9: 209, 209, 351, 24: 209, 209, 42: 209, 48: 376, 50: 209, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 552},
		// 60
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 551},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 550},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 464},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 549},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 548},
		// 65
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 547},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 546},
		{349, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 350},
		{10: 150, 24: 150},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 111, 111, 111, 26: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 49: 398, 107: 432, 158: 434, 181: 433},
		// 70
		{352, 248, 3: 248, 248, 248, 356, 355, 248, 248, 248, 351, 24: 248, 248, 42: 248, 48: 248, 50: 248, 353, 53: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 354, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 430},
		{337, 342, 330, 341, 343, 344, 340, 339, 273, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 424, 189: 425, 426},
		{2: 423},
		{2: 422},
		// 75
		{259, 259, 3: 259, 259, 259, 259, 259, 259, 259, 259, 259, 24: 259, 259, 42: 259, 48: 259, 50: 259, 259, 53: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{258, 258, 3: 258, 258, 258, 258, 258, 258, 258, 258, 258, 24: 258, 258, 42: 258, 48: 258, 50: 258, 258, 53: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 421},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 420},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 419},
		// 80
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 418},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 417},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 416},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 415},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 414},
		// 85
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 413},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 412},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 411},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 410},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 409},
		// 90
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 408},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 407},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 406},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 405},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 404},
		// 95
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 399},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 397},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 396},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 395},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 394},
		// 100
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 393},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 392},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 391},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 390},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 389},
		// 105
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 388},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 387},
		{352, 357, 3: 370, 360, 361, 356, 355, 216, 216, 216, 351, 24: 216, 216, 42: 216, 48: 376, 50: 216, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 217, 217, 217, 351, 24: 217, 217, 42: 217, 48: 376, 50: 217, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 218, 218, 218, 351, 24: 218, 218, 42: 218, 48: 376, 50: 218, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		// 110
		{352, 357, 3: 370, 360, 361, 356, 355, 219, 219, 219, 351, 24: 219, 219, 42: 219, 48: 376, 50: 219, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 220, 220, 220, 351, 24: 220, 220, 42: 220, 48: 376, 50: 220, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 221, 221, 221, 351, 24: 221, 221, 42: 221, 48: 376, 50: 221, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 222, 222, 222, 351, 24: 222, 222, 42: 222, 48: 376, 50: 222, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 223, 223, 223, 351, 24: 223, 223, 42: 223, 48: 376, 50: 223, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		// 115
		{352, 357, 3: 370, 360, 361, 356, 355, 224, 224, 224, 351, 24: 224, 224, 42: 224, 48: 376, 50: 224, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 225, 225, 225, 351, 24: 225, 225, 42: 225, 48: 376, 50: 225, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 226, 226, 226, 351, 24: 226, 226, 42: 226, 48: 376, 50: 226, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 213, 213, 213, 351, 25: 213, 48: 376, 50: 213, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{10: 401, 25: 400},
		// 120
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 403},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 402},
		{352, 357, 3: 370, 360, 361, 356, 355, 212, 212, 212, 351, 25: 212, 48: 376, 50: 212, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{352, 357, 3: 370, 360, 361, 356, 355, 227, 227, 227, 351, 24: 227, 227, 42: 227, 48: 227, 50: 227, 353, 53: 359, 358, 364, 365, 227, 371, 372, 227, 373, 227, 354, 227, 368, 367, 366, 362, 227, 227, 227, 369, 227, 374, 363, 227, 227, 227},
		{352, 357, 3: 370, 360, 361, 356, 355, 228, 228, 228, 351, 24: 228, 228, 42: 228, 48: 228, 50: 228, 353, 53: 359, 358, 364, 365, 228, 371, 372, 228, 373, 228, 354, 228, 368, 367, 366, 362, 228, 228, 228, 369, 228, 228, 363, 228, 228, 228},
		// 125
		{352, 357, 3: 370, 360, 361, 356, 355, 229, 229, 229, 351, 24: 229, 229, 42: 229, 48: 229, 50: 229, 353, 53: 359, 358, 364, 365, 229, 371, 372, 229, 229, 229, 354, 229, 368, 367, 366, 362, 229, 229, 229, 369, 229, 229, 363, 229, 229, 229},
		{352, 357, 3: 370, 360, 361, 356, 355, 230, 230, 230, 351, 24: 230, 230, 42: 230, 48: 230, 50: 230, 353, 53: 359, 358, 364, 365, 230, 371, 230, 230, 230, 230, 354, 230, 368, 367, 366, 362, 230, 230, 230, 369, 230, 230, 363, 230, 230, 230},
		{352, 357, 3: 370, 360, 361, 356, 355, 231, 231, 231, 351, 24: 231, 231, 42: 231, 48: 231, 50: 231, 353, 53: 359, 358, 364, 365, 231, 231, 231, 231, 231, 231, 354, 231, 368, 367, 366, 362, 231, 231, 231, 369, 231, 231, 363, 231, 231, 231},
		{352, 357, 3: 232, 360, 361, 356, 355, 232, 232, 232, 351, 24: 232, 232, 42: 232, 48: 232, 50: 232, 353, 53: 359, 358, 364, 365, 232, 232, 232, 232, 232, 232, 354, 232, 368, 367, 366, 362, 232, 232, 232, 369, 232, 232, 363, 232, 232, 232},
		{352, 357, 3: 233, 360, 361, 356, 355, 233, 233, 233, 351, 24: 233, 233, 42: 233, 48: 233, 50: 233, 353, 53: 359, 358, 364, 365, 233, 233, 233, 233, 233, 233, 354, 233, 233, 367, 366, 362, 233, 233, 233, 233, 233, 233, 363, 233, 233, 233},
		// 130
		{352, 357, 3: 234, 360, 361, 356, 355, 234, 234, 234, 351, 24: 234, 234, 42: 234, 48: 234, 50: 234, 353, 53: 359, 358, 364, 365, 234, 234, 234, 234, 234, 234, 354, 234, 234, 367, 366, 362, 234, 234, 234, 234, 234, 234, 363, 234, 234, 234},
		{352, 357, 3: 235, 360, 361, 356, 355, 235, 235, 235, 351, 24: 235, 235, 42: 235, 48: 235, 50: 235, 353, 53: 359, 358, 235, 235, 235, 235, 235, 235, 235, 235, 354, 235, 235, 235, 235, 362, 235, 235, 235, 235, 235, 235, 363, 235, 235, 235},
		{352, 357, 3: 236, 360, 361, 356, 355, 236, 236, 236, 351, 24: 236, 236, 42: 236, 48: 236, 50: 236, 353, 53: 359, 358, 236, 236, 236, 236, 236, 236, 236, 236, 354, 236, 236, 236, 236, 362, 236, 236, 236, 236, 236, 236, 363, 236, 236, 236},
		{352, 357, 3: 237, 360, 361, 356, 355, 237, 237, 237, 351, 24: 237, 237, 42: 237, 48: 237, 50: 237, 353, 53: 359, 358, 237, 237, 237, 237, 237, 237, 237, 237, 354, 237, 237, 237, 237, 362, 237, 237, 237, 237, 237, 237, 363, 237, 237, 237},
		{352, 357, 3: 238, 360, 361, 356, 355, 238, 238, 238, 351, 24: 238, 238, 42: 238, 48: 238, 50: 238, 353, 53: 359, 358, 238, 238, 238, 238, 238, 238, 238, 238, 354, 238, 238, 238, 238, 362, 238, 238, 238, 238, 238, 238, 363, 238, 238, 238},
		// 135
		{352, 357, 3: 239, 360, 361, 356, 355, 239, 239, 239, 351, 24: 239, 239, 42: 239, 48: 239, 50: 239, 353, 53: 359, 358, 239, 239, 239, 239, 239, 239, 239, 239, 354, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239, 239},
		{352, 357, 3: 240, 360, 361, 356, 355, 240, 240, 240, 351, 24: 240, 240, 42: 240, 48: 240, 50: 240, 353, 53: 359, 358, 240, 240, 240, 240, 240, 240, 240, 240, 354, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240},
		{352, 357, 3: 241, 241, 241, 356, 355, 241, 241, 241, 351, 24: 241, 241, 42: 241, 48: 241, 50: 241, 353, 53: 359, 358, 241, 241, 241, 241, 241, 241, 241, 241, 354, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241},
		{352, 357, 3: 242, 242, 242, 356, 355, 242, 242, 242, 351, 24: 242, 242, 42: 242, 48: 242, 50: 242, 353, 53: 359, 358, 242, 242, 242, 242, 242, 242, 242, 242, 354, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242},
		{352, 243, 3: 243, 243, 243, 356, 355, 243, 243, 243, 351, 24: 243, 243, 42: 243, 48: 243, 50: 243, 353, 53: 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 354, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		// 140
		{352, 244, 3: 244, 244, 244, 356, 355, 244, 244, 244, 351, 24: 244, 244, 42: 244, 48: 244, 50: 244, 353, 53: 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 354, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244, 244},
		{352, 245, 3: 245, 245, 245, 356, 355, 245, 245, 245, 351, 24: 245, 245, 42: 245, 48: 245, 50: 245, 353, 53: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 354, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		{260, 260, 3: 260, 260, 260, 260, 260, 260, 260, 260, 260, 24: 260, 260, 42: 260, 48: 260, 50: 260, 260, 53: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{261, 261, 3: 261, 261, 261, 261, 261, 261, 261, 261, 261, 24: 261, 261, 42: 261, 48: 261, 50: 261, 261, 53: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		{352, 357, 3: 370, 360, 361, 356, 355, 275, 10: 275, 351, 48: 376, 51: 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		// 145
		{8: 272, 10: 428},
		{8: 427},
		{262, 262, 3: 262, 262, 262, 262, 262, 262, 262, 262, 262, 24: 262, 262, 42: 262, 48: 262, 50: 262, 262, 53: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 429},
		{352, 357, 3: 370, 360, 361, 356, 355, 274, 10: 274, 351, 48: 376, 51: 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		// 150
		{10: 401, 50: 431},
		{263, 263, 3: 263, 263, 263, 263, 263, 263, 263, 263, 263, 24: 263, 263, 42: 263, 48: 263, 50: 263, 263, 53: 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263, 263},
		{8: 545, 10: 401},
		{8: 519},
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 106: 436, 111: 308, 310, 307, 435, 141: 437},
		// 155
		{164, 164, 164, 8: 164, 11: 164, 21: 314, 315, 316, 25: 164, 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 106: 436, 111: 308, 310, 307, 435, 141: 517, 178: 518},
		{164, 164, 164, 8: 164, 11: 164, 21: 314, 315, 316, 25: 164, 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 106: 436, 111: 308, 310, 307, 435, 141: 517, 178: 516},
		{131, 438, 8: 107, 11: 131, 127: 439, 441, 143: 442, 160: 440},
		{127, 127, 127, 8: 127, 10: 127, 127, 21: 314, 315, 316, 106: 449, 142: 453, 147: 514},
		{130, 2: 130, 8: 109, 10: 109, 130},
		// 160
		{8: 110},
		{444, 11: 95, 163: 443, 445},
		{8: 106, 10: 106},
		{510, 8: 108, 10: 108, 94},
		{131, 438, 8: 99, 11: 131, 21: 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 127: 439, 441, 143: 466, 159: 467},
		// 165
		{11: 446},
		{337, 448, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 314, 315, 316, 41: 452, 49: 447, 215, 106: 449, 142: 450, 152: 451},
		{352, 357, 3: 370, 360, 361, 356, 355, 11: 351, 48: 376, 50: 214, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 464, 465},
		{129, 129, 129, 129, 129, 129, 129, 129, 129, 10: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 41: 129, 50: 129},
		// 170
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 314, 315, 316, 41: 460, 49: 447, 215, 106: 457, 152: 459},
		{50: 458},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 314, 315, 316, 106: 449, 142: 453, 147: 454},
		{126, 126, 126, 126, 126, 126, 126, 126, 126, 10: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 314, 315, 316, 106: 457},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 455},
		// 175
		{352, 357, 3: 370, 360, 361, 356, 355, 11: 351, 48: 376, 50: 456, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{102, 8: 102, 10: 102, 102},
		{128, 128, 128, 128, 128, 128, 128, 128, 128, 10: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 41: 128, 50: 128},
		{104, 8: 104, 10: 104, 104},
		{50: 463},
		// 180
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 461},
		{352, 357, 3: 370, 360, 361, 356, 355, 11: 351, 48: 376, 50: 462, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{101, 8: 101, 10: 101, 101},
		{103, 8: 103, 10: 103, 103},
		{352, 253, 3: 253, 253, 253, 356, 355, 253, 253, 253, 351, 24: 253, 253, 42: 253, 48: 253, 50: 253, 353, 53: 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 354, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253, 253},
		// 185
		{100, 8: 100, 10: 100, 100},
		{8: 509},
		{8: 123, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 471, 290, 287, 146: 470, 155: 468, 469, 177: 472},
		{8: 125, 10: 506},
		{8: 122},
		// 190
		{8: 121, 10: 121},
		{131, 438, 131, 8: 107, 10: 107, 131, 127: 439, 474, 130: 475, 143: 442, 160: 476},
		{8: 473},
		{98, 8: 98, 10: 98, 98},
		{479, 2: 478, 11: 95, 163: 443, 445, 477},
		// 195
		{8: 119, 10: 119},
		{8: 118, 10: 118},
		{483, 8: 145, 145, 145, 482, 21: 145, 145, 145, 25: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 43: 145, 145, 145, 145, 145, 145, 52: 145},
		{142, 8: 142, 142, 142, 142, 21: 142, 142, 142, 25: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 43: 142, 142, 142, 142, 142, 142, 52: 142},
		{131, 438, 131, 8: 99, 11: 131, 21: 99, 99, 99, 26: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 127: 439, 474, 130: 480, 143: 466, 159: 467},
		// 200
		{8: 481},
		{141, 8: 141, 141, 141, 141, 21: 141, 141, 141, 25: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 43: 141, 141, 141, 141, 141, 141, 52: 141},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 314, 315, 316, 41: 494, 50: 127, 106: 449, 142: 495, 147: 493},
		{2: 486, 8: 115, 21: 136, 136, 136, 26: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 43: 136, 136, 136, 136, 136, 171: 487, 485, 186: 484},
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 471, 290, 287, 146: 470, 155: 468, 491},
		// 205
		{8: 490},
		{8: 117, 10: 117},
		{8: 114, 10: 488},
		{2: 489},
		{8: 116, 10: 116},
		// 210
		{134, 8: 134, 134, 134, 134, 21: 134, 134, 134, 25: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 43: 134, 134, 134, 134, 134, 134, 52: 134},
		{8: 492},
		{135, 8: 135, 135, 135, 135, 21: 135, 135, 135, 25: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 43: 135, 135, 135, 135, 135, 135, 52: 135},
		{337, 502, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 447, 215, 152: 503},
		{127, 127, 127, 127, 127, 127, 127, 127, 12: 127, 127, 127, 127, 127, 127, 127, 127, 127, 314, 315, 316, 106: 449, 142: 453, 147: 499},
		// 215
		{126, 126, 126, 126, 126, 126, 126, 126, 12: 126, 126, 126, 126, 126, 126, 126, 126, 126, 314, 315, 316, 41: 496, 50: 126, 106: 457},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 497},
		{352, 357, 3: 370, 360, 361, 356, 355, 11: 351, 48: 376, 50: 498, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{138, 8: 138, 138, 138, 138, 21: 138, 138, 138, 25: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 43: 138, 138, 138, 138, 138, 138, 52: 138},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 500},
		// 220
		{352, 357, 3: 370, 360, 361, 356, 355, 11: 351, 48: 376, 50: 501, 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		{139, 8: 139, 139, 139, 139, 21: 139, 139, 139, 25: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 43: 139, 139, 139, 139, 139, 139, 52: 139},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 464, 505},
		{50: 504},
		{140, 8: 140, 140, 140, 140, 21: 140, 140, 140, 25: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 43: 140, 140, 140, 140, 140, 140, 52: 140},
		// 225
		{137, 8: 137, 137, 137, 137, 21: 137, 137, 137, 25: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 43: 137, 137, 137, 137, 137, 137, 52: 137},
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 471, 290, 287, 139: 507, 146: 508},
		{8: 124},
		{8: 120, 10: 120},
		{105, 8: 105, 10: 105, 105},
		// 230
		{8: 97, 21: 97, 97, 97, 26: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 43: 97, 97, 97, 97, 97, 187: 511},
		{8: 123, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 106: 289, 111: 308, 310, 307, 288, 117: 471, 290, 287, 146: 470, 155: 468, 469, 177: 512},
		{8: 513},
		{96, 8: 96, 10: 96, 96},
		{133, 438, 133, 8: 133, 10: 133, 133, 127: 515},
		// 235
		{132, 2: 132, 8: 132, 10: 132, 132},
		{165, 165, 165, 8: 165, 11: 165, 25: 165},
		{163, 163, 163, 8: 163, 11: 163, 25: 163},
		{166, 166, 166, 8: 166, 11: 166, 25: 166},
		{337, 247, 330, 247, 247, 247, 340, 339, 247, 247, 247, 247, 346, 345, 331, 332, 333, 334, 335, 347, 336, 24: 247, 247, 42: 247, 48: 247, 520, 247, 247, 521, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		// 240
		{352, 246, 3: 246, 246, 246, 356, 355, 246, 246, 246, 351, 24: 246, 246, 42: 246, 48: 246, 50: 246, 353, 53: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 354, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 527, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 528, 88, 145: 526, 149: 525, 523, 524, 176: 522},
		{10: 538, 24: 158, 148: 543},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 534, 52: 535, 154: 536},
		{11: 527, 48: 532, 51: 528, 145: 533},
		// 245
		{87, 87, 87, 87, 87, 87, 87, 87, 12: 87, 87, 87, 87, 87, 87, 87, 87, 87, 52: 87},
		{11: 86, 48: 86, 51: 86},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 338, 138: 530},
		{2: 529},
		{11: 83, 48: 83, 51: 83},
		// 250
		{50: 531},
		{11: 84, 48: 84, 51: 84},
		{89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 89, 89, 89, 89, 89, 89, 89, 89, 52: 89},
		{11: 85, 48: 85, 51: 85},
		{352, 357, 3: 370, 360, 361, 356, 355, 9: 93, 93, 351, 24: 93, 48: 376, 51: 353, 53: 359, 358, 364, 365, 375, 371, 372, 380, 373, 384, 354, 378, 368, 367, 366, 362, 382, 379, 377, 369, 386, 374, 363, 383, 381, 385},
		// 255
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 527, 88, 88, 88, 88, 88, 88, 88, 88, 88, 51: 528, 88, 145: 526, 149: 525, 523, 524, 176: 537},
		{10: 91, 24: 91},
		{10: 538, 24: 158, 148: 539},
		{88, 88, 88, 88, 88, 88, 88, 88, 11: 527, 88, 88, 88, 88, 88, 88, 88, 88, 88, 24: 157, 51: 528, 88, 145: 526, 149: 525, 541, 524},
		{24: 540},
		// 260
		{9: 92, 92, 24: 92},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 534, 52: 535, 154: 542},
		{10: 90, 24: 90},
		{24: 544},
		{257, 257, 3: 257, 257, 257, 257, 257, 257, 257, 257, 257, 24: 257, 257, 42: 257, 48: 257, 50: 257, 257, 53: 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257, 257},
		// 265
		{264, 264, 3: 264, 264, 264, 264, 264, 264, 264, 264, 264, 24: 264, 264, 42: 264, 48: 264, 50: 264, 264, 53: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		{352, 249, 3: 249, 249, 249, 356, 355, 249, 249, 249, 351, 24: 249, 249, 42: 249, 48: 249, 50: 249, 353, 53: 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 354, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249, 249},
		{352, 250, 3: 250, 250, 250, 356, 355, 250, 250, 250, 351, 24: 250, 250, 42: 250, 48: 250, 50: 250, 353, 53: 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 354, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{352, 251, 3: 251, 251, 251, 356, 355, 251, 251, 251, 351, 24: 251, 251, 42: 251, 48: 251, 50: 251, 353, 53: 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 354, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251, 251},
		{352, 252, 3: 252, 252, 252, 356, 355, 252, 252, 252, 351, 24: 252, 252, 42: 252, 48: 252, 50: 252, 353, 53: 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 354, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252, 252},
		// 270
		{352, 254, 3: 254, 254, 254, 356, 355, 254, 254, 254, 351, 24: 254, 254, 42: 254, 48: 254, 50: 254, 353, 53: 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 354, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254},
		{352, 255, 3: 255, 255, 255, 356, 355, 255, 255, 255, 351, 24: 255, 255, 42: 255, 48: 255, 50: 255, 353, 53: 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 354, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{352, 256, 3: 256, 256, 256, 356, 355, 256, 256, 256, 351, 24: 256, 256, 42: 256, 48: 256, 50: 256, 353, 53: 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 354, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256, 256},
		{8: 554},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 520, 52: 521},
		// 275
		{2: 325, 24: 157, 167: 328, 558},
		{24: 557},
		{155, 155, 155, 8: 155, 155, 155, 155, 21: 155, 155, 155, 25: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 155, 155, 155, 155},
		{10: 152, 24: 152},
		{52: 175, 183: 561},
		// 280
		{172, 172, 172, 8: 172, 172, 172, 172, 21: 172, 172, 172, 25: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 43: 172, 172, 172, 172, 172, 52: 112},
		{52: 562},
		{21: 174, 174, 174, 26: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 184: 563},
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 106: 436, 111: 308, 310, 307, 435, 141: 566, 179: 565, 206: 564},
		{21: 314, 315, 316, 579, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 106: 436, 111: 308, 310, 307, 435, 141: 566, 179: 580},
		// 285
		{21: 169, 169, 169, 169, 26: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169},
		{131, 438, 131, 25: 144, 127: 572, 571, 130: 569, 162: 570, 180: 568, 207: 567},
		{9: 576, 577},
		{9: 162, 162},
		{9: 160, 160, 25: 143},
		// 290
		{25: 574},
		{573, 2: 478, 165: 477},
		{130, 2: 130},
		{131, 438, 131, 127: 572, 571, 130: 480},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 338, 138: 575},
		// 295
		{9: 159, 159},
		{21: 167, 167, 167, 167, 26: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167},
		{131, 438, 131, 25: 144, 127: 572, 571, 130: 569, 162: 570, 180: 578},
		{9: 161, 161},
		{173, 173, 173, 8: 173, 173, 173, 173, 21: 173, 173, 173, 25: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173},
		// 300
		{21: 168, 168, 168, 168, 26: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168},
		{204, 204, 204, 8: 204, 204, 204, 204},
		{202, 202, 202, 8: 202, 202, 202, 202},
		{205, 205, 205, 8: 205, 205, 205, 205},
		{206, 206, 206, 8: 206, 206, 206, 206},
		// 305
		{207, 207, 207, 8: 207, 207, 207, 207},
		{9: 680},
		{9: 201, 201},
		{9: 198, 678},
		{9: 197, 197, 21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 196, 52: 45, 106: 289, 111: 308, 310, 307, 288, 117: 590, 290, 287, 135: 593, 157: 591, 194: 594, 592},
		// 310
		{131, 438, 131, 9: 199, 127: 572, 571, 130: 677, 153: 587, 174: 588, 586},
		{48: 675},
		{52: 49, 182: 596},
		{21: 47, 47, 47, 26: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 52: 47},
		{21: 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 52: 44, 106: 289, 111: 308, 310, 307, 288, 117: 590, 290, 287, 135: 595},
		// 315
		{21: 46, 46, 46, 26: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 52: 46},
		{52: 597, 120: 598},
		{73, 73, 73, 73, 73, 73, 73, 73, 9: 73, 12: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 26: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 43: 73, 73, 73, 73, 73, 52: 73, 82: 73, 73, 73, 73, 73, 73, 73, 73, 73, 92: 73, 73, 188: 599},
		{21: 48, 48, 48, 26: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 314, 315, 316, 69, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 106: 289, 601, 111: 308, 310, 307, 288, 615, 117: 590, 290, 287, 603, 604, 606, 607, 602, 605, 614, 135: 613, 161: 611, 191: 612, 610},
		// 320
		{271, 271, 3: 271, 271, 271, 271, 271, 9: 271, 271, 271, 25: 673, 48: 271, 51: 271, 53: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{8: 210, 210, 401},
		{82, 82, 82, 82, 82, 82, 82, 82, 9: 82, 12: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 26: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 52: 82, 82: 82, 82, 82, 82, 82, 82, 82, 82, 82, 92: 82, 82, 108: 82},
		{81, 81, 81, 81, 81, 81, 81, 81, 9: 81, 12: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 26: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 52: 81, 82: 81, 81, 81, 81, 81, 81, 81, 81, 81, 92: 81, 81, 108: 81},
		{80, 80, 80, 80, 80, 80, 80, 80, 9: 80, 12: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 26: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 52: 80, 82: 80, 80, 80, 80, 80, 80, 80, 80, 80, 92: 80, 80, 108: 80},
		// 325
		{79, 79, 79, 79, 79, 79, 79, 79, 9: 79, 12: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 26: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 52: 79, 82: 79, 79, 79, 79, 79, 79, 79, 79, 79, 92: 79, 79, 108: 79},
		{78, 78, 78, 78, 78, 78, 78, 78, 9: 78, 12: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 26: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 52: 78, 82: 78, 78, 78, 78, 78, 78, 78, 78, 78, 92: 78, 78, 108: 78},
		{77, 77, 77, 77, 77, 77, 77, 77, 9: 77, 12: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 26: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 52: 77, 82: 77, 77, 77, 77, 77, 77, 77, 77, 77, 92: 77, 77, 108: 77},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 338, 138: 670},
		{25: 668},
		// 330
		{24: 667},
		{71, 71, 71, 71, 71, 71, 71, 71, 9: 71, 12: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 26: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 43: 71, 71, 71, 71, 71, 52: 71, 82: 71, 71, 71, 71, 71, 71, 71, 71, 71, 92: 71, 71},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 314, 315, 316, 68, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 106: 289, 601, 111: 308, 310, 307, 288, 615, 117: 590, 290, 287, 603, 604, 606, 607, 602, 605, 614, 135: 613, 161: 666},
		{67, 67, 67, 67, 67, 67, 67, 67, 9: 67, 12: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 26: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 52: 67, 82: 67, 67, 67, 67, 67, 67, 67, 67, 67, 92: 67, 67},
		{66, 66, 66, 66, 66, 66, 66, 66, 9: 66, 12: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 26: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 52: 66, 82: 66, 66, 66, 66, 66, 66, 66, 66, 66, 92: 66, 66},
		// 335
		{9: 665},
		{659},
		{655},
		{651},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 645},
		// 340
		{631},
		{2: 629},
		{9: 628},
		{9: 627},
		{337, 342, 330, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 601, 115: 625},
		// 345
		{9: 626},
		{54, 54, 54, 54, 54, 54, 54, 54, 9: 54, 12: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 26: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 43: 54, 54, 54, 54, 54, 52: 54, 82: 54, 54, 54, 54, 54, 54, 54, 54, 54, 92: 54, 54, 108: 54},
		{55, 55, 55, 55, 55, 55, 55, 55, 9: 55, 12: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 26: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 43: 55, 55, 55, 55, 55, 52: 55, 82: 55, 55, 55, 55, 55, 55, 55, 55, 55, 92: 55, 55, 108: 55},
		{56, 56, 56, 56, 56, 56, 56, 56, 9: 56, 12: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 26: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 43: 56, 56, 56, 56, 56, 52: 56, 82: 56, 56, 56, 56, 56, 56, 56, 56, 56, 92: 56, 56, 108: 56},
		{9: 630},
		// 350
		{57, 57, 57, 57, 57, 57, 57, 57, 9: 57, 12: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 26: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 43: 57, 57, 57, 57, 57, 52: 57, 82: 57, 57, 57, 57, 57, 57, 57, 57, 57, 92: 57, 57, 108: 57},
		{337, 342, 330, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 314, 315, 316, 26: 305, 297, 306, 302, 313, 301, 299, 300, 298, 303, 311, 309, 312, 304, 296, 293, 43: 294, 292, 317, 295, 291, 49: 398, 106: 289, 601, 111: 308, 310, 307, 288, 632, 117: 590, 290, 287, 135: 633},
		{9: 639},
		{337, 342, 330, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 601, 115: 634},
		{9: 635},
		// 355
		{337, 342, 330, 341, 343, 344, 340, 339, 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 601, 115: 636},
		{8: 637},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 638},
		{58, 58, 58, 58, 58, 58, 58, 58, 9: 58, 12: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 26: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 43: 58, 58, 58, 58, 58, 52: 58, 82: 58, 58, 58, 58, 58, 58, 58, 58, 58, 92: 58, 58, 108: 58},
		{337, 342, 330, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 601, 115: 640},
		// 360
		{9: 641},
		{337, 342, 330, 341, 343, 344, 340, 339, 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 601, 115: 642},
		{8: 643},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 644},
		{59, 59, 59, 59, 59, 59, 59, 59, 9: 59, 12: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 26: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 43: 59, 59, 59, 59, 59, 52: 59, 82: 59, 59, 59, 59, 59, 59, 59, 59, 59, 92: 59, 59, 108: 59},
		// 365
		{82: 646},
		{647},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 648},
		{8: 649, 10: 401},
		{9: 650},
		// 370
		{60, 60, 60, 60, 60, 60, 60, 60, 9: 60, 12: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 26: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 43: 60, 60, 60, 60, 60, 52: 60, 82: 60, 60, 60, 60, 60, 60, 60, 60, 60, 92: 60, 60, 108: 60},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 652},
		{8: 653, 10: 401},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 654},
		{61, 61, 61, 61, 61, 61, 61, 61, 9: 61, 12: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 26: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 43: 61, 61, 61, 61, 61, 52: 61, 82: 61, 61, 61, 61, 61, 61, 61, 61, 61, 92: 61, 61, 108: 61},
		// 375
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 656},
		{8: 657, 10: 401},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 658},
		{62, 62, 62, 62, 62, 62, 62, 62, 9: 62, 12: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 26: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 43: 62, 62, 62, 62, 62, 52: 62, 82: 62, 62, 62, 62, 62, 62, 62, 62, 62, 92: 62, 62, 108: 62},
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 107: 660},
		// 380
		{8: 661, 10: 401},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 662},
		{64, 64, 64, 64, 64, 64, 64, 64, 9: 64, 12: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 26: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 43: 64, 64, 64, 64, 64, 52: 64, 82: 64, 64, 64, 64, 64, 64, 64, 64, 64, 92: 64, 64, 108: 663},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 664},
		{63, 63, 63, 63, 63, 63, 63, 63, 9: 63, 12: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 26: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 43: 63, 63, 63, 63, 63, 52: 63, 82: 63, 63, 63, 63, 63, 63, 63, 63, 63, 92: 63, 63, 108: 63},
		// 385
		{65, 65, 65, 65, 65, 65, 65, 65, 9: 65, 12: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 26: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 43: 65, 65, 65, 65, 65, 52: 65, 82: 65, 65, 65, 65, 65, 65, 65, 65, 65, 92: 65, 65, 108: 65},
		{70, 70, 70, 70, 70, 70, 70, 70, 9: 70, 12: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 26: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 43: 70, 70, 70, 70, 70, 52: 70, 82: 70, 70, 70, 70, 70, 70, 70, 70, 70, 92: 70, 70},
		{72, 72, 72, 72, 72, 72, 72, 72, 9: 72, 12: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 26: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 52: 72, 82: 72, 72, 72, 72, 72, 72, 72, 72, 72, 92: 72, 72, 108: 72},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 669},
		{74, 74, 74, 74, 74, 74, 74, 74, 9: 74, 12: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 26: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 43: 74, 74, 74, 74, 74, 52: 74, 82: 74, 74, 74, 74, 74, 74, 74, 74, 74, 92: 74, 74, 108: 74},
		// 390
		{25: 671},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 672},
		{75, 75, 75, 75, 75, 75, 75, 75, 9: 75, 12: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 26: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 52: 75, 82: 75, 75, 75, 75, 75, 75, 75, 75, 75, 92: 75, 75, 108: 75},
		{337, 342, 600, 341, 343, 344, 340, 339, 9: 211, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 398, 52: 597, 82: 618, 623, 608, 622, 609, 619, 620, 621, 616, 92: 624, 617, 107: 601, 115: 615, 120: 603, 604, 606, 607, 602, 605, 674},
		{76, 76, 76, 76, 76, 76, 76, 76, 9: 76, 12: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 26: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 52: 76, 82: 76, 76, 76, 76, 76, 76, 76, 76, 76, 92: 76, 76, 108: 76},
		// 395
		{337, 342, 330, 341, 343, 344, 340, 339, 12: 346, 345, 331, 332, 333, 334, 335, 347, 336, 49: 534, 52: 535, 154: 676},
		{9: 195, 195},
		{9: 197, 197, 48: 196, 157: 591},
		{131, 438, 131, 127: 572, 571, 130: 677, 153: 679},
		{9: 200, 200},
		// 400
		{208, 208, 208, 208, 208, 208, 208, 208, 9: 208, 12: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 26: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 52: 208, 82: 208, 208, 208, 208, 208, 208, 208, 208, 208, 92: 208, 208},
		{21: 52, 52, 52, 26: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52},
		{42: 278},
		{42: 279},
		{42: 43, 79: 703, 705, 95: 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 746, 690, 687, 689},
		// 405
		{42: 42, 79: 42, 42, 42, 91: 42, 94: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{42: 38, 79: 38, 38, 38, 91: 38, 94: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{42: 37, 79: 37, 37, 37, 91: 37, 94: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{80: 705, 109: 768, 702},
		{42: 35, 79: 35, 35, 35, 91: 35, 94: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 410
		{81: 28, 91: 28, 94: 756, 166: 754, 196: 755, 753},
		{80: 705, 109: 750, 702},
		{2: 747},
		{2: 742},
		{2: 718, 79: 720, 202: 719},
		// 415
		{79: 703, 705, 109: 704, 702, 116: 717},
		{42: 16, 79: 16, 16, 16, 91: 16, 94: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{80: 705, 109: 715, 702},
		{80: 705, 109: 713, 702},
		{79: 703, 705, 109: 704, 702, 116: 712},
		// 420
		{2: 708},
		{42: 7, 79: 7, 7, 7, 91: 7, 94: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{79: 5, 707},
		{42: 4, 79: 4, 4, 4, 91: 4, 94: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{79: 706},
		// 425
		{79: 2, 2},
		{42: 3, 79: 3, 3, 3, 91: 3, 94: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{79: 1, 1},
		{79: 709, 705, 109: 710, 702},
		{42: 12, 79: 12, 12, 12, 91: 12, 94: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		// 430
		{79: 711},
		{42: 8, 79: 8, 8, 8, 91: 8, 94: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{42: 13, 79: 13, 13, 13, 91: 13, 94: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{79: 714},
		{42: 14, 79: 14, 14, 14, 91: 14, 94: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		// 435
		{79: 716},
		{42: 15, 79: 15, 15, 15, 91: 15, 94: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{42: 17, 79: 17, 17, 17, 91: 17, 94: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{79: 703, 705, 109: 704, 702, 116: 727, 137: 741},
		{2: 721, 8: 115, 139: 723, 171: 722, 724},
		// 440
		{42: 9, 79: 9, 9, 9, 91: 9, 94: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{8: 117, 10: 117, 139: 738},
		{8: 114, 10: 730},
		{8: 728},
		{8: 725},
		// 445
		{79: 703, 705, 109: 704, 702, 116: 727, 137: 726},
		{42: 18, 79: 18, 18, 18, 91: 18, 94: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{42: 6, 79: 6, 6, 6, 91: 6, 94: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{79: 703, 705, 109: 704, 702, 116: 727, 137: 729},
		{42: 20, 79: 20, 20, 20, 91: 20, 94: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 450
		{2: 731, 139: 732},
		{8: 116, 10: 116, 139: 735},
		{8: 733},
		{79: 703, 705, 109: 704, 702, 116: 727, 137: 734},
		{42: 19, 79: 19, 19, 19, 91: 19, 94: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 455
		{8: 736},
		{79: 703, 705, 109: 704, 702, 116: 727, 137: 737},
		{42: 10, 79: 10, 10, 10, 91: 10, 94: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{8: 739},
		{79: 703, 705, 109: 704, 702, 116: 727, 137: 740},
		// 460
		{42: 11, 79: 11, 11, 11, 91: 11, 94: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{42: 21, 79: 21, 21, 21, 91: 21, 94: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{79: 743},
		{79: 703, 705, 40, 91: 40, 94: 40, 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 685, 690, 687, 689, 136: 744, 140: 745},
		{79: 703, 705, 39, 91: 39, 94: 39, 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 746, 690, 687, 689},
		// 465
		{81: 31, 91: 31, 94: 31},
		{42: 41, 79: 41, 41, 41, 91: 41, 94: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{79: 748},
		{79: 703, 705, 40, 91: 40, 94: 40, 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 685, 690, 687, 689, 136: 744, 140: 749},
		{81: 32, 91: 32, 94: 32},
		// 470
		{79: 751},
		{79: 703, 705, 40, 91: 40, 94: 40, 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 685, 690, 687, 689, 136: 744, 140: 752},
		{81: 33, 91: 33, 94: 33},
		{81: 24, 91: 762, 198: 763, 761},
		{81: 30, 91: 30, 94: 30},
		// 475
		{81: 27, 91: 27, 94: 756, 166: 760},
		{80: 705, 109: 757, 702},
		{79: 758},
		{79: 703, 705, 40, 91: 40, 94: 40, 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 685, 690, 687, 689, 136: 744, 140: 759},
		{81: 26, 91: 26, 94: 26},
		// 480
		{81: 29, 91: 29, 94: 29},
		{81: 767, 200: 766},
		{79: 764},
		{81: 23},
		{79: 703, 705, 40, 95: 694, 695, 696, 691, 692, 693, 697, 698, 688, 699, 700, 109: 704, 702, 116: 701, 129: 686, 131: 685, 690, 687, 689, 136: 744, 140: 765},
		// 485
		{81: 25},
		{42: 34, 79: 34, 34, 34, 91: 34, 94: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{42: 22, 79: 22, 22, 22, 91: 22, 94: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{79: 769},
		{42: 36, 79: 36, 36, 36, 91: 36, 94: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), PrettyString(lval.Token): %v\n", yySymName(n), n, n, PrettyString(lval.Token))
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 212

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if !ok || msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = yyS[yypt-0].node.(*PreprocessingFile)
		}
	case 2:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = yyS[yypt-0].node.(*ConstantExpression)
		}
	case 3:
		{
			lx := yylex.(*lexer)
			if lx.report.Errors(false) == nil && lx.scope.kind != ScopeFile {
				panic("internal error")
			}

			lx.translationUnit = yyS[yypt-0].node.(*TranslationUnit).reverse()
			lx.translationUnit.Declarations = lx.scope
		}
	case 4:
		{
			yyVAL.node = &EnumerationConstant{
				Token: yyS[yypt-0].Token,
			}
		}
	case 5:
		{
			yyVAL.node = &ArgumentExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 6:
		{
			yyVAL.node = &ArgumentExpressionList{
				Case: 1,
				ArgumentExpressionList: yyS[yypt-2].node.(*ArgumentExpressionList),
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 7:
		{
			yyVAL.node = (*ArgumentExpressionListOpt)(nil)
		}
	case 8:
		{
			yyVAL.node = &ArgumentExpressionListOpt{
				ArgumentExpressionList: yyS[yypt-0].node.(*ArgumentExpressionList).reverse(),
			}
		}
	case 9:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
		}
	case 10:
		{
			yyVAL.node = &Expression{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 11:
		{
			yyVAL.node = &Expression{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 12:
		{
			yyVAL.node = &Expression{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 13:
		{
			yyVAL.node = &Expression{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 14:
		{
			yyVAL.node = &Expression{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 15:
		{
			yyVAL.node = &Expression{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
		}
	case 16:
		{
			yyVAL.node = &Expression{
				Case:           7,
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 17:
		{
			yyVAL.node = &Expression{
				Case:           8,
				Expression:     yyS[yypt-3].node.(*Expression),
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 18:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:       9,
				Expression: yyS[yypt-3].node.(*Expression),
				Token:      yyS[yypt-2].Token,
				ArgumentExpressionListOpt: yyS[yypt-1].node.(*ArgumentExpressionListOpt),
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
	case 19:
		{
			yyVAL.node = &Expression{
				Case:       10,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 20:
		{
			yyVAL.node = &Expression{
				Case:       11,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 21:
		{
			yyVAL.node = &Expression{
				Case:       12,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 22:
		{
			yyVAL.node = &Expression{
				Case:       13,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 23:
		{
			yyVAL.node = &Expression{
				Case:            14,
				Token:           yyS[yypt-6].Token,
				TypeName:        yyS[yypt-5].node.(*TypeName),
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token4:          yyS[yypt-0].Token,
			}
		}
	case 24:
		{
			yyVAL.node = &Expression{
				Case:       15,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 25:
		{
			yyVAL.node = &Expression{
				Case:       16,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 26:
		{
			yyVAL.node = &Expression{
				Case:       17,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 27:
		{
			yyVAL.node = &Expression{
				Case:       18,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 28:
		{
			yyVAL.node = &Expression{
				Case:       19,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 29:
		{
			yyVAL.node = &Expression{
				Case:       20,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 30:
		{
			yyVAL.node = &Expression{
				Case:       21,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 31:
		{
			yyVAL.node = &Expression{
				Case:       22,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 32:
		{
			yyVAL.node = &Expression{
				Case:       23,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 33:
		{
			yyVAL.node = &Expression{
				Case:     24,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 34:
		{
			yyVAL.node = &Expression{
				Case:       25,
				Token:      yyS[yypt-3].Token,
				TypeName:   yyS[yypt-2].node.(*TypeName),
				Token2:     yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 35:
		{
			yyVAL.node = &Expression{
				Case:        26,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 36:
		{
			yyVAL.node = &Expression{
				Case:        27,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 37:
		{
			yyVAL.node = &Expression{
				Case:        28,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 38:
		{
			yyVAL.node = &Expression{
				Case:        29,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 39:
		{
			yyVAL.node = &Expression{
				Case:        30,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 40:
		{
			yyVAL.node = &Expression{
				Case:        31,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 41:
		{
			yyVAL.node = &Expression{
				Case:        32,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 42:
		{
			yyVAL.node = &Expression{
				Case:        33,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 43:
		{
			yyVAL.node = &Expression{
				Case:        34,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 44:
		{
			yyVAL.node = &Expression{
				Case:        35,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 45:
		{
			yyVAL.node = &Expression{
				Case:        36,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 46:
		{
			yyVAL.node = &Expression{
				Case:        37,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 47:
		{
			yyVAL.node = &Expression{
				Case:        38,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 48:
		{
			yyVAL.node = &Expression{
				Case:        39,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 49:
		{
			yyVAL.node = &Expression{
				Case:        40,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 50:
		{
			yyVAL.node = &Expression{
				Case:        41,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 51:
		{
			yyVAL.node = &Expression{
				Case:        42,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 52:
		{
			yyVAL.node = &Expression{
				Case:        43,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 53:
		{
			yyVAL.node = &Expression{
				Case:           44,
				Expression:     yyS[yypt-4].node.(*Expression),
				Token:          yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-1].Token,
				Expression2:    yyS[yypt-0].node.(*Expression),
			}
		}
	case 54:
		{
			yyVAL.node = &Expression{
				Case:        45,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 55:
		{
			yyVAL.node = &Expression{
				Case:        46,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 56:
		{
			yyVAL.node = &Expression{
				Case:        47,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 57:
		{
			yyVAL.node = &Expression{
				Case:        48,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 58:
		{
			yyVAL.node = &Expression{
				Case:        49,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 59:
		{
			yyVAL.node = &Expression{
				Case:        50,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 60:
		{
			yyVAL.node = &Expression{
				Case:        51,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 61:
		{
			yyVAL.node = &Expression{
				Case:        52,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 62:
		{
			yyVAL.node = &Expression{
				Case:        53,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 63:
		{
			yyVAL.node = &Expression{
				Case:        54,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 64:
		{
			yyVAL.node = &Expression{
				Case:        55,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 65:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 66:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx.model)
		}
	case 67:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 68:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 69:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 70:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx.model)
		}
	case 71:
		{
			lx := yylex.(*lexer)
			lhs := &ConstantExpression{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Value, lhs.Type = lhs.Expression.eval(lx.model)
			if lhs.Value == nil {
				lx.report.Err(lhs.Pos(), "not a constant expression")
			}
		}
	case 72:
		{
			lx := yylex.(*lexer)
			lhs := &Declaration{
				DeclarationSpecifiers: yyS[yypt-2].node.(*DeclarationSpecifiers),
				InitDeclaratorListOpt: yyS[yypt-1].node.(*InitDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			o := lhs.InitDeclaratorListOpt
			if o != nil {
				break
			}

			s := lhs.DeclarationSpecifiers
			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			lhs.declarator = d
		}
	case 73:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				StorageClassSpecifier:    yyS[yypt-1].node.(*StorageClassSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
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

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.StorageClassSpecifier.Case != 0 /* "typedef" */ && lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
			}
		}
	case 74:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     1,
				TypeSpecifier:            yyS[yypt-1].node.(*TypeSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
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
	case 75:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     2,
				TypeQualifier:            yyS[yypt-1].node.(*TypeQualifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
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

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid type qualifier")
			}
		}
	case 76:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     3,
				FunctionSpecifier:        yyS[yypt-1].node.(*FunctionSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
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

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid function specifier")
			}
		}
	case 77:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 78:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 79:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 80:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 81:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 82:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 83:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 84:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 85:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Case:        1,
				Declarator:  yyS[yypt-3].node.(*Declarator),
				Token:       yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
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
	case 86:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 87:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 88:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 89:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 90:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 91:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 92:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 93:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 94:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 95:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 96:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 97:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 104:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpecifier{
				Case:  13,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypedefName)
			_, lhs.scope = lx.scope.Lookup2(NSIdentifiers, lhs.Token.Val)
		}
	case 105:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
		}
	case 106:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 107:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				StructOrUnion: yyS[yypt-6].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-5].node.(*IdentifierOpt),
				Token:         yyS[yypt-3].Token,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList).reverse(),
				Token2:                yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
							if x.bitOffset == 0 {
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
					d.padding = lhs.sizeOf - off
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
	case 108:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          1,
				StructOrUnion: yyS[yypt-1].node.(*StructOrUnion),
				Token:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareStructTag(lhs.Token, lx.report)
			lhs.scope = lx.scope
		}
	case 109:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 110:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 111:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 112:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 113:
		{
			yyVAL.node = &StructDeclaration{
				SpecifierQualifierList: yyS[yypt-2].node.(*SpecifierQualifierList),
				StructDeclaratorList:   yyS[yypt-1].node.(*StructDeclaratorList).reverse(),
				Token:                  yyS[yypt-0].Token,
			}
		}
	case 114:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				TypeSpecifier:             yyS[yypt-1].node.(*TypeSpecifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
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
	case 115:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				Case:                      1,
				TypeQualifier:             yyS[yypt-1].node.(*TypeQualifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
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

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
		}
	case 116:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 117:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 118:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 119:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 120:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 121:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Case:               1,
				DeclaratorOpt:      yyS[yypt-2].node.(*DeclaratorOpt),
				Token:              yyS[yypt-1].Token,
				ConstantExpression: yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
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
	case 122:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 123:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 124:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 125:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Token:          yyS[yypt-6].Token,
				IdentifierOpt:  yyS[yypt-5].node.(*IdentifierOpt),
				Token2:         yyS[yypt-3].Token,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList).reverse(),
				CommaOpt:       yyS[yypt-1].node.(*CommaOpt),
				Token3:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineEnumTag(o.Token, lhs, lx.report)
			}
		}
	case 126:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareEnumTag(lhs.Token2, lx.report)
		}
	case 127:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 128:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 129:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				EnumerationConstant: yyS[yypt-0].node.(*EnumerationConstant),
			}
			yyVAL.node = lhs
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, lx.iota)
		}
	case 130:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				Case:                1,
				EnumerationConstant: yyS[yypt-2].node.(*EnumerationConstant),
				Token:               yyS[yypt-1].Token,
				ConstantExpression:  yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
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
	case 131:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 132:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 133:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 134:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 135:
		{
			lx := yylex.(*lexer)
			lhs := &Declarator{
				PointerOpt:       yyS[yypt-1].node.(*PointerOpt),
				DirectDeclarator: yyS[yypt-0].node.(*DirectDeclarator),
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lhs.DirectDeclarator.declarator = lhs
		}
	case 136:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 137:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 138:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lx.scope.declareIdentifier(lhs.Token, lhs, lx.report)
			lhs.idScope = lx.scope
		}
	case 139:
		{
			lhs := &DirectDeclarator{
				Case:       1,
				Token:      yyS[yypt-2].Token,
				Declarator: yyS[yypt-1].node.(*Declarator),
				Token2:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Declarator.specifier = nil
			lhs.Declarator.DirectDeclarator.parent = lhs
		}
	case 140:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 2,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				ExpressionOpt:        yyS[yypt-1].node.(*ExpressionOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
	case 141:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 3,
				DirectDeclarator:     yyS[yypt-5].node.(*DirectDeclarator),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx.model)
			var err error
			if lhs.elements, err = elements(lhs.Expression.Value); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 142:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              4,
				DirectDeclarator:  yyS[yypt-5].node.(*DirectDeclarator),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx.model)
			var err error
			if lhs.elements, err = elements(lhs.Expression.Value); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 143:
		{
			lhs := &DirectDeclarator{
				Case:                 5,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Token2:               yyS[yypt-1].Token,
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
			lhs.elements = -1
		}
	case 144:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 145:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              6,
				DirectDeclarator:  yyS[yypt-4].node.(*DirectDeclarator),
				Token:             yyS[yypt-3].Token,
				ParameterTypeList: yyS[yypt-1].node.(*ParameterTypeList),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			lhs.DirectDeclarator.parent = lhs
		}
	case 146:
		{
			lhs := &DirectDeclarator{
				Case:              7,
				DirectDeclarator:  yyS[yypt-3].node.(*DirectDeclarator),
				Token:             yyS[yypt-2].Token,
				IdentifierListOpt: yyS[yypt-1].node.(*IdentifierListOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
		}
	case 147:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 148:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 149:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 150:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 151:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 152:
		{
			lx := yylex.(*lexer)
			lhs := &TypeQualifierList{
				Case:              1,
				TypeQualifierList: yyS[yypt-1].node.(*TypeQualifierList),
				TypeQualifier:     yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			a := lhs.TypeQualifierList
			b := lhs.TypeQualifier
			if a.attr&b.attr != 0 {
				lx.report.Err(b.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
		}
	case 153:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 154:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 155:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 156:
		{
			lhs := &ParameterTypeList{
				Case:          1,
				ParameterList: yyS[yypt-2].node.(*ParameterList).reverse(),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 157:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 158:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 159:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 160:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 161:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.declarator = lhs.Declarator
		}
	case 162:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				Case: 1,
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				AbstractDeclaratorOpt: yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
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
	case 163:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 164:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 165:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 166:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 167:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 168:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 169:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 170:
		{
			lx := yylex.(*lexer)
			lhs := &TypeName{
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				AbstractDeclaratorOpt:  yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
			} else {
				d := &Declarator{
					specifier: lhs.SpecifierQualifierList,
					DirectDeclarator: &DirectDeclarator{
						Case:    0, // IDENTIFIER
						idScope: lx.scope,
					},
				}
				d.DirectDeclarator.declarator = d
				lhs.declarator = d
			}
			lhs.Type = lhs.declarator.setFull(lx)
			lx.popScope(xc.Token{})
		}
	case 171:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier: lx.scope.specifier,
				PointerOpt: &PointerOpt{
					Pointer: lhs.Pointer,
				},
				DirectDeclarator: &DirectDeclarator{
					Case:    0, // IDENTIFIER
					idScope: lx.scope,
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 172:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Case:                     1,
				PointerOpt:               yyS[yypt-1].node.(*PointerOpt),
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier:        lx.scope.specifier,
				PointerOpt:       lhs.PointerOpt,
				DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 173:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 174:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 175:
		{
			lhs := &DirectAbstractDeclarator{
				Token:              yyS[yypt-2].Token,
				AbstractDeclarator: yyS[yypt-1].node.(*AbstractDeclarator),
				Token2:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.AbstractDeclarator.declarator.specifier = nil
			lhs.directDeclarator = &DirectDeclarator{
				Case:       1, // '(' Declarator ')'
				Declarator: lhs.AbstractDeclarator.declarator,
			}
			lhs.AbstractDeclarator.declarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 176:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 1,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:         yyS[yypt-2].Token,
				ExpressionOpt: yyS[yypt-1].node.(*ExpressionOpt),
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
				Case:             2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator: dd,
				ExpressionOpt:    lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 177:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 2,
				DirectAbstractDeclaratorOpt: yyS[yypt-4].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-3].Token,
				TypeQualifierList: yyS[yypt-2].node.(*TypeQualifierList).reverse(),
				ExpressionOpt:     yyS[yypt-1].node.(*ExpressionOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
				Case:                 2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: &TypeQualifierListOpt{lhs.TypeQualifierList},
				ExpressionOpt:        lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 178:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 3,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
				Case:                 2, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: lhs.TypeQualifierListOpt,
				Expression:           lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 179:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 4,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
				Case:              4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
				DirectDeclarator:  dd,
				TypeQualifierList: lhs.TypeQualifierList,
				Expression:        lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 180:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 5,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
				Case:             5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
				DirectDeclarator: dd,
			}
			dd.parent = lhs.directDeclarator
		}
	case 181:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 182:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case:                 6,
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
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
	case 183:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 184:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 7,
				DirectAbstractDeclarator: yyS[yypt-4].node.(*DirectAbstractDeclarator),
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case:              6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator:  lhs.DirectAbstractDeclarator.directDeclarator,
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case:             7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 185:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 186:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 187:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx.model)
		}
	case 188:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 189:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 190:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 191:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 192:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 193:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 194:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 195:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 196:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 197:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 198:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 199:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 200:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 201:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 202:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 203:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 204:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 205:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 206:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 207:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 208:
		{
			lx := yylex.(*lexer)
			lhs := &CompoundStatement{
				Token:            yyS[yypt-3].Token,
				BlockItemListOpt: yyS[yypt-1].node.(*BlockItemListOpt),
				Token2:           yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lhs.Token2)
		}
	case 209:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 210:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 211:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 212:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 213:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 214:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 215:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 216:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx.model)
		}
	case 217:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Token2:         yyS[yypt-5].Token,
				ExpressionList: yyS[yypt-4].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-3].Token,
				Statement:      yyS[yypt-2].node.(*Statement),
				Token4:         yyS[yypt-1].Token,
				Statement2:     yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx.model)
		}
	case 218:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           2,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx.model)
		}
	case 219:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx.model)
		}
	case 220:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Statement:      yyS[yypt-5].node.(*Statement),
				Token2:         yyS[yypt-4].Token,
				Token3:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token4:         yyS[yypt-1].Token,
				Token5:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx.model)
		}
	case 221:
		{
			yyVAL.node = &IterationStatement{
				Case:               2,
				Token:              yyS[yypt-8].Token,
				Token2:             yyS[yypt-7].Token,
				ExpressionListOpt:  yyS[yypt-6].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-5].Token,
				ExpressionListOpt2: yyS[yypt-4].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-3].Token,
				ExpressionListOpt3: yyS[yypt-2].node.(*ExpressionListOpt),
				Token5:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 222:
		{
			yyVAL.node = &IterationStatement{
				Case:               3,
				Token:              yyS[yypt-7].Token,
				Token2:             yyS[yypt-6].Token,
				Declaration:        yyS[yypt-5].node.(*Declaration),
				ExpressionListOpt:  yyS[yypt-4].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-3].Token,
				ExpressionListOpt2: yyS[yypt-2].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 223:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 224:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 225:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 226:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 227:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 228:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 229:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 230:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 231:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-1].node.(*Declarator)
			d.setFull(lx)
			if k := d.Type.Kind(); k != Function {
				lx.report.Err(d.Pos(), "declarator is not a function (have '%s': %v)", d.Type, k)
			}
			lx.scope.mergeScope = d.DirectDeclarator.paramsScope
		}
	case 232:
		{
			lx := yylex.(*lexer)
			lhs := &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				CompoundStatement:     yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
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
	case 233:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 234:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 235:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 236:
		{
			yyVAL.node = &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
		}
	case 237:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 238:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 239:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 240:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 241:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 242:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(node)
		}
	case 243:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(node)
		}
	case 244:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 245:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 246:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 247:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 248:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 249:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 250:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 251:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 252:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 253:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 254:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 255:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 256:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 257:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 258:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 259:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 260:
		{
			yyVAL.node = &ControlLine{
				Case:            1,
				Token:           yyS[yypt-4].Token,
				Token2:          yyS[yypt-3].Token,
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 261:
		{
			yyVAL.node = &ControlLine{
				Case:            2,
				Token:           yyS[yypt-6].Token,
				Token2:          yyS[yypt-5].Token,
				IdentifierList:  yyS[yypt-4].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 262:
		{
			yyVAL.node = &ControlLine{
				Case:              3,
				Token:             yyS[yypt-4].Token,
				Token2:            yyS[yypt-3].Token,
				IdentifierListOpt: yyS[yypt-2].node.(*IdentifierListOpt),
				Token3:            yyS[yypt-1].Token,
				ReplacementList:   yyS[yypt-0].toks,
			}
		}
	case 263:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 264:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 266:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 267:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 268:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 269:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            10,
				Token:           yyS[yypt-5].Token,
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token4, "missing comma before \"...\"")
			}
		}
	case 270:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            11,
				Token:           yyS[yypt-7].Token,
				Token2:          yyS[yypt-6].Token,
				IdentifierList:  yyS[yypt-5].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-4].Token,
				Token4:          yyS[yypt-3].Token,
				Token5:          yyS[yypt-2].Token,
				Token6:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token6, "missing comma before \"...\"")
			}
		}
	case 271:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:   12,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDefine {
				lx.report.ErrTok(lhs.Token2, "expected identifier")
			}
		}
	case 272:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:        13,
				Token:       yyS[yypt-3].Token,
				Token2:      yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token3:      yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableUndefExtraTokens {
				lx.report.ErrTok(decodeTokens(lhs.PPTokenList, nil)[0], "extra tokens after #undef argument")
			}
		}
	case 275:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 276:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
