// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
)

func ExampleAbstractDeclarator() {
	fmt.Println(exampleAST(194, "\U00100001 ( _Bool * )"))
	// Output:
	// &cc.AbstractDeclarator{
	// · Pointer: &cc.Pointer{
	// · · Token: example194.c:1:14: '*',
	// · },
	// }
}

func ExampleAbstractDeclarator_case1() {
	fmt.Println(exampleAST(195, "\U00100001 ( _Bool ( ) )"))
	// Output:
	// &cc.AbstractDeclarator{
	// · Case: 1,
	// · DirectAbstractDeclarator: &cc.DirectAbstractDeclarator{
	// · · Case: 6,
	// · · Token: example195.c:1:14: '(',
	// · · Token2: example195.c:1:16: ')',
	// · },
	// }
}

func ExampleAbstractDeclaratorOpt() {
	fmt.Println(exampleAST(196, "\U00100001 ( _Bool )") == (*AbstractDeclaratorOpt)(nil))
	// Output:
	// true
}

func ExampleAbstractDeclaratorOpt_case1() {
	fmt.Println(exampleAST(197, "\U00100001 ( _Bool * )"))
	// Output:
	// &cc.AbstractDeclaratorOpt{
	// · AbstractDeclarator: &cc.AbstractDeclarator{
	// · · Pointer: &cc.Pointer{
	// · · · Token: example197.c:1:14: '*',
	// · · },
	// · },
	// }
}

func ExampleAdditiveExpression() {
	fmt.Println(exampleAST(48, "\U00100001 'a'"))
	// Output:
	// &cc.AdditiveExpression{
	// · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · CastExpression: &cc.CastExpression{
	// · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · Case: 1,
	// · · · · · · Constant: &cc.Constant{
	// · · · · · · · Token: example48.c:1:6: CHARCONST "'a'",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleAdditiveExpression_case1() {
	fmt.Println(exampleAST(49, "\U00100001 'a' + 'b'"))
	// Output:
	// &cc.AdditiveExpression{
	// · AdditiveExpression: &cc.AdditiveExpression{
	// · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · CastExpression: &cc.CastExpression{
	// · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · Case: 1,
	// · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · Token: example49.c:1:6: CHARCONST "'a'",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 1,
	// · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · CastExpression: &cc.CastExpression{
	// · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · Case: 1,
	// · · · · · · Constant: &cc.Constant{
	// · · · · · · · Token: example49.c:1:12: CHARCONST "'b'",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example49.c:1:10: '+',
	// }
}

func ExampleAdditiveExpression_case2() {
	fmt.Println(exampleAST(50, "\U00100001 'a' - 'b'"))
	// Output:
	// &cc.AdditiveExpression{
	// · AdditiveExpression: &cc.AdditiveExpression{
	// · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · CastExpression: &cc.CastExpression{
	// · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · Case: 1,
	// · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · Token: example50.c:1:6: CHARCONST "'a'",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 2,
	// · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · CastExpression: &cc.CastExpression{
	// · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · Case: 1,
	// · · · · · · Constant: &cc.Constant{
	// · · · · · · · Token: example50.c:1:12: CHARCONST "'b'",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example50.c:1:10: '-',
	// }
}

func ExampleAndExpression() {
	fmt.Println(exampleAST(62, "\U00100001 'a'"))
	// Output:
	// &cc.AndExpression{
	// · EqualityExpression: &cc.EqualityExpression{
	// · · RelationalExpression: &cc.RelationalExpression{
	// · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · Case: 1,
	// · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · Token: example62.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleAndExpression_case1() {
	fmt.Println(exampleAST(63, "\U00100001 'a' & 'b'"))
	// Output:
	// &cc.AndExpression{
	// · AndExpression: &cc.AndExpression{
	// · · EqualityExpression: &cc.EqualityExpression{
	// · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · Token: example63.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 1,
	// · EqualityExpression: &cc.EqualityExpression{
	// · · RelationalExpression: &cc.RelationalExpression{
	// · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · Case: 1,
	// · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · Token: example63.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example63.c:1:10: '&',
	// }
}

