package translator

import (
	"bytes"
	"fmt"

	"github.com/cznic/c/internal/cc"
	"github.com/cznic/c/internal/xc"
)

type (
	assignmentOperator cc.AssignmentOperator
	unaryOperator      cc.UnaryOperator
	constantValue      cc.Constant
)

func (t *Translator) ExpandAssignmentExpression(ex *cc.AssignmentExpression) Expression {
	if ex == nil {
		return nil
	}
	switch ex.Case {
	case 0: // ConditionalExpression
		return t.ExpandConditionalExpression(ex.ConditionalExpression)
	case 1: // UnaryExpression AssignmentOperator AssignmentExpression
		buf := t.ExpandUnaryExpression(ex.UnaryExpression)
		buf = append(buf, (*assignmentOperator)(ex.AssignmentOperator).Expand()...)
		buf = append(buf, t.ExpandConditionalExpression(ex.ConditionalExpression)...)
		return buf
	default:
		return nil
	}
}

func (t *Translator) ExpandConditionalExpression(ex *cc.ConditionalExpression) Expression {
	switch ex.Case {
	case 0: // LogicalOrExpression
		return t.ExpandLogicalOrExpression(ex.LogicalOrExpression)
	case 1: // LogicalOrExpression '?' ExpressionList ':' ConditionalExpression
		// not supported in go, falling back to false
		return t.ExpandConditionalExpression(ex.ConditionalExpression)
	default:
		return nil
	}
}

func (t *Translator) ExpandLogicalOrExpression(ex *cc.LogicalOrExpression) Expression {
	switch ex.Case {
	case 0: // LogicalAndExpression
		return t.ExpandLogicalAndExpression(ex.LogicalAndExpression)
	case 1: // LogicalOrExpression "||" LogicalAndExpression
		return bytesJoin(
			t.ExpandLogicalOrExpression(ex.LogicalOrExpression),
			t.ExpandLogicalAndExpression(ex.LogicalAndExpression), "||")
		return nil
	default:
		return nil
	}
}

func (t *Translator) ExpandLogicalAndExpression(ex *cc.LogicalAndExpression) Expression {
	switch ex.Case {
	case 0: // InclusiveOrExpression
		return t.ExpandInclusiveOrExpression(ex.InclusiveOrExpression)
	case 1: // LogicalAndExpression "&&" InclusiveOrExpression
		return bytesJoin(
			t.ExpandLogicalAndExpression(ex.LogicalAndExpression),
			t.ExpandInclusiveOrExpression(ex.InclusiveOrExpression), "&&")
	default:
		return nil
	}
}

func (t *Translator) ExpandInclusiveOrExpression(ex *cc.InclusiveOrExpression) Expression {
	switch ex.Case {
	case 0: // ExclusiveOrExpression
		return t.ExpandExclusiveOrExpression(ex.ExclusiveOrExpression)
	case 1: // InclusiveOrExpression '|' ExclusiveOrExpression
		return bytesJoin(
			t.ExpandInclusiveOrExpression(ex.InclusiveOrExpression),
			t.ExpandExclusiveOrExpression(ex.ExclusiveOrExpression), "|")
	default:
		return nil
	}
}

func (t *Translator) ExpandExclusiveOrExpression(ex *cc.ExclusiveOrExpression) Expression {
	switch ex.Case {
	case 0: // AndExpression
		return t.ExpandAndExpression(ex.AndExpression)
	case 1: // ExclusiveOrExpression '^' AndExpression
		return bytesJoin(
			t.ExpandExclusiveOrExpression(ex.ExclusiveOrExpression),
			t.ExpandAndExpression(ex.AndExpression), "^")
	default:
		return nil
	}
}

func (t *Translator) ExpandAndExpression(ex *cc.AndExpression) Expression {
	switch ex.Case {
	case 0: // EqualityExpression
		return t.ExpandEqualityExpression(ex.EqualityExpression)
	case 1: // AndExpression '&' EqualityExpression
		return bytesJoin(
			t.ExpandAndExpression(ex.AndExpression),
			t.ExpandEqualityExpression(ex.EqualityExpression), "&")
	default:
		return nil
	}
}

func (t *Translator) ExpandEqualityExpression(ex *cc.EqualityExpression) Expression {
	switch ex.Case {
	case 0: // RelationalExpression
		return t.ExpandRelationalExpression(ex.RelationalExpression)
	case 1: // EqualityExpression "==" RelationalExpression
		return bytesJoin(
			t.ExpandEqualityExpression(ex.EqualityExpression),
			t.ExpandRelationalExpression(ex.RelationalExpression), "==")
	case 2: // EqualityExpression "!=" RelationalExpression
		return bytesJoin(
			t.ExpandEqualityExpression(ex.EqualityExpression),
			t.ExpandRelationalExpression(ex.RelationalExpression), "!=")
	default:
		return nil
	}
}

