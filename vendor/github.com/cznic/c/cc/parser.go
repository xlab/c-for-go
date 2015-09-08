// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.10.

package cc

import __yyfmt__ "fmt"

import (
	"github.com/cznic/c/xc"
	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
)

type yySymType struct {
	yys   int
	item  interface{}
	Token xc.Token
	toks  PpTokenList
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault           = 57446
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ANDAND              = 57347
	ANDASSIGN           = 57348
	ARROW               = 57349
	AUTO                = 57350
	BOOL                = 57351
	BREAK               = 57352
	CASE                = 57353
	CHAR                = 57354
	CHARCONST           = 57355
	COMPLEX             = 57356
	CONST               = 57357
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57358
	DDD                 = 57359
	DEC                 = 57360
	DEFAULT             = 57361
	DEFINED             = 57362
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
	MACRO_ARG           = 57388
	MACRO_ARGS          = 1048579
	MACRO_ARG_EMPTY     = 57389
	MODASSIGN           = 57390
	MULASSIGN           = 57391
	NEQ                 = 57392
	NOELSE              = 57393
	ORASSIGN            = 57394
	OROR                = 57395
	PPASSERT            = 57396
	PPDEFINE            = 57397
	PPELIF              = 57398
	PPELSE              = 57399
	PPENDIF             = 57400
	PPERROR             = 57401
	PPHASH_NL           = 57402
	PPHEADER_NAME       = 57403
	PPIDENT             = 57404
	PPIF                = 57405
	PPIFDEF             = 57406
	PPIFNDEF            = 57407
	PPIMPORT            = 57408
	PPINCLUDE           = 57409
	PPINCLUDE_NEXT      = 57410
	PPLINE              = 57411
	PPNONDIRECTIVE      = 57412
	PPNUMBER            = 57413
	PPOTHER             = 57414
	PPPASTE             = 57415
	PPPRAGMA            = 57416
	PPUNASSERT          = 57417
	PPUNDEF             = 57418
	PPWARNING           = 57419
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57420
	RESTRICT            = 57421
	RETURN              = 57422
	RSH                 = 57423
	RSHASSIGN           = 57424
	SHORT               = 57425
	SIGNED              = 57426
	SIZEOF              = 57427
	STATIC              = 57428
	STRINGLITERAL       = 57429
	STRUCT              = 57430
	SUBASSIGN           = 57431
	SWITCH              = 57432
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57433
	TYPEDEFNAME         = 57434
	UNION               = 57435
	UNSIGNED            = 57436
	VOID                = 57437
	VOLATILE            = 57438
	WHILE               = 57439
	XORASSIGN           = 57440
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -310
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (238x)
		42:      1,   // '*' (223x)
		57375:   2,   // IDENTIFIER (210x)
		41:      3,   // ')' (192x)
		59:      4,   // ';' (192x)
		38:      5,   // '&' (185x)
		44:      6,   // ',' (183x)
		43:      7,   // '+' (173x)
		45:      8,   // '-' (173x)
		57360:   9,   // DEC (153x)
		57378:   10,  // INC (153x)
		33:      11,  // '!' (133x)
		126:     12,  // '~' (133x)
		57355:   13,  // CHARCONST (133x)
		57362:   14,  // DEFINED (133x)
		57371:   15,  // FLOATCONST (133x)
		57381:   16,  // INTCONST (133x)
		57384:   17,  // LONGCHARCONST (133x)
		57385:   18,  // LONGSTRINGLITERAL (133x)
		57427:   19,  // SIZEOF (133x)
		57429:   20,  // STRINGLITERAL (133x)
		125:     21,  // '}' (113x)
		57357:   22,  // CONST (113x)
		57421:   23,  // RESTRICT (113x)
		57438:   24,  // VOLATILE (113x)
		58:      25,  // ':' (110x)
		91:      26,  // '[' (106x)
		57344:   27,  // $end (103x)
		57351:   28,  // BOOL (103x)
		57354:   29,  // CHAR (103x)
		57356:   30,  // COMPLEX (103x)
		57365:   31,  // DOUBLE (103x)
		57367:   32,  // ENUM (103x)
		57370:   33,  // FLOAT (103x)
		57380:   34,  // INT (103x)
		57383:   35,  // LONG (103x)
		57425:   36,  // SHORT (103x)
		57426:   37,  // SIGNED (103x)
		57430:   38,  // STRUCT (103x)
		57434:   39,  // TYPEDEFNAME (103x)
		57435:   40,  // UNION (103x)
		57436:   41,  // UNSIGNED (103x)
		57437:   42,  // VOID (103x)
		57428:   43,  // STATIC (98x)
		57350:   44,  // AUTO (92x)
		57369:   45,  // EXTERN (92x)
		57379:   46,  // INLINE (92x)
		57420:   47,  // REGISTER (92x)
		57433:   48,  // TYPEDEF (92x)
		93:      49,  // ']' (86x)
		123:     50,  // '{' (78x)
		57462:   51,  // Constant (70x)
		57527:   52,  // PostfixExpression (70x)
		57532:   53,  // PrimaryExpression (70x)
		57557:   54,  // UnaryExpression (70x)
		57558:   55,  // UnaryOperator (70x)
		57459:   56,  // CastExpression (67x)
		57520:   57,  // MultiplicativeExpression (62x)
		57414:   58,  // PPOTHER (62x)
		57449:   59,  // AdditiveExpression (60x)
		63:      60,  // '?' (59x)
		57395:   61,  // OROR (59x)
		57347:   62,  // ANDAND (58x)
		57536:   63,  // ShiftExpression (58x)
		124:     64,  // '|' (56x)
		10:      65,  // '\n' (55x)
		94:      66,  // '^' (54x)
		57533:   67,  // RelationalExpression (54x)
		57491:   68,  // EqualityExpression (52x)
		57450:   69,  // AndExpression (51x)
		57368:   70,  // EQ (50x)
		57492:   71,  // ExclusiveOrExpression (50x)
		57392:   72,  // NEQ (50x)
		61:      73,  // '=' (49x)
		57507:   74,  // InclusiveOrExpression (49x)
		57400:   75,  // PPENDIF (49x)
		60:      76,  // '<' (48x)
		62:      77,  // '>' (48x)
		57373:   78,  // GEQ (48x)
		57382:   79,  // LEQ (48x)
		57516:   80,  // LogicalAndExpression (48x)
		57461:   81,  // ConditionalExpression (47x)
		57517:   82,  // LogicalOrExpression (47x)
		57386:   83,  // LSH (45x)
		57399:   84,  // PPELSE (45x)
		57423:   85,  // RSH (45x)
		57398:   86,  // PPELIF (44x)
		57453:   87,  // AssignmentExpression (41x)
		57439:   88,  // WHILE (41x)
		57352:   89,  // BREAK (40x)
		57353:   90,  // CASE (40x)
		57358:   91,  // CONTINUE (40x)
		57361:   92,  // DEFAULT (40x)
		57364:   93,  // DO (40x)
		57372:   94,  // FOR (40x)
		57374:   95,  // GOTO (40x)
		57377:   96,  // IF (40x)
		57422:   97,  // RETURN (40x)
		57432:   98,  // SWITCH (40x)
		57396:   99,  // PPASSERT (39x)
		57397:   100, // PPDEFINE (39x)
		57401:   101, // PPERROR (39x)
		57402:   102, // PPHASH_NL (39x)
		57404:   103, // PPIDENT (39x)
		57405:   104, // PPIF (39x)
		57406:   105, // PPIFDEF (39x)
		57407:   106, // PPIFNDEF (39x)
		57408:   107, // PPIMPORT (39x)
		57409:   108, // PPINCLUDE (39x)
		57410:   109, // PPINCLUDE_NEXT (39x)
		57411:   110, // PPLINE (39x)
		57412:   111, // PPNONDIRECTIVE (39x)
		57416:   112, // PPPRAGMA (39x)
		57417:   113, // PPUNASSERT (39x)
		57418:   114, // PPUNDEF (39x)
		57419:   115, // PPWARNING (39x)
		37:      116, // '%' (37x)
		47:      117, // '/' (37x)
		57553:   118, // TypeQualifier (31x)
		57346:   119, // ADDASSIGN (30x)
		57348:   120, // ANDASSIGN (30x)
		57363:   121, // DIVASSIGN (30x)
		57387:   122, // LSHASSIGN (30x)
		57390:   123, // MODASSIGN (30x)
		57391:   124, // MULASSIGN (30x)
		57394:   125, // ORASSIGN (30x)
		57424:   126, // RSHASSIGN (30x)
		57431:   127, // SUBASSIGN (30x)
		57440:   128, // XORASSIGN (30x)
		46:      129, // '.' (29x)
		57528:   130, // PpTokenList (28x)
		57530:   131, // PpTokens (28x)
		57493:   132, // ExpressionList (27x)
		57486:   133, // EnumSpecifier (23x)
		57487:   134, // EnumSpecifier0 (23x)
		57547:   135, // StructOrUnion (23x)
		57548:   136, // StructOrUnionSpecifier (23x)
		57549:   137, // StructOrUnionSpecifier0 (23x)
		57556:   138, // TypeSpecifier (23x)
		57366:   139, // ELSE (22x)
		57349:   140, // ARROW (20x)
		57494:   141, // ExpressionOpt (18x)
		57529:   142, // PpTokenListOpt (17x)
		57468:   143, // DeclarationSpecifiers (16x)
		57498:   144, // FunctionSpecifier (16x)
		57541:   145, // StorageClassSpecifier (16x)
		57460:   146, // CompoundStatement (13x)
		57495:   147, // ExpressionStatement (12x)
		57513:   148, // IterationStatement (12x)
		57514:   149, // JumpStatement (12x)
		57515:   150, // LabeledStatement (12x)
		57535:   151, // SelectionStatement (12x)
		57540:   152, // Statement (12x)
		57525:   153, // Pointer (11x)
		57526:   154, // PointerOpt (10x)
		57464:   155, // ControlLine (8x)
		57470:   156, // Declarator (8x)
		57501:   157, // GroupPart (8x)
		57505:   158, // IfGroup (8x)
		57506:   159, // IfSection (8x)
		57550:   160, // TextLine (8x)
		57465:   161, // Declaration (7x)
		57537:   162, // SpecifierQualifierList (7x)
		57499:   163, // GroupList (6x)
		57534:   164, // ReplacementList (6x)
		57463:   165, // ConstantExpression (5x)
		57359:   166, // DDD (5x)
		57474:   167, // Designator (5x)
		57500:   168, // GroupListOpt (5x)
		57388:   169, // MACRO_ARG (5x)
		57521:   170, // ParameterDeclaration (5x)
		57554:   171, // TypeQualifierList (5x)
		57447:   172, // AbstractDeclarator (4x)
		57469:   173, // DeclarationSpecifiersOpt (4x)
		57472:   174, // Designation (4x)
		57473:   175, // DesignationOpt (4x)
		57475:   176, // DesignatorList (4x)
		57522:   177, // ParameterList (4x)
		57523:   178, // ParameterTypeList (4x)
		57555:   179, // TypeQualifierListOpt (4x)
		57454:   180, // AssignmentExpressionOpt (3x)
		57508:   181, // InitDeclarator (3x)
		57511:   182, // Initializer (3x)
		57524:   183, // ParameterTypeListOpt (3x)
		57552:   184, // TypeName (3x)
		57448:   185, // AbstractDeclaratorOpt (2x)
		57456:   186, // BlockItem (2x)
		57471:   187, // DeclaratorOpt (2x)
		57476:   188, // DirectAbstractDeclarator (2x)
		57477:   189, // DirectAbstractDeclaratorOpt (2x)
		57478:   190, // DirectDeclarator (2x)
		57480:   191, // ElifGroup (2x)
		57488:   192, // EnumerationConstant (2x)
		57489:   193, // Enumerator (2x)
		57496:   194, // ExternalDeclaration (2x)
		57497:   195, // FunctionDefinition (2x)
		57502:   196, // IdentifierList (2x)
		57503:   197, // IdentifierListOpt (2x)
		57504:   198, // IdentifierOpt (2x)
		57509:   199, // InitDeclaratorList (2x)
		57510:   200, // InitDeclaratorListOpt (2x)
		57512:   201, // InitializerList (2x)
		57518:   202, // MacroArgList (2x)
		57538:   203, // SpecifierQualifierListOpt (2x)
		57542:   204, // StructDeclaration (2x)
		57544:   205, // StructDeclarator (2x)
		57441:   206, // $@1 (1x)
		57442:   207, // $@2 (1x)
		57443:   208, // $@3 (1x)
		57444:   209, // $@4 (1x)
		57445:   210, // $@5 (1x)
		57451:   211, // ArgumentExpressionList (1x)
		57452:   212, // ArgumentExpressionListOpt (1x)
		57455:   213, // AssignmentOperator (1x)
		57457:   214, // BlockItemList (1x)
		57458:   215, // BlockItemListOpt (1x)
		1048577: 216, // CONSTANT_EXPRESSION (1x)
		57466:   217, // DeclarationList (1x)
		57467:   218, // DeclarationListOpt (1x)
		57479:   219, // DirectDeclarator2 (1x)
		57481:   220, // ElifGroupList (1x)
		57482:   221, // ElifGroupListOpt (1x)
		57483:   222, // ElseGroup (1x)
		57484:   223, // ElseGroupOpt (1x)
		57485:   224, // EndifLine (1x)
		57490:   225, // EnumeratorList (1x)
		57376:   226, // IDENTIFIER_LPAREN (1x)
		1048579: 227, // MACRO_ARGS (1x)
		57519:   228, // MacroArgsList (1x)
		1048576: 229, // PREPROCESSING_FILE (1x)
		57531:   230, // PreprocessingFile (1x)
		57539:   231, // Start (1x)
		57543:   232, // StructDeclarationList (1x)
		57545:   233, // StructDeclaratorList (1x)
		57546:   234, // StructDeclaratorListOpt (1x)
		1048578: 235, // TRANSLATION_UNIT (1x)
		57551:   236, // TranslationUnit (1x)
		57446:   237, // $default (0x)
		57345:   238, // error (0x)
		57389:   239, // MACRO_ARG_EMPTY (0x)
		57393:   240, // NOELSE (0x)
		57403:   241, // PPHEADER_NAME (0x)
		57413:   242, // PPNUMBER (0x)
		57415:   243, // PPPASTE (0x)
	}

	yySymNames = []string{
		"'('",
		"'*'",
		"IDENTIFIER",
		"')'",
		"';'",
		"'&'",
		"','",
		"'+'",
		"'-'",
		"DEC",
		"INC",
		"'!'",
		"'~'",
		"CHARCONST",
		"DEFINED",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"STRINGLITERAL",
		"'}'",
		"CONST",
		"RESTRICT",
		"VOLATILE",
		"':'",
		"'['",
		"$end",
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
		"AUTO",
		"EXTERN",
		"INLINE",
		"REGISTER",
		"TYPEDEF",
		"']'",
		"'{'",
		"Constant",
		"PostfixExpression",
		"PrimaryExpression",
		"UnaryExpression",
		"UnaryOperator",
		"CastExpression",
		"MultiplicativeExpression",
		"PPOTHER",
		"AdditiveExpression",
		"'?'",
		"OROR",
		"ANDAND",
		"ShiftExpression",
		"'|'",
		"'\\n'",
		"'^'",
		"RelationalExpression",
		"EqualityExpression",
		"AndExpression",
		"EQ",
		"ExclusiveOrExpression",
		"NEQ",
		"'='",
		"InclusiveOrExpression",
		"PPENDIF",
		"'<'",
		"'>'",
		"GEQ",
		"LEQ",
		"LogicalAndExpression",
		"ConditionalExpression",
		"LogicalOrExpression",
		"LSH",
		"PPELSE",
		"RSH",
		"PPELIF",
		"AssignmentExpression",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"GOTO",
		"IF",
		"RETURN",
		"SWITCH",
		"PPASSERT",
		"PPDEFINE",
		"PPERROR",
		"PPHASH_NL",
		"PPIDENT",
		"PPIF",
		"PPIFDEF",
		"PPIFNDEF",
		"PPIMPORT",
		"PPINCLUDE",
		"PPINCLUDE_NEXT",
		"PPLINE",
		"PPNONDIRECTIVE",
		"PPPRAGMA",
		"PPUNASSERT",
		"PPUNDEF",
		"PPWARNING",
		"'%'",
		"'/'",
		"TypeQualifier",
		"ADDASSIGN",
		"ANDASSIGN",
		"DIVASSIGN",
		"LSHASSIGN",
		"MODASSIGN",
		"MULASSIGN",
		"ORASSIGN",
		"RSHASSIGN",
		"SUBASSIGN",
		"XORASSIGN",
		"'.'",
		"PpTokenList",
		"PpTokens",
		"ExpressionList",
		"EnumSpecifier",
		"EnumSpecifier0",
		"StructOrUnion",
		"StructOrUnionSpecifier",
		"StructOrUnionSpecifier0",
		"TypeSpecifier",
		"ELSE",
		"ARROW",
		"ExpressionOpt",
		"PpTokenListOpt",
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
		"SpecifierQualifierList",
		"GroupList",
		"ReplacementList",
		"ConstantExpression",
		"DDD",
		"Designator",
		"GroupListOpt",
		"MACRO_ARG",
		"ParameterDeclaration",
		"TypeQualifierList",
		"AbstractDeclarator",
		"DeclarationSpecifiersOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ParameterList",
		"ParameterTypeList",
		"TypeQualifierListOpt",
		"AssignmentExpressionOpt",
		"InitDeclarator",
		"Initializer",
		"ParameterTypeListOpt",
		"TypeName",
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
		"MacroArgList",
		"SpecifierQualifierListOpt",
		"StructDeclaration",
		"StructDeclarator",
		"$@1",
		"$@2",
		"$@3",
		"$@4",
		"$@5",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"AssignmentOperator",
		"BlockItemList",
		"BlockItemListOpt",
		"CONSTANT_EXPRESSION",
		"DeclarationList",
		"DeclarationListOpt",
		"DirectDeclarator2",
		"ElifGroupList",
		"ElifGroupListOpt",
		"ElseGroup",
		"ElseGroupOpt",
		"EndifLine",
		"EnumeratorList",
		"IDENTIFIER_LPAREN",
		"MACRO_ARGS",
		"MacroArgsList",
		"PREPROCESSING_FILE",
		"PreprocessingFile",
		"Start",
		"StructDeclarationList",
		"StructDeclaratorList",
		"StructDeclaratorListOpt",
		"TRANSLATION_UNIT",
		"TranslationUnit",
		"$default",
		"error",
		"MACRO_ARG_EMPTY",
		"NOELSE",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {231, 2},
		2:   {231, 2},
		3:   {231, 2},
		4:   {231, 4},
		5:   {192, 1},
		6:   {53, 1},
		7:   {53, 1},
		8:   {53, 3},
		9:   {51, 1},
		10:  {51, 1},
		11:  {51, 1},
		12:  {51, 1},
		13:  {51, 1},
		14:  {51, 1},
		15:  {52, 1},
		16:  {52, 4},
		17:  {52, 4},
		18:  {52, 3},
		19:  {52, 3},
		20:  {52, 2},
		21:  {52, 2},
		22:  {52, 6},
		23:  {52, 7},
		24:  {211, 1},
		25:  {211, 3},
		26:  {212, 0},
		27:  {212, 1},
		28:  {54, 1},
		29:  {54, 2},
		30:  {54, 2},
		31:  {54, 2},
		32:  {54, 2},
		33:  {54, 4},
		34:  {54, 2},
		35:  {54, 4},
		36:  {55, 1},
		37:  {55, 1},
		38:  {55, 1},
		39:  {55, 1},
		40:  {55, 1},
		41:  {55, 1},
		42:  {56, 1},
		43:  {56, 4},
		44:  {57, 1},
		45:  {57, 3},
		46:  {57, 3},
		47:  {57, 3},
		48:  {59, 1},
		49:  {59, 3},
		50:  {59, 3},
		51:  {63, 1},
		52:  {63, 3},
		53:  {63, 3},
		54:  {67, 1},
		55:  {67, 3},
		56:  {67, 3},
		57:  {67, 3},
		58:  {67, 3},
		59:  {68, 1},
		60:  {68, 3},
		61:  {68, 3},
		62:  {69, 1},
		63:  {69, 3},
		64:  {71, 1},
		65:  {71, 3},
		66:  {74, 1},
		67:  {74, 3},
		68:  {80, 1},
		69:  {80, 3},
		70:  {82, 1},
		71:  {82, 3},
		72:  {81, 1},
		73:  {81, 5},
		74:  {87, 1},
		75:  {87, 3},
		76:  {180, 0},
		77:  {180, 1},
		78:  {213, 1},
		79:  {213, 1},
		80:  {213, 1},
		81:  {213, 1},
		82:  {213, 1},
		83:  {213, 1},
		84:  {213, 1},
		85:  {213, 1},
		86:  {213, 1},
		87:  {213, 1},
		88:  {213, 1},
		89:  {132, 1},
		90:  {132, 3},
		91:  {141, 0},
		92:  {141, 1},
		93:  {165, 1},
		94:  {161, 3},
		95:  {143, 2},
		96:  {206, 0},
		97:  {143, 3},
		98:  {143, 2},
		99:  {143, 2},
		100: {173, 0},
		101: {173, 1},
		102: {199, 1},
		103: {199, 3},
		104: {200, 0},
		105: {200, 1},
		106: {181, 1},
		107: {181, 3},
		108: {145, 1},
		109: {145, 1},
		110: {145, 1},
		111: {145, 1},
		112: {145, 1},
		113: {138, 1},
		114: {138, 1},
		115: {138, 1},
		116: {138, 1},
		117: {138, 1},
		118: {138, 1},
		119: {138, 1},
		120: {138, 1},
		121: {138, 1},
		122: {138, 1},
		123: {138, 1},
		124: {138, 1},
		125: {138, 1},
		126: {138, 1},
		127: {137, 2},
		128: {136, 4},
		129: {136, 2},
		130: {135, 1},
		131: {135, 1},
		132: {232, 1},
		133: {232, 2},
		134: {204, 3},
		135: {207, 0},
		136: {162, 3},
		137: {162, 2},
		138: {203, 0},
		139: {203, 1},
		140: {233, 1},
		141: {233, 3},
		142: {234, 0},
		143: {234, 1},
		144: {205, 1},
		145: {205, 3},
		146: {134, 2},
		147: {133, 4},
		148: {133, 5},
		149: {133, 2},
		150: {225, 1},
		151: {225, 3},
		152: {193, 1},
		153: {193, 3},
		154: {118, 1},
		155: {118, 1},
		156: {118, 1},
		157: {144, 1},
		158: {156, 2},
		159: {187, 0},
		160: {187, 1},
		161: {190, 1},
		162: {190, 3},
		163: {190, 5},
		164: {190, 6},
		165: {190, 6},
		166: {190, 5},
		167: {208, 0},
		168: {190, 4},
		169: {219, 2},
		170: {219, 2},
		171: {153, 2},
		172: {153, 3},
		173: {154, 0},
		174: {154, 1},
		175: {171, 1},
		176: {171, 2},
		177: {179, 0},
		178: {179, 1},
		179: {178, 1},
		180: {178, 3},
		181: {183, 0},
		182: {183, 1},
		183: {177, 1},
		184: {177, 3},
		185: {170, 2},
		186: {170, 2},
		187: {196, 1},
		188: {196, 3},
		189: {197, 0},
		190: {197, 1},
		191: {198, 0},
		192: {198, 1},
		193: {184, 2},
		194: {172, 1},
		195: {172, 2},
		196: {185, 0},
		197: {185, 1},
		198: {188, 3},
		199: {188, 4},
		200: {188, 5},
		201: {188, 6},
		202: {188, 6},
		203: {188, 4},
		204: {188, 3},
		205: {188, 4},
		206: {189, 0},
		207: {189, 1},
		208: {182, 1},
		209: {182, 3},
		210: {182, 4},
		211: {201, 2},
		212: {201, 4},
		213: {174, 2},
		214: {175, 0},
		215: {175, 1},
		216: {176, 1},
		217: {176, 2},
		218: {167, 3},
		219: {167, 2},
		220: {152, 1},
		221: {152, 1},
		222: {152, 1},
		223: {152, 1},
		224: {152, 1},
		225: {152, 1},
		226: {150, 3},
		227: {150, 4},
		228: {150, 3},
		229: {209, 0},
		230: {146, 4},
		231: {214, 1},
		232: {214, 2},
		233: {215, 0},
		234: {215, 1},
		235: {186, 1},
		236: {186, 1},
		237: {147, 2},
		238: {151, 5},
		239: {151, 7},
		240: {151, 5},
		241: {148, 5},
		242: {148, 7},
		243: {148, 9},
		244: {148, 8},
		245: {149, 3},
		246: {149, 2},
		247: {149, 2},
		248: {149, 3},
		249: {236, 1},
		250: {236, 2},
		251: {194, 1},
		252: {194, 1},
		253: {210, 0},
		254: {195, 5},
		255: {217, 1},
		256: {217, 2},
		257: {218, 0},
		258: {218, 1},
		259: {230, 1},
		260: {163, 1},
		261: {163, 2},
		262: {168, 0},
		263: {168, 1},
		264: {157, 1},
		265: {157, 1},
		266: {157, 2},
		267: {157, 1},
		268: {159, 4},
		269: {158, 3},
		270: {158, 4},
		271: {158, 4},
		272: {220, 1},
		273: {220, 2},
		274: {221, 0},
		275: {221, 1},
		276: {191, 3},
		277: {222, 3},
		278: {223, 0},
		279: {223, 1},
		280: {224, 2},
		281: {155, 3},
		282: {155, 5},
		283: {155, 7},
		284: {155, 5},
		285: {155, 2},
		286: {155, 1},
		287: {155, 2},
		288: {155, 2},
		289: {155, 2},
		290: {155, 3},
		291: {155, 2},
		292: {155, 6},
		293: {155, 8},
		294: {155, 2},
		295: {155, 2},
		296: {155, 2},
		297: {155, 2},
		298: {155, 2},
		299: {160, 1},
		300: {164, 1},
		301: {130, 2},
		302: {142, 1},
		303: {142, 1},
		304: {131, 1},
		305: {131, 2},
		306: {228, 1},
		307: {228, 3},
		308: {202, 0},
		309: {202, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 27}:   "invalid empty input",
		yyXError{508, -1}: "expected #endif",
		yyXError{510, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{10, -1}:  "expected $end",
		yyXError{422, -1}: "expected $end",
		yyXError{423, -1}: "expected $end",
		yyXError{5, -1}:   "expected '('",
		yyXError{357, -1}: "expected '('",
		yyXError{358, -1}: "expected '('",
		yyXError{359, -1}: "expected '('",
		yyXError{361, -1}: "expected '('",
		yyXError{387, -1}: "expected '('",
		yyXError{153, -1}: "expected ')'",
		yyXError{158, -1}: "expected ')'",
		yyXError{164, -1}: "expected ')'",
		yyXError{190, -1}: "expected ')'",
		yyXError{193, -1}: "expected ')'",
		yyXError{194, -1}: "expected ')'",
		yyXError{203, -1}: "expected ')'",
		yyXError{209, -1}: "expected ')'",
		yyXError{210, -1}: "expected ')'",
		yyXError{231, -1}: "expected ')'",
		yyXError{234, -1}: "expected ')'",
		yyXError{272, -1}: "expected ')'",
		yyXError{283, -1}: "expected ')'",
		yyXError{291, -1}: "expected ')'",
		yyXError{377, -1}: "expected ')'",
		yyXError{383, -1}: "expected ')'",
		yyXError{470, -1}: "expected ')'",
		yyXError{471, -1}: "expected ')'",
		yyXError{479, -1}: "expected ')'",
		yyXError{482, -1}: "expected ')'",
		yyXError{485, -1}: "expected ')'",
		yyXError{308, -1}: "expected ':'",
		yyXError{350, -1}: "expected ':'",
		yyXError{411, -1}: "expected ':'",
		yyXError{304, -1}: "expected ';'",
		yyXError{327, -1}: "expected ';'",
		yyXError{356, -1}: "expected ';'",
		yyXError{363, -1}: "expected ';'",
		yyXError{364, -1}: "expected ';'",
		yyXError{366, -1}: "expected ';'",
		yyXError{370, -1}: "expected ';'",
		yyXError{373, -1}: "expected ';'",
		yyXError{375, -1}: "expected ';'",
		yyXError{381, -1}: "expected ';'",
		yyXError{390, -1}: "expected ';'",
		yyXError{169, -1}: "expected '['",
		yyXError{460, -1}: "expected '\\\\n'",
		yyXError{489, -1}: "expected '\\\\n'",
		yyXError{494, -1}: "expected '\\\\n'",
		yyXError{507, -1}: "expected '\\\\n'",
		yyXError{172, -1}: "expected ']'",
		yyXError{175, -1}: "expected ']'",
		yyXError{179, -1}: "expected ']'",
		yyXError{183, -1}: "expected ']'",
		yyXError{185, -1}: "expected ']'",
		yyXError{221, -1}: "expected ']'",
		yyXError{224, -1}: "expected ']'",
		yyXError{227, -1}: "expected ']'",
		yyXError{252, -1}: "expected ']'",
		yyXError{39, -1}:  "expected '{'",
		yyXError{43, -1}:  "expected '{'",
		yyXError{273, -1}: "expected '{'",
		yyXError{298, -1}: "expected '{'",
		yyXError{319, -1}: "expected '{'",
		yyXError{351, -1}: "expected '}'",
		yyXError{206, -1}: "expected DirectDeclarator2 or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{207, -1}: "expected DirectDeclarator2 or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, macro arguments prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{303, -1}: "expected StructDeclaratorListOpt or one of ['(', '*', ':', ';', identifier]",
		yyXError{202, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{168, -1}: "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{133, -1}: "expected additive expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{134, -1}: "expected additive expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{121, -1}: "expected and expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{100, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{103, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{178, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{184, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{220, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{223, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{286, -1}: "expected assignment expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{176, -1}: "expected assignment expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{218, -1}: "expected assignment expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{95, -1}:  "expected assignment operator or one of [!=, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{353, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{292, -1}: "expected cast expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected cast expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{139, -1}: "expected cast expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{140, -1}: "expected cast expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{141, -1}: "expected cast expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{334, -1}: "expected compound statement or '{'",
		yyXError{330, -1}: "expected compound statement or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{332, -1}: "expected compound statement or optional declaration list or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{99, -1}:  "expected conditional expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{56, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{249, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{312, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{349, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{336, -1}: "expected declaration or one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{372, -1}: "expected declaration or optional expression or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, defined, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{311, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{192, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{14, -1}:  "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{246, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{197, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{165, -1}: "expected direct abstract declarator or one of ['(', '[']",
		yyXError{309, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{501, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{506, -1}: "expected endif line or #endif",
		yyXError{430, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{499, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{51, -1}:  "expected enumerator list or identifier",
		yyXError{295, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{123, -1}: "expected equality expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{94, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{275, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{388, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{392, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{396, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{400, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{59, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, defined, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{155, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, defined, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{270, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Bool, _Complex, char, character constant, const, defined, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected external declaration or one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{491, -1}: "expected group part or one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{424, -1}: "expected group part or one of [#, #assert, #define, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{152, -1}: "expected identifier",
		yyXError{213, -1}: "expected identifier",
		yyXError{250, -1}: "expected identifier",
		yyXError{277, -1}: "expected identifier",
		yyXError{278, -1}: "expected identifier",
		yyXError{362, -1}: "expected identifier",
		yyXError{432, -1}: "expected identifier",
		yyXError{433, -1}: "expected identifier",
		yyXError{440, -1}: "expected identifier",
		yyXError{467, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{117, -1}: "expected inclusive-or expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{418, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{243, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{257, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{245, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{263, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{331, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{261, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{266, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{93, -1}:  "expected logical-and expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{11, -1}:  "expected macro argument list or one of [')', ',', macro argument]",
		yyXError{6, -1}:   "expected macro arguments list or one of [')', ',', macro argument]",
		yyXError{136, -1}: "expected multiplicative expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{137, -1}: "expected multiplicative expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{64, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{66, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{67, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{265, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{267, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{268, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{281, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{282, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{288, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{290, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{242, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '{', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{79, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{151, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{156, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{269, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{271, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{293, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', ']', '^', '|', '}', *=, +=, -=, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{80, -1}:  "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{81, -1}:  "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{138, -1}: "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{142, -1}: "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{143, -1}: "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{144, -1}: "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{145, -1}: "expected one of [!=, $end, &&, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{82, -1}:  "expected one of [!=, $end, &&, '&', ')', '+', ',', '-', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{135, -1}: "expected one of [!=, $end, &&, '&', ')', '+', ',', '-', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{146, -1}: "expected one of [!=, $end, &&, '&', ')', '+', ',', '-', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{83, -1}:  "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{132, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{147, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{148, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{149, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <<, <=, ==, >=, >>, ||]",
		yyXError{84, -1}:  "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <=, ==, >=, ||]",
		yyXError{127, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <=, ==, >=, ||]",
		yyXError{150, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '<', '>', '?', ']', '^', '|', '}', <=, ==, >=, ||]",
		yyXError{85, -1}:  "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '?', ']', '^', '|', '}', ==, ||]",
		yyXError{124, -1}: "expected one of [!=, $end, &&, '&', ')', ',', ':', ';', '?', ']', '^', '|', '}', ==, ||]",
		yyXError{341, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{425, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{426, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{427, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{429, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{436, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{447, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{449, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{450, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{452, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{454, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{455, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{456, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{457, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{458, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{459, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{461, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{462, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{463, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{464, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{465, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{473, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{474, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{476, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{481, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{484, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{487, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{488, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{493, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{511, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{513, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{514, -1}: "expected one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, $end, '\\\\n', ppother]",
		yyXError{492, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{496, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{498, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{500, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{504, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{505, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{86, -1}:  "expected one of [$end, &&, '&', ')', ',', ':', ';', '?', ']', '^', '|', '}', ||]",
		yyXError{122, -1}: "expected one of [$end, &&, '&', ')', ',', ':', ';', '?', ']', '^', '|', '}', ||]",
		yyXError{87, -1}:  "expected one of [$end, &&, ')', ',', ':', ';', '?', ']', '^', '|', '}', ||]",
		yyXError{120, -1}: "expected one of [$end, &&, ')', ',', ':', ';', '?', ']', '^', '|', '}', ||]",
		yyXError{88, -1}:  "expected one of [$end, &&, ')', ',', ':', ';', '?', ']', '|', '}', ||]",
		yyXError{118, -1}: "expected one of [$end, &&, ')', ',', ':', ';', '?', ']', '|', '}', ||]",
		yyXError{89, -1}:  "expected one of [$end, &&, ')', ',', ':', ';', '?', ']', '}', ||]",
		yyXError{116, -1}: "expected one of [$end, &&, ')', ',', ':', ';', '?', ']', '}', ||]",
		yyXError{408, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{420, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{90, -1}:  "expected one of [$end, ')', ',', ':', ';', '?', ']', '}', ||]",
		yyXError{102, -1}: "expected one of [$end, ')', ',', ':', ';', ']', '}']",
		yyXError{91, -1}:  "expected one of [$end, ',', ':', ';', ']', '}']",
		yyXError{48, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{49, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{50, -1}:  "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{339, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{421, -1}: "expected one of [$end, _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{44, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, defined, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{45, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, defined, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{46, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Bool, _Complex, auto, char, character constant, const, defined, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{173, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{181, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{343, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{344, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{345, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{346, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{347, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{348, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{367, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{368, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{369, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{371, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{379, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{385, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{391, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{395, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{399, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{406, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{410, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{413, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{415, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{352, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{354, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{355, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{407, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{171, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{226, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{247, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{254, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{73, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{74, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{75, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{76, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{77, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{78, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{104, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{105, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{106, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{107, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{108, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{109, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{110, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{111, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{112, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{113, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{114, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{299, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{320, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{31, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{32, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{34, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{35, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{36, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{294, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{296, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{317, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{19, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{20, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{21, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{22, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{23, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{47, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{321, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{322, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{323, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{325, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{326, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{238, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{241, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{200, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{201, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{204, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{208, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{215, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{216, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{222, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{225, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{228, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{229, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{163, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{237, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{167, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{180, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{182, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{186, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{187, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{188, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{195, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{196, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{235, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{72, -1}:  "expected one of ['(', identifier]",
		yyXError{310, -1}: "expected one of ['(', identifier]",
		yyXError{97, -1}:  "expected one of [')', ',', ':', ';', ']', '}']",
		yyXError{115, -1}: "expected one of [')', ',', ':', ';', ']', '}']",
		yyXError{98, -1}:  "expected one of [')', ',', ':', ';', ']']",
		yyXError{101, -1}: "expected one of [')', ',', ':', ';', ']']",
		yyXError{342, -1}: "expected one of [')', ',', ';']",
		yyXError{468, -1}: "expected one of [')', ',', ...]",
		yyXError{478, -1}: "expected one of [')', ',', ...]",
		yyXError{8, -1}:   "expected one of [')', ',', macro argument]",
		yyXError{9, -1}:   "expected one of [')', ',', macro argument]",
		yyXError{12, -1}:  "expected one of [')', ',', macro argument]",
		yyXError{7, -1}:   "expected one of [')', ',']",
		yyXError{157, -1}: "expected one of [')', ',']",
		yyXError{166, -1}: "expected one of [')', ',']",
		yyXError{189, -1}: "expected one of [')', ',']",
		yyXError{191, -1}: "expected one of [')', ',']",
		yyXError{198, -1}: "expected one of [')', ',']",
		yyXError{199, -1}: "expected one of [')', ',']",
		yyXError{211, -1}: "expected one of [')', ',']",
		yyXError{212, -1}: "expected one of [')', ',']",
		yyXError{214, -1}: "expected one of [')', ',']",
		yyXError{232, -1}: "expected one of [')', ',']",
		yyXError{284, -1}: "expected one of [')', ',']",
		yyXError{285, -1}: "expected one of [')', ',']",
		yyXError{287, -1}: "expected one of [')', ',']",
		yyXError{389, -1}: "expected one of [')', ',']",
		yyXError{393, -1}: "expected one of [')', ',']",
		yyXError{397, -1}: "expected one of [')', ',']",
		yyXError{401, -1}: "expected one of [')', ',']",
		yyXError{469, -1}: "expected one of [')', ',']",
		yyXError{307, -1}: "expected one of [',', ':', ';']",
		yyXError{96, -1}:  "expected one of [',', ':']",
		yyXError{416, -1}: "expected one of [',', ';', '=']",
		yyXError{256, -1}: "expected one of [',', ';', '}']",
		yyXError{260, -1}: "expected one of [',', ';', '}']",
		yyXError{262, -1}: "expected one of [',', ';', '}']",
		yyXError{305, -1}: "expected one of [',', ';']",
		yyXError{306, -1}: "expected one of [',', ';']",
		yyXError{313, -1}: "expected one of [',', ';']",
		yyXError{315, -1}: "expected one of [',', ';']",
		yyXError{328, -1}: "expected one of [',', ';']",
		yyXError{329, -1}: "expected one of [',', ';']",
		yyXError{417, -1}: "expected one of [',', ';']",
		yyXError{419, -1}: "expected one of [',', ';']",
		yyXError{52, -1}:  "expected one of [',', '=', '}']",
		yyXError{55, -1}:  "expected one of [',', '=', '}']",
		yyXError{289, -1}: "expected one of [',', ']']",
		yyXError{53, -1}:  "expected one of [',', '}']",
		yyXError{54, -1}:  "expected one of [',', '}']",
		yyXError{92, -1}:  "expected one of [',', '}']",
		yyXError{244, -1}: "expected one of [',', '}']",
		yyXError{258, -1}: "expected one of [',', '}']",
		yyXError{259, -1}: "expected one of [',', '}']",
		yyXError{264, -1}: "expected one of [',', '}']",
		yyXError{297, -1}: "expected one of [',', '}']",
		yyXError{248, -1}: "expected one of ['.', '=', '[']",
		yyXError{251, -1}: "expected one of ['.', '=', '[']",
		yyXError{253, -1}: "expected one of ['.', '=', '[']",
		yyXError{255, -1}: "expected one of ['.', '=', '[']",
		yyXError{448, -1}: "expected one of ['\\\\n', ppother]",
		yyXError{451, -1}: "expected one of ['\\\\n', ppother]",
		yyXError{453, -1}: "expected one of ['\\\\n', ppother]",
		yyXError{335, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{337, -1}: "expected one of ['{', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of ['{', identifier]",
		yyXError{41, -1}:  "expected one of ['{', identifier]",
		yyXError{302, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{316, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{318, -1}: "expected one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{477, -1}: "expected one of [..., identifier]",
		yyXError{434, -1}: "expected one of [identifier, identifier ]",
		yyXError{161, -1}: "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{276, -1}: "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{217, -1}: "expected optional assignment expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{205, -1}: "expected optional assignment expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{170, -1}: "expected optional assignment expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{174, -1}: "expected optional assignment expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{338, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{340, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Bool, _Complex, auto, break, case, char, character constant, const, continue, default, defined, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, union, unsigned, void, volatile, while]",
		yyXError{15, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{18, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{324, -1}: "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{376, -1}: "expected optional expression or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{382, -1}: "expected optional expression or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{365, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{374, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{380, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{490, -1}: "expected optional group list or one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{495, -1}: "expected optional group list or one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{497, -1}: "expected optional group list or one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{503, -1}: "expected optional group list or one of [#, #assert, #define, #elif, #else, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{509, -1}: "expected optional group list or one of [#, #assert, #define, #endif, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{38, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{42, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{333, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{233, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{159, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{160, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{240, -1}: "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{435, -1}: "expected optional token list or one of ['\\\\n', ppother]",
		yyXError{439, -1}: "expected optional token list or one of ['\\\\n', ppother]",
		yyXError{512, -1}: "expected optional token list or one of ['\\\\n', ppother]",
		yyXError{162, -1}: "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{230, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{236, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #assert, #define, #error, #foo, #ident, #if, #ifdef, #ifndef, #import, #include, #include_next, #line, #pragma, #unassert, #undef, #warning, '\\\\n', ppother]",
		yyXError{125, -1}: "expected relational expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{126, -1}: "expected relational expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{466, -1}: "expected replacement token list or one of ['\\\\n', ppother]",
		yyXError{472, -1}: "expected replacement token list or one of ['\\\\n', ppother]",
		yyXError{475, -1}: "expected replacement token list or one of ['\\\\n', ppother]",
		yyXError{480, -1}: "expected replacement token list or one of ['\\\\n', ppother]",
		yyXError{483, -1}: "expected replacement token list or one of ['\\\\n', ppother]",
		yyXError{486, -1}: "expected replacement token list or one of ['\\\\n', ppother]",
		yyXError{128, -1}: "expected shift expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{129, -1}: "expected shift expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{130, -1}: "expected shift expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{131, -1}: "expected shift expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{360, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{378, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{384, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{394, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{398, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{402, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{404, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{409, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{412, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{414, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, break, case, character constant, continue, default, defined, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{300, -1}: "expected struct declaration list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{301, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, union, unsigned, void, volatile]",
		yyXError{314, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{428, -1}: "expected token list or ppother",
		yyXError{431, -1}: "expected token list or ppother",
		yyXError{437, -1}: "expected token list or ppother",
		yyXError{438, -1}: "expected token list or ppother",
		yyXError{441, -1}: "expected token list or ppother",
		yyXError{442, -1}: "expected token list or ppother",
		yyXError{443, -1}: "expected token list or ppother",
		yyXError{444, -1}: "expected token list or ppother",
		yyXError{445, -1}: "expected token list or ppother",
		yyXError{446, -1}: "expected token list or ppother",
		yyXError{502, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [_Bool, _Complex, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, union, unsigned, void, volatile]",
		yyXError{177, -1}: "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{219, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, character constant, const, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{68, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{71, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{386, -1}: "expected while",
		yyXError{119, -1}: "expected xor expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, character constant, defined, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{3, 27}:   "constant expression: unexpected EOF",
		yyXError{5, 27}:   "macro arguments: unexpected EOF",
		yyXError{2, 27}:   "preprocessing file: unexpected EOF",
		yyXError{4, 27}:   "translation unit: unexpected EOF",
	}

	yyParseTab = [515][]uint16{
		// 0
		{216: 313, 227: 315, 229: 312, 231: 311, 235: 314},
		{27: 310},
		{58: 761, 65: 759, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 735, 740, 737, 739, 163: 734, 230: 733},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 401, 400, 165: 732},
		{22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 324, 328, 325, 161: 360, 194: 358, 359, 236: 323},
		// 5
		{316},
		{3: 2, 6: 2, 169: 2, 202: 318, 228: 317},
		{3: 320, 6: 321},
		{3: 4, 6: 4, 169: 319},
		{3: 1, 6: 1, 169: 1},
		// 10
		{27: 306},
		{3: 2, 6: 2, 169: 2, 202: 322},
		{3: 3, 6: 3, 169: 319},
		{22: 354, 355, 356, 27: 307, 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 324, 328, 325, 161: 360, 194: 731, 359},
		{137, 472, 137, 4: 206, 153: 620, 619, 156: 640, 181: 638, 199: 639, 637},
		// 15
		{210, 210, 210, 210, 210, 6: 210, 22: 354, 355, 356, 26: 210, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 632, 328, 325, 173: 636},
		{214, 214, 214, 214, 214, 6: 214, 22: 214, 214, 214, 26: 214, 28: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 206: 634},
		{210, 210, 210, 210, 210, 6: 210, 22: 354, 355, 356, 26: 210, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 632, 328, 325, 173: 633},
		{210, 210, 210, 210, 210, 6: 210, 22: 354, 355, 356, 26: 210, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 632, 328, 325, 173: 631},
		{202, 202, 202, 202, 202, 6: 202, 22: 202, 202, 202, 26: 202, 28: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202},
		// 20
		{201, 201, 201, 201, 201, 6: 201, 22: 201, 201, 201, 26: 201, 28: 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201, 201},
		{200, 200, 200, 200, 200, 6: 200, 22: 200, 200, 200, 26: 200, 28: 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200},
		{199, 199, 199, 199, 199, 6: 199, 22: 199, 199, 199, 26: 199, 28: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199},
		{198, 198, 198, 198, 198, 6: 198, 22: 198, 198, 198, 26: 198, 28: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198},
		{197, 197, 197, 197, 197, 6: 197, 22: 197, 197, 197, 197, 197, 28: 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197},
		// 25
		{196, 196, 196, 196, 196, 6: 196, 22: 196, 196, 196, 196, 196, 28: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196},
		{195, 195, 195, 195, 195, 6: 195, 22: 195, 195, 195, 195, 195, 28: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195},
		{194, 194, 194, 194, 194, 6: 194, 22: 194, 194, 194, 194, 194, 28: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194},
		{193, 193, 193, 193, 193, 6: 193, 22: 193, 193, 193, 193, 193, 28: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193},
		{192, 192, 192, 192, 192, 6: 192, 22: 192, 192, 192, 192, 192, 28: 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192, 192},
		// 30
		{191, 191, 191, 191, 191, 6: 191, 22: 191, 191, 191, 191, 191, 28: 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191, 191},
		{190, 190, 190, 190, 190, 6: 190, 22: 190, 190, 190, 190, 190, 28: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190},
		{189, 189, 189, 189, 189, 6: 189, 22: 189, 189, 189, 189, 189, 28: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189},
		{188, 188, 188, 188, 188, 6: 188, 22: 188, 188, 188, 188, 188, 28: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188},
		{187, 187, 187, 187, 187, 6: 187, 22: 187, 187, 187, 187, 187, 28: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187},
		// 35
		{186, 186, 186, 186, 186, 6: 186, 22: 186, 186, 186, 186, 186, 28: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186},
		{185, 185, 185, 185, 185, 6: 185, 22: 185, 185, 185, 185, 185, 28: 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185, 185},
		{184, 184, 184, 184, 184, 6: 184, 22: 184, 184, 184, 184, 184, 28: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184},
		{2: 630, 50: 119, 198: 629},
		{50: 610},
		// 40
		{2: 180, 50: 180},
		{2: 179, 50: 179},
		{2: 609, 50: 119, 198: 608},
		{50: 361},
		{156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 22: 156, 156, 156, 156, 156, 28: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156},
		// 45
		{155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 22: 155, 155, 155, 155, 155, 28: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155},
		{154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 22: 154, 154, 154, 154, 154, 28: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154},
		{153, 153, 153, 153, 153, 6: 153, 22: 153, 153, 153, 26: 153, 28: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153},
		{22: 61, 61, 61, 27: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61},
		{22: 59, 59, 59, 27: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59},
		// 50
		{22: 58, 58, 58, 27: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58},
		{2: 362, 192: 365, 364, 225: 363},
		{6: 305, 21: 305, 73: 305},
		{6: 605, 21: 604},
		{6: 160, 21: 160},
		// 55
		{6: 158, 21: 158, 73: 366},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 401, 400, 165: 402},
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 21: 304, 25: 304, 304, 304, 49: 304, 60: 304, 304, 304, 64: 304, 66: 304, 70: 304, 72: 304, 304, 76: 304, 304, 304, 304, 83: 304, 85: 304, 116: 304, 304, 119: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 140: 304},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 21: 303, 25: 303, 303, 303, 49: 303, 60: 303, 303, 303, 64: 303, 66: 303, 70: 303, 72: 303, 303, 76: 303, 303, 303, 303, 83: 303, 85: 303, 116: 303, 303, 119: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 140: 303},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 118: 470, 132: 467, 346, 353, 348, 345, 349, 469, 162: 471, 184: 601},
		// 60
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 21: 301, 25: 301, 301, 301, 49: 301, 60: 301, 301, 301, 64: 301, 66: 301, 70: 301, 72: 301, 301, 76: 301, 301, 301, 301, 83: 301, 85: 301, 116: 301, 301, 119: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 140: 301},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 21: 300, 25: 300, 300, 300, 49: 300, 60: 300, 300, 300, 64: 300, 66: 300, 70: 300, 72: 300, 300, 76: 300, 300, 300, 300, 83: 300, 85: 300, 116: 300, 300, 119: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 140: 300},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 21: 299, 25: 299, 299, 299, 49: 299, 60: 299, 299, 299, 64: 299, 66: 299, 70: 299, 72: 299, 299, 76: 299, 299, 299, 299, 83: 299, 85: 299, 116: 299, 299, 119: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 140: 299},
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 21: 298, 25: 298, 298, 298, 49: 298, 60: 298, 298, 298, 64: 298, 66: 298, 70: 298, 72: 298, 298, 76: 298, 298, 298, 298, 83: 298, 85: 298, 116: 298, 298, 119: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 140: 298},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 21: 297, 25: 297, 297, 297, 49: 297, 60: 297, 297, 297, 64: 297, 66: 297, 70: 297, 72: 297, 297, 76: 297, 297, 297, 297, 83: 297, 85: 297, 116: 297, 297, 119: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 140: 297},
		// 65
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 21: 296, 25: 296, 296, 296, 49: 296, 60: 296, 296, 296, 64: 296, 66: 296, 70: 296, 72: 296, 296, 76: 296, 296, 296, 296, 83: 296, 85: 296, 116: 296, 296, 119: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 140: 296},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 21: 295, 25: 295, 295, 295, 49: 295, 60: 295, 295, 295, 64: 295, 66: 295, 70: 295, 72: 295, 295, 76: 295, 295, 295, 295, 83: 295, 85: 295, 116: 295, 295, 119: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 140: 295},
		{586, 282, 3: 282, 282, 282, 282, 282, 282, 590, 589, 21: 282, 25: 282, 585, 282, 49: 282, 60: 282, 282, 282, 64: 282, 66: 282, 70: 282, 72: 282, 282, 76: 282, 282, 282, 282, 83: 282, 85: 282, 116: 282, 282, 119: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 587, 140: 588},
		{580, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 584, 380},
		{580, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 581, 380},
		// 70
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 579},
		{465, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 466, 380},
		{462, 2: 461},
		{274, 274, 274, 5: 274, 7: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		{273, 273, 273, 5: 273, 7: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273},
		// 75
		{272, 272, 272, 5: 272, 7: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272},
		{271, 271, 271, 5: 271, 7: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271},
		{270, 270, 270, 5: 270, 7: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270},
		{269, 269, 269, 5: 269, 7: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269},
		{1: 268, 3: 268, 268, 268, 268, 268, 268, 21: 268, 25: 268, 27: 268, 49: 268, 60: 268, 268, 268, 64: 268, 66: 268, 70: 268, 72: 268, 268, 76: 268, 268, 268, 268, 83: 268, 85: 268, 116: 268, 268, 119: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		// 80
		{1: 266, 3: 266, 266, 266, 266, 266, 266, 21: 266, 25: 266, 27: 266, 49: 266, 60: 266, 266, 266, 64: 266, 66: 266, 70: 266, 72: 266, 76: 266, 266, 266, 266, 83: 266, 85: 266, 116: 266, 266},
		{1: 449, 3: 262, 262, 262, 262, 262, 262, 21: 262, 25: 262, 27: 262, 49: 262, 60: 262, 262, 262, 64: 262, 66: 262, 70: 262, 72: 262, 76: 262, 262, 262, 262, 83: 262, 85: 262, 116: 451, 450},
		{3: 259, 259, 259, 259, 446, 447, 21: 259, 25: 259, 27: 259, 49: 259, 60: 259, 259, 259, 64: 259, 66: 259, 70: 259, 72: 259, 76: 259, 259, 259, 259, 83: 259, 85: 259},
		{3: 256, 256, 256, 256, 21: 256, 25: 256, 27: 256, 49: 256, 60: 256, 256, 256, 64: 256, 66: 256, 70: 256, 72: 256, 76: 256, 256, 256, 256, 83: 443, 85: 444},
		{3: 251, 251, 251, 251, 21: 251, 25: 251, 27: 251, 49: 251, 60: 251, 251, 251, 64: 251, 66: 251, 70: 251, 72: 251, 76: 438, 439, 441, 440},
		// 85
		{3: 248, 248, 248, 248, 21: 248, 25: 248, 27: 248, 49: 248, 60: 248, 248, 248, 64: 248, 66: 248, 70: 435, 72: 436},
		{3: 246, 246, 433, 246, 21: 246, 25: 246, 27: 246, 49: 246, 60: 246, 246, 246, 64: 246, 66: 246},
		{3: 244, 244, 6: 244, 21: 244, 25: 244, 27: 244, 49: 244, 60: 244, 244, 244, 64: 244, 66: 431},
		{3: 242, 242, 6: 242, 21: 242, 25: 242, 27: 242, 49: 242, 60: 242, 242, 242, 64: 429},
		{3: 240, 240, 6: 240, 21: 240, 25: 240, 27: 240, 49: 240, 60: 240, 240, 427},
		// 90
		{3: 238, 238, 6: 238, 21: 238, 25: 238, 27: 238, 49: 238, 60: 404, 403},
		{4: 217, 6: 217, 21: 217, 25: 217, 27: 217, 49: 217},
		{6: 157, 21: 157},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 426},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 406},
		// 95
		{1: 268, 3: 268, 268, 268, 268, 268, 268, 21: 268, 25: 268, 49: 268, 60: 268, 268, 268, 64: 268, 66: 268, 70: 268, 72: 268, 414, 76: 268, 268, 268, 268, 83: 268, 85: 268, 116: 268, 268, 119: 418, 422, 416, 420, 417, 415, 424, 421, 419, 423, 213: 413},
		{6: 410, 25: 409},
		{3: 236, 236, 6: 236, 21: 236, 25: 236, 49: 236},
		{3: 221, 221, 6: 221, 25: 221, 49: 221},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 412, 400},
		// 100
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 411},
		{3: 220, 220, 6: 220, 25: 220, 49: 220},
		{3: 237, 237, 6: 237, 21: 237, 25: 237, 27: 237, 49: 237},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 425},
		{232, 232, 232, 5: 232, 7: 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232},
		// 105
		{231, 231, 231, 5: 231, 7: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231},
		{230, 230, 230, 5: 230, 7: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230},
		{229, 229, 229, 5: 229, 7: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229},
		{228, 228, 228, 5: 228, 7: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228},
		{227, 227, 227, 5: 227, 7: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		// 110
		{226, 226, 226, 5: 226, 7: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{225, 225, 225, 5: 225, 7: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		{224, 224, 224, 5: 224, 7: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 5: 223, 7: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 5: 222, 7: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		// 115
		{3: 235, 235, 6: 235, 21: 235, 25: 235, 49: 235},
		{3: 239, 239, 6: 239, 21: 239, 25: 239, 27: 239, 49: 239, 60: 239, 239, 427},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 428},
		{3: 241, 241, 6: 241, 21: 241, 25: 241, 27: 241, 49: 241, 60: 241, 241, 241, 64: 429},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 430},
		// 120
		{3: 243, 243, 6: 243, 21: 243, 25: 243, 27: 243, 49: 243, 60: 243, 243, 243, 64: 243, 66: 431},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 432},
		{3: 245, 245, 433, 245, 21: 245, 25: 245, 27: 245, 49: 245, 60: 245, 245, 245, 64: 245, 66: 245},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 434},
		{3: 247, 247, 247, 247, 21: 247, 25: 247, 27: 247, 49: 247, 60: 247, 247, 247, 64: 247, 66: 247, 70: 435, 72: 436},
		// 125
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 460},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 437},
		{3: 249, 249, 249, 249, 21: 249, 25: 249, 27: 249, 49: 249, 60: 249, 249, 249, 64: 249, 66: 249, 70: 249, 72: 249, 76: 438, 439, 441, 440},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 459},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 458},
		// 130
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 457},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 442},
		{3: 252, 252, 252, 252, 21: 252, 25: 252, 27: 252, 49: 252, 60: 252, 252, 252, 64: 252, 66: 252, 70: 252, 72: 252, 76: 252, 252, 252, 252, 83: 443, 85: 444},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 456},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 445},
		// 135
		{3: 257, 257, 257, 257, 446, 447, 21: 257, 25: 257, 27: 257, 49: 257, 60: 257, 257, 257, 64: 257, 66: 257, 70: 257, 72: 257, 76: 257, 257, 257, 257, 83: 257, 85: 257},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 455},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 448},
		{1: 449, 3: 260, 260, 260, 260, 260, 260, 21: 260, 25: 260, 27: 260, 49: 260, 60: 260, 260, 260, 64: 260, 66: 260, 70: 260, 72: 260, 76: 260, 260, 260, 260, 83: 260, 85: 260, 116: 451, 450},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 454},
		// 140
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 453},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 452},
		{1: 263, 3: 263, 263, 263, 263, 263, 263, 21: 263, 25: 263, 27: 263, 49: 263, 60: 263, 263, 263, 64: 263, 66: 263, 70: 263, 72: 263, 76: 263, 263, 263, 263, 83: 263, 85: 263, 116: 263, 263},
		{1: 264, 3: 264, 264, 264, 264, 264, 264, 21: 264, 25: 264, 27: 264, 49: 264, 60: 264, 264, 264, 64: 264, 66: 264, 70: 264, 72: 264, 76: 264, 264, 264, 264, 83: 264, 85: 264, 116: 264, 264},
		{1: 265, 3: 265, 265, 265, 265, 265, 265, 21: 265, 25: 265, 27: 265, 49: 265, 60: 265, 265, 265, 64: 265, 66: 265, 70: 265, 72: 265, 76: 265, 265, 265, 265, 83: 265, 85: 265, 116: 265, 265},
		// 145
		{1: 449, 3: 261, 261, 261, 261, 261, 261, 21: 261, 25: 261, 27: 261, 49: 261, 60: 261, 261, 261, 64: 261, 66: 261, 70: 261, 72: 261, 76: 261, 261, 261, 261, 83: 261, 85: 261, 116: 451, 450},
		{3: 258, 258, 258, 258, 446, 447, 21: 258, 25: 258, 27: 258, 49: 258, 60: 258, 258, 258, 64: 258, 66: 258, 70: 258, 72: 258, 76: 258, 258, 258, 258, 83: 258, 85: 258},
		{3: 253, 253, 253, 253, 21: 253, 25: 253, 27: 253, 49: 253, 60: 253, 253, 253, 64: 253, 66: 253, 70: 253, 72: 253, 76: 253, 253, 253, 253, 83: 443, 85: 444},
		{3: 254, 254, 254, 254, 21: 254, 25: 254, 27: 254, 49: 254, 60: 254, 254, 254, 64: 254, 66: 254, 70: 254, 72: 254, 76: 254, 254, 254, 254, 83: 443, 85: 444},
		{3: 255, 255, 255, 255, 21: 255, 25: 255, 27: 255, 49: 255, 60: 255, 255, 255, 64: 255, 66: 255, 70: 255, 72: 255, 76: 255, 255, 255, 255, 83: 443, 85: 444},
		// 150
		{3: 250, 250, 250, 250, 21: 250, 25: 250, 27: 250, 49: 250, 60: 250, 250, 250, 64: 250, 66: 250, 70: 250, 72: 250, 76: 438, 439, 441, 440},
		{1: 276, 3: 276, 276, 276, 276, 276, 276, 21: 276, 25: 276, 27: 276, 49: 276, 60: 276, 276, 276, 64: 276, 66: 276, 70: 276, 72: 276, 276, 76: 276, 276, 276, 276, 83: 276, 85: 276, 116: 276, 276, 119: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{2: 463},
		{3: 464},
		{1: 275, 3: 275, 275, 275, 275, 275, 275, 21: 275, 25: 275, 27: 275, 49: 275, 60: 275, 275, 275, 64: 275, 66: 275, 70: 275, 72: 275, 275, 76: 275, 275, 275, 275, 83: 275, 85: 275, 116: 275, 275, 119: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		// 155
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 118: 470, 132: 467, 346, 353, 348, 345, 349, 469, 162: 471, 184: 468},
		{1: 278, 3: 278, 278, 278, 278, 278, 278, 21: 278, 25: 278, 27: 278, 49: 278, 60: 278, 278, 278, 64: 278, 66: 278, 70: 278, 72: 278, 278, 76: 278, 278, 278, 278, 83: 278, 85: 278, 116: 278, 278, 119: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{3: 578, 6: 410},
		{3: 552},
		{175, 175, 175, 175, 175, 22: 175, 175, 175, 175, 175, 28: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 207: 550},
		// 160
		{172, 172, 172, 172, 172, 22: 354, 355, 356, 172, 172, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 118: 470, 133: 346, 353, 348, 345, 349, 469, 162: 549, 203: 548},
		{137, 472, 3: 114, 26: 137, 153: 473, 475, 172: 476, 185: 474},
		{133, 133, 133, 133, 6: 133, 22: 354, 355, 356, 26: 133, 118: 483, 171: 487, 179: 546},
		{136, 2: 136, 116, 6: 116, 26: 136},
		{3: 117},
		// 165
		{478, 26: 104, 188: 477, 479},
		{3: 113, 6: 113},
		{543, 3: 115, 6: 115, 26: 103},
		{137, 472, 3: 129, 22: 354, 355, 356, 26: 137, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 502, 328, 325, 153: 473, 475, 170: 501, 172: 503, 177: 499, 500, 183: 504},
		{26: 480},
		// 170
		{369, 481, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 22: 354, 355, 356, 43: 486, 49: 234, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 482, 118: 483, 171: 484, 180: 485},
		{273, 273, 273, 5: 273, 7: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 49: 498},
		{49: 233},
		{135, 135, 135, 135, 5: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 22: 135, 135, 135, 26: 135, 43: 135, 49: 135},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 22: 354, 355, 356, 43: 494, 49: 234, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 482, 118: 491, 180: 493},
		// 175
		{49: 492},
		{133, 133, 133, 5: 133, 7: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 22: 354, 355, 356, 118: 483, 171: 487, 179: 488},
		{132, 132, 132, 132, 5: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 22: 354, 355, 356, 26: 132, 118: 491},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 489},
		{49: 490},
		// 180
		{109, 3: 109, 6: 109, 26: 109},
		{134, 134, 134, 134, 5: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 22: 134, 134, 134, 26: 134, 43: 134, 49: 134},
		{111, 3: 111, 6: 111, 26: 111},
		{49: 497},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 495},
		// 185
		{49: 496},
		{108, 3: 108, 6: 108, 26: 108},
		{110, 3: 110, 6: 110, 26: 110},
		{107, 3: 107, 6: 107, 26: 107},
		{3: 131, 6: 540},
		// 190
		{3: 128},
		{3: 127, 6: 127},
		{137, 472, 137, 114, 6: 114, 26: 137, 153: 473, 507, 156: 508, 172: 476, 185: 509},
		{3: 506},
		{3: 505},
		// 195
		{106, 3: 106, 6: 106, 26: 106},
		{112, 3: 112, 6: 112, 26: 112},
		{512, 2: 511, 26: 104, 188: 477, 479, 510},
		{3: 125, 6: 125},
		{3: 124, 6: 124},
		// 200
		{516, 3: 152, 152, 6: 152, 22: 152, 152, 152, 152, 515, 28: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 50: 152, 73: 152},
		{149, 3: 149, 149, 6: 149, 22: 149, 149, 149, 149, 149, 28: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 50: 149, 73: 149},
		{137, 472, 137, 129, 22: 354, 355, 356, 26: 137, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 502, 328, 325, 153: 473, 507, 156: 513, 170: 501, 172: 503, 177: 499, 500, 183: 504},
		{3: 514},
		{148, 3: 148, 148, 6: 148, 22: 148, 148, 148, 148, 148, 28: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 50: 148, 73: 148},
		// 205
		{133, 133, 133, 5: 133, 7: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 22: 354, 355, 356, 43: 528, 49: 133, 118: 483, 171: 529, 179: 527},
		{2: 143, 143, 22: 143, 143, 143, 28: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 208: 517},
		{2: 521, 121, 22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 502, 328, 325, 170: 501, 177: 499, 519, 196: 522, 520, 219: 518},
		{142, 3: 142, 142, 6: 142, 22: 142, 142, 142, 142, 142, 28: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 50: 142, 73: 142},
		{3: 526},
		// 210
		{3: 525},
		{3: 123, 6: 123},
		{3: 120, 6: 523},
		{2: 524},
		{3: 122, 6: 122},
		// 215
		{140, 3: 140, 140, 6: 140, 22: 140, 140, 140, 140, 140, 28: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 50: 140, 73: 140},
		{141, 3: 141, 141, 6: 141, 22: 141, 141, 141, 141, 141, 28: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 50: 141, 73: 141},
		{369, 536, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 49: 234, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 482, 180: 537},
		{133, 133, 133, 5: 133, 7: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 22: 354, 355, 356, 118: 483, 171: 487, 179: 533},
		{132, 132, 132, 5: 132, 7: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 22: 354, 355, 356, 43: 530, 49: 132, 118: 491},
		// 220
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 531},
		{49: 532},
		{145, 3: 145, 145, 6: 145, 22: 145, 145, 145, 145, 145, 28: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 50: 145, 73: 145},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 534},
		{49: 535},
		// 225
		{146, 3: 146, 146, 6: 146, 22: 146, 146, 146, 146, 146, 28: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 50: 146, 73: 146},
		{273, 273, 273, 5: 273, 7: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 49: 539},
		{49: 538},
		{147, 3: 147, 147, 6: 147, 22: 147, 147, 147, 147, 147, 28: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 50: 147, 73: 147},
		{144, 3: 144, 144, 6: 144, 22: 144, 144, 144, 144, 144, 28: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 50: 144, 73: 144},
		// 230
		{22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 502, 328, 325, 166: 541, 170: 542},
		{3: 130},
		{3: 126, 6: 126},
		{3: 129, 22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 502, 328, 325, 170: 501, 177: 499, 500, 183: 544},
		{3: 545},
		// 235
		{105, 3: 105, 6: 105, 26: 105},
		{139, 472, 139, 139, 6: 139, 26: 139, 153: 547},
		{138, 2: 138, 138, 6: 138, 26: 138},
		{173, 173, 173, 173, 173, 25: 173, 173},
		{171, 171, 171, 171, 171, 25: 171, 171},
		// 240
		{172, 172, 172, 172, 172, 22: 354, 355, 356, 172, 172, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 118: 470, 133: 346, 353, 348, 345, 349, 469, 162: 549, 203: 551},
		{174, 174, 174, 174, 174, 25: 174, 174},
		{1: 277, 3: 277, 277, 277, 277, 277, 277, 21: 277, 25: 277, 27: 277, 49: 277, 553, 60: 277, 277, 277, 64: 277, 66: 277, 70: 277, 72: 277, 277, 76: 277, 277, 277, 277, 83: 277, 85: 277, 116: 277, 277, 119: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{96, 96, 96, 5: 96, 7: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 26: 559, 50: 96, 129: 560, 167: 558, 174: 557, 555, 556, 201: 554},
		{6: 576, 21: 575},
		// 245
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 567, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 566, 182: 568},
		{26: 559, 73: 564, 129: 560, 167: 565},
		{95, 95, 95, 5: 95, 7: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 50: 95},
		{26: 94, 73: 94, 129: 94},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 401, 400, 165: 562},
		// 250
		{2: 561},
		{26: 91, 73: 91, 129: 91},
		{49: 563},
		{26: 92, 73: 92, 129: 92},
		{97, 97, 97, 5: 97, 7: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 50: 97},
		// 255
		{26: 93, 73: 93, 129: 93},
		{4: 102, 6: 102, 21: 102},
		{96, 96, 96, 5: 96, 7: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 26: 559, 50: 96, 129: 560, 167: 558, 174: 557, 555, 556, 201: 569},
		{6: 99, 21: 99},
		{6: 571, 21: 570},
		// 260
		{4: 101, 6: 101, 21: 101},
		{96, 96, 96, 5: 96, 7: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 572, 26: 559, 50: 96, 129: 560, 167: 558, 174: 557, 573, 556},
		{4: 100, 6: 100, 21: 100},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 567, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 566, 182: 574},
		{6: 98, 21: 98},
		// 265
		{288, 288, 3: 288, 288, 288, 288, 288, 288, 288, 288, 21: 288, 25: 288, 288, 288, 49: 288, 60: 288, 288, 288, 64: 288, 66: 288, 70: 288, 72: 288, 288, 76: 288, 288, 288, 288, 83: 288, 85: 288, 116: 288, 288, 119: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 140: 288},
		{96, 96, 96, 5: 96, 7: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 577, 26: 559, 50: 96, 129: 560, 167: 558, 174: 557, 573, 556},
		{287, 287, 3: 287, 287, 287, 287, 287, 287, 287, 287, 21: 287, 25: 287, 287, 287, 49: 287, 60: 287, 287, 287, 64: 287, 66: 287, 70: 287, 72: 287, 287, 76: 287, 287, 287, 287, 83: 287, 85: 287, 116: 287, 287, 119: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 140: 287},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 21: 302, 25: 302, 302, 302, 49: 302, 60: 302, 302, 302, 64: 302, 66: 302, 70: 302, 72: 302, 302, 76: 302, 302, 302, 302, 83: 302, 85: 302, 116: 302, 302, 119: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 140: 302},
		{1: 279, 3: 279, 279, 279, 279, 279, 279, 21: 279, 25: 279, 27: 279, 49: 279, 60: 279, 279, 279, 64: 279, 66: 279, 70: 279, 72: 279, 279, 76: 279, 279, 279, 279, 83: 279, 85: 279, 116: 279, 279, 119: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		// 270
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 118: 470, 132: 467, 346, 353, 348, 345, 349, 469, 162: 471, 184: 582},
		{1: 280, 3: 280, 280, 280, 280, 280, 280, 21: 280, 25: 280, 27: 280, 49: 280, 60: 280, 280, 280, 64: 280, 66: 280, 70: 280, 72: 280, 280, 76: 280, 280, 280, 280, 83: 280, 85: 280, 116: 280, 280, 119: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{3: 583},
		{50: 553},
		{1: 281, 3: 281, 281, 281, 281, 281, 281, 21: 281, 25: 281, 27: 281, 49: 281, 60: 281, 281, 281, 64: 281, 66: 281, 70: 281, 72: 281, 281, 76: 281, 281, 281, 281, 83: 281, 85: 281, 116: 281, 281, 119: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		// 275
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 599},
		{369, 384, 367, 284, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 594, 211: 595, 593},
		{2: 592},
		{2: 591},
		{290, 290, 3: 290, 290, 290, 290, 290, 290, 290, 290, 21: 290, 25: 290, 290, 290, 49: 290, 60: 290, 290, 290, 64: 290, 66: 290, 70: 290, 72: 290, 290, 76: 290, 290, 290, 290, 83: 290, 85: 290, 116: 290, 290, 119: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 140: 290},
		// 280
		{289, 289, 3: 289, 289, 289, 289, 289, 289, 289, 289, 21: 289, 25: 289, 289, 289, 49: 289, 60: 289, 289, 289, 64: 289, 66: 289, 70: 289, 72: 289, 289, 76: 289, 289, 289, 289, 83: 289, 85: 289, 116: 289, 289, 119: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 140: 289},
		{291, 291, 3: 291, 291, 291, 291, 291, 291, 291, 291, 21: 291, 25: 291, 291, 291, 49: 291, 60: 291, 291, 291, 64: 291, 66: 291, 70: 291, 72: 291, 291, 76: 291, 291, 291, 291, 83: 291, 85: 291, 116: 291, 291, 119: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 140: 291},
		{292, 292, 3: 292, 292, 292, 292, 292, 292, 292, 292, 21: 292, 25: 292, 292, 292, 49: 292, 60: 292, 292, 292, 64: 292, 66: 292, 70: 292, 72: 292, 292, 76: 292, 292, 292, 292, 83: 292, 85: 292, 116: 292, 292, 119: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 140: 292},
		{3: 598},
		{3: 286, 6: 286},
		// 285
		{3: 283, 6: 596},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 597},
		{3: 285, 6: 285},
		{293, 293, 3: 293, 293, 293, 293, 293, 293, 293, 293, 21: 293, 25: 293, 293, 293, 49: 293, 60: 293, 293, 293, 64: 293, 66: 293, 70: 293, 72: 293, 293, 76: 293, 293, 293, 293, 83: 293, 85: 293, 116: 293, 293, 119: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 140: 293},
		{6: 410, 49: 600},
		// 290
		{294, 294, 3: 294, 294, 294, 294, 294, 294, 294, 294, 21: 294, 25: 294, 294, 294, 49: 294, 60: 294, 294, 294, 64: 294, 66: 294, 70: 294, 72: 294, 294, 76: 294, 294, 294, 294, 83: 294, 85: 294, 116: 294, 294, 119: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 140: 294},
		{3: 602},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 553, 368, 377, 376, 389, 380, 603},
		{1: 267, 3: 267, 267, 267, 267, 267, 267, 21: 267, 25: 267, 27: 267, 49: 267, 60: 267, 267, 267, 64: 267, 66: 267, 70: 267, 72: 267, 267, 76: 267, 267, 267, 267, 83: 267, 85: 267, 116: 267, 267, 119: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		{163, 163, 163, 163, 163, 6: 163, 22: 163, 163, 163, 163, 163, 28: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163},
		// 295
		{2: 362, 21: 606, 192: 365, 607},
		{162, 162, 162, 162, 162, 6: 162, 22: 162, 162, 162, 162, 162, 28: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162},
		{6: 159, 21: 159},
		{50: 164},
		{161, 161, 161, 161, 161, 6: 161, 22: 161, 161, 161, 161, 161, 28: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 50: 118},
		// 300
		{22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 118: 470, 133: 346, 353, 348, 345, 349, 469, 162: 613, 204: 612, 232: 611},
		{21: 627, 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 118: 470, 133: 346, 353, 348, 345, 349, 469, 162: 613, 204: 628},
		{21: 178, 178, 178, 178, 28: 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178},
		{137, 472, 137, 4: 168, 25: 151, 153: 620, 619, 156: 617, 187: 618, 205: 615, 233: 616, 614},
		{4: 626},
		// 305
		{4: 170, 6: 170},
		{4: 167, 6: 624},
		{4: 166, 6: 166, 25: 150},
		{25: 622},
		{621, 2: 511, 190: 510},
		// 310
		{136, 2: 136},
		{137, 472, 137, 153: 620, 619, 156: 513},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 401, 400, 165: 623},
		{4: 165, 6: 165},
		{137, 472, 137, 25: 151, 153: 620, 619, 156: 617, 187: 618, 205: 625},
		// 315
		{4: 169, 6: 169},
		{21: 176, 176, 176, 176, 28: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176},
		{182, 182, 182, 182, 182, 6: 182, 22: 182, 182, 182, 182, 182, 28: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182},
		{21: 177, 177, 177, 177, 28: 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177},
		{50: 183},
		// 320
		{181, 181, 181, 181, 181, 6: 181, 22: 181, 181, 181, 181, 181, 28: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 50: 118},
		{211, 211, 211, 211, 211, 6: 211, 26: 211},
		{209, 209, 209, 209, 209, 6: 209, 26: 209},
		{212, 212, 212, 212, 212, 6: 212, 26: 212},
		{210, 210, 210, 210, 210, 6: 210, 22: 354, 355, 356, 26: 210, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 632, 328, 325, 173: 635},
		// 325
		{213, 213, 213, 213, 213, 6: 213, 26: 213},
		{215, 215, 215, 215, 215, 6: 215, 26: 215},
		{4: 730},
		{4: 208, 6: 208},
		{4: 205, 6: 728},
		// 330
		{4: 204, 6: 204, 22: 57, 57, 57, 28: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 50: 57, 73: 641, 210: 642},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 567, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 566, 182: 727},
		{22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 50: 53, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 643, 328, 325, 161: 645, 217: 646, 644},
		{137, 472, 137, 4: 206, 153: 620, 619, 156: 726, 181: 638, 199: 639, 637},
		{50: 648, 146: 649},
		// 335
		{22: 55, 55, 55, 28: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 50: 55},
		{22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 50: 52, 118: 327, 133: 346, 353, 348, 345, 349, 326, 143: 643, 328, 325, 161: 647},
		{22: 54, 54, 54, 28: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 50: 54},
		{81, 81, 81, 4: 81, 81, 7: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 28: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 50: 81, 88: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 209: 650},
		{22: 56, 56, 56, 27: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56},
		// 340
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 77, 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 118: 327, 132: 652, 346, 353, 348, 345, 349, 326, 141: 666, 143: 643, 328, 325, 654, 655, 657, 658, 653, 656, 665, 161: 664, 186: 662, 214: 663, 661},
		{304, 304, 4: 304, 304, 304, 304, 304, 304, 304, 25: 724, 304, 60: 304, 304, 304, 64: 304, 66: 304, 70: 304, 72: 304, 304, 76: 304, 304, 304, 304, 83: 304, 85: 304, 116: 304, 304, 119: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 140: 304},
		{3: 218, 218, 6: 410},
		{90, 90, 90, 4: 90, 90, 7: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 28: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 50: 90, 88: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 139: 90},
		{89, 89, 89, 4: 89, 89, 7: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 28: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 50: 89, 88: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 139: 89},
		// 345
		{88, 88, 88, 4: 88, 88, 7: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 28: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 50: 88, 88: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 139: 88},
		{87, 87, 87, 4: 87, 87, 7: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 28: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 50: 87, 88: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 139: 87},
		{86, 86, 86, 4: 86, 86, 7: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 28: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 50: 86, 88: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 139: 86},
		{85, 85, 85, 4: 85, 85, 7: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 28: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 50: 85, 88: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 139: 85},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 389, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 401, 400, 165: 721},
		// 350
		{25: 719},
		{21: 718},
		{79, 79, 79, 4: 79, 79, 7: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 28: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 50: 79, 88: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 76, 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 118: 327, 132: 652, 346, 353, 348, 345, 349, 326, 141: 666, 143: 643, 328, 325, 654, 655, 657, 658, 653, 656, 665, 161: 664, 186: 717},
		{75, 75, 75, 4: 75, 75, 7: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 28: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 50: 75, 88: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75},
		// 355
		{74, 74, 74, 4: 74, 74, 7: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 28: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 50: 74, 88: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74},
		{4: 716},
		{710},
		{706},
		{702},
		// 360
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 696},
		{682},
		{2: 680},
		{4: 679},
		{4: 678},
		// 365
		{369, 384, 367, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 652, 141: 676},
		{4: 677},
		{62, 62, 62, 4: 62, 62, 7: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 28: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 50: 62, 88: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 139: 62},
		{63, 63, 63, 4: 63, 63, 7: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 28: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 50: 63, 88: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 139: 63},
		{64, 64, 64, 4: 64, 64, 7: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 28: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 50: 64, 88: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 139: 64},
		// 370
		{4: 681},
		{65, 65, 65, 4: 65, 65, 7: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 28: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 50: 65, 88: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 139: 65},
		{369, 384, 367, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 22: 354, 355, 356, 28: 343, 335, 344, 340, 352, 339, 337, 338, 336, 341, 350, 347, 351, 342, 334, 331, 332, 330, 357, 333, 329, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 118: 327, 132: 652, 346, 353, 348, 345, 349, 326, 141: 683, 143: 643, 328, 325, 161: 684},
		{4: 690},
		{369, 384, 367, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 652, 141: 685},
		// 375
		{4: 686},
		{369, 384, 367, 219, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 652, 141: 687},
		{3: 688},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 689},
		{66, 66, 66, 4: 66, 66, 7: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 28: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 50: 66, 88: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 139: 66},
		// 380
		{369, 384, 367, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 652, 141: 691},
		{4: 692},
		{369, 384, 367, 219, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 652, 141: 693},
		{3: 694},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 695},
		// 385
		{67, 67, 67, 4: 67, 67, 7: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 28: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 50: 67, 88: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 139: 67},
		{88: 697},
		{698},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 699},
		{3: 700, 6: 410},
		// 390
		{4: 701},
		{68, 68, 68, 4: 68, 68, 7: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 28: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 50: 68, 88: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 139: 68},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 703},
		{3: 704, 6: 410},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 705},
		// 395
		{69, 69, 69, 4: 69, 69, 7: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 28: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 50: 69, 88: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 139: 69},
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 707},
		{3: 708, 6: 410},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 709},
		{70, 70, 70, 4: 70, 70, 7: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 28: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 50: 70, 88: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 139: 70},
		// 400
		{369, 384, 367, 5: 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 51: 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 132: 711},
		{3: 712, 6: 410},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 713},
		{72, 72, 72, 4: 72, 72, 7: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 28: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 50: 72, 88: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 139: 714},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 715},
		// 405
		{71, 71, 71, 4: 71, 71, 7: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 28: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 50: 71, 88: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 139: 71},
		{73, 73, 73, 4: 73, 73, 7: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 28: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 50: 73, 88: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 139: 73},
		{78, 78, 78, 4: 78, 78, 7: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 28: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 50: 78, 88: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78},
		{80, 80, 80, 4: 80, 80, 7: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 27: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 50: 80, 88: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 139: 80},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 720},
		// 410
		{82, 82, 82, 4: 82, 82, 7: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 28: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 50: 82, 88: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 139: 82},
		{25: 722},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 723},
		{83, 83, 83, 4: 83, 83, 7: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 28: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 50: 83, 88: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 139: 83},
		{369, 384, 651, 4: 219, 383, 7: 385, 386, 379, 378, 388, 387, 370, 382, 371, 372, 373, 374, 381, 375, 50: 648, 368, 377, 376, 405, 380, 390, 391, 59: 392, 63: 393, 67: 394, 395, 396, 71: 397, 74: 398, 80: 399, 407, 400, 87: 408, 669, 674, 659, 673, 660, 670, 671, 672, 667, 675, 668, 132: 652, 141: 666, 146: 654, 655, 657, 658, 653, 656, 725},
		// 415
		{84, 84, 84, 4: 84, 84, 7: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 28: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 50: 84, 88: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 139: 84},
		{4: 204, 6: 204, 73: 641},
		{4: 203, 6: 203},
		{137, 472, 137, 153: 620, 619, 156: 726, 181: 729},
		{4: 207, 6: 207},
		// 420
		{216, 216, 216, 4: 216, 216, 7: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 27: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 50: 216, 88: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{22: 60, 60, 60, 27: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60},
		{27: 308},
		{27: 309},
		{27: 51, 58: 761, 65: 759, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 803, 740, 737, 739},
		// 425
		{27: 50, 58: 50, 65: 50, 75: 50, 84: 50, 86: 50, 99: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		{27: 46, 58: 46, 65: 46, 75: 46, 84: 46, 86: 46, 99: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46},
		{27: 45, 58: 45, 65: 45, 75: 45, 84: 45, 86: 45, 99: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45},
		{58: 761, 130: 824, 758},
		{27: 43, 58: 43, 65: 43, 75: 43, 84: 43, 86: 43, 99: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43},
		// 430
		{75: 36, 84: 36, 86: 812, 191: 810, 220: 811, 809},
		{58: 761, 130: 807, 758},
		{2: 804},
		{2: 799},
		{2: 776, 226: 777},
		// 435
		{58: 761, 65: 759, 130: 760, 758, 142: 775},
		{27: 24, 58: 24, 65: 24, 75: 24, 84: 24, 86: 24, 99: 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		{58: 761, 130: 774, 758},
		{58: 761, 130: 773, 758},
		{58: 761, 65: 759, 130: 760, 758, 142: 772},
		// 440
		{2: 770},
		{58: 761, 130: 769, 758},
		{58: 761, 130: 768, 758},
		{58: 761, 130: 767, 758},
		{58: 761, 130: 766, 758},
		// 445
		{58: 761, 130: 765, 758},
		{58: 761, 130: 764, 758},
		{27: 11, 58: 11, 65: 11, 75: 11, 84: 11, 86: 11, 99: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{58: 763, 65: 762},
		{27: 8, 58: 8, 65: 8, 75: 8, 84: 8, 86: 8, 99: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		// 450
		{27: 7, 58: 7, 65: 7, 75: 7, 84: 7, 86: 7, 99: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{58: 6, 65: 6},
		{27: 9, 58: 9, 65: 9, 75: 9, 84: 9, 86: 9, 99: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{58: 5, 65: 5},
		{27: 12, 58: 12, 65: 12, 75: 12, 84: 12, 86: 12, 99: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		// 455
		{27: 13, 58: 13, 65: 13, 75: 13, 84: 13, 86: 13, 99: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{27: 14, 58: 14, 65: 14, 75: 14, 84: 14, 86: 14, 99: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{27: 15, 58: 15, 65: 15, 75: 15, 84: 15, 86: 15, 99: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{27: 16, 58: 16, 65: 16, 75: 16, 84: 16, 86: 16, 99: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{27: 19, 58: 19, 65: 19, 75: 19, 84: 19, 86: 19, 99: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 460
		{65: 771},
		{27: 20, 58: 20, 65: 20, 75: 20, 84: 20, 86: 20, 99: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{27: 21, 58: 21, 65: 21, 75: 21, 84: 21, 86: 21, 99: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{27: 22, 58: 22, 65: 22, 75: 22, 84: 22, 86: 22, 99: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{27: 23, 58: 23, 65: 23, 75: 23, 84: 23, 86: 23, 99: 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		// 465
		{27: 25, 58: 25, 65: 25, 75: 25, 84: 25, 86: 25, 99: 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		{58: 761, 65: 759, 130: 760, 758, 142: 784, 164: 798},
		{2: 778, 121, 166: 780, 196: 779, 781},
		{3: 123, 6: 123, 166: 795},
		{3: 120, 6: 787},
		// 470
		{3: 785},
		{3: 782},
		{58: 761, 65: 759, 130: 760, 758, 142: 784, 164: 783},
		{27: 26, 58: 26, 65: 26, 75: 26, 84: 26, 86: 26, 99: 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{27: 10, 58: 10, 65: 10, 75: 10, 84: 10, 86: 10, 99: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 475
		{58: 761, 65: 759, 130: 760, 758, 142: 784, 164: 786},
		{27: 28, 58: 28, 65: 28, 75: 28, 84: 28, 86: 28, 99: 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{2: 788, 166: 789},
		{3: 122, 6: 122, 166: 792},
		{3: 790},
		// 480
		{58: 761, 65: 759, 130: 760, 758, 142: 784, 164: 791},
		{27: 27, 58: 27, 65: 27, 75: 27, 84: 27, 86: 27, 99: 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{3: 793},
		{58: 761, 65: 759, 130: 760, 758, 142: 784, 164: 794},
		{27: 17, 58: 17, 65: 17, 75: 17, 84: 17, 86: 17, 99: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		// 485
		{3: 796},
		{58: 761, 65: 759, 130: 760, 758, 142: 784, 164: 797},
		{27: 18, 58: 18, 65: 18, 75: 18, 84: 18, 86: 18, 99: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{27: 29, 58: 29, 65: 29, 75: 29, 84: 29, 86: 29, 99: 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		{65: 800},
		// 490
		{58: 761, 65: 759, 75: 48, 84: 48, 86: 48, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 735, 740, 737, 739, 163: 801, 168: 802},
		{58: 761, 65: 759, 75: 47, 84: 47, 86: 47, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 803, 740, 737, 739},
		{75: 39, 84: 39, 86: 39},
		{27: 49, 58: 49, 65: 49, 75: 49, 84: 49, 86: 49, 99: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49},
		{65: 805},
		// 495
		{58: 761, 65: 759, 75: 48, 84: 48, 86: 48, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 735, 740, 737, 739, 163: 801, 168: 806},
		{75: 40, 84: 40, 86: 40},
		{58: 761, 65: 759, 75: 48, 84: 48, 86: 48, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 735, 740, 737, 739, 163: 801, 168: 808},
		{75: 41, 84: 41, 86: 41},
		{75: 32, 84: 817, 222: 818, 816},
		// 500
		{75: 38, 84: 38, 86: 38},
		{75: 35, 84: 35, 86: 812, 191: 815},
		{58: 761, 130: 813, 758},
		{58: 761, 65: 759, 75: 48, 84: 48, 86: 48, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 735, 740, 737, 739, 163: 801, 168: 814},
		{75: 34, 84: 34, 86: 34},
		// 505
		{75: 37, 84: 37, 86: 37},
		{75: 822, 224: 821},
		{65: 819},
		{75: 31},
		{58: 761, 65: 759, 75: 48, 99: 751, 744, 745, 746, 752, 741, 742, 743, 753, 747, 754, 748, 738, 749, 755, 750, 756, 130: 760, 758, 142: 757, 155: 736, 157: 735, 740, 737, 739, 163: 801, 168: 820},
		// 510
		{75: 33},
		{27: 42, 58: 42, 65: 42, 75: 42, 84: 42, 86: 42, 99: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{58: 761, 65: 759, 130: 760, 758, 142: 823},
		{27: 30, 58: 30, 65: 30, 75: 30, 84: 30, 86: 30, 99: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{27: 44, 58: 44, 65: 44, 75: 44, 84: 44, 86: 44, 99: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
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
	const yyError = 238

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
	case 2:
		{
			lx := yylex.(*lexer)
			lx.constExpr = yyS[yypt-0].item.(*ConstantExpression)
		}
	case 3:
		{
			lx := yylex.(*lexer)
			tu := yyS[yypt-0].item.(*TranslationUnit).reverse()
			tu.Declarations = lx.tu.Declarations
			lx.tu = tu
			if compilation.Errors(false) == nil && (lx.scope.Type != ScopeFile || lx.compoundStmt != 0) {
				panic("internal error")
			}
		}
	case 5:
		{
			yyVAL.item = &EnumerationConstant{
				Token: yyS[yypt-0].Token,
			}
		}
	case 6:
		{
			yyVAL.item = &PrimaryExpression{
				Token: yyS[yypt-0].Token,
			}
		}
	case 7:
		{
			yyVAL.item = &PrimaryExpression{
				Case:     1,
				Constant: yyS[yypt-0].item.(*Constant),
			}
		}
	case 8:
		{
			yyVAL.item = &PrimaryExpression{
				Case:           2,
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].item.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 9:
		{
			yyVAL.item = &Constant{
				Token: yyS[yypt-0].Token,
			}
		}
	case 10:
		{
			yyVAL.item = &Constant{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 11:
		{
			yyVAL.item = &Constant{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 12:
		{
			yyVAL.item = &Constant{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 13:
		{
			yyVAL.item = &Constant{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 14:
		{
			yyVAL.item = &Constant{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 15:
		{
			yyVAL.item = &PostfixExpression{
				PrimaryExpression: yyS[yypt-0].item.(*PrimaryExpression),
			}
		}
	case 16:
		{
			yyVAL.item = &PostfixExpression{
				Case:              1,
				PostfixExpression: yyS[yypt-3].item.(*PostfixExpression),
				Token:             yyS[yypt-2].Token,
				ExpressionList:    yyS[yypt-1].item.(*ExpressionList).reverse(),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 17:
		{
			yyVAL.item = &PostfixExpression{
				Case:              2,
				PostfixExpression: yyS[yypt-3].item.(*PostfixExpression),
				Token:             yyS[yypt-2].Token,
				ArgumentExpressionListOpt: yyS[yypt-1].item.(*ArgumentExpressionListOpt),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 18:
		{
			yyVAL.item = &PostfixExpression{
				Case:              3,
				PostfixExpression: yyS[yypt-2].item.(*PostfixExpression),
				Token:             yyS[yypt-1].Token,
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 19:
		{
			yyVAL.item = &PostfixExpression{
				Case:              4,
				PostfixExpression: yyS[yypt-2].item.(*PostfixExpression),
				Token:             yyS[yypt-1].Token,
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 20:
		{
			yyVAL.item = &PostfixExpression{
				Case:              5,
				PostfixExpression: yyS[yypt-1].item.(*PostfixExpression),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 21:
		{
			yyVAL.item = &PostfixExpression{
				Case:              6,
				PostfixExpression: yyS[yypt-1].item.(*PostfixExpression),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 22:
		{
			yyVAL.item = &PostfixExpression{
				Case:            7,
				Token:           yyS[yypt-5].Token,
				TypeName:        yyS[yypt-4].item.(*TypeName),
				Token2:          yyS[yypt-3].Token,
				Token3:          yyS[yypt-2].Token,
				InitializerList: yyS[yypt-1].item.(*InitializerList).reverse(),
				Token4:          yyS[yypt-0].Token,
			}
		}
	case 23:
		{
			yyVAL.item = &PostfixExpression{
				Case:            8,
				Token:           yyS[yypt-6].Token,
				TypeName:        yyS[yypt-5].item.(*TypeName),
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].item.(*InitializerList).reverse(),
				Token4:          yyS[yypt-1].Token,
				Token5:          yyS[yypt-0].Token,
			}
		}
	case 24:
		{
			yyVAL.item = &ArgumentExpressionList{
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 25:
		{
			yyVAL.item = &ArgumentExpressionList{
				Case: 1,
				ArgumentExpressionList: yyS[yypt-2].item.(*ArgumentExpressionList),
				Token:                yyS[yypt-1].Token,
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 26:
		{
			yyVAL.item = (*ArgumentExpressionListOpt)(nil)
		}
	case 27:
		{
			yyVAL.item = &ArgumentExpressionListOpt{
				ArgumentExpressionList: yyS[yypt-0].item.(*ArgumentExpressionList).reverse(),
			}
		}
	case 28:
		{
			yyVAL.item = &UnaryExpression{
				PostfixExpression: yyS[yypt-0].item.(*PostfixExpression),
			}
		}
	case 29:
		{
			yyVAL.item = &UnaryExpression{
				Case:            1,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].item.(*UnaryExpression),
			}
		}
	case 30:
		{
			yyVAL.item = &UnaryExpression{
				Case:            2,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].item.(*UnaryExpression),
			}
		}
	case 31:
		{
			yyVAL.item = &UnaryExpression{
				Case:           3,
				UnaryOperator:  yyS[yypt-1].item.(*UnaryOperator),
				CastExpression: yyS[yypt-0].item.(*CastExpression),
			}
		}
	case 32:
		{
			yyVAL.item = &UnaryExpression{
				Case:            4,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].item.(*UnaryExpression),
			}
		}
	case 33:
		{
			yyVAL.item = &UnaryExpression{
				Case:     5,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].item.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 34:
		{
			yyVAL.item = &UnaryExpression{
				Case:   6,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 35:
		{
			yyVAL.item = &UnaryExpression{
				Case:   7,
				Token:  yyS[yypt-3].Token,
				Token2: yyS[yypt-2].Token,
				Token3: yyS[yypt-1].Token,
				Token4: yyS[yypt-0].Token,
			}
		}
	case 36:
		{
			yyVAL.item = &UnaryOperator{
				Token: yyS[yypt-0].Token,
			}
		}
	case 37:
		{
			yyVAL.item = &UnaryOperator{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 38:
		{
			yyVAL.item = &UnaryOperator{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 39:
		{
			yyVAL.item = &UnaryOperator{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 40:
		{
			yyVAL.item = &UnaryOperator{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 41:
		{
			yyVAL.item = &UnaryOperator{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 42:
		{
			yyVAL.item = &CastExpression{
				UnaryExpression: yyS[yypt-0].item.(*UnaryExpression),
			}
		}
	case 43:
		{
			yyVAL.item = &CastExpression{
				Case:           1,
				Token:          yyS[yypt-3].Token,
				TypeName:       yyS[yypt-2].item.(*TypeName),
				Token2:         yyS[yypt-1].Token,
				CastExpression: yyS[yypt-0].item.(*CastExpression),
			}
		}
	case 44:
		{
			yyVAL.item = &MultiplicativeExpression{
				CastExpression: yyS[yypt-0].item.(*CastExpression),
			}
		}
	case 45:
		{
			yyVAL.item = &MultiplicativeExpression{
				Case: 1,
				MultiplicativeExpression: yyS[yypt-2].item.(*MultiplicativeExpression),
				Token:          yyS[yypt-1].Token,
				CastExpression: yyS[yypt-0].item.(*CastExpression),
			}
		}
	case 46:
		{
			yyVAL.item = &MultiplicativeExpression{
				Case: 2,
				MultiplicativeExpression: yyS[yypt-2].item.(*MultiplicativeExpression),
				Token:          yyS[yypt-1].Token,
				CastExpression: yyS[yypt-0].item.(*CastExpression),
			}
		}
	case 47:
		{
			yyVAL.item = &MultiplicativeExpression{
				Case: 3,
				MultiplicativeExpression: yyS[yypt-2].item.(*MultiplicativeExpression),
				Token:          yyS[yypt-1].Token,
				CastExpression: yyS[yypt-0].item.(*CastExpression),
			}
		}
	case 48:
		{
			yyVAL.item = &AdditiveExpression{
				MultiplicativeExpression: yyS[yypt-0].item.(*MultiplicativeExpression),
			}
		}
	case 49:
		{
			yyVAL.item = &AdditiveExpression{
				Case:               1,
				AdditiveExpression: yyS[yypt-2].item.(*AdditiveExpression),
				Token:              yyS[yypt-1].Token,
				MultiplicativeExpression: yyS[yypt-0].item.(*MultiplicativeExpression),
			}
		}
	case 50:
		{
			yyVAL.item = &AdditiveExpression{
				Case:               2,
				AdditiveExpression: yyS[yypt-2].item.(*AdditiveExpression),
				Token:              yyS[yypt-1].Token,
				MultiplicativeExpression: yyS[yypt-0].item.(*MultiplicativeExpression),
			}
		}
	case 51:
		{
			yyVAL.item = &ShiftExpression{
				AdditiveExpression: yyS[yypt-0].item.(*AdditiveExpression),
			}
		}
	case 52:
		{
			yyVAL.item = &ShiftExpression{
				Case:               1,
				ShiftExpression:    yyS[yypt-2].item.(*ShiftExpression),
				Token:              yyS[yypt-1].Token,
				AdditiveExpression: yyS[yypt-0].item.(*AdditiveExpression),
			}
		}
	case 53:
		{
			yyVAL.item = &ShiftExpression{
				Case:               2,
				ShiftExpression:    yyS[yypt-2].item.(*ShiftExpression),
				Token:              yyS[yypt-1].Token,
				AdditiveExpression: yyS[yypt-0].item.(*AdditiveExpression),
			}
		}
	case 54:
		{
			yyVAL.item = &RelationalExpression{
				ShiftExpression: yyS[yypt-0].item.(*ShiftExpression),
			}
		}
	case 55:
		{
			yyVAL.item = &RelationalExpression{
				Case:                 1,
				RelationalExpression: yyS[yypt-2].item.(*RelationalExpression),
				Token:                yyS[yypt-1].Token,
				ShiftExpression:      yyS[yypt-0].item.(*ShiftExpression),
			}
		}
	case 56:
		{
			yyVAL.item = &RelationalExpression{
				Case:                 2,
				RelationalExpression: yyS[yypt-2].item.(*RelationalExpression),
				Token:                yyS[yypt-1].Token,
				ShiftExpression:      yyS[yypt-0].item.(*ShiftExpression),
			}
		}
	case 57:
		{
			yyVAL.item = &RelationalExpression{
				Case:                 3,
				RelationalExpression: yyS[yypt-2].item.(*RelationalExpression),
				Token:                yyS[yypt-1].Token,
				ShiftExpression:      yyS[yypt-0].item.(*ShiftExpression),
			}
		}
	case 58:
		{
			yyVAL.item = &RelationalExpression{
				Case:                 4,
				RelationalExpression: yyS[yypt-2].item.(*RelationalExpression),
				Token:                yyS[yypt-1].Token,
				ShiftExpression:      yyS[yypt-0].item.(*ShiftExpression),
			}
		}
	case 59:
		{
			yyVAL.item = &EqualityExpression{
				RelationalExpression: yyS[yypt-0].item.(*RelationalExpression),
			}
		}
	case 60:
		{
			yyVAL.item = &EqualityExpression{
				Case:                 1,
				EqualityExpression:   yyS[yypt-2].item.(*EqualityExpression),
				Token:                yyS[yypt-1].Token,
				RelationalExpression: yyS[yypt-0].item.(*RelationalExpression),
			}
		}
	case 61:
		{
			yyVAL.item = &EqualityExpression{
				Case:                 2,
				EqualityExpression:   yyS[yypt-2].item.(*EqualityExpression),
				Token:                yyS[yypt-1].Token,
				RelationalExpression: yyS[yypt-0].item.(*RelationalExpression),
			}
		}
	case 62:
		{
			yyVAL.item = &AndExpression{
				EqualityExpression: yyS[yypt-0].item.(*EqualityExpression),
			}
		}
	case 63:
		{
			yyVAL.item = &AndExpression{
				Case:               1,
				AndExpression:      yyS[yypt-2].item.(*AndExpression),
				Token:              yyS[yypt-1].Token,
				EqualityExpression: yyS[yypt-0].item.(*EqualityExpression),
			}
		}
	case 64:
		{
			yyVAL.item = &ExclusiveOrExpression{
				AndExpression: yyS[yypt-0].item.(*AndExpression),
			}
		}
	case 65:
		{
			yyVAL.item = &ExclusiveOrExpression{
				Case: 1,
				ExclusiveOrExpression: yyS[yypt-2].item.(*ExclusiveOrExpression),
				Token:         yyS[yypt-1].Token,
				AndExpression: yyS[yypt-0].item.(*AndExpression),
			}
		}
	case 66:
		{
			yyVAL.item = &InclusiveOrExpression{
				ExclusiveOrExpression: yyS[yypt-0].item.(*ExclusiveOrExpression),
			}
		}
	case 67:
		{
			yyVAL.item = &InclusiveOrExpression{
				Case: 1,
				InclusiveOrExpression: yyS[yypt-2].item.(*InclusiveOrExpression),
				Token: yyS[yypt-1].Token,
				ExclusiveOrExpression: yyS[yypt-0].item.(*ExclusiveOrExpression),
			}
		}
	case 68:
		{
			yyVAL.item = &LogicalAndExpression{
				InclusiveOrExpression: yyS[yypt-0].item.(*InclusiveOrExpression),
			}
		}
	case 69:
		{
			yyVAL.item = &LogicalAndExpression{
				Case:                 1,
				LogicalAndExpression: yyS[yypt-2].item.(*LogicalAndExpression),
				Token:                yyS[yypt-1].Token,
				InclusiveOrExpression: yyS[yypt-0].item.(*InclusiveOrExpression),
			}
		}
	case 70:
		{
			yyVAL.item = &LogicalOrExpression{
				LogicalAndExpression: yyS[yypt-0].item.(*LogicalAndExpression),
			}
		}
	case 71:
		{
			yyVAL.item = &LogicalOrExpression{
				Case:                 1,
				LogicalOrExpression:  yyS[yypt-2].item.(*LogicalOrExpression),
				Token:                yyS[yypt-1].Token,
				LogicalAndExpression: yyS[yypt-0].item.(*LogicalAndExpression),
			}
		}
	case 72:
		{
			yyVAL.item = &ConditionalExpression{
				LogicalOrExpression: yyS[yypt-0].item.(*LogicalOrExpression),
			}
		}
	case 73:
		{
			yyVAL.item = &ConditionalExpression{
				Case:                  1,
				LogicalOrExpression:   yyS[yypt-4].item.(*LogicalOrExpression),
				Token:                 yyS[yypt-3].Token,
				ExpressionList:        yyS[yypt-2].item.(*ExpressionList).reverse(),
				Token2:                yyS[yypt-1].Token,
				ConditionalExpression: yyS[yypt-0].item.(*ConditionalExpression),
			}
		}
	case 74:
		{
			yyVAL.item = &AssignmentExpression{
				ConditionalExpression: yyS[yypt-0].item.(*ConditionalExpression),
			}
		}
	case 75:
		{
			yyVAL.item = &AssignmentExpression{
				Case:                 1,
				UnaryExpression:      yyS[yypt-2].item.(*UnaryExpression),
				AssignmentOperator:   yyS[yypt-1].item.(*AssignmentOperator),
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 76:
		{
			yyVAL.item = (*AssignmentExpressionOpt)(nil)
		}
	case 77:
		{
			yyVAL.item = &AssignmentExpressionOpt{
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 78:
		{
			yyVAL.item = &AssignmentOperator{
				Token: yyS[yypt-0].Token,
			}
		}
	case 79:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 80:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 81:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 82:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 83:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 84:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
		}
	case 85:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
		}
	case 86:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
		}
	case 87:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
		}
	case 88:
		{
			yyVAL.item = &AssignmentOperator{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
		}
	case 89:
		{
			yyVAL.item = &ExpressionList{
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 90:
		{
			yyVAL.item = &ExpressionList{
				Case:                 1,
				ExpressionList:       yyS[yypt-2].item.(*ExpressionList),
				Token:                yyS[yypt-1].Token,
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 91:
		{
			yyVAL.item = (*ExpressionOpt)(nil)
		}
	case 92:
		{
			yyVAL.item = &ExpressionOpt{
				ExpressionList: yyS[yypt-0].item.(*ExpressionList).reverse(),
			}
		}
	case 93:
		{
			yyVAL.item = &ConstantExpression{
				ConditionalExpression: yyS[yypt-0].item.(*ConditionalExpression),
			}
		}
	case 94:
		{
			lx := yylex.(*lexer)
			lhs := &Declaration{
				DeclarationSpecifiers: yyS[yypt-2].item.(*DeclarationSpecifiers),
				InitDeclaratorListOpt: yyS[yypt-1].item.(*InitDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.IsFileScope = lx.scope.Type == ScopeFile
			sc := lx.scope
			sc.isTypedef = false
			o := lhs.InitDeclaratorListOpt
			if o == nil {
				break
			}

			for l := o.InitDeclaratorList; l != nil; l = l.InitDeclaratorList {
				d := l.InitDeclarator.Declarator
				d.DeclarationSpecifiers = lhs.DeclarationSpecifiers
				lhs.IsTypedef = d.IsTypedef
			}
		}
	case 95:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				StorageClassSpecifier:    yyS[yypt-1].item.(*StorageClassSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].item.(*DeclarationSpecifiersOpt),
			}
			yyVAL.item = lhs
			lhs.sum(lhs.DeclarationSpecifiersOpt)
			switch lhs.StorageClassSpecifier.Case {
			case 0: // "typedef"
				lhs.IsTypedef = true
			case 1: // "extern"
				lhs.IsExtern = true
			case 2: // "static"
				lhs.IsStatic = true
			case 3: // "auto"
				lhs.IsAuto = true
			case 4: // "register"
				lhs.IsRegister = true
			default:
				panic("internal error")
			}
			lx.scope.specifier = (*declarationSpecifiers)(lhs).Type()
		}
	case 96:
		{
			lx := yylex.(*lexer)
			yyS[yypt-0].item.(*TypeSpecifier).bindings = lx.scope
		}
	case 97:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     1,
				TypeSpecifier:            yyS[yypt-2].item.(*TypeSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].item.(*DeclarationSpecifiersOpt),
			}
			yyVAL.item = lhs
			lhs.sum(lhs.DeclarationSpecifiersOpt)
			ts := lhs.TypeSpecifier
			if lhs.typeSpecifiers > 0xffffff {
				compilation.ErrTok(ts.Token, "invalid type specifier")
				lhs.typ = tsVoid
				break
			}

			c := ts.Case
			if c == tsStructOrUnion {
				c = ts.case2
			}
			lhs.typeSpecifiers = lhs.typeSpecifiers<<8 | c
			lhs.typeSum()
			lx.scope.specifier = (*declarationSpecifiers)(lhs).Type()
		}
	case 98:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     2,
				TypeQualifier:            yyS[yypt-1].item.(*TypeQualifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].item.(*DeclarationSpecifiersOpt),
			}
			yyVAL.item = lhs
			lhs.sum(lhs.DeclarationSpecifiersOpt)
			switch lhs.TypeQualifier.Case {
			case 0: // "const"
				lhs.IsConst = true
			case 1: // "restrict"
				lhs.IsRestrict = true
			case 2: // "volatile"
				lhs.IsVolatile = true
			default:
				panic("internal error")
			}
			lx.scope.specifier = (*declarationSpecifiers)(lhs).Type()
		}
	case 99:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     3,
				FunctionSpecifier:        yyS[yypt-1].item.(*FunctionSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].item.(*DeclarationSpecifiersOpt),
			}
			yyVAL.item = lhs
			lhs.sum(lhs.DeclarationSpecifiersOpt)
			lhs.IsInline = true
			lx.scope.specifier = (*declarationSpecifiers)(lhs).Type()
		}
	case 100:
		{
			yyVAL.item = (*DeclarationSpecifiersOpt)(nil)
		}
	case 101:
		{
			yyVAL.item = &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].item.(*DeclarationSpecifiers),
			}
		}
	case 102:
		{
			yyVAL.item = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].item.(*InitDeclarator),
			}
		}
	case 103:
		{
			yyVAL.item = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].item.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].item.(*InitDeclarator),
			}
		}
	case 104:
		{
			yyVAL.item = (*InitDeclaratorListOpt)(nil)
		}
	case 105:
		{
			yyVAL.item = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].item.(*InitDeclaratorList).reverse(),
			}
		}
	case 106:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].item.(*Declarator),
			}
			yyVAL.item = lhs
			lhs.Declarator.insert(lx.scope, NSIdentifiers, false)
		}
	case 107:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Case:        1,
				Declarator:  yyS[yypt-2].item.(*Declarator),
				Token:       yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].item.(*Initializer),
			}
			yyVAL.item = lhs
			d := lhs.Declarator
			d.Initializer = lhs.Initializer
			d.insert(lx.scope, NSIdentifiers, true)
		}
	case 108:
		{
			lx := yylex.(*lexer)
			yyVAL.item = &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			lx.scope.isTypedef = true
		}
	case 109:
		{
			yyVAL.item = &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 110:
		{
			yyVAL.item = &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 111:
		{
			yyVAL.item = &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 112:
		{
			yyVAL.item = &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 113:
		{
			yyVAL.item = &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
		}
	case 114:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 115:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 116:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 117:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 118:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 119:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
		}
	case 120:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
		}
	case 121:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
		}
	case 122:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
		}
	case 123:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
		}
	case 124:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].item.(*StructOrUnionSpecifier),
			}
			yyVAL.item = lhs
			if lhs.StructOrUnionSpecifier.isUnion {
				lhs.case2 = tsUnion
				break
			}

			lhs.case2 = tsStruct
		}
	case 125:
		{
			yyVAL.item = &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].item.(*EnumSpecifier),
			}
		}
	case 126:
		{
			yyVAL.item = &TypeSpecifier{
				Case:  13,
				Token: yyS[yypt-0].Token,
			}
		}
	case 127:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier0{
				StructOrUnion: yyS[yypt-1].item.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-0].item.(*IdentifierOpt),
			}
			yyVAL.item = lhs
			lx.pushScope(ScopeMembers)
			lx.scope.SUSpecifier0 = lhs
			lx.scope.isUnion = lhs.StructOrUnion.Token.Val == idUnion
			lx.scope.maxFldAlign = 1
		}
	case 128:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				StructOrUnionSpecifier0: yyS[yypt-3].item.(*StructOrUnionSpecifier0),
				Token: yyS[yypt-2].Token,
				StructDeclarationList: yyS[yypt-1].item.(*StructDeclarationList).reverse(),
				Token2:                yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			s0 := lhs.StructOrUnionSpecifier0
			if io := s0.IdentifierOpt; io != nil {
				lx.fileScope.insert(NSTags, io.Token, lhs)
			}
			s0.SUSpecifier = lhs
			pos := s0.StructOrUnion.Token.Pos()
			lhs.align.pos = pos
			lhs.size.pos = pos
			sc := lx.scope
			lhs.isUnion = sc.isUnion
			switch {
			case lhs.isUnion:
				lhs.align.set(sc.maxFldAlign)
				lhs.size.set(sc.maxFldSize)
			default:
				lhs.align.set(sc.maxFldAlign)
				lhs.size.set(fieldOffset(sc.fldOffset, sc.maxFldAlign))
			}
			lhs.Members = lx.popScope(yyS[yypt-0].Token)
		}
	case 129:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          1,
				StructOrUnion: yyS[yypt-1].item.(*StructOrUnion),
				Token:         yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lx.fileScope.insert(NSTags, lhs.Token, lhs)
			lhs.isUnion = lhs.StructOrUnion.Token.Val == idUnion
			lhs.align.set(maxAlignment)
			lhs.size.set(0)
			lhs.bindings = lx.scope
		}
	case 130:
		{
			yyVAL.item = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 131:
		{
			yyVAL.item = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 132:
		{
			yyVAL.item = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].item.(*StructDeclaration),
			}
		}
	case 133:
		{
			yyVAL.item = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].item.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].item.(*StructDeclaration),
			}
		}
	case 134:
		{
			yyVAL.item = &StructDeclaration{
				SpecifierQualifierList:  yyS[yypt-2].item.(*SpecifierQualifierList),
				StructDeclaratorListOpt: yyS[yypt-1].item.(*StructDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
		}
	case 135:
		{
			lx := yylex.(*lexer)
			yyS[yypt-0].item.(*TypeSpecifier).bindings = lx.scope
		}
	case 136:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				TypeSpecifier:             yyS[yypt-2].item.(*TypeSpecifier),
				SpecifierQualifierListOpt: yyS[yypt-0].item.(*SpecifierQualifierListOpt),
			}
			yyVAL.item = lhs
			lhs.sum(lhs.SpecifierQualifierListOpt)
			ts := lhs.TypeSpecifier
			ts.bindings = lx.scope
			if lhs.typeSpecifiers > 0xffffff {
				compilation.ErrTok(ts.Token, "invalid type specifier")
				lhs.typ = tsVoid
				break
			}

			c := ts.Case
			if c == tsStructOrUnion {
				c = ts.case2
			}
			lhs.typeSpecifiers = lhs.typeSpecifiers<<8 | c
			lhs.typeSum()
			lx.scope.specifier = (*specifierQualifierList)(lhs)
		}
	case 137:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				Case:                      1,
				TypeQualifier:             yyS[yypt-1].item.(*TypeQualifier),
				SpecifierQualifierListOpt: yyS[yypt-0].item.(*SpecifierQualifierListOpt),
			}
			yyVAL.item = lhs
			lhs.sum(lhs.SpecifierQualifierListOpt)
			switch lhs.TypeQualifier.Case {
			case 0: // "const"
				lhs.IsConst = true
			case 1: // "restrict"
				lhs.IsRestrict = true
			case 2: // "volatile"
				lhs.IsVolatile = true
			default:
				panic("internal error")
			}
			lx.scope.specifier = (*specifierQualifierList)(lhs)
		}
	case 138:
		{
			yyVAL.item = (*SpecifierQualifierListOpt)(nil)
		}
	case 139:
		{
			yyVAL.item = &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].item.(*SpecifierQualifierList),
			}
		}
	case 140:
		{
			yyVAL.item = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].item.(*StructDeclarator),
			}
		}
	case 141:
		{
			yyVAL.item = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].item.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].item.(*StructDeclarator),
			}
		}
	case 142:
		{
			yyVAL.item = (*StructDeclaratorListOpt)(nil)
		}
	case 143:
		{
			yyVAL.item = &StructDeclaratorListOpt{
				StructDeclaratorList: yyS[yypt-0].item.(*StructDeclaratorList).reverse(),
			}
		}
	case 144:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].item.(*Declarator),
			}
			yyVAL.item = lhs
			d := lhs.Declarator
			pos := d.Ident().Pos()
			lhs.align.pos = pos
			lhs.offset.pos = pos
			lhs.size.pos = pos
			sc := lx.scope
			t := d.Type()
			sz := t.Sizeof()
			sc.maxFldSize = mathutil.Max(sc.maxFldSize, sz)
			lhs.size.set(sz)
			al := t.Alignof()
			sc.maxFldAlign = mathutil.Max(sc.maxFldAlign, al)
			lhs.align.set(al)
			fldOffset := fieldOffset(sc.fldOffset, al)
			if sc.isUnion {
				fldOffset = 0
			}
			lhs.offset.set(fldOffset)
			if !sc.isUnion {
				sc.fldOffset = fldOffset + sz
			}
			lhs.Declarator.insert(sc, NSMembers, true)
		}
	case 145:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Case:               1,
				DeclaratorOpt:      yyS[yypt-2].item.(*DeclaratorOpt),
				Token:              yyS[yypt-1].Token,
				ConstantExpression: yyS[yypt-0].item.(*ConstantExpression),
			}
			yyVAL.item = lhs
			sc := lx.scope
			pos := lhs.Token.Pos()
			lhs.align.pos = pos
			lhs.offset.pos = pos
			lhs.size.pos = pos
			var t Type
			if o := lhs.DeclaratorOpt; o != nil {
				d := o.Declarator
				pos = d.Ident().Pos()
				d.insert(sc, NSMembers, true)
				t = d.Type()
			}
			t = newBitField(t, int(intT(lhs.ConstantExpression.eval()).(int32)))
			lhs.Bits = t
			al := t.Alignof()
			sz := t.Sizeof()
			sc.maxFldSize = mathutil.Max(sc.maxFldSize, sz)
			lhs.size.set(sz)
			sc.maxFldAlign = mathutil.Max(sc.maxFldAlign, al)
			lhs.align.set(al)
			fldOffset := fieldOffset(sc.fldOffset, al)
			if sc.isUnion {
				fldOffset = 0
			}
			lhs.offset.set(fldOffset)
			if !sc.isUnion {
				sc.fldOffset = fldOffset + sz
			}
		}
	case 146:
		{
			lx := yylex.(*lexer)
			yyVAL.item = &EnumSpecifier0{
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].item.(*IdentifierOpt),
			}
			lx.pushScope(ScopeMembers)
		}
	case 147:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				EnumSpecifier0: yyS[yypt-3].item.(*EnumSpecifier0),
				Token:          yyS[yypt-2].Token,
				EnumeratorList: yyS[yypt-1].item.(*EnumeratorList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if io := lhs.EnumSpecifier0.IdentifierOpt; io != nil {
				lx.fileScope.insert(NSTags, io.Token, lhs)
			}
			lx.popScope(yyS[yypt-0].Token)
		}
	case 148:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Case:           1,
				EnumSpecifier0: yyS[yypt-4].item.(*EnumSpecifier0),
				Token:          yyS[yypt-3].Token,
				EnumeratorList: yyS[yypt-2].item.(*EnumeratorList).reverse(),
				Token2:         yyS[yypt-1].Token,
				Token3:         yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if io := lhs.EnumSpecifier0.IdentifierOpt; io != nil {
				lx.fileScope.insert(NSTags, io.Token, lhs)
			}
			lx.popScope(yyS[yypt-0].Token)
		}
	case 149:
		{
			yyVAL.item = &EnumSpecifier{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 150:
		{
			yyVAL.item = &EnumeratorList{
				Enumerator: yyS[yypt-0].item.(*Enumerator),
			}
		}
	case 151:
		{
			yyVAL.item = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].item.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].item.(*Enumerator),
			}
		}
	case 152:
		{
			yyVAL.item = &Enumerator{
				EnumerationConstant: yyS[yypt-0].item.(*EnumerationConstant),
			}
		}
	case 153:
		{
			yyVAL.item = &Enumerator{
				Case:                1,
				EnumerationConstant: yyS[yypt-2].item.(*EnumerationConstant),
				Token:               yyS[yypt-1].Token,
				ConstantExpression:  yyS[yypt-0].item.(*ConstantExpression),
			}
		}
	case 154:
		{
			yyVAL.item = &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
		}
	case 155:
		{
			yyVAL.item = &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 156:
		{
			yyVAL.item = &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 157:
		{
			yyVAL.item = &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
		}
	case 158:
		{
			lx := yylex.(*lexer)
			lhs := &Declarator{
				PointerOpt:       yyS[yypt-1].item.(*PointerOpt),
				DirectDeclarator: yyS[yypt-0].item.(*DirectDeclarator),
			}
			yyVAL.item = lhs
			lx.declaratorSerial++
			lhs.Serial = lx.declaratorSerial
			lhs.DirectDeclarator.indirection = lhs.PointerOpt.indirection()
			sc := lx.scope
			lhs.IsTypedef = sc.isTypedef
			lhs.SUSpecifier0 = sc.SUSpecifier0
			lhs.Scope = lx.scope
		}
	case 159:
		{
			yyVAL.item = (*DeclaratorOpt)(nil)
		}
	case 160:
		{
			yyVAL.item = &DeclaratorOpt{
				Declarator: yyS[yypt-0].item.(*Declarator),
			}
		}
	case 161:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.specifier = lx.scope.specifier
		}
	case 162:
		{
			yyVAL.item = &DirectDeclarator{
				Case:       1,
				Token:      yyS[yypt-2].Token,
				Declarator: yyS[yypt-1].item.(*Declarator),
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 163:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                    2,
				DirectDeclarator:        yyS[yypt-4].item.(*DirectDeclarator),
				Token:                   yyS[yypt-3].Token,
				TypeQualifierListOpt:    yyS[yypt-2].item.(*TypeQualifierListOpt),
				AssignmentExpressionOpt: yyS[yypt-1].item.(*AssignmentExpressionOpt),
				Token2:                  yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.postProc(lx.scope)
		}
	case 164:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 3,
				DirectDeclarator:     yyS[yypt-5].item.(*DirectDeclarator),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].item.(*TypeQualifierListOpt),
				AssignmentExpression: yyS[yypt-1].item.(*AssignmentExpression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.postProc(lx.scope)
		}
	case 165:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 4,
				DirectDeclarator:     yyS[yypt-5].item.(*DirectDeclarator),
				Token:                yyS[yypt-4].Token,
				TypeQualifierList:    yyS[yypt-3].item.(*TypeQualifierList).reverse(),
				Token2:               yyS[yypt-2].Token,
				AssignmentExpression: yyS[yypt-1].item.(*AssignmentExpression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.postProc(lx.scope)
		}
	case 166:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 5,
				DirectDeclarator:     yyS[yypt-4].item.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].item.(*TypeQualifierListOpt),
				Token2:               yyS[yypt-1].Token,
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.postProc(lx.scope)
		}
	case 167:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeFnParams)
		}
	case 168:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              6,
				DirectDeclarator:  yyS[yypt-3].item.(*DirectDeclarator),
				Token:             yyS[yypt-2].Token,
				DirectDeclarator2: yyS[yypt-0].item.(*DirectDeclarator2),
			}
			yyVAL.item = lhs
			p := lx.scope
			lx.popScope(lhs.DirectDeclarator2.Token)
			lx.scope.params = p
			lhs.postProc(lx.scope)
		}
	case 169:
		{
			yyVAL.item = &DirectDeclarator2{
				ParameterTypeList: yyS[yypt-1].item.(*ParameterTypeList),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 170:
		{
			yyVAL.item = &DirectDeclarator2{
				Case:              1,
				IdentifierListOpt: yyS[yypt-1].item.(*IdentifierListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 171:
		{
			lhs := &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].item.(*TypeQualifierListOpt),
			}
			yyVAL.item = lhs
			lhs.indirection = 1
		}
	case 172:
		{
			lhs := &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].item.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].item.(*Pointer),
			}
			yyVAL.item = lhs
			lhs.indirection = lhs.Pointer.indirection + 1
		}
	case 173:
		{
			yyVAL.item = (*PointerOpt)(nil)
		}
	case 174:
		{
			yyVAL.item = &PointerOpt{
				Pointer: yyS[yypt-0].item.(*Pointer),
			}
		}
	case 175:
		{
			yyVAL.item = &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].item.(*TypeQualifier),
			}
		}
	case 176:
		{
			yyVAL.item = &TypeQualifierList{
				Case:              1,
				TypeQualifierList: yyS[yypt-1].item.(*TypeQualifierList),
				TypeQualifier:     yyS[yypt-0].item.(*TypeQualifier),
			}
		}
	case 177:
		{
			yyVAL.item = (*TypeQualifierListOpt)(nil)
		}
	case 178:
		{
			yyVAL.item = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].item.(*TypeQualifierList).reverse(),
			}
		}
	case 179:
		{
			yyVAL.item = &ParameterTypeList{
				ParameterList: yyS[yypt-0].item.(*ParameterList).reverse(),
			}
		}
	case 180:
		{
			yyVAL.item = &ParameterTypeList{
				Case:          1,
				ParameterList: yyS[yypt-2].item.(*ParameterList).reverse(),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
		}
	case 181:
		{
			yyVAL.item = (*ParameterTypeListOpt)(nil)
		}
	case 182:
		{
			yyVAL.item = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].item.(*ParameterTypeList),
			}
		}
	case 183:
		{
			yyVAL.item = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].item.(*ParameterDeclaration),
			}
		}
	case 184:
		{
			yyVAL.item = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].item.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].item.(*ParameterDeclaration),
			}
		}
	case 185:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				DeclarationSpecifiers: yyS[yypt-1].item.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-0].item.(*Declarator),
			}
			yyVAL.item = lhs
			lhs.Declarator.insert(lx.scope, NSIdentifiers, true)
		}
	case 186:
		{
			yyVAL.item = &ParameterDeclaration{
				Case: 1,
				DeclarationSpecifiers: yyS[yypt-1].item.(*DeclarationSpecifiers),
				AbstractDeclaratorOpt: yyS[yypt-0].item.(*AbstractDeclaratorOpt),
			}
		}
	case 187:
		{
			yyVAL.item = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 188:
		{
			yyVAL.item = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].item.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 189:
		{
			yyVAL.item = (*IdentifierListOpt)(nil)
		}
	case 190:
		{
			yyVAL.item = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].item.(*IdentifierList).reverse(),
			}
		}
	case 191:
		{
			yyVAL.item = (*IdentifierOpt)(nil)
		}
	case 192:
		{
			yyVAL.item = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 193:
		{
			lhs := &TypeName{
				SpecifierQualifierList: yyS[yypt-1].item.(*SpecifierQualifierList),
				AbstractDeclaratorOpt:  yyS[yypt-0].item.(*AbstractDeclaratorOpt),
			}
			yyVAL.item = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				o.AbstractDeclarator.specifier = (*specifierQualifierList)(lhs.SpecifierQualifierList)
			}
		}
	case 194:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Pointer: yyS[yypt-0].item.(*Pointer),
			}
			yyVAL.item = lhs
			lhs.specifier = lx.scope.specifier
			lhs.indirection = lhs.Pointer.indirection
		}
	case 195:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Case:                     1,
				PointerOpt:               yyS[yypt-1].item.(*PointerOpt),
				DirectAbstractDeclarator: yyS[yypt-0].item.(*DirectAbstractDeclarator),
			}
			yyVAL.item = lhs
			dad := lhs.DirectAbstractDeclarator
			dad.specifier = lx.scope.specifier
			dad.indirection = lhs.PointerOpt.indirection()
		}
	case 196:
		{
			yyVAL.item = (*AbstractDeclaratorOpt)(nil)
		}
	case 197:
		{
			yyVAL.item = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].item.(*AbstractDeclarator),
			}
		}
	case 198:
		{
			yyVAL.item = &DirectAbstractDeclarator{
				Token:              yyS[yypt-2].Token,
				AbstractDeclarator: yyS[yypt-1].item.(*AbstractDeclarator),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 199:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 1,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].item.(*DirectAbstractDeclaratorOpt),
				Token: yyS[yypt-2].Token,
				AssignmentExpressionOpt: yyS[yypt-1].item.(*AssignmentExpressionOpt),
				Token2:                  yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if !isExample {
				// fmt.Fprintf(os.Stderr, "TODO: DirectAbstractDeclarator case %v\n%s", lhs.Case, PrettyString(lhs))
			}
		}
	case 200:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 2,
				DirectAbstractDeclaratorOpt: yyS[yypt-4].item.(*DirectAbstractDeclaratorOpt),
				Token:                   yyS[yypt-3].Token,
				TypeQualifierList:       yyS[yypt-2].item.(*TypeQualifierList).reverse(),
				AssignmentExpressionOpt: yyS[yypt-1].item.(*AssignmentExpressionOpt),
				Token2:                  yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if !isExample {
				// fmt.Fprintf(os.Stderr, "TODO: DirectAbstractDeclarator case %v\n%s", lhs.Case, PrettyString(lhs))
			}
		}
	case 201:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 3,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].item.(*DirectAbstractDeclaratorOpt),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].item.(*TypeQualifierListOpt),
				AssignmentExpression: yyS[yypt-1].item.(*AssignmentExpression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if !isExample {
				// fmt.Fprintf(os.Stderr, "TODO: DirectAbstractDeclarator case %v\n%s", lhs.Case, PrettyString(lhs))
			}
		}
	case 202:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 4,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].item.(*DirectAbstractDeclaratorOpt),
				Token:                yyS[yypt-4].Token,
				TypeQualifierList:    yyS[yypt-3].item.(*TypeQualifierList).reverse(),
				Token2:               yyS[yypt-2].Token,
				AssignmentExpression: yyS[yypt-1].item.(*AssignmentExpression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if !isExample {
				// fmt.Fprintf(os.Stderr, "TODO: DirectAbstractDeclarator case %v\n%s", lhs.Case, PrettyString(lhs))
			}
		}
	case 203:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 5,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].item.(*DirectAbstractDeclaratorOpt),
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if !isExample {
				// fmt.Fprintf(os.Stderr, "TODO: DirectAbstractDeclarator case %v\n%s", lhs.Case, PrettyString(lhs))
			}
		}
	case 204:
		{
			lhs := &DirectAbstractDeclarator{
				Case:                 6,
				Token:                yyS[yypt-2].Token,
				ParameterTypeListOpt: yyS[yypt-1].item.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			if !isExample {
				// fmt.Fprintf(os.Stderr, "TODO: DirectAbstractDeclarator case %v\n%s", lhs.Case, PrettyString(lhs))
			}
		}
	case 205:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 7,
				DirectAbstractDeclarator: yyS[yypt-3].item.(*DirectAbstractDeclarator),
				Token:                yyS[yypt-2].Token,
				ParameterTypeListOpt: yyS[yypt-1].item.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			dad := lhs.DirectAbstractDeclarator
			dad.specifier = (*directAbstractDeclarator)(lhs)
			if dad.Case == 0 { //  '(' AbstractDeclarator ')'
				switch ad := dad.AbstractDeclarator; ad.Case {
				case 0: // Pointer
					ad.specifier = (*directAbstractDeclarator)(lhs)
				case 1: // PointerOpt DirectAbstractDeclarator
					ad.DirectAbstractDeclarator.specifier = (*directAbstractDeclarator)(lhs)
				}
			}
		}
	case 206:
		{
			yyVAL.item = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 207:
		{
			yyVAL.item = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].item.(*DirectAbstractDeclarator),
			}
		}
	case 208:
		{
			yyVAL.item = &Initializer{
				AssignmentExpression: yyS[yypt-0].item.(*AssignmentExpression),
			}
		}
	case 209:
		{
			yyVAL.item = &Initializer{
				Case:            1,
				Token:           yyS[yypt-2].Token,
				InitializerList: yyS[yypt-1].item.(*InitializerList).reverse(),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 210:
		{
			yyVAL.item = &Initializer{
				Case:            2,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].item.(*InitializerList).reverse(),
				Token2:          yyS[yypt-1].Token,
				Token3:          yyS[yypt-0].Token,
			}
		}
	case 211:
		{
			yyVAL.item = &InitializerList{
				DesignationOpt: yyS[yypt-1].item.(*DesignationOpt),
				Initializer:    yyS[yypt-0].item.(*Initializer),
			}
		}
	case 212:
		{
			yyVAL.item = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].item.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].item.(*DesignationOpt),
				Initializer:     yyS[yypt-0].item.(*Initializer),
			}
		}
	case 213:
		{
			yyVAL.item = &Designation{
				DesignatorList: yyS[yypt-1].item.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 214:
		{
			yyVAL.item = (*DesignationOpt)(nil)
		}
	case 215:
		{
			yyVAL.item = &DesignationOpt{
				Designation: yyS[yypt-0].item.(*Designation),
			}
		}
	case 216:
		{
			yyVAL.item = &DesignatorList{
				Designator: yyS[yypt-0].item.(*Designator),
			}
		}
	case 217:
		{
			yyVAL.item = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].item.(*DesignatorList),
				Designator:     yyS[yypt-0].item.(*Designator),
			}
		}
	case 218:
		{
			yyVAL.item = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].item.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 219:
		{
			yyVAL.item = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 220:
		{
			yyVAL.item = &Statement{
				LabeledStatement: yyS[yypt-0].item.(*LabeledStatement),
			}
		}
	case 221:
		{
			yyVAL.item = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].item.(*CompoundStatement),
			}
		}
	case 222:
		{
			yyVAL.item = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].item.(*ExpressionStatement),
			}
		}
	case 223:
		{
			yyVAL.item = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].item.(*SelectionStatement),
			}
		}
	case 224:
		{
			yyVAL.item = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].item.(*IterationStatement),
			}
		}
	case 225:
		{
			yyVAL.item = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].item.(*JumpStatement),
			}
		}
	case 226:
		{
			yyVAL.item = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].item.(*Statement),
			}
		}
	case 227:
		{
			yyVAL.item = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].item.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].item.(*Statement),
			}
		}
	case 228:
		{
			yyVAL.item = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].item.(*Statement),
			}
		}
	case 229:
		{
			lx := yylex.(*lexer)
			lx.compoundStmt++
			if lx.compoundStmt != 1 {
				lx.pushScope(ScopeBlock)
			}
		}
	case 230:
		{
			lx := yylex.(*lexer)
			lhs := &CompoundStatement{
				Token:            yyS[yypt-3].Token,
				BlockItemListOpt: yyS[yypt-1].item.(*BlockItemListOpt),
				Token2:           yyS[yypt-0].Token,
			}
			yyVAL.item = lhs
			lhs.Declarations = lx.scope
			lx.compoundStmt--
			if lx.compoundStmt != 0 {
				lx.popScope(yyS[yypt-0].Token)
			}

		}
	case 231:
		{
			yyVAL.item = &BlockItemList{
				BlockItem: yyS[yypt-0].item.(*BlockItem),
			}
		}
	case 232:
		{
			yyVAL.item = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].item.(*BlockItemList),
				BlockItem:     yyS[yypt-0].item.(*BlockItem),
			}
		}
	case 233:
		{
			yyVAL.item = (*BlockItemListOpt)(nil)
		}
	case 234:
		{
			yyVAL.item = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].item.(*BlockItemList).reverse(),
			}
		}
	case 235:
		{
			yyVAL.item = &BlockItem{
				Declaration: yyS[yypt-0].item.(*Declaration),
			}
		}
	case 236:
		{
			yyVAL.item = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].item.(*Statement),
			}
		}
	case 237:
		{
			yyVAL.item = &ExpressionStatement{
				ExpressionOpt: yyS[yypt-1].item.(*ExpressionOpt),
				Token:         yyS[yypt-0].Token,
			}
		}
	case 238:
		{
			yyVAL.item = &SelectionStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].item.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].item.(*Statement),
			}
		}
	case 239:
		{
			yyVAL.item = &SelectionStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Token2:         yyS[yypt-5].Token,
				ExpressionList: yyS[yypt-4].item.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-3].Token,
				Statement:      yyS[yypt-2].item.(*Statement),
				Token4:         yyS[yypt-1].Token,
				Statement2:     yyS[yypt-0].item.(*Statement),
			}
		}
	case 240:
		{
			yyVAL.item = &SelectionStatement{
				Case:           2,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].item.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].item.(*Statement),
			}
		}
	case 241:
		{
			yyVAL.item = &IterationStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].item.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].item.(*Statement),
			}
		}
	case 242:
		{
			yyVAL.item = &IterationStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Statement:      yyS[yypt-5].item.(*Statement),
				Token2:         yyS[yypt-4].Token,
				Token3:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].item.(*ExpressionList).reverse(),
				Token4:         yyS[yypt-1].Token,
				Token5:         yyS[yypt-0].Token,
			}
		}
	case 243:
		{
			yyVAL.item = &IterationStatement{
				Case:           2,
				Token:          yyS[yypt-8].Token,
				Token2:         yyS[yypt-7].Token,
				ExpressionOpt:  yyS[yypt-6].item.(*ExpressionOpt),
				Token3:         yyS[yypt-5].Token,
				ExpressionOpt2: yyS[yypt-4].item.(*ExpressionOpt),
				Token4:         yyS[yypt-3].Token,
				ExpressionOpt3: yyS[yypt-2].item.(*ExpressionOpt),
				Token5:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].item.(*Statement),
			}
		}
	case 244:
		{
			yyVAL.item = &IterationStatement{
				Case:           3,
				Token:          yyS[yypt-7].Token,
				Token2:         yyS[yypt-6].Token,
				Declaration:    yyS[yypt-5].item.(*Declaration),
				ExpressionOpt:  yyS[yypt-4].item.(*ExpressionOpt),
				Token3:         yyS[yypt-3].Token,
				ExpressionOpt2: yyS[yypt-2].item.(*ExpressionOpt),
				Token4:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].item.(*Statement),
			}
		}
	case 245:
		{
			yyVAL.item = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 246:
		{
			yyVAL.item = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 247:
		{
			yyVAL.item = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 248:
		{
			yyVAL.item = &JumpStatement{
				Case:          3,
				Token:         yyS[yypt-2].Token,
				ExpressionOpt: yyS[yypt-1].item.(*ExpressionOpt),
				Token2:        yyS[yypt-0].Token,
			}
		}
	case 249:
		{
			yyVAL.item = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].item.(*ExternalDeclaration),
			}
		}
	case 250:
		{
			yyVAL.item = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].item.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].item.(*ExternalDeclaration),
			}
		}
	case 251:
		{
			yyVAL.item = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].item.(*FunctionDefinition),
			}
		}
	case 252:
		{
			yyVAL.item = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].item.(*Declaration),
			}
		}
	case 253:
		{
			lx := yylex.(*lexer)
			parScope := lx.scope.params
			lx.pushScope(ScopeFunction).copy(parScope)
			lx.compoundStmt = 0
		}
	case 254:
		{
			lx := yylex.(*lexer)
			lhs := &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].item.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].item.(*Declarator),
				DeclarationListOpt:    yyS[yypt-1].item.(*DeclarationListOpt),
				CompoundStatement:     yyS[yypt-0].item.(*CompoundStatement),
			}
			yyVAL.item = lhs
			lhs.Declarations = lx.popScope(lhs.CompoundStatement.Token2)
			d := lhs.Declarator
			d.IsDefinition = true
			lx.scope.insert(NSIdentifiers, d.Ident(), lhs)
		}
	case 255:
		{
			yyVAL.item = &DeclarationList{
				Declaration: yyS[yypt-0].item.(*Declaration),
			}
		}
	case 256:
		{
			yyVAL.item = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].item.(*DeclarationList),
				Declaration:     yyS[yypt-0].item.(*Declaration),
			}
		}
	case 257:
		{
			yyVAL.item = (*DeclarationListOpt)(nil)
		}
	case 258:
		{
			yyVAL.item = &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].item.(*DeclarationList).reverse(),
			}
		}
	case 259:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].item.(*GroupList).reverse(),
			}
			yyVAL.item = lhs
			lx.ast = lhs
			lhs.file = lx.file
		}
	case 260:
		{
			switch e := yyS[yypt-0].item.(*GroupPart); {
			case e != nil:
				yyVAL.item = &GroupList{
					GroupPart: e,
				}
			default:
				yyVAL.item = (*GroupList)(nil)
			}
		}
	case 261:
		{
			switch l, e := yyS[yypt-1].item.(*GroupList), yyS[yypt-0].item.(*GroupPart); {
			case e == nil:
				yyVAL.item = l
			default:
				yyVAL.item = &GroupList{
					GroupList: l,
					GroupPart: e,
				}
			}
		}
	case 262:
		{
			yyVAL.item = (*GroupListOpt)(nil)
		}
	case 263:
		{
			yyVAL.item = &GroupListOpt{
				GroupList: yyS[yypt-0].item.(*GroupList).reverse(),
			}
		}
	case 264:
		{
			yyVAL.item = &GroupPart{
				ControlLine: yyS[yypt-0].item.(*ControlLine),
			}
		}
	case 265:
		{
			yyVAL.item = &GroupPart{
				IfSection: yyS[yypt-0].item.(*IfSection),
			}
		}
	case 266:
		{
			yyVAL.item = &GroupPart{
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 267:
		{
			if yyS[yypt-0].toks == 0 {
				yyVAL.item = (*GroupPart)(nil)
				break
			}

			yyVAL.item = &GroupPart{
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 268:
		{
			yyVAL.item = &IfSection{
				IfGroup:          yyS[yypt-3].item.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].item.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].item.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].item.(*EndifLine),
			}
		}
	case 269:
		{
			yyVAL.item = &IfGroup{
				Token:        yyS[yypt-2].Token,
				PpTokenList:  yyS[yypt-1].toks,
				GroupListOpt: yyS[yypt-0].item.(*GroupListOpt),
			}
		}
	case 270:
		{
			yyVAL.item = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].item.(*GroupListOpt),
			}
		}
	case 271:
		{
			yyVAL.item = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].item.(*GroupListOpt),
			}
		}
	case 272:
		{
			yyVAL.item = &ElifGroupList{
				ElifGroup: yyS[yypt-0].item.(*ElifGroup),
			}
		}
	case 273:
		{
			yyVAL.item = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].item.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].item.(*ElifGroup),
			}
		}
	case 274:
		{
			yyVAL.item = (*ElifGroupListOpt)(nil)
		}
	case 275:
		{
			yyVAL.item = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].item.(*ElifGroupList).reverse(),
			}
		}
	case 276:
		{
			yyVAL.item = &ElifGroup{
				Token:        yyS[yypt-2].Token,
				PpTokenList:  yyS[yypt-1].toks,
				GroupListOpt: yyS[yypt-0].item.(*GroupListOpt),
			}
		}
	case 277:
		{
			yyVAL.item = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].item.(*GroupListOpt),
			}
		}
	case 278:
		{
			yyVAL.item = (*ElseGroupOpt)(nil)
		}
	case 279:
		{
			yyVAL.item = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].item.(*ElseGroup),
			}
		}
	case 280:
		{
			yyVAL.item = &EndifLine{
				Token:          yyS[yypt-1].Token,
				PpTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 281:
		{
			yyVAL.item = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 282:
		{
			yyVAL.item = &ControlLine{
				Case:            1,
				Token:           yyS[yypt-4].Token,
				Token2:          yyS[yypt-3].Token,
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 283:
		{
			yyVAL.item = &ControlLine{
				Case:            2,
				Token:           yyS[yypt-6].Token,
				Token2:          yyS[yypt-5].Token,
				IdentifierList:  yyS[yypt-4].item.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 284:
		{
			yyVAL.item = &ControlLine{
				Case:              3,
				Token:             yyS[yypt-4].Token,
				Token2:            yyS[yypt-3].Token,
				IdentifierListOpt: yyS[yypt-2].item.(*IdentifierListOpt),
				Token3:            yyS[yypt-1].Token,
				ReplacementList:   yyS[yypt-0].toks,
			}
		}
	case 285:
		{
			yyVAL.item = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PpTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 286:
		{
			yyVAL.item = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 287:
		{
			yyVAL.item = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 288:
		{
			yyVAL.item = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 289:
		{
			yyVAL.item = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PpTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 290:
		{
			yyVAL.item = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 291:
		{
			yyVAL.item = &ControlLine{
				Case:        10,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 292:
		{
			yyVAL.item = &ControlLine{
				Case:            11,
				Token:           yyS[yypt-5].Token,
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 293:
		{
			yyVAL.item = &ControlLine{
				Case:            12,
				Token:           yyS[yypt-7].Token,
				Token2:          yyS[yypt-6].Token,
				IdentifierList:  yyS[yypt-5].item.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-4].Token,
				Token4:          yyS[yypt-3].Token,
				Token5:          yyS[yypt-2].Token,
				Token6:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 294:
		{
			yyVAL.item = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 295:
		{
			yyVAL.item = &ControlLine{
				Case:        14,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 296:
		{
			yyVAL.item = &ControlLine{
				Case:        15,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 297:
		{
			yyVAL.item = &ControlLine{
				Case:        16,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 298:
		{
			yyVAL.item = &ControlLine{
				Case:        17,
				Token:       yyS[yypt-1].Token,
				PpTokenList: yyS[yypt-0].toks,
			}
		}
	case 301:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PpTokenList(db.putTokens(lx.zipToks))
		}
	case 302:
		{
			yyVAL.toks = 0
		}
	case 304:
		{
			lx := yylex.(*lexer)
			lx.zipToks = append(lx.zipToks[:0], yyS[yypt-0].Token)
		}
	case 305:
		{
			lx := yylex.(*lexer)
			lx.zipToks = append(lx.zipToks, yyS[yypt-0].Token)
		}
	case 306:
		{
			lx := yylex.(*lexer)
			n := len(lx.macroArg)
			last := lx.macroArg[n-1]
			if last.Rune != MACRO_ARG_EMPTY {
				lx.macroArgs = append(lx.macroArgs, lx.macroArg)
			}
			lx.macroArg = nil
		}
	case 307:
		{
			lx := yylex.(*lexer)
			lx.macroArgs = append(lx.macroArgs, lx.macroArg)
			lx.macroArg = nil
		}
	case 308:
		{
			lx := yylex.(*lexer)
			lx.macroArg = append(lx.macroArg, xc.Token{Char: lex.NewChar(lx.last.Pos(), MACRO_ARG_EMPTY)})
		}
	case 309:
		{
			lx := yylex.(*lexer)
			lx.macroArg = append(lx.macroArg, yyS[yypt-0].Token)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