func ExampleArgumentExpressionList() {
	fmt.Println(exampleAST(24, "\U00100001 'a' ( 'b' )"))
	// Output:
	// &cc.ArgumentExpressionList{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example24.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleArgumentExpressionList_case1() {
	fmt.Println(exampleAST(25, "\U00100001 'a' ( 'b' , 'c' )"))
	// Output:
	// &cc.ArgumentExpressionList{
	// · ArgumentExpressionList: &cc.ArgumentExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example25.c:1:18: CHARCONST "'c'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Case: 1,
	// · · Token: example25.c:1:16: ',',
	// · },
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example25.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleArgumentExpressionListOpt() {
	fmt.Println(exampleAST(26, "\U00100001 'a' ( )") == (*ArgumentExpressionListOpt)(nil))
	// Output:
	// true
}

func ExampleArgumentExpressionListOpt_case1() {
	fmt.Println(exampleAST(27, "\U00100001 'a' ( 'b' )"))
	// Output:
	// &cc.ArgumentExpressionListOpt{
	// · ArgumentExpressionList: &cc.ArgumentExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example27.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleAssignmentExpression() {
	fmt.Println(exampleAST(74, "\U00100001 ( 'a' )"))
	// Output:
	// &cc.AssignmentExpression{
	// · ConditionalExpression: &cc.ConditionalExpression{
	// · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · Token: example74.c:1:8: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleAssignmentExpression_case1() {
	fmt.Println(exampleAST(75, "\U00100001 ( 'a' = 'b' )"))
	// Output:
	// &cc.AssignmentExpression{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example75.c:1:14: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · AssignmentOperator: &cc.AssignmentOperator{
	// · · Token: example75.c:1:12: '=',
	// · },
	// · Case: 1,
	// · UnaryExpression: &cc.UnaryExpression{
	// · · PostfixExpression: &cc.PostfixExpression{
	// · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · Case: 1,
	// · · · · Constant: &cc.Constant{
	// · · · · · Token: example75.c:1:8: CHARCONST "'a'",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleAssignmentExpressionOpt() {
	fmt.Println(exampleAST(76, "\U00100001 ( _Bool [ ]") == (*AssignmentExpressionOpt)(nil))
	// Output:
	// true
}

func ExampleAssignmentExpressionOpt_case1() {
	fmt.Println(exampleAST(77, "\U00100001 ( _Bool [ 'a' ]"))
	// Output:
	// &cc.AssignmentExpressionOpt{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example77.c:1:16: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleAssignmentOperator() {
	fmt.Println(exampleAST(78, "\U00100001 ( 'a' = !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Token: example78.c:1:12: '=',
	// }
}

func ExampleAssignmentOperator_case01() {
	fmt.Println(exampleAST(79, "\U00100001 ( 'a' *= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 1,
	// · Token: example79.c:1:12: MULASSIGN,
	// }
}

func ExampleAssignmentOperator_case02() {
	fmt.Println(exampleAST(80, "\U00100001 ( 'a' /= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 2,
	// · Token: example80.c:1:12: DIVASSIGN,
	// }
}

func ExampleAssignmentOperator_case03() {
	fmt.Println(exampleAST(81, "\U00100001 ( 'a' %= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 3,
	// · Token: example81.c:1:12: MODASSIGN,
	// }
}

func ExampleAssignmentOperator_case04() {
	fmt.Println(exampleAST(82, "\U00100001 ( 'a' += !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 4,
	// · Token: example82.c:1:12: ADDASSIGN,
	// }
}

func ExampleAssignmentOperator_case05() {
	fmt.Println(exampleAST(83, "\U00100001 ( 'a' -= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 5,
	// · Token: example83.c:1:12: SUBASSIGN,
	// }
}

func ExampleAssignmentOperator_case06() {
	fmt.Println(exampleAST(84, "\U00100001 ( 'a' <<= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 6,
	// · Token: example84.c:1:12: LSHASSIGN,
	// }
}

func ExampleAssignmentOperator_case07() {
	fmt.Println(exampleAST(85, "\U00100001 ( 'a' >>= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 7,
	// · Token: example85.c:1:12: RSHASSIGN,
	// }
}

func ExampleAssignmentOperator_case08() {
	fmt.Println(exampleAST(86, "\U00100001 ( 'a' &= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 8,
	// · Token: example86.c:1:12: ANDASSIGN,
	// }
}

func ExampleAssignmentOperator_case09() {
	fmt.Println(exampleAST(87, "\U00100001 ( 'a' ^= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 9,
	// · Token: example87.c:1:12: XORASSIGN,
	// }
}

func ExampleAssignmentOperator_case10() {
	fmt.Println(exampleAST(88, "\U00100001 ( 'a' |= !"))
	// Output:
	// &cc.AssignmentOperator{
	// · Case: 10,
	// · Token: example88.c:1:12: ORASSIGN,
	// }
}

func ExampleBlockItem() {
	fmt.Println(exampleAST(235, "\U00100002 auto a { auto ; !"))
	// Output:
	// &cc.BlockItem{
	// · Declaration: &cc.Declaration{
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example235.c:1:15: AUTO "auto",
	// · · · },
	// · · },
	// · · Token: example235.c:1:20: ';',
	// · },
	// }
}

func ExampleBlockItem_case1() {
	fmt.Println(exampleAST(236, "\U00100002 auto a { ; !"))
	// Output:
	// &cc.BlockItem{
	// · Case: 1,
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example236.c:1:15: ';',
	// · · },
	// · },
	// }
}

func ExampleBlockItemList() {
	fmt.Println(exampleAST(231, "\U00100002 auto a { ; !"))
	// Output:
	// &cc.BlockItemList{
	// · BlockItem: &cc.BlockItem{
	// · · Case: 1,
	// · · Statement: &cc.Statement{
	// · · · Case: 2,
	// · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · Token: example231.c:1:15: ';',
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleBlockItemList_case1() {
	fmt.Println(exampleAST(232, "\U00100002 auto a { ; ; !"))
	// Output:
	// &cc.BlockItemList{
	// · BlockItem: &cc.BlockItem{
	// · · Case: 1,
	// · · Statement: &cc.Statement{
	// · · · Case: 2,
	// · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · Token: example232.c:1:15: ';',
	// · · · },
	// · · },
	// · },
	// · BlockItemList: &cc.BlockItemList{
	// · · BlockItem: &cc.BlockItem{
	// · · · Case: 1,
	// · · · Statement: &cc.Statement{
	// · · · · Case: 2,
	// · · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · · Token: example232.c:1:17: ';',
	// · · · · },
	// · · · },
	// · · },
	// · · Case: 1,
	// · },
	// }
}

func ExampleBlockItemListOpt() {
	fmt.Println(exampleAST(233, "\U00100002 auto a { }") == (*BlockItemListOpt)(nil))
	// Output:
	// true
}

func ExampleBlockItemListOpt_case1() {
	fmt.Println(exampleAST(234, "\U00100002 auto a { ; }"))
	// Output:
	// &cc.BlockItemListOpt{
	// · BlockItemList: &cc.BlockItemList{
	// · · BlockItem: &cc.BlockItem{
	// · · · Case: 1,
	// · · · Statement: &cc.Statement{
	// · · · · Case: 2,
	// · · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · · Token: example234.c:1:15: ';',
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleCastExpression() {
	fmt.Println(exampleAST(42, "\U00100001 'a'"))
	// Output:
	// &cc.CastExpression{
	// · UnaryExpression: &cc.UnaryExpression{
	// · · PostfixExpression: &cc.PostfixExpression{
	// · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · Case: 1,
	// · · · · Constant: &cc.Constant{
	// · · · · · Token: example42.c:1:6: CHARCONST "'a'",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleCastExpression_case1() {
	fmt.Println(exampleAST(43, "\U00100001 ( _Bool ) 'a'"))
	// Output:
	// &cc.CastExpression{
	// · Case: 1,
	// · CastExpression: &cc.CastExpression{
	// · · UnaryExpression: &cc.UnaryExpression{
	// · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · Case: 1,
	// · · · · · Constant: &cc.Constant{
	// · · · · · · Token: example43.c:1:16: CHARCONST "'a'",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example43.c:1:6: '(',
	// · Token2: example43.c:1:14: ')',
	// · TypeName: &cc.TypeName{
	// · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · Case: 9,
	// · · · · Token: example43.c:1:8: BOOL "_Bool",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleCompoundStatement() {
	fmt.Println(exampleAST(230, "\U00100002 auto a { }"))
	// Output:
	// &cc.CompoundStatement{
	// · Declarations: &cc.Bindings{
	// · · Type: 3,
	// · },
	// · Token: example230.c:1:13: '{',
	// · Token2: example230.c:1:15: '}',
	// }
}

func ExampleConditionalExpression() {
	fmt.Println(exampleAST(72, "\U00100001 'a'"))
	// Output:
	// &cc.ConditionalExpression{
	// · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · Token: example72.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleConditionalExpression_case1() {
	fmt.Println(exampleAST(73, "\U00100001 'a' ? 'b' : 'c'"))
	// Output:
	// &cc.ConditionalExpression{
	// · Case: 1,
	// · ConditionalExpression: &cc.ConditionalExpression{
	// · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · Token: example73.c:1:18: CHARCONST "'c'",
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example73.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · Token: example73.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example73.c:1:10: '?',
	// · Token2: example73.c:1:16: ':',
	// }
}

func ExampleConstant() {
	fmt.Println(exampleAST(9, "\U00100001 'a'"))
	// Output:
	// &cc.Constant{
	// · Token: example9.c:1:6: CHARCONST "'a'",
	// }
}

func ExampleConstant_case1() {
	fmt.Println(exampleAST(10, "\U00100001 1.97"))
	// Output:
	// &cc.Constant{
	// · Case: 1,
	// · Token: example10.c:1:6: FLOATCONST "1.97",
	// }
}

func ExampleConstant_case2() {
	fmt.Println(exampleAST(11, "\U00100001 97"))
	// Output:
	// &cc.Constant{
	// · Case: 2,
	// · Token: example11.c:1:6: INTCONST "97",
	// }
}

func ExampleConstant_case3() {
	fmt.Println(exampleAST(12, "\U00100001 L'a'"))
	// Output:
	// &cc.Constant{
	// · Case: 3,
	// · Token: example12.c:1:6: LONGCHARCONST "L'a'",
	// }
}

func ExampleConstant_case4() {
	fmt.Println(exampleAST(13, "\U00100001 L\"a\""))
	// Output:
	// &cc.Constant{
	// · Case: 4,
	// · Token: example13.c:1:6: LONGSTRINGLITERAL "L\"a\"",
	// }
}

func ExampleConstant_case5() {
	fmt.Println(exampleAST(14, "\U00100001 \"a\""))
	// Output:
	// &cc.Constant{
	// · Case: 5,
	// · Token: example14.c:1:6: STRINGLITERAL "\"a\"",
	// }
}

func ExampleConstantExpression() {
	fmt.Println(exampleAST(93, "\U00100001 'a'"))
	// Output:
	// &cc.ConstantExpression{
	// · ConditionalExpression: &cc.ConditionalExpression{
	// · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · Token: example93.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleControlLine() {
	fmt.Println(exampleAST(281, "\U00100000 \n#define a "))
	// Output:
	// &cc.ControlLine{
	// · Token: example281.c:2:2: PPDEFINE,
	// · Token2: example281.c:2:9: IDENTIFIER "a",
	// }
}

func ExampleControlLine_case01() {
	fmt.Println(exampleAST(282, "\U00100000 \n#define a( ... ) "))
	// Output:
	// &cc.ControlLine{
	// · Case: 1,
	// · Token: example282.c:2:2: PPDEFINE,
	// · Token2: example282.c:2:9: IDENTIFIER_LPAREN "a",
	// · Token3: example282.c:2:12: DDD,
	// · Token4: example282.c:2:16: ')',
	// }
}

func ExampleControlLine_case02() {
	fmt.Println(exampleAST(283, "\U00100000 \n#define a( b , ... ) "))
	// Output:
	// &cc.ControlLine{
	// · Case: 2,
	// · IdentifierList: &cc.IdentifierList{
	// · · Token: example283.c:2:12: IDENTIFIER "b",
	// · },
	// · Token: example283.c:2:2: PPDEFINE,
	// · Token2: example283.c:2:9: IDENTIFIER_LPAREN "a",
	// · Token3: example283.c:2:14: ',',
	// · Token4: example283.c:2:16: DDD,
	// · Token5: example283.c:2:20: ')',
	// }
}

func ExampleControlLine_case03() {
	fmt.Println(exampleAST(284, "\U00100000 \n#define a( ) "))
	// Output:
	// &cc.ControlLine{
	// · Case: 3,
	// · Token: example284.c:2:2: PPDEFINE,
	// · Token2: example284.c:2:9: IDENTIFIER_LPAREN "a",
	// · Token3: example284.c:2:12: ')',
	// }
}

func ExampleControlLine_case04() {
	fmt.Println(exampleAST(285, "\U00100000 \n#error "))
	// Output:
	// &cc.ControlLine{
	// · Case: 4,
	// · Token: example285.c:2:2: PPERROR,
	// }
}

func ExampleControlLine_case05() {
	fmt.Println(exampleAST(286, "\U00100000 \n#"))
	// Output:
	// &cc.ControlLine{
	// · Case: 5,
	// · Token: example286.c:2:1: '#',
	// }
}

func ExampleControlLine_case06() {
	fmt.Println(exampleAST(287, "\U00100000 \n#include ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 6,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example287.c:2:10: IDENTIFIER "ppother",
	// },
	// · Token: example287.c:2:2: PPINCLUDE,
	// }
}

func ExampleControlLine_case07() {
	fmt.Println(exampleAST(288, "\U00100000 \n#line ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 7,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example288.c:2:7: IDENTIFIER "ppother",
	// },
	// · Token: example288.c:2:2: PPLINE,
	// }
}

func ExampleControlLine_case08() {
	fmt.Println(exampleAST(289, "\U00100000 \n#pragma "))
	// Output:
	// &cc.ControlLine{
	// · Case: 8,
	// · Token: example289.c:2:2: PPPRAGMA,
	// }
}

func ExampleControlLine_case09() {
	fmt.Println(exampleAST(290, "\U00100000 \n#undef a "))
	// Output:
	// &cc.ControlLine{
	// · Case: 9,
	// · Token: example290.c:2:2: PPUNDEF,
	// · Token2: example290.c:2:8: IDENTIFIER "a",
	// · Token3: example290.c:2:10: '\n',
	// }
}

func ExampleControlLine_case10() {
	fmt.Println(exampleAST(291, "\U00100000 \n#assert ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 10,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example291.c:2:9: IDENTIFIER "ppother",
	// },
	// · Token: example291.c:2:2: PPASSERT,
	// }
}

func ExampleControlLine_case11() {
	fmt.Println(exampleAST(292, "\U00100000 \n#define a( b ... ) "))
	// Output:
	// &cc.ControlLine{
	// · Case: 11,
	// · Token: example292.c:2:2: PPDEFINE,
	// · Token2: example292.c:2:9: IDENTIFIER_LPAREN "a",
	// · Token3: example292.c:2:12: IDENTIFIER "b",
	// · Token4: example292.c:2:14: DDD,
	// · Token5: example292.c:2:18: ')',
	// }
}

func ExampleControlLine_case12() {
	fmt.Println(exampleAST(293, "\U00100000 \n#define a( b , c ... ) "))
	// Output:
	// &cc.ControlLine{
	// · Case: 12,
	// · IdentifierList: &cc.IdentifierList{
	// · · Token: example293.c:2:12: IDENTIFIER "b",
	// · },
	// · Token: example293.c:2:2: PPDEFINE,
	// · Token2: example293.c:2:9: IDENTIFIER_LPAREN "a",
	// · Token3: example293.c:2:14: ',',
	// · Token4: example293.c:2:16: IDENTIFIER "c",
	// · Token5: example293.c:2:18: DDD,
	// · Token6: example293.c:2:22: ')',
	// }
}

func ExampleControlLine_case13() {
	fmt.Println(exampleAST(294, "\U00100000 \n#ident ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 13,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example294.c:2:8: IDENTIFIER "ppother",
	// },
	// · Token: example294.c:2:2: PPIDENT,
	// }
}

func ExampleControlLine_case14() {
	fmt.Println(exampleAST(295, "\U00100000 \n#import ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 14,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example295.c:2:9: IDENTIFIER "ppother",
	// },
	// · Token: example295.c:2:2: PPIMPORT,
	// }
}

func ExampleControlLine_case15() {
	fmt.Println(exampleAST(296, "\U00100000 \n#include_next ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 15,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example296.c:2:15: IDENTIFIER "ppother",
	// },
	// · Token: example296.c:2:2: PPINCLUDE_NEXT,
	// }
}

func ExampleControlLine_case16() {
	fmt.Println(exampleAST(297, "\U00100000 \n#unassert ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 16,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example297.c:2:11: IDENTIFIER "ppother",
	// },
	// · Token: example297.c:2:2: PPUNASSERT,
	// }
}

func ExampleControlLine_case17() {
	fmt.Println(exampleAST(298, "\U00100000 \n#warning ppother "))
	// Output:
	// &cc.ControlLine{
	// · Case: 17,
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example298.c:2:10: IDENTIFIER "ppother",
	// },
	// · Token: example298.c:2:2: PPWARNING,
	// }
}

func ExampleDeclaration() {
	fmt.Println(exampleAST(94, "\U00100002 auto ;"))
	// Output:
	// &cc.Declaration{
	// · IsFileScope: true,
	// · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · IsAuto: true,
	// · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · Case: 3,
	// · · · Token: example94.c:1:6: AUTO "auto",
	// · · },
	// · },
	// · Token: example94.c:1:11: ';',
	// }
}

func ExampleDeclarationList() {
	fmt.Println(exampleAST(255, "\U00100002 auto a auto ; {"))
	// Output:
	// &cc.DeclarationList{
	// · Declaration: &cc.Declaration{
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example255.c:1:13: AUTO "auto",
	// · · · },
	// · · },
	// · · Token: example255.c:1:18: ';',
	// · },
	// }
}

func ExampleDeclarationList_case1() {
	fmt.Println(exampleAST(256, "\U00100002 auto a auto ; auto ; {"))
	// Output:
	// &cc.DeclarationList{
	// · Declaration: &cc.Declaration{
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example256.c:1:13: AUTO "auto",
	// · · · },
	// · · },
	// · · Token: example256.c:1:18: ';',
	// · },
	// · DeclarationList: &cc.DeclarationList{
	// · · Case: 1,
	// · · Declaration: &cc.Declaration{
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example256.c:1:20: AUTO "auto",
	// · · · · },
	// · · · },
	// · · · Token: example256.c:1:25: ';',
	// · · },
	// · },
	// }
}

func ExampleDeclarationListOpt() {
	fmt.Println(exampleAST(257, "\U00100002 auto a {") == (*DeclarationListOpt)(nil))
	// Output:
	// true
}

func ExampleDeclarationListOpt_case1() {
	fmt.Println(exampleAST(258, "\U00100002 auto a auto ; {"))
	// Output:
	// &cc.DeclarationListOpt{
	// · DeclarationList: &cc.DeclarationList{
	// · · Declaration: &cc.Declaration{
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example258.c:1:13: AUTO "auto",
	// · · · · },
	// · · · },
	// · · · Token: example258.c:1:18: ';',
	// · · },
	// · },
	// }
}

func ExampleDeclarationSpecifiers() {
	fmt.Println(exampleAST(95, "\U00100002 auto ("))
	// Output:
	// &cc.DeclarationSpecifiers{
	// · IsAuto: true,
	// · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · Case: 3,
	// · · Token: example95.c:1:6: AUTO "auto",
	// · },
	// }
}

func ExampleDeclarationSpecifiers_case1() {
	fmt.Println(exampleAST(97, "\U00100002 _Bool ("))
	// Output:
	// &cc.DeclarationSpecifiers{
	// · Case: 1,
	// · TypeSpecifier: &cc.TypeSpecifier{
	// · · Case: 9,
	// · · Token: example97.c:1:6: BOOL "_Bool",
	// · },
	// }
}

func ExampleDeclarationSpecifiers_case2() {
	fmt.Println(exampleAST(98, "\U00100002 const ("))
	// Output:
	// &cc.DeclarationSpecifiers{
	// · IsConst: true,
	// · Case: 2,
	// · TypeQualifier: &cc.TypeQualifier{
	// · · Token: example98.c:1:6: CONST "const",
	// · },
	// }
}

func ExampleDeclarationSpecifiers_case3() {
	fmt.Println(exampleAST(99, "\U00100002 inline ("))
	// Output:
	// &cc.DeclarationSpecifiers{
	// · IsInline: true,
	// · Case: 3,
	// · FunctionSpecifier: &cc.FunctionSpecifier{
	// · · Token: example99.c:1:6: INLINE "inline",
	// · },
	// }
}

func ExampleDeclarationSpecifiersOpt() {
	fmt.Println(exampleAST(100, "\U00100002 auto (") == (*DeclarationSpecifiersOpt)(nil))
	// Output:
	// true
}

func ExampleDeclarationSpecifiersOpt_case1() {
	fmt.Println(exampleAST(101, "\U00100002 auto auto ("))
	// Output:
	// &cc.DeclarationSpecifiersOpt{
	// · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · IsAuto: true,
	// · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · Case: 3,
	// · · · Token: example101.c:1:11: AUTO "auto",
	// · · },
	// · },
	// }
}

func ExampleDeclarator() {
	fmt.Println(exampleAST(158, "\U00100002 auto a )"))
	// Output:
	// &cc.Declarator{
	// · DirectDeclarator: &cc.DirectDeclarator{
	// · · Token: example158.c:1:11: IDENTIFIER "a",
	// · },
	// }
}

func ExampleDeclaratorOpt() {
	fmt.Println(exampleAST(159, "\U00100002 struct { _Bool :") == (*DeclaratorOpt)(nil))
	// Output:
	// true
}

func ExampleDeclaratorOpt_case1() {
	fmt.Println(exampleAST(160, "\U00100002 struct { _Bool a :"))
	// Output:
	// &cc.DeclaratorOpt{
	// · Declarator: &cc.Declarator{
	// · · SUSpecifier0: &cc.StructOrUnionSpecifier0{
	// · · · StructOrUnion: &cc.StructOrUnion{
	// · · · · Token: example160.c:1:6: STRUCT "struct",
	// · · · },
	// · · },
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example160.c:1:21: IDENTIFIER "a",
	// · · },
	// · },
	// }
}

func ExampleDesignation() {
	fmt.Println(exampleAST(213, "\U00100001 ( _Bool ) { . a = !"))
	// Output:
	// &cc.Designation{
	// · DesignatorList: &cc.DesignatorList{
	// · · Designator: &cc.Designator{
	// · · · Case: 1,
	// · · · Token: example213.c:1:18: '.',
	// · · · Token2: example213.c:1:20: IDENTIFIER "a",
	// · · },
	// · },
	// · Token: example213.c:1:22: '=',
	// }
}

func ExampleDesignationOpt() {
	fmt.Println(exampleAST(214, "\U00100001 ( _Bool ) { !") == (*DesignationOpt)(nil))
	// Output:
	// true
}

func ExampleDesignationOpt_case1() {
	fmt.Println(exampleAST(215, "\U00100001 ( _Bool ) { . a = !"))
	// Output:
	// &cc.DesignationOpt{
	// · Designation: &cc.Designation{
	// · · DesignatorList: &cc.DesignatorList{
	// · · · Designator: &cc.Designator{
	// · · · · Case: 1,
	// · · · · Token: example215.c:1:18: '.',
	// · · · · Token2: example215.c:1:20: IDENTIFIER "a",
	// · · · },
	// · · },
	// · · Token: example215.c:1:22: '=',
	// · },
	// }
}

func ExampleDesignator() {
	fmt.Println(exampleAST(218, "\U00100001 ( _Bool ) { [ 'a' ] ."))
	// Output:
	// &cc.Designator{
	// · ConstantExpression: &cc.ConstantExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example218.c:1:20: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example218.c:1:18: '[',
	// · Token2: example218.c:1:24: ']',
	// }
}

func ExampleDesignator_case1() {
	fmt.Println(exampleAST(219, "\U00100001 ( _Bool ) { . a ."))
	// Output:
	// &cc.Designator{
	// · Case: 1,
	// · Token: example219.c:1:18: '.',
	// · Token2: example219.c:1:20: IDENTIFIER "a",
	// }
}

func ExampleDesignatorList() {
	fmt.Println(exampleAST(216, "\U00100001 ( _Bool ) { . a ."))
	// Output:
	// &cc.DesignatorList{
	// · Designator: &cc.Designator{
	// · · Case: 1,
	// · · Token: example216.c:1:18: '.',
	// · · Token2: example216.c:1:20: IDENTIFIER "a",
	// · },
	// }
}

func ExampleDesignatorList_case1() {
	fmt.Println(exampleAST(217, "\U00100001 ( _Bool ) { . a . b ."))
	// Output:
	// &cc.DesignatorList{
	// · Designator: &cc.Designator{
	// · · Case: 1,
	// · · Token: example217.c:1:18: '.',
	// · · Token2: example217.c:1:20: IDENTIFIER "a",
	// · },
	// · DesignatorList: &cc.DesignatorList{
	// · · Case: 1,
	// · · Designator: &cc.Designator{
	// · · · Case: 1,
	// · · · Token: example217.c:1:22: '.',
	// · · · Token2: example217.c:1:24: IDENTIFIER "b",
	// · · },
	// · },
	// }
}

func ExampleDirectAbstractDeclarator() {
	fmt.Println(exampleAST(198, "\U00100001 ( _Bool ( * ) ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · AbstractDeclarator: &cc.AbstractDeclarator{
	// · · Pointer: &cc.Pointer{
	// · · · Token: example198.c:1:16: '*',
	// · · },
	// · },
	// · Token: example198.c:1:14: '(',
	// · Token2: example198.c:1:18: ')',
	// }
}

func ExampleDirectAbstractDeclarator_case1() {
	fmt.Println(exampleAST(199, "\U00100001 ( _Bool [ ] ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · Case: 1,
	// · Token: example199.c:1:14: '[',
	// · Token2: example199.c:1:16: ']',
	// }
}

func ExampleDirectAbstractDeclarator_case2() {
	fmt.Println(exampleAST(200, "\U00100001 ( _Bool [ const ] ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · Case: 2,
	// · Token: example200.c:1:14: '[',
	// · Token2: example200.c:1:22: ']',
	// · TypeQualifierList: &cc.TypeQualifierList{
	// · · TypeQualifier: &cc.TypeQualifier{
	// · · · Token: example200.c:1:16: CONST "const",
	// · · },
	// · },
	// }
}

func ExampleDirectAbstractDeclarator_case3() {
	fmt.Println(exampleAST(201, "\U00100001 ( _Bool [ static 'a' ] ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example201.c:1:23: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 3,
	// · Token: example201.c:1:14: '[',
	// · Token2: example201.c:1:16: STATIC "static",
	// · Token3: example201.c:1:27: ']',
	// }
}

func ExampleDirectAbstractDeclarator_case4() {
	fmt.Println(exampleAST(202, "\U00100001 ( _Bool [ const static 'a' ] ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example202.c:1:29: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 4,
	// · Token: example202.c:1:14: '[',
	// · Token2: example202.c:1:22: STATIC "static",
	// · Token3: example202.c:1:33: ']',
	// · TypeQualifierList: &cc.TypeQualifierList{
	// · · TypeQualifier: &cc.TypeQualifier{
	// · · · Token: example202.c:1:16: CONST "const",
	// · · },
	// · },
	// }
}

func ExampleDirectAbstractDeclarator_case5() {
	fmt.Println(exampleAST(203, "\U00100001 ( _Bool [ * ] ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · Case: 5,
	// · Token: example203.c:1:14: '[',
	// · Token2: example203.c:1:16: '*',
	// · Token3: example203.c:1:18: ']',
	// }
}

func ExampleDirectAbstractDeclarator_case6() {
	fmt.Println(exampleAST(204, "\U00100001 ( _Bool ( ) ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · Case: 6,
	// · Token: example204.c:1:14: '(',
	// · Token2: example204.c:1:16: ')',
	// }
}

func ExampleDirectAbstractDeclarator_case7() {
	fmt.Println(exampleAST(205, "\U00100001 ( _Bool ( ) ( ) ("))
	// Output:
	// &cc.DirectAbstractDeclarator{
	// · Case: 7,
	// · DirectAbstractDeclarator: &cc.DirectAbstractDeclarator{
	// · · Case: 6,
	// · · Token: example205.c:1:14: '(',
	// · · Token2: example205.c:1:16: ')',
	// · },
	// · Token: example205.c:1:18: '(',
	// · Token2: example205.c:1:20: ')',
	// }
}

func ExampleDirectAbstractDeclaratorOpt() {
	fmt.Println(exampleAST(206, "\U00100001 ( _Bool [") == (*DirectAbstractDeclaratorOpt)(nil))
	// Output:
	// true
}

func ExampleDirectAbstractDeclaratorOpt_case1() {
	fmt.Println(exampleAST(207, "\U00100001 ( _Bool ( ) ["))
	// Output:
	// &cc.DirectAbstractDeclaratorOpt{
	// · DirectAbstractDeclarator: &cc.DirectAbstractDeclarator{
	// · · Case: 6,
	// · · Token: example207.c:1:14: '(',
	// · · Token2: example207.c:1:16: ')',
	// · },
	// }
}

func ExampleDirectDeclarator() {
	fmt.Println(exampleAST(161, "\U00100002 auto a ("))
	// Output:
	// &cc.DirectDeclarator{
	// · Token: example161.c:1:11: IDENTIFIER "a",
	// }
}

func ExampleDirectDeclarator_case1() {
	fmt.Println(exampleAST(162, "\U00100002 auto ( a ) ("))
	// Output:
	// &cc.DirectDeclarator{
	// · Case: 1,
	// · Declarator: &cc.Declarator{
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example162.c:1:13: IDENTIFIER "a",
	// · · },
	// · },
	// · Token: example162.c:1:11: '(',
	// · Token2: example162.c:1:15: ')',
	// }
}

func ExampleDirectDeclarator_case2() {
	fmt.Println(exampleAST(163, "\U00100002 auto a [ ] ("))
	// Output:
	// &cc.DirectDeclarator{
	// · Case: 2,
	// · DirectDeclarator: &cc.DirectDeclarator{
	// · · Token: example163.c:1:11: IDENTIFIER "a",
	// · },
	// · Token: example163.c:1:13: '[',
	// · Token2: example163.c:1:15: ']',
	// }
}

func ExampleDirectDeclarator_case3() {
	fmt.Println(exampleAST(164, "\U00100002 auto a [ static 'b' ] ("))
	// Output:
	// &cc.DirectDeclarator{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example164.c:1:22: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 3,
	// · DirectDeclarator: &cc.DirectDeclarator{
	// · · Token: example164.c:1:11: IDENTIFIER "a",
	// · },
	// · Token: example164.c:1:13: '[',
	// · Token2: example164.c:1:15: STATIC "static",
	// · Token3: example164.c:1:26: ']',
	// }
}

func ExampleDirectDeclarator_case4() {
	fmt.Println(exampleAST(165, "\U00100002 auto a [ const static 'b' ] ("))
	// Output:
	// &cc.DirectDeclarator{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example165.c:1:28: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 4,
	// · DirectDeclarator: &cc.DirectDeclarator{
	// · · Token: example165.c:1:11: IDENTIFIER "a",
	// · },
	// · Token: example165.c:1:13: '[',
	// · Token2: example165.c:1:21: STATIC "static",
	// · Token3: example165.c:1:32: ']',
	// · TypeQualifierList: &cc.TypeQualifierList{
	// · · TypeQualifier: &cc.TypeQualifier{
	// · · · Token: example165.c:1:15: CONST "const",
	// · · },
	// · },
	// }
}

func ExampleDirectDeclarator_case5() {
	fmt.Println(exampleAST(166, "\U00100002 auto a [ * ] ("))
	// Output:
	// &cc.DirectDeclarator{
	// · Case: 5,
	// · DirectDeclarator: &cc.DirectDeclarator{
	// · · Token: example166.c:1:11: IDENTIFIER "a",
	// · },
	// · Token: example166.c:1:13: '[',
	// · Token2: example166.c:1:15: '*',
	// · Token3: example166.c:1:17: ']',
	// }
}

func ExampleDirectDeclarator_case6() {
	fmt.Println(exampleAST(168, "\U00100002 auto a ( ) ("))
	// Output:
	// &cc.DirectDeclarator{
	// · Case: 6,
	// · DirectDeclarator: &cc.DirectDeclarator{
	// · · Token: example168.c:1:11: IDENTIFIER "a",
	// · },
	// · DirectDeclarator2: &cc.DirectDeclarator2{
	// · · Case: 1,
	// · · Token: example168.c:1:15: ')',
	// · },
	// · Token: example168.c:1:13: '(',
	// }
}

func ExampleDirectDeclarator2() {
	fmt.Println(exampleAST(169, "\U00100002 auto a ( auto ) ("))
	// Output:
	// &cc.DirectDeclarator2{
	// · ParameterTypeList: &cc.ParameterTypeList{
	// · · ParameterList: &cc.ParameterList{
	// · · · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · · · Case: 1,
	// · · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · · IsAuto: true,
	// · · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · · Case: 3,
	// · · · · · · Token: example169.c:1:15: AUTO "auto",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example169.c:1:20: ')',
	// }
}

func ExampleDirectDeclarator2_case1() {
	fmt.Println(exampleAST(170, "\U00100002 auto a ( ) ("))
	// Output:
	// &cc.DirectDeclarator2{
	// · Case: 1,
	// · Token: example170.c:1:15: ')',
	// }
}

func ExampleElifGroup() {
	fmt.Println(exampleAST(276, "\U00100000 \n#if ppother  \n#elif ppother  \n#elif"))
	// Output:
	// &cc.ElifGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example276.c:3:7: IDENTIFIER "ppother",
	// },
	// · Token: example276.c:3:2: PPELIF,
	// }
}

func ExampleElifGroupList() {
	fmt.Println(exampleAST(272, "\U00100000 \n#if ppother  \n#elif ppother  \n#elif"))
	// Output:
	// &cc.ElifGroupList{
	// · ElifGroup: &cc.ElifGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example272.c:3:7: IDENTIFIER "ppother",
	// },
	// · · Token: example272.c:3:2: PPELIF,
	// · },
	// }
}

func ExampleElifGroupList_case1() {
	fmt.Println(exampleAST(273, "\U00100000 \n#if ppother  \n#elif ppother  \n#elif ppother  \n#elif"))
	// Output:
	// &cc.ElifGroupList{
	// · ElifGroup: &cc.ElifGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example273.c:3:7: IDENTIFIER "ppother",
	// },
	// · · Token: example273.c:3:2: PPELIF,
	// · },
	// · ElifGroupList: &cc.ElifGroupList{
	// · · Case: 1,
	// · · ElifGroup: &cc.ElifGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example273.c:4:7: IDENTIFIER "ppother",
	// },
	// · · · Token: example273.c:4:2: PPELIF,
	// · · },
	// · },
	// }
}

func ExampleElifGroupListOpt() {
	fmt.Println(exampleAST(274, "\U00100000 \n#if ppother  \n#else") == (*ElifGroupListOpt)(nil))
	// Output:
	// true
}

func ExampleElifGroupListOpt_case1() {
	fmt.Println(exampleAST(275, "\U00100000 \n#if ppother  \n#elif ppother  \n#else"))
	// Output:
	// &cc.ElifGroupListOpt{
	// · ElifGroupList: &cc.ElifGroupList{
	// · · ElifGroup: &cc.ElifGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example275.c:3:7: IDENTIFIER "ppother",
	// },
	// · · · Token: example275.c:3:2: PPELIF,
	// · · },
	// · },
	// }
}

func ExampleElseGroup() {
	fmt.Println(exampleAST(277, "\U00100000 \n#if ppother  \n#else  \n#endif"))
	// Output:
	// &cc.ElseGroup{
	// · Token: example277.c:3:2: PPELSE,
	// · Token2: example277.c:3:8: '\n',
	// }
}

func ExampleElseGroupOpt() {
	fmt.Println(exampleAST(278, "\U00100000 \n#if ppother  \n#endif") == (*ElseGroupOpt)(nil))
	// Output:
	// true
}

func ExampleElseGroupOpt_case1() {
	fmt.Println(exampleAST(279, "\U00100000 \n#if ppother  \n#else  \n#endif"))
	// Output:
	// &cc.ElseGroupOpt{
	// · ElseGroup: &cc.ElseGroup{
	// · · Token: example279.c:3:2: PPELSE,
	// · · Token2: example279.c:3:8: '\n',
	// · },
	// }
}

func ExampleEndifLine() {
	fmt.Println(exampleAST(280, "\U00100000 \n#if ppother  \n#endif "))
	// Output:
	// &cc.EndifLine{
	// · Token: example280.c:3:2: PPENDIF,
	// }
}

func ExampleEnumSpecifier() {
	fmt.Println(exampleAST(147, "\U00100002 enum { a } ("))
	// Output:
	// &cc.EnumSpecifier{
	// · EnumSpecifier0: &cc.EnumSpecifier0{
	// · · Token: example147.c:1:6: ENUM "enum",
	// · },
	// · EnumeratorList: &cc.EnumeratorList{
	// · · Enumerator: &cc.Enumerator{
	// · · · EnumerationConstant: &cc.EnumerationConstant{
	// · · · · Token: example147.c:1:13: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// · Token: example147.c:1:11: '{',
	// · Token2: example147.c:1:15: '}',
	// }
}

func ExampleEnumSpecifier_case1() {
	fmt.Println(exampleAST(148, "\U00100002 enum { a , } ("))
	// Output:
	// &cc.EnumSpecifier{
	// · Case: 1,
	// · EnumSpecifier0: &cc.EnumSpecifier0{
	// · · Token: example148.c:1:6: ENUM "enum",
	// · },
	// · EnumeratorList: &cc.EnumeratorList{
	// · · Enumerator: &cc.Enumerator{
	// · · · EnumerationConstant: &cc.EnumerationConstant{
	// · · · · Token: example148.c:1:13: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// · Token: example148.c:1:11: '{',
	// · Token2: example148.c:1:15: ',',
	// · Token3: example148.c:1:17: '}',
	// }
}

func ExampleEnumSpecifier_case2() {
	fmt.Println(exampleAST(149, "\U00100002 enum a ("))
	// Output:
	// &cc.EnumSpecifier{
	// · Case: 2,
	// · Token: example149.c:1:6: ENUM "enum",
	// · Token2: example149.c:1:11: IDENTIFIER "a",
	// }
}

func ExampleEnumSpecifier0() {
	fmt.Println(exampleAST(146, "\U00100002 enum {"))
	// Output:
	// &cc.EnumSpecifier0{
	// · Token: example146.c:1:6: ENUM "enum",
	// }
}

func ExampleEnumerationConstant() {
	fmt.Println(exampleAST(5, "\U00100002 enum { a ,"))
	// Output:
	// &cc.EnumerationConstant{
	// · Token: example5.c:1:13: IDENTIFIER "a",
	// }
}

func ExampleEnumerator() {
	fmt.Println(exampleAST(152, "\U00100002 enum { a ,"))
	// Output:
	// &cc.Enumerator{
	// · EnumerationConstant: &cc.EnumerationConstant{
	// · · Token: example152.c:1:13: IDENTIFIER "a",
	// · },
	// }
}

func ExampleEnumerator_case1() {
	fmt.Println(exampleAST(153, "\U00100002 enum { a = 'b' ,"))
	// Output:
	// &cc.Enumerator{
	// · Case: 1,
	// · ConstantExpression: &cc.ConstantExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example153.c:1:17: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · EnumerationConstant: &cc.EnumerationConstant{
	// · · Token: example153.c:1:13: IDENTIFIER "a",
	// · },
	// · Token: example153.c:1:15: '=',
	// }
}

func ExampleEnumeratorList() {
	fmt.Println(exampleAST(150, "\U00100002 enum { a ,"))
	// Output:
	// &cc.EnumeratorList{
	// · Enumerator: &cc.Enumerator{
	// · · EnumerationConstant: &cc.EnumerationConstant{
	// · · · Token: example150.c:1:13: IDENTIFIER "a",
	// · · },
	// · },
	// }
}

func ExampleEnumeratorList_case1() {
	fmt.Println(exampleAST(151, "\U00100002 enum { a , b ,"))
	// Output:
	// &cc.EnumeratorList{
	// · Enumerator: &cc.Enumerator{
	// · · EnumerationConstant: &cc.EnumerationConstant{
	// · · · Token: example151.c:1:13: IDENTIFIER "a",
	// · · },
	// · },
	// · EnumeratorList: &cc.EnumeratorList{
	// · · Case: 1,
	// · · Enumerator: &cc.Enumerator{
	// · · · EnumerationConstant: &cc.EnumerationConstant{
	// · · · · Token: example151.c:1:17: IDENTIFIER "b",
	// · · · },
	// · · },
	// · · Token: example151.c:1:15: ',',
	// · },
	// }
}

func ExampleEqualityExpression() {
	fmt.Println(exampleAST(59, "\U00100001 'a'"))
	// Output:
	// &cc.EqualityExpression{
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example59.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleEqualityExpression_case1() {
	fmt.Println(exampleAST(60, "\U00100001 'a' == 'b'"))
	// Output:
	// &cc.EqualityExpression{
	// · Case: 1,
	// · EqualityExpression: &cc.EqualityExpression{
	// · · RelationalExpression: &cc.RelationalExpression{
	// · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · Case: 1,
	// · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · Token: example60.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example60.c:1:13: CHARCONST "'b'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example60.c:1:10: EQ,
	// }
}

func ExampleEqualityExpression_case2() {
	fmt.Println(exampleAST(61, "\U00100001 'a' != 'b'"))
	// Output:
	// &cc.EqualityExpression{
	// · Case: 2,
	// · EqualityExpression: &cc.EqualityExpression{
	// · · RelationalExpression: &cc.RelationalExpression{
	// · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · Case: 1,
	// · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · Token: example61.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example61.c:1:13: CHARCONST "'b'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example61.c:1:10: NEQ,
	// }
}

func ExampleExclusiveOrExpression() {
	fmt.Println(exampleAST(64, "\U00100001 'a'"))
	// Output:
	// &cc.ExclusiveOrExpression{
	// · AndExpression: &cc.AndExpression{
	// · · EqualityExpression: &cc.EqualityExpression{
	// · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · Token: example64.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleExclusiveOrExpression_case1() {
	fmt.Println(exampleAST(65, "\U00100001 'a' ^ 'b'"))
	// Output:
	// &cc.ExclusiveOrExpression{
	// · AndExpression: &cc.AndExpression{
	// · · EqualityExpression: &cc.EqualityExpression{
	// · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · Token: example65.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 1,
	// · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · AndExpression: &cc.AndExpression{
	// · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · Token: example65.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example65.c:1:10: '^',
	// }
}

func ExampleExpressionList() {
	fmt.Println(exampleAST(89, "\U00100001 ( 'a' )"))
	// Output:
	// &cc.ExpressionList{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example89.c:1:8: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleExpressionList_case1() {
	fmt.Println(exampleAST(90, "\U00100001 ( 'a' , 'b' )"))
	// Output:
	// &cc.ExpressionList{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example90.c:1:8: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example90.c:1:14: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Case: 1,
	// · · Token: example90.c:1:12: ',',
	// · },
	// }
}

func ExampleExpressionOpt() {
	fmt.Println(exampleAST(91, "\U00100002 auto a { ;") == (*ExpressionOpt)(nil))
	// Output:
	// true
}

func ExampleExpressionOpt_case1() {
	fmt.Println(exampleAST(92, "\U00100002 auto a { 'b' )"))
	// Output:
	// &cc.ExpressionOpt{
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example92.c:1:15: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleExpressionStatement() {
	fmt.Println(exampleAST(237, "\U00100002 auto a { ; !"))
	// Output:
	// &cc.ExpressionStatement{
	// · Token: example237.c:1:15: ';',
	// }
}

func ExampleExternalDeclaration() {
	fmt.Println(exampleAST(251, "\U00100002 auto a { }"))
	// Output:
	// &cc.ExternalDeclaration{
	// · FunctionDefinition: &cc.FunctionDefinition{
	// · · CompoundStatement: &cc.CompoundStatement{
	// · · · Declarations: &cc.Bindings{
	// · · · · Type: 3,
	// · · · },
	// · · · Token: example251.c:1:13: '{',
	// · · · Token2: example251.c:1:15: '}',
	// · · },
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example251.c:1:6: AUTO "auto",
	// · · · },
	// · · },
	// · · Declarator: &cc.Declarator{
	// · · · IsDefinition: true,
	// · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · Token: example251.c:1:11: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleExternalDeclaration_case1() {
	fmt.Println(exampleAST(252, "\U00100002 auto ;"))
	// Output:
	// &cc.ExternalDeclaration{
	// · Case: 1,
	// · Declaration: &cc.Declaration{
	// · · IsFileScope: true,
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example252.c:1:6: AUTO "auto",
	// · · · },
	// · · },
	// · · Token: example252.c:1:11: ';',
	// · },
	// }
}

func ExampleFunctionDefinition() {
	fmt.Println(exampleAST(254, "\U00100002 auto a { }"))
	// Output:
	// &cc.FunctionDefinition{
	// · CompoundStatement: &cc.CompoundStatement{
	// · · Declarations: &cc.Bindings{
	// · · · Type: 3,
	// · · },
	// · · Token: example254.c:1:13: '{',
	// · · Token2: example254.c:1:15: '}',
	// · },
	// · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · IsAuto: true,
	// · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · Case: 3,
	// · · · Token: example254.c:1:6: AUTO "auto",
	// · · },
	// · },
	// · Declarator: &cc.Declarator{
	// · · IsDefinition: true,
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example254.c:1:11: IDENTIFIER "a",
	// · · },
	// · },
	// }
}

func ExampleFunctionSpecifier() {
	fmt.Println(exampleAST(157, "\U00100002 inline ("))
	// Output:
	// &cc.FunctionSpecifier{
	// · Token: example157.c:1:6: INLINE "inline",
	// }
}

func ExampleGroupListOpt() {
	fmt.Println(exampleAST(262, "\U00100000 \n#ifndef a  \n#elif") == (*GroupListOpt)(nil))
	// Output:
	// true
}

func ExampleGroupListOpt_case1() {
	fmt.Println(exampleAST(263, "\U00100000 \n#if 1 \n a \n#elif"))
	// Output:
	// &cc.GroupListOpt{
	// · GroupList: &cc.GroupList{
	// · · GroupPart: &cc.GroupPart{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example263.c:3:2: IDENTIFIER "a",
	// },
	// · · },
	// · },
	// }
}

func ExampleIdentifierList() {
	fmt.Println(exampleAST(187, "\U00100002 auto a ( b )"))
	// Output:
	// &cc.IdentifierList{
	// · Token: example187.c:1:15: IDENTIFIER "b",
	// }
}

func ExampleIdentifierList_case1() {
	fmt.Println(exampleAST(188, "\U00100002 auto a ( b , c )"))
	// Output:
	// &cc.IdentifierList{
	// · IdentifierList: &cc.IdentifierList{
	// · · Case: 1,
	// · · Token: example188.c:1:17: ',',
	// · · Token2: example188.c:1:19: IDENTIFIER "c",
	// · },
	// · Token: example188.c:1:15: IDENTIFIER "b",
	// }
}

func ExampleIdentifierListOpt() {
	fmt.Println(exampleAST(189, "\U00100002 auto a ( )") == (*IdentifierListOpt)(nil))
	// Output:
	// true
}

func ExampleIdentifierListOpt_case1() {
	fmt.Println(exampleAST(190, "\U00100002 auto a ( b )"))
	// Output:
	// &cc.IdentifierListOpt{
	// · IdentifierList: &cc.IdentifierList{
	// · · Token: example190.c:1:15: IDENTIFIER "b",
	// · },
	// }
}

func ExampleIdentifierOpt() {
	fmt.Println(exampleAST(191, "\U00100002 struct {") == (*IdentifierOpt)(nil))
	// Output:
	// true
}

func ExampleIdentifierOpt_case1() {
	fmt.Println(exampleAST(192, "\U00100002 enum a {"))
	// Output:
	// &cc.IdentifierOpt{
	// · Token: example192.c:1:11: IDENTIFIER "a",
	// }
}

func ExampleIfGroup() {
	fmt.Println(exampleAST(269, "\U00100000 \n#if ppother  \n#elif"))
	// Output:
	// &cc.IfGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example269.c:2:5: IDENTIFIER "ppother",
	// },
	// · Token: example269.c:2:2: PPIF,
	// }
}

func ExampleIfGroup_case1() {
	fmt.Println(exampleAST(270, "\U00100000 \n#ifdef a  \n#elif"))
	// Output:
	// &cc.IfGroup{
	// · Case: 1,
	// · Token: example270.c:2:2: PPIFDEF,
	// · Token2: example270.c:2:8: IDENTIFIER "a",
	// · Token3: example270.c:2:11: '\n',
	// }
}

func ExampleIfGroup_case2() {
	fmt.Println(exampleAST(271, "\U00100000 \n#ifndef a  \n#elif"))
	// Output:
	// &cc.IfGroup{
	// · Case: 2,
	// · Token: example271.c:2:2: PPIFNDEF,
	// · Token2: example271.c:2:9: IDENTIFIER "a",
	// · Token3: example271.c:2:12: '\n',
	// }
}

func ExampleIfSection() {
	fmt.Println(exampleAST(268, "\U00100000 \n#if ppother  \n#endif "))
	// Output:
	// &cc.IfSection{
	// · EndifLine: &cc.EndifLine{
	// · · Token: example268.c:3:2: PPENDIF,
	// · },
	// · IfGroup: &cc.IfGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example268.c:2:5: IDENTIFIER "ppother",
	// },
	// · · Token: example268.c:2:2: PPIF,
	// · },
	// }
}

func ExampleInclusiveOrExpression() {
	fmt.Println(exampleAST(66, "\U00100001 'a'"))
	// Output:
	// &cc.InclusiveOrExpression{
	// · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · AndExpression: &cc.AndExpression{
	// · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · Token: example66.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleInclusiveOrExpression_case1() {
	fmt.Println(exampleAST(67, "\U00100001 'a' | 'b'"))
	// Output:
	// &cc.InclusiveOrExpression{
	// · Case: 1,
	// · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · AndExpression: &cc.AndExpression{
	// · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · Token: example67.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · AndExpression: &cc.AndExpression{
	// · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · Token: example67.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example67.c:1:10: '|',
	// }
}

func ExampleInitDeclarator() {
	fmt.Println(exampleAST(106, "\U00100002 auto a ,"))
	// Output:
	// &cc.InitDeclarator{
	// · Declarator: &cc.Declarator{
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example106.c:1:11: IDENTIFIER "a",
	// · · },
	// · },
	// }
}

func ExampleInitDeclarator_case1() {
	fmt.Println(exampleAST(107, "\U00100002 auto a = 'b' ,"))
	// Output:
	// &cc.InitDeclarator{
	// · Case: 1,
	// · Declarator: &cc.Declarator{
	// · · IsDefinition: true,
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example107.c:1:11: IDENTIFIER "a",
	// · · },
	// · },
	// · Initializer: &cc.Initializer{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example107.c:1:15: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example107.c:1:13: '=',
	// }
}

func ExampleInitDeclaratorList() {
	fmt.Println(exampleAST(102, "\U00100002 auto a ,"))
	// Output:
	// &cc.InitDeclaratorList{
	// · InitDeclarator: &cc.InitDeclarator{
	// · · Declarator: &cc.Declarator{
	// · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · Token: example102.c:1:11: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleInitDeclaratorList_case1() {
	fmt.Println(exampleAST(103, "\U00100002 auto a , b ,"))
	// Output:
	// &cc.InitDeclaratorList{
	// · InitDeclarator: &cc.InitDeclarator{
	// · · Declarator: &cc.Declarator{
	// · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · Token: example103.c:1:11: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// · InitDeclaratorList: &cc.InitDeclaratorList{
	// · · Case: 1,
	// · · InitDeclarator: &cc.InitDeclarator{
	// · · · Declarator: &cc.Declarator{
	// · · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · · Token: example103.c:1:15: IDENTIFIER "b",
	// · · · · },
	// · · · },
	// · · },
	// · · Token: example103.c:1:13: ',',
	// · },
	// }
}

func ExampleInitDeclaratorListOpt() {
	fmt.Println(exampleAST(104, "\U00100002 auto ;") == (*InitDeclaratorListOpt)(nil))
	// Output:
	// true
}

func ExampleInitDeclaratorListOpt_case1() {
	fmt.Println(exampleAST(105, "\U00100002 auto a ;"))
	// Output:
	// &cc.InitDeclaratorListOpt{
	// · InitDeclaratorList: &cc.InitDeclaratorList{
	// · · InitDeclarator: &cc.InitDeclarator{
	// · · · Declarator: &cc.Declarator{
	// · · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · · Token: example105.c:1:11: IDENTIFIER "a",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleInitializer() {
	fmt.Println(exampleAST(208, "\U00100002 auto a = 'b' ,"))
	// Output:
	// &cc.Initializer{
	// · AssignmentExpression: &cc.AssignmentExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example208.c:1:15: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleInitializer_case1() {
	fmt.Println(exampleAST(209, "\U00100002 auto a = { 'b' } ,"))
	// Output:
	// &cc.Initializer{
	// · Case: 1,
	// · InitializerList: &cc.InitializerList{
	// · · Initializer: &cc.Initializer{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example209.c:1:17: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example209.c:1:15: '{',
	// · Token2: example209.c:1:21: '}',
	// }
}

func ExampleInitializer_case2() {
	fmt.Println(exampleAST(210, "\U00100002 auto a = { 'b' , } ,"))
	// Output:
	// &cc.Initializer{
	// · Case: 2,
	// · InitializerList: &cc.InitializerList{
	// · · Initializer: &cc.Initializer{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example210.c:1:17: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example210.c:1:15: '{',
	// · Token2: example210.c:1:21: ',',
	// · Token3: example210.c:1:23: '}',
	// }
}

func ExampleInitializerList() {
	fmt.Println(exampleAST(211, "\U00100001 ( _Bool ) { 'a' ,"))
	// Output:
	// &cc.InitializerList{
	// · Initializer: &cc.Initializer{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example211.c:1:18: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleInitializerList_case1() {
	fmt.Println(exampleAST(212, "\U00100002 auto a = { 'b' , 'c' ,"))
	// Output:
	// &cc.InitializerList{
	// · Initializer: &cc.Initializer{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example212.c:1:17: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · InitializerList: &cc.InitializerList{
	// · · Case: 1,
	// · · Initializer: &cc.Initializer{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example212.c:1:23: CHARCONST "'c'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Token: example212.c:1:21: ',',
	// · },
	// }
}

func ExampleIterationStatement() {
	fmt.Println(exampleAST(241, "\U00100002 auto a { while ( 'b' ) ; !"))
	// Output:
	// &cc.IterationStatement{
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example241.c:1:23: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example241.c:1:29: ';',
	// · · },
	// · },
	// · Token: example241.c:1:15: WHILE "while",
	// · Token2: example241.c:1:21: '(',
	// · Token3: example241.c:1:27: ')',
	// }
}

func ExampleIterationStatement_case1() {
	fmt.Println(exampleAST(242, "\U00100002 auto a { do ; while ( 'b' ) ; !"))
	// Output:
	// &cc.IterationStatement{
	// · Case: 1,
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example242.c:1:28: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example242.c:1:18: ';',
	// · · },
	// · },
	// · Token: example242.c:1:15: DO "do",
	// · Token2: example242.c:1:20: WHILE "while",
	// · Token3: example242.c:1:26: '(',
	// · Token4: example242.c:1:32: ')',
	// · Token5: example242.c:1:34: ';',
	// }
}

func ExampleIterationStatement_case2() {
	fmt.Println(exampleAST(243, "\U00100002 auto a { for ( ; ; ) ; !"))
	// Output:
	// &cc.IterationStatement{
	// · Case: 2,
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example243.c:1:27: ';',
	// · · },
	// · },
	// · Token: example243.c:1:15: FOR "for",
	// · Token2: example243.c:1:19: '(',
	// · Token3: example243.c:1:21: ';',
	// · Token4: example243.c:1:23: ';',
	// · Token5: example243.c:1:25: ')',
	// }
}

func ExampleIterationStatement_case3() {
	fmt.Println(exampleAST(244, "\U00100002 auto a { for ( auto ; ; ) ; !"))
	// Output:
	// &cc.IterationStatement{
	// · Case: 3,
	// · Declaration: &cc.Declaration{
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example244.c:1:21: AUTO "auto",
	// · · · },
	// · · },
	// · · Token: example244.c:1:26: ';',
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example244.c:1:32: ';',
	// · · },
	// · },
	// · Token: example244.c:1:15: FOR "for",
	// · Token2: example244.c:1:19: '(',
	// · Token3: example244.c:1:28: ';',
	// · Token4: example244.c:1:30: ')',
	// }
}

func ExampleJumpStatement() {
	fmt.Println(exampleAST(245, "\U00100002 auto a { goto b ; !"))
	// Output:
	// &cc.JumpStatement{
	// · Token: example245.c:1:15: GOTO "goto",
	// · Token2: example245.c:1:20: IDENTIFIER "b",
	// · Token3: example245.c:1:22: ';',
	// }
}

func ExampleJumpStatement_case1() {
	fmt.Println(exampleAST(246, "\U00100002 auto a { continue ; !"))
	// Output:
	// &cc.JumpStatement{
	// · Case: 1,
	// · Token: example246.c:1:15: CONTINUE "continue",
	// · Token2: example246.c:1:24: ';',
	// }
}

func ExampleJumpStatement_case2() {
	fmt.Println(exampleAST(247, "\U00100002 auto a { break ; !"))
	// Output:
	// &cc.JumpStatement{
	// · Case: 2,
	// · Token: example247.c:1:15: BREAK "break",
	// · Token2: example247.c:1:21: ';',
	// }
}

func ExampleJumpStatement_case3() {
	fmt.Println(exampleAST(248, "\U00100002 auto a { return ; !"))
	// Output:
	// &cc.JumpStatement{
	// · Case: 3,
	// · Token: example248.c:1:15: RETURN "return",
	// · Token2: example248.c:1:22: ';',
	// }
}

func ExampleLabeledStatement() {
	fmt.Println(exampleAST(226, "\U00100002 auto a { b : ; !"))
	// Output:
	// &cc.LabeledStatement{
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example226.c:1:19: ';',
	// · · },
	// · },
	// · Token: example226.c:1:15: IDENTIFIER "b",
	// · Token2: example226.c:1:17: ':',
	// }
}

func ExampleLabeledStatement_case1() {
	fmt.Println(exampleAST(227, "\U00100002 auto a { case 'b' : ; !"))
	// Output:
	// &cc.LabeledStatement{
	// · Case: 1,
	// · ConstantExpression: &cc.ConstantExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example227.c:1:20: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example227.c:1:26: ';',
	// · · },
	// · },
	// · Token: example227.c:1:15: CASE "case",
	// · Token2: example227.c:1:24: ':',
	// }
}

func ExampleLabeledStatement_case2() {
	fmt.Println(exampleAST(228, "\U00100002 auto a { default : ; !"))
	// Output:
	// &cc.LabeledStatement{
	// · Case: 2,
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example228.c:1:25: ';',
	// · · },
	// · },
	// · Token: example228.c:1:15: DEFAULT "default",
	// · Token2: example228.c:1:23: ':',
	// }
}

func ExampleLogicalAndExpression() {
	fmt.Println(exampleAST(68, "\U00100001 'a'"))
	// Output:
	// &cc.LogicalAndExpression{
	// · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · AndExpression: &cc.AndExpression{
	// · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · Token: example68.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleLogicalAndExpression_case1() {
	fmt.Println(exampleAST(69, "\U00100001 'a' && 'b'"))
	// Output:
	// &cc.LogicalAndExpression{
	// · Case: 1,
	// · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · AndExpression: &cc.AndExpression{
	// · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · Token: example69.c:1:13: CHARCONST "'b'",
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · AndExpression: &cc.AndExpression{
	// · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · Token: example69.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example69.c:1:10: ANDAND,
	// }
}

func ExampleLogicalOrExpression() {
	fmt.Println(exampleAST(70, "\U00100001 'a'"))
	// Output:
	// &cc.LogicalOrExpression{
	// · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · AndExpression: &cc.AndExpression{
	// · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · Token: example70.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleLogicalOrExpression_case1() {
	fmt.Println(exampleAST(71, "\U00100001 'a' || 'b'"))
	// Output:
	// &cc.LogicalOrExpression{
	// · Case: 1,
	// · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · AndExpression: &cc.AndExpression{
	// · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · Token: example71.c:1:13: CHARCONST "'b'",
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · Token: example71.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example71.c:1:10: OROR,
	// }
}

func ExampleMultiplicativeExpression() {
	fmt.Println(exampleAST(44, "\U00100001 'a'"))
	// Output:
	// &cc.MultiplicativeExpression{
	// · CastExpression: &cc.CastExpression{
	// · · UnaryExpression: &cc.UnaryExpression{
	// · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · Case: 1,
	// · · · · · Constant: &cc.Constant{
	// · · · · · · Token: example44.c:1:6: CHARCONST "'a'",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleMultiplicativeExpression_case1() {
	fmt.Println(exampleAST(45, "\U00100001 'a' * 'b'"))
	// Output:
	// &cc.MultiplicativeExpression{
	// · Case: 1,
	// · CastExpression: &cc.CastExpression{
	// · · UnaryExpression: &cc.UnaryExpression{
	// · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · Case: 1,
	// · · · · · Constant: &cc.Constant{
	// · · · · · · Token: example45.c:1:12: CHARCONST "'b'",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · CastExpression: &cc.CastExpression{
	// · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · Case: 1,
	// · · · · · · Constant: &cc.Constant{
	// · · · · · · · Token: example45.c:1:6: CHARCONST "'a'",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example45.c:1:10: '*',
	// }
}

func ExampleMultiplicativeExpression_case2() {
	fmt.Println(exampleAST(46, "\U00100001 'a' / 'b'"))
	// Output:
	// &cc.MultiplicativeExpression{
	// · Case: 2,
	// · CastExpression: &cc.CastExpression{
	// · · UnaryExpression: &cc.UnaryExpression{
	// · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · Case: 1,
	// · · · · · Constant: &cc.Constant{
	// · · · · · · Token: example46.c:1:12: CHARCONST "'b'",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · CastExpression: &cc.CastExpression{
	// · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · Case: 1,
	// · · · · · · Constant: &cc.Constant{
	// · · · · · · · Token: example46.c:1:6: CHARCONST "'a'",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example46.c:1:10: '/',
	// }
}

func ExampleMultiplicativeExpression_case3() {
	fmt.Println(exampleAST(47, "\U00100001 'a' % 'b'"))
	// Output:
	// &cc.MultiplicativeExpression{
	// · Case: 3,
	// · CastExpression: &cc.CastExpression{
	// · · UnaryExpression: &cc.UnaryExpression{
	// · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · Case: 1,
	// · · · · · Constant: &cc.Constant{
	// · · · · · · Token: example47.c:1:12: CHARCONST "'b'",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · CastExpression: &cc.CastExpression{
	// · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · Case: 1,
	// · · · · · · Constant: &cc.Constant{
	// · · · · · · · Token: example47.c:1:6: CHARCONST "'a'",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example47.c:1:10: '%',
	// }
}

func ExampleParameterDeclaration() {
	fmt.Println(exampleAST(185, "\U00100001 ( _Bool ( auto a )"))
	// Output:
	// &cc.ParameterDeclaration{
	// · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · IsAuto: true,
	// · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · Case: 3,
	// · · · Token: example185.c:1:16: AUTO "auto",
	// · · },
	// · },
	// · Declarator: &cc.Declarator{
	// · · IsDefinition: true,
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example185.c:1:21: IDENTIFIER "a",
	// · · },
	// · },
	// }
}

func ExampleParameterDeclaration_case1() {
	fmt.Println(exampleAST(186, "\U00100001 ( _Bool ( auto )"))
	// Output:
	// &cc.ParameterDeclaration{
	// · Case: 1,
	// · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · IsAuto: true,
	// · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · Case: 3,
	// · · · Token: example186.c:1:16: AUTO "auto",
	// · · },
	// · },
	// }
}

func ExampleParameterList() {
	fmt.Println(exampleAST(183, "\U00100001 ( _Bool ( auto )"))
	// Output:
	// &cc.ParameterList{
	// · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · Case: 1,
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example183.c:1:16: AUTO "auto",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleParameterList_case1() {
	fmt.Println(exampleAST(184, "\U00100001 ( _Bool ( auto , auto )"))
	// Output:
	// &cc.ParameterList{
	// · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · Case: 1,
	// · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · IsAuto: true,
	// · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · Case: 3,
	// · · · · Token: example184.c:1:16: AUTO "auto",
	// · · · },
	// · · },
	// · },
	// · ParameterList: &cc.ParameterList{
	// · · Case: 1,
	// · · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · · Case: 1,
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example184.c:1:23: AUTO "auto",
	// · · · · },
	// · · · },
	// · · },
	// · · Token: example184.c:1:21: ',',
	// · },
	// }
}

func ExampleParameterTypeList() {
	fmt.Println(exampleAST(179, "\U00100001 ( _Bool ( auto )"))
	// Output:
	// &cc.ParameterTypeList{
	// · ParameterList: &cc.ParameterList{
	// · · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · · Case: 1,
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example179.c:1:16: AUTO "auto",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleParameterTypeList_case1() {
	fmt.Println(exampleAST(180, "\U00100001 ( _Bool ( auto , ... )"))
	// Output:
	// &cc.ParameterTypeList{
	// · Case: 1,
	// · ParameterList: &cc.ParameterList{
	// · · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · · Case: 1,
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example180.c:1:16: AUTO "auto",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example180.c:1:21: ',',
	// · Token2: example180.c:1:23: DDD,
	// }
}

func ExampleParameterTypeListOpt() {
	fmt.Println(exampleAST(181, "\U00100001 ( _Bool ( )") == (*ParameterTypeListOpt)(nil))
	// Output:
	// true
}

func ExampleParameterTypeListOpt_case1() {
	fmt.Println(exampleAST(182, "\U00100001 ( _Bool ( auto )"))
	// Output:
	// &cc.ParameterTypeListOpt{
	// · ParameterTypeList: &cc.ParameterTypeList{
	// · · ParameterList: &cc.ParameterList{
	// · · · ParameterDeclaration: &cc.ParameterDeclaration{
	// · · · · Case: 1,
	// · · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · · IsAuto: true,
	// · · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · · Case: 3,
	// · · · · · · Token: example182.c:1:16: AUTO "auto",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExamplePointer() {
	fmt.Println(exampleAST(171, "\U00100002 auto * ("))
	// Output:
	// &cc.Pointer{
	// · Token: example171.c:1:11: '*',
	// }
}

func ExamplePointer_case1() {
	fmt.Println(exampleAST(172, "\U00100002 auto * * ("))
	// Output:
	// &cc.Pointer{
	// · Case: 1,
	// · Pointer: &cc.Pointer{
	// · · Token: example172.c:1:13: '*',
	// · },
	// · Token: example172.c:1:11: '*',
	// }
}

func ExamplePointerOpt() {
	fmt.Println(exampleAST(173, "\U00100002 auto (") == (*PointerOpt)(nil))
	// Output:
	// true
}

func ExamplePointerOpt_case1() {
	fmt.Println(exampleAST(174, "\U00100001 ( _Bool * ("))
	// Output:
	// &cc.PointerOpt{
	// · Pointer: &cc.Pointer{
	// · · Token: example174.c:1:14: '*',
	// · },
	// }
}

func ExamplePostfixExpression() {
	fmt.Println(exampleAST(15, "\U00100001 'a'"))
	// Output:
	// &cc.PostfixExpression{
	// · PrimaryExpression: &cc.PrimaryExpression{
	// · · Case: 1,
	// · · Constant: &cc.Constant{
	// · · · Token: example15.c:1:6: CHARCONST "'a'",
	// · · },
	// · },
	// }
}

func ExamplePostfixExpression_case1() {
	fmt.Println(exampleAST(16, "\U00100001 'a' [ 'b' ]"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 1,
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example16.c:1:12: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example16.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// · Token: example16.c:1:10: '[',
	// · Token2: example16.c:1:16: ']',
	// }
}

func ExamplePostfixExpression_case2() {
	fmt.Println(exampleAST(17, "\U00100001 'a' ( )"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 2,
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example17.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// · Token: example17.c:1:10: '(',
	// · Token2: example17.c:1:12: ')',
	// }
}

func ExamplePostfixExpression_case3() {
	fmt.Println(exampleAST(18, "\U00100001 'a' . b"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 3,
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example18.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// · Token: example18.c:1:10: '.',
	// · Token2: example18.c:1:12: IDENTIFIER "b",
	// }
}

func ExamplePostfixExpression_case4() {
	fmt.Println(exampleAST(19, "\U00100001 'a' -> b"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 4,
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example19.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// · Token: example19.c:1:10: ARROW,
	// · Token2: example19.c:1:13: IDENTIFIER "b",
	// }
}

func ExamplePostfixExpression_case5() {
	fmt.Println(exampleAST(20, "\U00100001 'a' ++"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 5,
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example20.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// · Token: example20.c:1:10: INC,
	// }
}

func ExamplePostfixExpression_case6() {
	fmt.Println(exampleAST(21, "\U00100001 'a' --"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 6,
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example21.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// · Token: example21.c:1:10: DEC,
	// }
}

func ExamplePostfixExpression_case7() {
	fmt.Println(exampleAST(22, "\U00100001 ( _Bool ) { 'a' }"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 7,
	// · InitializerList: &cc.InitializerList{
	// · · Initializer: &cc.Initializer{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example22.c:1:18: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example22.c:1:6: '(',
	// · Token2: example22.c:1:14: ')',
	// · Token3: example22.c:1:16: '{',
	// · Token4: example22.c:1:22: '}',
	// · TypeName: &cc.TypeName{
	// · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · Case: 9,
	// · · · · Token: example22.c:1:8: BOOL "_Bool",
	// · · · },
	// · · },
	// · },
	// }
}

func ExamplePostfixExpression_case8() {
	fmt.Println(exampleAST(23, "\U00100001 ( _Bool ) { 'a' , }"))
	// Output:
	// &cc.PostfixExpression{
	// · Case: 8,
	// · InitializerList: &cc.InitializerList{
	// · · Initializer: &cc.Initializer{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example23.c:1:18: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example23.c:1:6: '(',
	// · Token2: example23.c:1:14: ')',
	// · Token3: example23.c:1:16: '{',
	// · Token4: example23.c:1:22: ',',
	// · Token5: example23.c:1:24: '}',
	// · TypeName: &cc.TypeName{
	// · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · Case: 9,
	// · · · · Token: example23.c:1:8: BOOL "_Bool",
	// · · · },
	// · · },
	// · },
	// }
}

func ExamplePreprocessingFile() {
	fmt.Println(exampleAST(259, "\U00100000 #if 0\n#endif"))
	// Output:
	// &cc.PreprocessingFile{
	// · GroupList: &cc.GroupList{
	// · · GroupPart: &cc.GroupPart{
	// · · · IfSection: &cc.IfSection{
	// · · · · EndifLine: &cc.EndifLine{
	// · · · · · Token: example259.c:2:2: PPENDIF,
	// · · · · },
	// · · · · IfGroup: &cc.IfGroup{
	// PpTokenList: []xc.Token{ // len 1
	// · 0: example259.c:1:10: INTCONST "0",
	// },
	// · · · · · Token: example259.c:1:7: PPIF,
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExamplePrimaryExpression() {
	fmt.Println(exampleAST(6, "\U00100001 a"))
	// Output:
	// &cc.PrimaryExpression{
	// · Token: example6.c:1:6: IDENTIFIER "a",
	// }
}

func ExamplePrimaryExpression_case1() {
	fmt.Println(exampleAST(7, "\U00100001 'a'"))
	// Output:
	// &cc.PrimaryExpression{
	// · Case: 1,
	// · Constant: &cc.Constant{
	// · · Token: example7.c:1:6: CHARCONST "'a'",
	// · },
	// }
}

func ExamplePrimaryExpression_case2() {
	fmt.Println(exampleAST(8, "\U00100001 ( 'a' )"))
	// Output:
	// &cc.PrimaryExpression{
	// · Case: 2,
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example8.c:1:8: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example8.c:1:6: '(',
	// · Token2: example8.c:1:12: ')',
	// }
}

func ExampleRelationalExpression() {
	fmt.Println(exampleAST(54, "\U00100001 'a'"))
	// Output:
	// &cc.RelationalExpression{
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example54.c:1:6: CHARCONST "'a'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleRelationalExpression_case1() {
	fmt.Println(exampleAST(55, "\U00100001 'a' < 'b'"))
	// Output:
	// &cc.RelationalExpression{
	// · Case: 1,
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example55.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example55.c:1:12: CHARCONST "'b'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example55.c:1:10: '<',
	// }
}

func ExampleRelationalExpression_case2() {
	fmt.Println(exampleAST(56, "\U00100001 'a' > 'b'"))
	// Output:
	// &cc.RelationalExpression{
	// · Case: 2,
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example56.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example56.c:1:12: CHARCONST "'b'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example56.c:1:10: '>',
	// }
}

func ExampleRelationalExpression_case3() {
	fmt.Println(exampleAST(57, "\U00100001 'a' <= 'b'"))
	// Output:
	// &cc.RelationalExpression{
	// · Case: 3,
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example57.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example57.c:1:13: CHARCONST "'b'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example57.c:1:10: LEQ,
	// }
}

func ExampleRelationalExpression_case4() {
	fmt.Println(exampleAST(58, "\U00100001 'a' >= 'b'"))
	// Output:
	// &cc.RelationalExpression{
	// · Case: 4,
	// · RelationalExpression: &cc.RelationalExpression{
	// · · ShiftExpression: &cc.ShiftExpression{
	// · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · Case: 1,
	// · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · Token: example58.c:1:6: CHARCONST "'a'",
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example58.c:1:13: CHARCONST "'b'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example58.c:1:10: GEQ,
	// }
}

func ExampleSelectionStatement() {
	fmt.Println(exampleAST(238, "\U00100002 auto a { if ( 'b' ) ; !"))
	// Output:
	// &cc.SelectionStatement{
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example238.c:1:20: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example238.c:1:26: ';',
	// · · },
	// · },
	// · Token: example238.c:1:15: IF "if",
	// · Token2: example238.c:1:18: '(',
	// · Token3: example238.c:1:24: ')',
	// }
}

func ExampleSelectionStatement_case1() {
	fmt.Println(exampleAST(239, "\U00100002 auto a { if ( 'b' ) ; else ; !"))
	// Output:
	// &cc.SelectionStatement{
	// · Case: 1,
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example239.c:1:20: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example239.c:1:26: ';',
	// · · },
	// · },
	// · Statement2: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example239.c:1:33: ';',
	// · · },
	// · },
	// · Token: example239.c:1:15: IF "if",
	// · Token2: example239.c:1:18: '(',
	// · Token3: example239.c:1:24: ')',
	// · Token4: example239.c:1:28: ELSE "else",
	// }
}

func ExampleSelectionStatement_case2() {
	fmt.Println(exampleAST(240, "\U00100002 auto a { switch ( 'b' ) ; !"))
	// Output:
	// &cc.SelectionStatement{
	// · Case: 2,
	// · ExpressionList: &cc.ExpressionList{
	// · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · Token: example240.c:1:24: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Statement: &cc.Statement{
	// · · Case: 2,
	// · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · Token: example240.c:1:30: ';',
	// · · },
	// · },
	// · Token: example240.c:1:15: SWITCH "switch",
	// · Token2: example240.c:1:22: '(',
	// · Token3: example240.c:1:28: ')',
	// }
}

func ExampleShiftExpression() {
	fmt.Println(exampleAST(51, "\U00100001 'a'"))
	// Output:
	// &cc.ShiftExpression{
	// · AdditiveExpression: &cc.AdditiveExpression{
	// · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · CastExpression: &cc.CastExpression{
	// · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · Case: 1,
	// · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · Token: example51.c:1:6: CHARCONST "'a'",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleShiftExpression_case1() {
	fmt.Println(exampleAST(52, "\U00100001 'a' << 'b'"))
	// Output:
	// &cc.ShiftExpression{
	// · AdditiveExpression: &cc.AdditiveExpression{
	// · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · CastExpression: &cc.CastExpression{
	// · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · Case: 1,
	// · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · Token: example52.c:1:13: CHARCONST "'b'",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 1,
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example52.c:1:6: CHARCONST "'a'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example52.c:1:10: LSH,
	// }
}

func ExampleShiftExpression_case2() {
	fmt.Println(exampleAST(53, "\U00100001 'a' >> 'b'"))
	// Output:
	// &cc.ShiftExpression{
	// · AdditiveExpression: &cc.AdditiveExpression{
	// · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · CastExpression: &cc.CastExpression{
	// · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · Case: 1,
	// · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · Token: example53.c:1:13: CHARCONST "'b'",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Case: 2,
	// · ShiftExpression: &cc.ShiftExpression{
	// · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · CastExpression: &cc.CastExpression{
	// · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · Case: 1,
	// · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · Token: example53.c:1:6: CHARCONST "'a'",
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example53.c:1:10: RSH,
	// }
}

func ExampleSpecifierQualifierList() {
	fmt.Println(exampleAST(136, "\U00100001 ( _Bool ("))
	// Output:
	// &cc.SpecifierQualifierList{
	// · TypeSpecifier: &cc.TypeSpecifier{
	// · · Case: 9,
	// · · Token: example136.c:1:8: BOOL "_Bool",
	// · },
	// }
}

func ExampleSpecifierQualifierList_case1() {
	fmt.Println(exampleAST(137, "\U00100001 ( const ("))
	// Output:
	// &cc.SpecifierQualifierList{
	// · IsConst: true,
	// · Case: 1,
	// · TypeQualifier: &cc.TypeQualifier{
	// · · Token: example137.c:1:8: CONST "const",
	// · },
	// }
}

func ExampleSpecifierQualifierListOpt() {
	fmt.Println(exampleAST(138, "\U00100001 ( const (") == (*SpecifierQualifierListOpt)(nil))
	// Output:
	// true
}

func ExampleSpecifierQualifierListOpt_case1() {
	fmt.Println(exampleAST(139, "\U00100001 ( const _Bool ("))
	// Output:
	// &cc.SpecifierQualifierListOpt{
	// · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · Case: 9,
	// · · · Token: example139.c:1:14: BOOL "_Bool",
	// · · },
	// · },
	// }
}

func ExampleStatement() {
	fmt.Println(exampleAST(220, "\U00100002 auto a { default : ; !"))
	// Output:
	// &cc.Statement{
	// · LabeledStatement: &cc.LabeledStatement{
	// · · Case: 2,
	// · · Statement: &cc.Statement{
	// · · · Case: 2,
	// · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · Token: example220.c:1:25: ';',
	// · · · },
	// · · },
	// · · Token: example220.c:1:15: DEFAULT "default",
	// · · Token2: example220.c:1:23: ':',
	// · },
	// }
}

func ExampleStatement_case1() {
	fmt.Println(exampleAST(221, "\U00100002 auto a { { } !"))
	// Output:
	// &cc.Statement{
	// · Case: 1,
	// · CompoundStatement: &cc.CompoundStatement{
	// · · Declarations: &cc.Bindings{
	// · · · Type: 3,
	// · · },
	// · · Token: example221.c:1:15: '{',
	// · · Token2: example221.c:1:17: '}',
	// · },
	// }
}

func ExampleStatement_case2() {
	fmt.Println(exampleAST(222, "\U00100002 auto a { ; !"))
	// Output:
	// &cc.Statement{
	// · Case: 2,
	// · ExpressionStatement: &cc.ExpressionStatement{
	// · · Token: example222.c:1:15: ';',
	// · },
	// }
}

func ExampleStatement_case3() {
	fmt.Println(exampleAST(223, "\U00100002 auto a { if ( 'b' ) ; !"))
	// Output:
	// &cc.Statement{
	// · Case: 3,
	// · SelectionStatement: &cc.SelectionStatement{
	// · · ExpressionList: &cc.ExpressionList{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example223.c:1:20: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Statement: &cc.Statement{
	// · · · Case: 2,
	// · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · Token: example223.c:1:26: ';',
	// · · · },
	// · · },
	// · · Token: example223.c:1:15: IF "if",
	// · · Token2: example223.c:1:18: '(',
	// · · Token3: example223.c:1:24: ')',
	// · },
	// }
}

func ExampleStatement_case4() {
	fmt.Println(exampleAST(224, "\U00100002 auto a { while ( 'b' ) ; !"))
	// Output:
	// &cc.Statement{
	// · Case: 4,
	// · IterationStatement: &cc.IterationStatement{
	// · · ExpressionList: &cc.ExpressionList{
	// · · · AssignmentExpression: &cc.AssignmentExpression{
	// · · · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · · · Token: example224.c:1:23: CHARCONST "'b'",
	// · · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Statement: &cc.Statement{
	// · · · Case: 2,
	// · · · ExpressionStatement: &cc.ExpressionStatement{
	// · · · · Token: example224.c:1:29: ';',
	// · · · },
	// · · },
	// · · Token: example224.c:1:15: WHILE "while",
	// · · Token2: example224.c:1:21: '(',
	// · · Token3: example224.c:1:27: ')',
	// · },
	// }
}

func ExampleStatement_case5() {
	fmt.Println(exampleAST(225, "\U00100002 auto a { break ; !"))
	// Output:
	// &cc.Statement{
	// · Case: 5,
	// · JumpStatement: &cc.JumpStatement{
	// · · Case: 2,
	// · · Token: example225.c:1:15: BREAK "break",
	// · · Token2: example225.c:1:21: ';',
	// · },
	// }
}

func ExampleStorageClassSpecifier() {
	fmt.Println(exampleAST(108, "\U00100002 typedef ("))
	// Output:
	// &cc.StorageClassSpecifier{
	// · Token: example108.c:1:6: TYPEDEF "typedef",
	// }
}

func ExampleStorageClassSpecifier_case1() {
	fmt.Println(exampleAST(109, "\U00100002 extern ("))
	// Output:
	// &cc.StorageClassSpecifier{
	// · Case: 1,
	// · Token: example109.c:1:6: EXTERN "extern",
	// }
}

func ExampleStorageClassSpecifier_case2() {
	fmt.Println(exampleAST(110, "\U00100002 static ("))
	// Output:
	// &cc.StorageClassSpecifier{
	// · Case: 2,
	// · Token: example110.c:1:6: STATIC "static",
	// }
}

func ExampleStorageClassSpecifier_case3() {
	fmt.Println(exampleAST(111, "\U00100002 auto ("))
	// Output:
	// &cc.StorageClassSpecifier{
	// · Case: 3,
	// · Token: example111.c:1:6: AUTO "auto",
	// }
}

func ExampleStorageClassSpecifier_case4() {
	fmt.Println(exampleAST(112, "\U00100002 register ("))
	// Output:
	// &cc.StorageClassSpecifier{
	// · Case: 4,
	// · Token: example112.c:1:6: REGISTER "register",
	// }
}

func ExampleStructDeclaration() {
	fmt.Println(exampleAST(134, "\U00100002 struct { _Bool ; }"))
	// Output:
	// &cc.StructDeclaration{
	// · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · Case: 9,
	// · · · Token: example134.c:1:15: BOOL "_Bool",
	// · · },
	// · },
	// · Token: example134.c:1:21: ';',
	// }
}

func ExampleStructDeclarationList() {
	fmt.Println(exampleAST(132, "\U00100002 struct { _Bool ; }"))
	// Output:
	// &cc.StructDeclarationList{
	// · StructDeclaration: &cc.StructDeclaration{
	// · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · Case: 9,
	// · · · · Token: example132.c:1:15: BOOL "_Bool",
	// · · · },
	// · · },
	// · · Token: example132.c:1:21: ';',
	// · },
	// }
}

func ExampleStructDeclarationList_case1() {
	fmt.Println(exampleAST(133, "\U00100002 struct { _Bool ; _Bool ; }"))
	// Output:
	// &cc.StructDeclarationList{
	// · StructDeclaration: &cc.StructDeclaration{
	// · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · Case: 9,
	// · · · · Token: example133.c:1:15: BOOL "_Bool",
	// · · · },
	// · · },
	// · · Token: example133.c:1:21: ';',
	// · },
	// · StructDeclarationList: &cc.StructDeclarationList{
	// · · Case: 1,
	// · · StructDeclaration: &cc.StructDeclaration{
	// · · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · · Case: 9,
	// · · · · · Token: example133.c:1:23: BOOL "_Bool",
	// · · · · },
	// · · · },
	// · · · Token: example133.c:1:29: ';',
	// · · },
	// · },
	// }
}

func ExampleStructDeclarator() {
	fmt.Println(exampleAST(144, "\U00100002 struct { _Bool a ,"))
	// Output:
	// &cc.StructDeclarator{
	// · align: cc.align{
	// · },
	// · offset: cc.offset{
	// · },
	// · size: cc.size{
	// · },
	// · Declarator: &cc.Declarator{
	// · · IsDefinition: true,
	// · · SUSpecifier0: &cc.StructOrUnionSpecifier0{
	// · · · StructOrUnion: &cc.StructOrUnion{
	// · · · · Token: example144.c:1:6: STRUCT "struct",
	// · · · },
	// · · },
	// · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · Token: example144.c:1:21: IDENTIFIER "a",
	// · · },
	// · },
	// }
}

func ExampleStructDeclarator_case1() {
	fmt.Println(exampleAST(145, "\U00100002 struct { _Bool : 'a' ,"))
	// Output:
	// &cc.StructDeclarator{
	// · align: cc.align{
	// · },
	// · offset: cc.offset{
	// · },
	// · size: cc.size{
	// · },
	// · Case: 1,
	// · ConstantExpression: &cc.ConstantExpression{
	// · · ConditionalExpression: &cc.ConditionalExpression{
	// · · · LogicalOrExpression: &cc.LogicalOrExpression{
	// · · · · LogicalAndExpression: &cc.LogicalAndExpression{
	// · · · · · InclusiveOrExpression: &cc.InclusiveOrExpression{
	// · · · · · · ExclusiveOrExpression: &cc.ExclusiveOrExpression{
	// · · · · · · · AndExpression: &cc.AndExpression{
	// · · · · · · · · EqualityExpression: &cc.EqualityExpression{
	// · · · · · · · · · RelationalExpression: &cc.RelationalExpression{
	// · · · · · · · · · · ShiftExpression: &cc.ShiftExpression{
	// · · · · · · · · · · · AdditiveExpression: &cc.AdditiveExpression{
	// · · · · · · · · · · · · MultiplicativeExpression: &cc.MultiplicativeExpression{
	// · · · · · · · · · · · · · CastExpression: &cc.CastExpression{
	// · · · · · · · · · · · · · · UnaryExpression: &cc.UnaryExpression{
	// · · · · · · · · · · · · · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · · · · · · · · · · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · · · · · · · · · · · · · Case: 1,
	// · · · · · · · · · · · · · · · · · Constant: &cc.Constant{
	// · · · · · · · · · · · · · · · · · · Token: example145.c:1:23: CHARCONST "'a'",
	// · · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · · },
	// · · · · · · · · · · · · · },
	// · · · · · · · · · · · · },
	// · · · · · · · · · · · },
	// · · · · · · · · · · },
	// · · · · · · · · · },
	// · · · · · · · · },
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Token: example145.c:1:21: ':',
	// }
}

func ExampleStructDeclaratorList() {
	fmt.Println(exampleAST(140, "\U00100002 struct { _Bool a ,"))
	// Output:
	// &cc.StructDeclaratorList{
	// · StructDeclarator: &cc.StructDeclarator{
	// · · align: cc.align{
	// · · },
	// · · offset: cc.offset{
	// · · },
	// · · size: cc.size{
	// · · },
	// · · Declarator: &cc.Declarator{
	// · · · IsDefinition: true,
	// · · · SUSpecifier0: &cc.StructOrUnionSpecifier0{
	// · · · · StructOrUnion: &cc.StructOrUnion{
	// · · · · · Token: example140.c:1:6: STRUCT "struct",
	// · · · · },
	// · · · },
	// · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · Token: example140.c:1:21: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleStructDeclaratorList_case1() {
	fmt.Println(exampleAST(141, "\U00100002 struct { _Bool a , b ,"))
	// Output:
	// &cc.StructDeclaratorList{
	// · StructDeclarator: &cc.StructDeclarator{
	// · · align: cc.align{
	// · · },
	// · · offset: cc.offset{
	// · · },
	// · · size: cc.size{
	// · · },
	// · · Declarator: &cc.Declarator{
	// · · · IsDefinition: true,
	// · · · SUSpecifier0: &cc.StructOrUnionSpecifier0{
	// · · · · StructOrUnion: &cc.StructOrUnion{
	// · · · · · Token: example141.c:1:6: STRUCT "struct",
	// · · · · },
	// · · · },
	// · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · Token: example141.c:1:21: IDENTIFIER "a",
	// · · · },
	// · · },
	// · },
	// · StructDeclaratorList: &cc.StructDeclaratorList{
	// · · Case: 1,
	// · · StructDeclarator: &cc.StructDeclarator{
	// · · · align: cc.align{
	// · · · },
	// · · · offset: cc.offset{
	// · · · },
	// · · · size: cc.size{
	// · · · },
	// · · · Declarator: &cc.Declarator{
	// · · · · IsDefinition: true,
	// · · · · SUSpecifier0: &cc.StructOrUnionSpecifier0{ /* recursive/repetitive pointee not shown */ },
	// · · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · · Token: example141.c:1:25: IDENTIFIER "b",
	// · · · · },
	// · · · },
	// · · },
	// · · Token: example141.c:1:23: ',',
	// · },
	// }
}

func ExampleStructDeclaratorListOpt() {
	fmt.Println(exampleAST(142, "\U00100002 struct { _Bool ;") == (*StructDeclaratorListOpt)(nil))
	// Output:
	// true
}

func ExampleStructDeclaratorListOpt_case1() {
	fmt.Println(exampleAST(143, "\U00100002 struct { _Bool a ;"))
	// Output:
	// &cc.StructDeclaratorListOpt{
	// · StructDeclaratorList: &cc.StructDeclaratorList{
	// · · StructDeclarator: &cc.StructDeclarator{
	// · · · align: cc.align{
	// · · · },
	// · · · offset: cc.offset{
	// · · · },
	// · · · size: cc.size{
	// · · · },
	// · · · Declarator: &cc.Declarator{
	// · · · · IsDefinition: true,
	// · · · · SUSpecifier0: &cc.StructOrUnionSpecifier0{
	// · · · · · StructOrUnion: &cc.StructOrUnion{
	// · · · · · · Token: example143.c:1:6: STRUCT "struct",
	// · · · · · },
	// · · · · },
	// · · · · DirectDeclarator: &cc.DirectDeclarator{
	// · · · · · Token: example143.c:1:21: IDENTIFIER "a",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleStructOrUnion() {
	fmt.Println(exampleAST(130, "\U00100002 struct {"))
	// Output:
	// &cc.StructOrUnion{
	// · Token: example130.c:1:6: STRUCT "struct",
	// }
}

func ExampleStructOrUnion_case1() {
	fmt.Println(exampleAST(131, "\U00100002 union {"))
	// Output:
	// &cc.StructOrUnion{
	// · Case: 1,
	// · Token: example131.c:1:6: UNION "union",
	// }
}

func ExampleStructOrUnionSpecifier() {
	fmt.Println(exampleAST(128, "\U00100002 struct { _Bool ; } ("))
	// Output:
	// &cc.StructOrUnionSpecifier{
	// · Members: &cc.Bindings{
	// · · SUSpecifier0: &cc.StructOrUnionSpecifier0{
	// · · · SUSpecifier: &cc.StructOrUnionSpecifier{ /* recursive/repetitive pointee not shown */ },
	// · · · StructOrUnion: &cc.StructOrUnion{
	// · · · · Token: example128.c:1:6: STRUCT "struct",
	// · · · },
	// · · },
	// · · Type: 4,
	// · },
	// · align: cc.align{
	// · },
	// · size: cc.size{
	// · },
	// · StructDeclarationList: &cc.StructDeclarationList{
	// · · StructDeclaration: &cc.StructDeclaration{
	// · · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · · Case: 9,
	// · · · · · Token: example128.c:1:15: BOOL "_Bool",
	// · · · · },
	// · · · },
	// · · · Token: example128.c:1:21: ';',
	// · · },
	// · },
	// · StructOrUnionSpecifier0: &cc.StructOrUnionSpecifier0{ /* recursive/repetitive pointee not shown */ },
	// · Token: example128.c:1:13: '{',
	// · Token2: example128.c:1:23: '}',
	// }
}

func ExampleStructOrUnionSpecifier_case1() {
	fmt.Println(exampleAST(129, "\U00100002 struct a ("))
	// Output:
	// &cc.StructOrUnionSpecifier{
	// · Case: 1,
	// · StructOrUnion: &cc.StructOrUnion{
	// · · Token: example129.c:1:6: STRUCT "struct",
	// · },
	// · Token: example129.c:1:13: IDENTIFIER "a",
	// }
}

func ExampleStructOrUnionSpecifier0() {
	fmt.Println(exampleAST(127, "\U00100002 struct {"))
	// Output:
	// &cc.StructOrUnionSpecifier0{
	// · StructOrUnion: &cc.StructOrUnion{
	// · · Token: example127.c:1:6: STRUCT "struct",
	// · },
	// }
}

func ExampleTranslationUnit() {
	fmt.Println(exampleAST(249, "\U00100002 auto ;"))
	// Output:
	// &cc.TranslationUnit{
	// · ExternalDeclaration: &cc.ExternalDeclaration{
	// · · Case: 1,
	// · · Declaration: &cc.Declaration{
	// · · · IsFileScope: true,
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example249.c:1:6: AUTO "auto",
	// · · · · },
	// · · · },
	// · · · Token: example249.c:1:11: ';',
	// · · },
	// · },
	// }
}

func ExampleTranslationUnit_case1() {
	fmt.Println(exampleAST(250, "\U00100002 auto ; auto ;"))
	// Output:
	// &cc.TranslationUnit{
	// · ExternalDeclaration: &cc.ExternalDeclaration{
	// · · Case: 1,
	// · · Declaration: &cc.Declaration{
	// · · · IsFileScope: true,
	// · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · IsAuto: true,
	// · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · Case: 3,
	// · · · · · Token: example250.c:1:6: AUTO "auto",
	// · · · · },
	// · · · },
	// · · · Token: example250.c:1:11: ';',
	// · · },
	// · },
	// · TranslationUnit: &cc.TranslationUnit{
	// · · Case: 1,
	// · · ExternalDeclaration: &cc.ExternalDeclaration{
	// · · · Case: 1,
	// · · · Declaration: &cc.Declaration{
	// · · · · IsFileScope: true,
	// · · · · DeclarationSpecifiers: &cc.DeclarationSpecifiers{
	// · · · · · IsAuto: true,
	// · · · · · StorageClassSpecifier: &cc.StorageClassSpecifier{
	// · · · · · · Case: 3,
	// · · · · · · Token: example250.c:1:13: AUTO "auto",
	// · · · · · },
	// · · · · },
	// · · · · Token: example250.c:1:18: ';',
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleTypeName() {
	fmt.Println(exampleAST(193, "\U00100001 ( _Bool )"))
	// Output:
	// &cc.TypeName{
	// · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · Case: 9,
	// · · · Token: example193.c:1:8: BOOL "_Bool",
	// · · },
	// · },
	// }
}

func ExampleTypeQualifier() {
	fmt.Println(exampleAST(154, "\U00100002 const !"))
	// Output:
	// &cc.TypeQualifier{
	// · Token: example154.c:1:6: CONST "const",
	// }
}

func ExampleTypeQualifier_case1() {
	fmt.Println(exampleAST(155, "\U00100002 restrict !"))
	// Output:
	// &cc.TypeQualifier{
	// · Case: 1,
	// · Token: example155.c:1:6: RESTRICT "restrict",
	// }
}

func ExampleTypeQualifier_case2() {
	fmt.Println(exampleAST(156, "\U00100002 volatile !"))
	// Output:
	// &cc.TypeQualifier{
	// · Case: 2,
	// · Token: example156.c:1:6: VOLATILE "volatile",
	// }
}

func ExampleTypeQualifierList() {
	fmt.Println(exampleAST(175, "\U00100002 auto * const !"))
	// Output:
	// &cc.TypeQualifierList{
	// · TypeQualifier: &cc.TypeQualifier{
	// · · Token: example175.c:1:13: CONST "const",
	// · },
	// }
}

func ExampleTypeQualifierList_case1() {
	fmt.Println(exampleAST(176, "\U00100002 auto * const const !"))
	// Output:
	// &cc.TypeQualifierList{
	// · TypeQualifier: &cc.TypeQualifier{
	// · · Token: example176.c:1:13: CONST "const",
	// · },
	// · TypeQualifierList: &cc.TypeQualifierList{
	// · · Case: 1,
	// · · TypeQualifier: &cc.TypeQualifier{
	// · · · Token: example176.c:1:19: CONST "const",
	// · · },
	// · },
	// }
}

func ExampleTypeQualifierListOpt() {
	fmt.Println(exampleAST(177, "\U00100002 auto * (") == (*TypeQualifierListOpt)(nil))
	// Output:
	// true
}

func ExampleTypeQualifierListOpt_case1() {
	fmt.Println(exampleAST(178, "\U00100002 auto * const !"))
	// Output:
	// &cc.TypeQualifierListOpt{
	// · TypeQualifierList: &cc.TypeQualifierList{
	// · · TypeQualifier: &cc.TypeQualifier{
	// · · · Token: example178.c:1:13: CONST "const",
	// · · },
	// · },
	// }
}

func ExampleTypeSpecifier() {
	fmt.Println(exampleAST(113, "\U00100002 void ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Token: example113.c:1:6: VOID "void",
	// }
}

func ExampleTypeSpecifier_case01() {
	fmt.Println(exampleAST(114, "\U00100002 char ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 1,
	// · Token: example114.c:1:6: CHAR "char",
	// }
}

func ExampleTypeSpecifier_case02() {
	fmt.Println(exampleAST(115, "\U00100002 short ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 2,
	// · Token: example115.c:1:6: SHORT "short",
	// }
}

func ExampleTypeSpecifier_case03() {
	fmt.Println(exampleAST(116, "\U00100002 int ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 3,
	// · Token: example116.c:1:6: INT "int",
	// }
}

func ExampleTypeSpecifier_case04() {
	fmt.Println(exampleAST(117, "\U00100002 long ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 4,
	// · Token: example117.c:1:6: LONG "long",
	// }
}

func ExampleTypeSpecifier_case05() {
	fmt.Println(exampleAST(118, "\U00100002 float ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 5,
	// · Token: example118.c:1:6: FLOAT "float",
	// }
}

func ExampleTypeSpecifier_case06() {
	fmt.Println(exampleAST(119, "\U00100002 double ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 6,
	// · Token: example119.c:1:6: DOUBLE "double",
	// }
}

func ExampleTypeSpecifier_case07() {
	fmt.Println(exampleAST(120, "\U00100002 signed ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 7,
	// · Token: example120.c:1:6: SIGNED "signed",
	// }
}

func ExampleTypeSpecifier_case08() {
	fmt.Println(exampleAST(121, "\U00100002 unsigned ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 8,
	// · Token: example121.c:1:6: UNSIGNED "unsigned",
	// }
}

func ExampleTypeSpecifier_case09() {
	fmt.Println(exampleAST(122, "\U00100002 _Bool ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 9,
	// · Token: example122.c:1:6: BOOL "_Bool",
	// }
}

func ExampleTypeSpecifier_case10() {
	fmt.Println(exampleAST(123, "\U00100002 _Complex ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 10,
	// · Token: example123.c:1:6: COMPLEX "_Complex",
	// }
}

func ExampleTypeSpecifier_case11() {
	fmt.Println(exampleAST(124, "\U00100002 struct a ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 11,
	// · StructOrUnionSpecifier: &cc.StructOrUnionSpecifier{
	// · · Case: 1,
	// · · StructOrUnion: &cc.StructOrUnion{
	// · · · Token: example124.c:1:6: STRUCT "struct",
	// · · },
	// · · Token: example124.c:1:13: IDENTIFIER "a",
	// · },
	// }
}

func ExampleTypeSpecifier_case12() {
	fmt.Println(exampleAST(125, "\U00100002 enum a ("))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 12,
	// · EnumSpecifier: &cc.EnumSpecifier{
	// · · Case: 2,
	// · · Token: example125.c:1:6: ENUM "enum",
	// · · Token2: example125.c:1:11: IDENTIFIER "a",
	// · },
	// }
}

func ExampleTypeSpecifier_case13() {
	fmt.Println(exampleAST(126, "\U00100002 typedef int i; i j;"))
	// Output:
	// &cc.TypeSpecifier{
	// · Case: 13,
	// · Token: example126.c:1:21: TYPEDEFNAME "i",
	// }
}

func ExampleUnaryExpression() {
	fmt.Println(exampleAST(28, "\U00100001 'a'"))
	// Output:
	// &cc.UnaryExpression{
	// · PostfixExpression: &cc.PostfixExpression{
	// · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · Case: 1,
	// · · · Constant: &cc.Constant{
	// · · · · Token: example28.c:1:6: CHARCONST "'a'",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleUnaryExpression_case1() {
	fmt.Println(exampleAST(29, "\U00100001 ++ 'a'"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 1,
	// · Token: example29.c:1:6: INC,
	// · UnaryExpression: &cc.UnaryExpression{
	// · · PostfixExpression: &cc.PostfixExpression{
	// · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · Case: 1,
	// · · · · Constant: &cc.Constant{
	// · · · · · Token: example29.c:1:9: CHARCONST "'a'",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleUnaryExpression_case2() {
	fmt.Println(exampleAST(30, "\U00100001 -- 'a'"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 2,
	// · Token: example30.c:1:6: DEC,
	// · UnaryExpression: &cc.UnaryExpression{
	// · · PostfixExpression: &cc.PostfixExpression{
	// · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · Case: 1,
	// · · · · Constant: &cc.Constant{
	// · · · · · Token: example30.c:1:9: CHARCONST "'a'",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleUnaryExpression_case3() {
	fmt.Println(exampleAST(31, "\U00100001 ! 'a'"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 3,
	// · CastExpression: &cc.CastExpression{
	// · · UnaryExpression: &cc.UnaryExpression{
	// · · · PostfixExpression: &cc.PostfixExpression{
	// · · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · · Case: 1,
	// · · · · · Constant: &cc.Constant{
	// · · · · · · Token: example31.c:1:8: CHARCONST "'a'",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · UnaryOperator: &cc.UnaryOperator{
	// · · Case: 5,
	// · · Token: example31.c:1:6: '!',
	// · },
	// }
}

func ExampleUnaryExpression_case4() {
	fmt.Println(exampleAST(32, "\U00100001 sizeof 'a'"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 4,
	// · Token: example32.c:1:6: SIZEOF "sizeof",
	// · UnaryExpression: &cc.UnaryExpression{
	// · · PostfixExpression: &cc.PostfixExpression{
	// · · · PrimaryExpression: &cc.PrimaryExpression{
	// · · · · Case: 1,
	// · · · · Constant: &cc.Constant{
	// · · · · · Token: example32.c:1:13: CHARCONST "'a'",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleUnaryExpression_case5() {
	fmt.Println(exampleAST(33, "\U00100001 sizeof ( _Bool )"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 5,
	// · Token: example33.c:1:6: SIZEOF "sizeof",
	// · Token2: example33.c:1:13: '(',
	// · Token3: example33.c:1:21: ')',
	// · TypeName: &cc.TypeName{
	// · · SpecifierQualifierList: &cc.SpecifierQualifierList{
	// · · · TypeSpecifier: &cc.TypeSpecifier{
	// · · · · Case: 9,
	// · · · · Token: example33.c:1:15: BOOL "_Bool",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleUnaryExpression_case6() {
	fmt.Println(exampleAST(34, "\U00100001 defined a"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 6,
	// · Token: example34.c:1:6: DEFINED "defined",
	// · Token2: example34.c:1:14: IDENTIFIER "a",
	// }
}

func ExampleUnaryExpression_case7() {
	fmt.Println(exampleAST(35, "\U00100001 defined ( a )"))
	// Output:
	// &cc.UnaryExpression{
	// · Case: 7,
	// · Token: example35.c:1:6: DEFINED "defined",
	// · Token2: example35.c:1:14: '(',
	// · Token3: example35.c:1:16: IDENTIFIER "a",
	// · Token4: example35.c:1:18: ')',
	// }
}

func ExampleUnaryOperator() {
	fmt.Println(exampleAST(36, "\U00100001 & !"))
	// Output:
	// &cc.UnaryOperator{
	// · Token: example36.c:1:6: '&',
	// }
}

func ExampleUnaryOperator_case1() {
	fmt.Println(exampleAST(37, "\U00100001 * !"))
	// Output:
	// &cc.UnaryOperator{
	// · Case: 1,
	// · Token: example37.c:1:6: '*',
	// }
}

func ExampleUnaryOperator_case2() {
	fmt.Println(exampleAST(38, "\U00100001 + !"))
	// Output:
	// &cc.UnaryOperator{
	// · Case: 2,
	// · Token: example38.c:1:6: '+',
	// }
}

func ExampleUnaryOperator_case3() {
	fmt.Println(exampleAST(39, "\U00100001 - !"))
	// Output:
	// &cc.UnaryOperator{
	// · Case: 3,
	// · Token: example39.c:1:6: '-',
	// }
}

func ExampleUnaryOperator_case4() {
	fmt.Println(exampleAST(40, "\U00100001 ~ !"))
	// Output:
	// &cc.UnaryOperator{
	// · Case: 4,
	// · Token: example40.c:1:6: '~',
	// }
}

func ExampleUnaryOperator_case5() {
	fmt.Println(exampleAST(41, "\U00100001 ! !"))
	// Output:
	// &cc.UnaryOperator{
	// · Case: 5,
	// · Token: example41.c:1:6: '!',
	// }
}