func (t *Translator) ExpandRelationalExpression(ex *cc.RelationalExpression) Expression {
	switch ex.Case {
	case 0: // ShiftExpression
		return t.ExpandShiftExpression(ex.ShiftExpression)
	case 1: // RelationalExpression '<' ShiftExpression
		return bytesJoin(
			t.ExpandRelationalExpression(ex.RelationalExpression),
			t.ExpandShiftExpression(ex.ShiftExpression), "<")
	case 2: // RelationalExpression '>' ShiftExpression
		return bytesJoin(
			t.ExpandRelationalExpression(ex.RelationalExpression),
			t.ExpandShiftExpression(ex.ShiftExpression), ">")
	case 3: // RelationalExpression "<=" ShiftExpression
		return bytesJoin(
			t.ExpandRelationalExpression(ex.RelationalExpression),
			t.ExpandShiftExpression(ex.ShiftExpression), "<=")
	case 4: // RelationalExpression ">=" ShiftExpression
		return bytesJoin(
			t.ExpandRelationalExpression(ex.RelationalExpression),
			t.ExpandShiftExpression(ex.ShiftExpression), ">=")
	default:
		return nil
	}
}

func (t *Translator) ExpandShiftExpression(ex *cc.ShiftExpression) Expression {
	switch ex.Case {
	case 0: // AdditiveExpression
		return t.ExpandAdditiveExpression(ex.AdditiveExpression)
	case 1: // ShiftExpression "<<" AdditiveExpression
		return bytesJoin(
			t.ExpandShiftExpression(ex.ShiftExpression),
			t.ExpandAdditiveExpression(ex.AdditiveExpression), "<<")
	case 2: // ShiftExpression ">>" AdditiveExpression
		return bytesJoin(
			t.ExpandShiftExpression(ex.ShiftExpression),
			t.ExpandAdditiveExpression(ex.AdditiveExpression), ">>")
	default:
		return nil
	}
}

func (t *Translator) ExpandAdditiveExpression(ex *cc.AdditiveExpression) Expression {
	switch ex.Case {
	case 0: // MultiplicativeExpression
		return t.ExpandMultiplicativeExpression(ex.MultiplicativeExpression)
	case 1: // AdditiveExpression '+' MultiplicativeExpression
		return bytesJoin(
			t.ExpandAdditiveExpression(ex.AdditiveExpression),
			t.ExpandMultiplicativeExpression(ex.MultiplicativeExpression), "+")
	case 2: // AdditiveExpression '-' MultiplicativeExpression
		return bytesJoin(
			t.ExpandAdditiveExpression(ex.AdditiveExpression),
			t.ExpandMultiplicativeExpression(ex.MultiplicativeExpression), "-")
	default:
		return nil
	}
}

func (t *Translator) ExpandMultiplicativeExpression(ex *cc.MultiplicativeExpression) Expression {
	switch ex.Case {
	case 0: // CastExpression
		return t.ExpandCastExpression(ex.CastExpression)
	case 1: // MultiplicativeExpression '*' CastExpression
		return bytesJoin(
			t.ExpandMultiplicativeExpression(ex.MultiplicativeExpression),
			t.ExpandCastExpression(ex.CastExpression), "*")
	case 2: // MultiplicativeExpression '/' CastExpression
		return bytesJoin(
			t.ExpandMultiplicativeExpression(ex.MultiplicativeExpression),
			t.ExpandCastExpression(ex.CastExpression), "/")
	case 3: // MultiplicativeExpression '%' CastExpression
		return bytesJoin(
			t.ExpandMultiplicativeExpression(ex.MultiplicativeExpression),
			t.ExpandCastExpression(ex.CastExpression), "%")
	default:
		return nil
	}
}

func (t *Translator) ExpandCastExpression(ex *cc.CastExpression) Expression {
	switch ex.Case {
	case 0: // UnaryExpression
		return t.ExpandUnaryExpression(ex.UnaryExpression)
	case 1: // '(' TypeName ')' CastExpression
		// ignore the type cast
		return t.ExpandCastExpression(ex.CastExpression)
	default:
		return nil
	}
}

