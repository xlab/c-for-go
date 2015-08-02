package translator

import (
	"bytes"

	"github.com/cznic/c/internal/cc"
)

type (
	assignmentExpression     cc.AssignmentExpression
	conditionalExpression    cc.ConditionalExpression
	logicalOrExpression      cc.LogicalOrExpression
	logicalAndExpression     cc.LogicalAndExpression
	inclusiveOrExpression    cc.InclusiveOrExpression
	exclusiveOrExpression    cc.ExclusiveOrExpression
	andExpression            cc.AndExpression
	equalityExpression       cc.EqualityExpression
	relationalExpression     cc.RelationalExpression
	shiftExpression          cc.ShiftExpression
	additiveExpression       cc.AdditiveExpression
	multiplicativeExpression cc.MultiplicativeExpression
	castExpression           cc.CastExpression
	unaryExpression          cc.UnaryExpression
	postfixExpression        cc.PostfixExpression
	primaryExpression        cc.PrimaryExpression
	//
	expressionList         cc.ExpressionList
	argumentExpressionList cc.ArgumentExpressionList
)

type (
	assignmentOperator cc.AssignmentOperator
	unaryOperator      cc.UnaryOperator
	constantValue      cc.Constant
)

func (ex *assignmentExpression) Eval() []byte {
	switch ex.Case {
	case 0: // ConditionalExpression
		return (*conditionalExpression)(ex.ConditionalExpression).Eval()
	case 1: // UnaryExpression AssignmentOperator AssignmentExpression
		buf := (*unaryExpression)(ex.UnaryExpression).Eval()
		buf = append(buf, (*assignmentOperator)(ex.AssignmentOperator).Eval()...)
		buf = append(buf, (*conditionalExpression)(ex.ConditionalExpression).Eval()...)
		return buf
	default:
		return nil
	}
}

func (ex *conditionalExpression) Eval() []byte {
	switch ex.Case {
	case 0: // LogicalOrExpression
		return (*logicalOrExpression)(ex.LogicalOrExpression).Eval()
	case 1: // LogicalOrExpression '?' ExpressionList ':' ConditionalExpression
		// not supported in go, falling back to false
		return (*conditionalExpression)(ex.ConditionalExpression).Eval()
	default:
		return nil
	}
}

func (ex *logicalOrExpression) Eval() []byte {
	switch ex.Case {
	case 0: // LogicalAndExpression
		return (*logicalAndExpression)(ex.LogicalAndExpression).Eval()
	case 1: // LogicalOrExpression "||" LogicalAndExpression
		return evalAndJoin(
			(*logicalOrExpression)(ex.LogicalOrExpression),
			(*logicalAndExpression)(ex.LogicalAndExpression), "||")
		return nil
	default:
		return nil
	}
}

func (ex *logicalAndExpression) Eval() []byte {
	switch ex.Case {
	case 0: // InclusiveOrExpression
		return (*inclusiveOrExpression)(ex.InclusiveOrExpression).Eval()
	case 1: // LogicalAndExpression "&&" InclusiveOrExpression
		return evalAndJoin(
			(*logicalAndExpression)(ex.LogicalAndExpression),
			(*inclusiveOrExpression)(ex.InclusiveOrExpression), "&&")
	default:
		return nil
	}
}

func (ex *inclusiveOrExpression) Eval() []byte {
	switch ex.Case {
	case 0: // ExclusiveOrExpression
		return (*exclusiveOrExpression)(ex.ExclusiveOrExpression).Eval()
	case 1: // InclusiveOrExpression '|' ExclusiveOrExpression
		return evalAndJoin(
			(*inclusiveOrExpression)(ex.InclusiveOrExpression),
			(*exclusiveOrExpression)(ex.ExclusiveOrExpression), "|")
	default:
		return nil
	}
}

func (ex *exclusiveOrExpression) Eval() []byte {
	switch ex.Case {
	case 0: // AndExpression
		return (*andExpression)(ex.AndExpression).Eval()
	case 1: // ExclusiveOrExpression '^' AndExpression
		return evalAndJoin(
			(*exclusiveOrExpression)(ex.ExclusiveOrExpression),
			(*andExpression)(ex.AndExpression), "^")
	default:
		return nil
	}
}

func (ex *andExpression) Eval() []byte {
	switch ex.Case {
	case 0: // EqualityExpression
		return (*equalityExpression)(ex.EqualityExpression).Eval()
	case 1: // AndExpression '&' EqualityExpression
		return evalAndJoin(
			(*andExpression)(ex.AndExpression),
			(*equalityExpression)(ex.EqualityExpression), "&")
	default:
		return nil
	}
}

