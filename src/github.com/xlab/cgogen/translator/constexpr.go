package translator

import (
	"strconv"
	"strings"

	"github.com/cznic/c/internal/cc"
)

// TODO: should be a full-featured constant expressin evaluator for cc.AssignmentExpression
func walkAssigmentExperession(expr *cc.AssignmentExpression) (value interface{}) {
	return nil
}

func walkPrimaryExpressionToInt64(expr *cc.PrimaryExpression) int64 {
	if expr == nil {
		return 0
	}
	switch expr.Case {
	case 0: // IDENTIFIER
		unresolvedIdentifierWarn(string(expr.Token.S()), expr.Token.Pos())
		return 0
	case 1: // Constant
		switch expr.Constant.Case {
		case 2: // INTCONST
			intStr := strings.TrimRight(string(expr.Constant.Token.S()), "luLU")
			i, err := strconv.ParseInt(intStr, 10, 64)
			if err != nil {
				return 0
			}
			return i
		default:
			return 0
		}
	case 2: // '(' ExpressionList ')'
		expr = walkAssigmentExperessionToPrimary(expr.ExpressionList.AssignmentExpression)
		return walkPrimaryExpressionToInt64(expr)
	}
	return 0
}

func walkAssigmentExperessionToPrimary(expr *cc.AssignmentExpression) *cc.PrimaryExpression {
	if expr.ConditionalExpression == nil {
		return nil
	}
	if expr.ConditionalExpression.LogicalOrExpression == nil {
		return nil
	}
	if expr.ConditionalExpression.LogicalOrExpression.LogicalAndExpression == nil {
		return nil
	}
	checkpoint1 := expr.ConditionalExpression.LogicalOrExpression.LogicalAndExpression.InclusiveOrExpression
	if checkpoint1 == nil {
		return nil
	}
	if checkpoint1.ExclusiveOrExpression == nil {
		return nil
	}
	if checkpoint1.ExclusiveOrExpression.AndExpression == nil {
		return nil
	}
	if checkpoint1.ExclusiveOrExpression.AndExpression.EqualityExpression == nil {
		return nil
	}
	if checkpoint1.ExclusiveOrExpression.AndExpression.EqualityExpression.RelationalExpression == nil {
		return nil
	}
	checkpoint2 := checkpoint1.ExclusiveOrExpression.AndExpression.EqualityExpression.RelationalExpression.ShiftExpression
	if checkpoint2 == nil {
		return nil
	}
	if checkpoint2.AdditiveExpression == nil {
		return nil
	}
	if checkpoint2.AdditiveExpression.MultiplicativeExpression == nil {
		return nil
	}
	if checkpoint2.AdditiveExpression.MultiplicativeExpression.CastExpression == nil {
		return nil
	}
	if checkpoint2.AdditiveExpression.MultiplicativeExpression.CastExpression.UnaryExpression == nil {
		return nil
	}
	checkpoint3 := checkpoint2.AdditiveExpression.MultiplicativeExpression.CastExpression.UnaryExpression.PostfixExpression
	if checkpoint3 == nil {
		return nil
	}
	return checkpoint3.PrimaryExpression
}