func (t *Translator) ExpandUnaryExpression(ex *cc.UnaryExpression) Expression {
	switch ex.Case {
	case 0: // PostfixExpression
		return t.ExpandPostfixExpression(ex.PostfixExpression)
	case 1: // "++" UnaryExpression
		return bytesJoin(t.ExpandUnaryExpression(ex.UnaryExpression), nil, "+1")
	case 2: // "--" UnaryExpression
		return bytesJoin(t.ExpandUnaryExpression(ex.UnaryExpression), nil, "-1")
	case 3: // UnaryOperator CastExpression
		return append((*unaryOperator)(ex.UnaryOperator).Expand(), t.ExpandCastExpression(ex.CastExpression)...)
	case 4, // "sizeof" UnaryExpression
		5: // "sizeof" '(' TypeName ')'
		// sizeof will always be 0, why not.
		return nil
	case 6, // "defined" IDENTIFIER
		7: // "defined" '(' IDENTIFIER ')'
		if xc.Dict.ID(ex.Token.S()) > 0 {
			return []byte("true")
		}
		return []byte("false")
	default:
		return nil
	}
}

func (t *Translator) ExpandPostfixExpression(ex *cc.PostfixExpression) Expression {
	switch ex.Case {
	case 0: // PrimaryExpression
		return t.ExpandPrimaryExpression(ex.PrimaryExpression)
	case 1: // PostfixExpression '[' ExpressionList ']'
		buf := t.ExpandPostfixExpression(ex.PostfixExpression)
		buf = append(buf, '[')
		buf = append(buf, t.ExpandExpressionList(ex.ExpressionList)...)
		buf = append(buf, ']')
		return buf
	case 2: // PostfixExpression '(' ArgumentExpressionListOpt ')'
		buf := t.ExpandPostfixExpression(ex.PostfixExpression)
		if ex.ArgumentExpressionListOpt == nil {
			return buf
		}
		buf = append(buf, '(')
		expressionList := ex.ArgumentExpressionListOpt.ArgumentExpressionList
		buf = append(buf, t.ExpandArgumentExpressionList(expressionList)...)
		buf = append(buf, ')')
		return buf
	case 3, // PostfixExpression '.' IDENTIFIER
		4: // PostfixExpression "->" IDENTIFIER
		return bytesJoin(
			t.ExpandPostfixExpression(ex.PostfixExpression), ex.Token.S(), ".")
	case 5, // PostfixExpression "++"
		6: // PostfixExpression "--"
		// not present in golang
		return t.ExpandPostfixExpression(ex.PostfixExpression)
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

func (t *Translator) ExpandPrimaryExpression(ex *cc.PrimaryExpression) Expression {
	switch ex.Case {
	case 0: // IDENTIFIER
		name := ex.Token.S()
		switch t.constRules[ConstDeclaration] {
		case ConstExpandFull:
			if expr, ok := t.exprMap[string(name)]; ok {
				return bytesWrap(expr, "(", ")")
			}
		case ConstEval:
			if v, ok := t.valueMap[string(name)]; ok {
				switch v := v.(type) {
				case int32, uint32, int64, uint64:
					return []byte(fmt.Sprintf("%d", v))
				case float32, float64:
					return []byte(fmt.Sprintf("%f", v))
				case string:
					return []byte(v)
				}
			}
		default:
			if _, ok := t.exprMap[string(name)]; ok {
				return name
			}
		}
		// just skip undefined
		return nil
	case 1: // Constant
		return (*constantValue)(ex.Constant).Expand()
	case 2:
		// get only the first one from expression list
		return t.ExpandExpressionList(ex.ExpressionList)
	default:
		return nil
	}
}

func (t *Translator) ExpandExpressionList(ex *cc.ExpressionList) Expression {
	switch ex.Case {
	case 0, // AssignmentExpression
		1: // ExpressionList ',' AssignmentExpression
		// get only the first one from the list
		return t.ExpandAssignmentExpression(ex.AssignmentExpression)
	default:
		return nil
	}
}

func (t *Translator) ExpandArgumentExpressionList(ex *cc.ArgumentExpressionList) Expression {
	switch ex.Case {
	case 0, // AssignmentExpression
		1: // ExpressionList ',' AssignmentExpression
		// get only the first one from the list
		return t.ExpandAssignmentExpression(ex.AssignmentExpression)
	default:
		return nil
	}
}

func (cv *constantValue) Expand() Expression {
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

func (op *assignmentOperator) Expand() Expression {
	switch op.Case {
	case 0: // "="
		return []byte{'='}
	default:
		unmanagedCaseWarn(op.Case, op.Token.Pos())
		return nil
	}
}

func (op *unaryOperator) Expand() Expression {
	switch op.Case {
	case 0: // '&'
		return []byte{'&'}
	case 1: // '*'
		return []byte{'*'}
	case 2: // '+'
		return []byte{'+'}
	case 3: // '-'
		return []byte{'-'}
	case 4: // '~'
		return []byte{'~'}
	case 5: // '!'
		return []byte{'!'}
	default:
		return nil
	}
}