func (ex *equalityExpression) Eval() []byte {
	switch ex.Case {
	case 0: // RelationalExpression
		return (*relationalExpression)(ex.RelationalExpression).Eval()
	case 1: // EqualityExpression "==" RelationalExpression
		return evalAndJoin(
			(*equalityExpression)(ex.EqualityExpression),
			(*relationalExpression)(ex.RelationalExpression), "==")
	case 2: // EqualityExpression "!=" RelationalExpression
		return evalAndJoin(
			(*equalityExpression)(ex.EqualityExpression),
			(*relationalExpression)(ex.RelationalExpression), "!=")
	default:
		return nil
	}
}

func (ex *relationalExpression) Eval() []byte {
	switch ex.Case {
	case 0: // ShiftExpression
		return (*shiftExpression)(ex.ShiftExpression).Eval()
	case 1: // RelationalExpression '<' ShiftExpression
		return evalAndJoin(
			(*relationalExpression)(ex.RelationalExpression),
			(*shiftExpression)(ex.ShiftExpression), "<")
	case 2: // RelationalExpression '>' ShiftExpression
		return evalAndJoin(
			(*relationalExpression)(ex.RelationalExpression),
			(*shiftExpression)(ex.ShiftExpression), ">")
	case 3: // RelationalExpression "<=" ShiftExpression
		return evalAndJoin(
			(*relationalExpression)(ex.RelationalExpression),
			(*shiftExpression)(ex.ShiftExpression), "<=")
	case 4: // RelationalExpression ">=" ShiftExpression
		return evalAndJoin(
			(*relationalExpression)(ex.RelationalExpression),
			(*shiftExpression)(ex.ShiftExpression), ">=")
	default:
		return nil
	}
}

func (ex *shiftExpression) Eval() []byte {
	switch ex.Case {
	case 0: // AdditiveExpression
		return (*additiveExpression)(ex.AdditiveExpression).Eval()
	case 1: // ShiftExpression "<<" AdditiveExpression
		return evalAndJoin(
			(*shiftExpression)(ex.ShiftExpression),
			(*additiveExpression)(ex.AdditiveExpression), "<<")
	case 2: // ShiftExpression ">>" AdditiveExpression
		return evalAndJoin(
			(*shiftExpression)(ex.ShiftExpression),
			(*additiveExpression)(ex.AdditiveExpression), ">>")
	default:
		return nil
	}
}

func (ex *additiveExpression) Eval() []byte {
	switch ex.Case {
	case 0: // MultiplicativeExpression
		return (*multiplicativeExpression)(ex.MultiplicativeExpression).Eval()
	case 1: // AdditiveExpression '+' MultiplicativeExpression
		return evalAndJoin(
			(*additiveExpression)(ex.AdditiveExpression),
			(*multiplicativeExpression)(ex.MultiplicativeExpression), "+")
	case 2: // AdditiveExpression '-' MultiplicativeExpression
		return evalAndJoin(
			(*additiveExpression)(ex.AdditiveExpression),
			(*multiplicativeExpression)(ex.MultiplicativeExpression), "-")
	default:
		return nil
	}
}

func (ex *multiplicativeExpression) Eval() []byte {
	switch ex.Case {
	case 0: // CastExpression
		return (*castExpression)(ex.CastExpression).Eval()
	case 1: // MultiplicativeExpression '*' CastExpression
		return evalAndJoin(
			(*multiplicativeExpression)(ex.MultiplicativeExpression),
			(*castExpression)(ex.CastExpression), "*")
	case 2: // MultiplicativeExpression '/' CastExpression
		return evalAndJoin(
			(*multiplicativeExpression)(ex.MultiplicativeExpression),
			(*castExpression)(ex.CastExpression), "/")
	case 3: // MultiplicativeExpression '%' CastExpression
		return evalAndJoin(
			(*multiplicativeExpression)(ex.MultiplicativeExpression),
			(*castExpression)(ex.CastExpression), "%")
	default:
		return nil
	}
}

func (ex *castExpression) Eval() []byte {
	switch ex.Case {
	case 0: // UnaryExpression
		return (*unaryExpression)(ex.UnaryExpression).Eval()
	case 1: // '(' TypeName ')' CastExpression
		// ignore the type cast
		return (*castExpression)(ex.CastExpression).Eval()
	default:
		return nil
	}
}

func (ex *unaryExpression) Eval() []byte {
	switch ex.Case {
	case 0: // PostfixExpression
		return (*postfixExpression)(ex.PostfixExpression).Eval()
	case 1: // "++" UnaryExpression
		return append([]byte("++"), (*unaryExpression)(ex.UnaryExpression).Eval()...)
	case 2: // "--" UnaryExpression
		return append([]byte("--"), (*unaryExpression)(ex.UnaryExpression).Eval()...)
	case 3: // UnaryOperator CastExpression
		return append((*unaryOperator)(ex.UnaryOperator).Eval(), (*castExpression)(ex.CastExpression).Eval()...)
	case 4, // "sizeof" UnaryExpression
		5: // "sizeof" '(' TypeName ')'
		// sizeof will always be 0, why not.
		return nil
	case 6, // "defined" IDENTIFIER
		7: // "defined" '(' IDENTIFIER ')'
		// what's defined by the way?
		return nil
	default:
		return nil
	}
}

func (ex *postfixExpression) Eval() []byte {
	switch ex.Case {
	case 0: // PrimaryExpression
		return (*primaryExpression)(ex.PrimaryExpression).Eval()
	case 1: // PostfixExpression '[' ExpressionList ']'
		buf := (*postfixExpression)(ex.PostfixExpression).Eval()
		buf = append(buf, '[')
		buf = append(buf, (*expressionList)(ex.ExpressionList).Eval()...)
		buf = append(buf, ']')
		return buf
	case 2: // PostfixExpression '(' ArgumentExpressionListOpt ')'
		buf := (*postfixExpression)(ex.PostfixExpression).Eval()
		if ex.ArgumentExpressionListOpt == nil {
			return buf
		}
		buf = append(buf, '(')
		expressionList := ex.ArgumentExpressionListOpt.ArgumentExpressionList
		buf = append(buf, (*argumentExpressionList)(expressionList).Eval()...)
		buf = append(buf, ')')
		return buf
	case 3, // PostfixExpression '.' IDENTIFIER
		4: // PostfixExpression "->" IDENTIFIER
		buf := (*postfixExpression)(ex.PostfixExpression).Eval()
		buf = append(buf, '.')
		buf = append(buf, (*postfixExpression)(ex.PostfixExpression).Eval()...)
		return buf
	case 5, // PostfixExpression "++"
		6: // PostfixExpression "--"
		// not present in golang
		return (*postfixExpression)(ex.PostfixExpression).Eval()
	case 7: // '(' TypeName ')' '{' InitializerList '}'
		// wat is dis?
		return nil
	case 8: // '(' TypeName ')' '{' InitializerList ',' '}'
		// wat is dis?
		return nil
	default:
		return nil
	}
}

func (ex *primaryExpression) Eval() []byte {
	switch ex.Case {
	case 0: // IDENTIFIER
		return ex.Token.S()
	case 1: // Constant
		return (*constantValue)(ex.Constant).Eval()
	case 2:
		// get only the first one from expression list
		return (*expressionList)(ex.ExpressionList).Eval()
	default:
		return nil
	}
}

func (ex *expressionList) Eval() []byte {
	switch ex.Case {
	case 0, // AssignmentExpression
		1: // ExpressionList ',' AssignmentExpression
		// get only the first one from the list
		return (*assignmentExpression)(ex.AssignmentExpression).Eval()
	default:
		return nil
	}
}

func (ex *argumentExpressionList) Eval() []byte {
	switch ex.Case {
	case 0, // AssignmentExpression
		1: // ExpressionList ',' AssignmentExpression
		// get only the first one from the list
		return (*assignmentExpression)(ex.AssignmentExpression).Eval()
	default:
		return nil
	}
}

func (cv *constantValue) Eval() []byte {
	switch cv.Case {
	case 0: // CHARCONST
		return cv.Token.S()
	case 1: // FLOATCONST
		return cv.Token.S()
	case 2: // INTCONST
		return bytes.TrimRight(cv.Token.S(), "luLU")
	case 3: // LONGCHARCONST
		return cv.Token.S()
	case 4: // LONGSTRINGLITERAL
		return cv.Token.S()
	case 5: // STRINGLITERAL
		return cv.Token.S()
	default:
		return nil
	}
}

func (op *assignmentOperator) Eval() []byte {
	switch op.Case {
	case 0: // "="
		return op.Token.S()
	case 1: // "*="
		return op.Token.S()
	case 2: // "/="
		return op.Token.S()
	case 3: // "%="
		return op.Token.S()
	case 4: // "+="
		return op.Token.S()
	case 5: // "-="
		return op.Token.S()
	case 6: // "<<="
		return op.Token.S()
	case 7: // ">>="
		return op.Token.S()
	case 8: // "&="
		return op.Token.S()
	case 9: // "^="
		return op.Token.S()
	case 10: // "|="
		return op.Token.S()
	default:
		return nil
	}
}

func (op *unaryOperator) Eval() []byte {
	// no caveats -> return it as-is
	return op.Token.S()
}
