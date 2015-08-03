package translator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cznic/c/internal/cc"
	"github.com/cznic/c/internal/xc"
)

type Value interface{}

func (t *Translator) TypeOf(ex *cc.AssignmentExpression) (GoTypeSpec, error) {
	switch x := t.EvalAssignmentExpression(ex).(type) {
	case int32:
		return Int32Spec, nil
	case int64:
		return Int64Spec, nil
	case uint32:
		return Uint32Spec, nil
	case uint64:
		return Uint64Spec, nil
	case float32:
		return Float32Spec, nil
	case float64:
		return Float64Spec, nil
	case string:
		return StringSpec, nil
	default:
		return GoTypeSpec{}, errors.New(fmt.Sprintf("cannot resolve type %T", x))
	}
}

func (t *Translator) EvalAssignmentExpression(ex *cc.AssignmentExpression) Value {
	switch ex.Case {
	case 0: // ConditionalExpression
		return t.EvalConditionalExpression(ex.ConditionalExpression)
	default:
		return nil
	}
}

func (t *Translator) EvalConditionalExpression(ex *cc.ConditionalExpression) Value {
	switch ex.Case {
	case 0: // LogicalOrExpression
		return t.EvalLogicalOrExpression(ex.LogicalOrExpression)
	case 1: // LogicalOrExpression '?' ExpressionList ':' ConditionalExpression
		if isZero(t.EvalLogicalOrExpression(ex.LogicalOrExpression)) {
			return t.EvalConditionalExpression(ex.ConditionalExpression)
		}
		return t.EvalExpressionList(ex.ExpressionList)
	default:
		return nil
	}
}

func (t *Translator) EvalLogicalOrExpression(ex *cc.LogicalOrExpression) Value {
	switch ex.Case {
	case 0: // LogicalAndExpression
		return t.EvalLogicalAndExpression(ex.LogicalAndExpression)
	case 1: // LogicalOrExpression "||" LogicalAndExpression
		ex1 := t.EvalLogicalOrExpression(ex.LogicalOrExpression)
		ex2 := t.EvalLogicalAndExpression(ex.LogicalAndExpression)
		if isZero(ex1) && isZero(ex2) {
			return mustBool(false)
		}
		return mustBool(true)
	default:
		return nil
	}
}

func (t *Translator) EvalLogicalAndExpression(ex *cc.LogicalAndExpression) Value {
	switch ex.Case {
	case 0: // InclusiveOrExpression
		return t.EvalInclusiveOrExpression(ex.InclusiveOrExpression)
	case 1: // LogicalAndExpression "&&" InclusiveOrExpression
		ex1 := t.EvalLogicalAndExpression(ex.LogicalAndExpression)
		ex2 := t.EvalInclusiveOrExpression(ex.InclusiveOrExpression)
		if isZero(ex1) || isZero(ex2) {
			return mustBool(false)
		}
		return mustBool(true)
	default:
		return nil
	}
}

func (t *Translator) EvalInclusiveOrExpression(ex *cc.InclusiveOrExpression) Value {
	switch ex.Case {
	case 0: // ExclusiveOrExpression
		return t.EvalExclusiveOrExpression(ex.ExclusiveOrExpression)
	case 1: // InclusiveOrExpression '|' ExclusiveOrExpression
		v := t.EvalInclusiveOrExpression(ex.InclusiveOrExpression)
		v2 := t.EvalExclusiveOrExpression(ex.ExclusiveOrExpression)

		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x | y
			case int64:
				v = int64(x) | y
			case uint32:
				v = uint32(x) | y
			case uint64:
				v = uint64(x) | y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x | int64(y)
			case int64:
				v = x | y
			case uint32:
				v = uint64(x) | uint64(y)
			case uint64:
				v = uint64(x) | y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x | uint32(y)
			case int64:
				v = uint64(x) | uint64(y)
			case uint32:
				v = uint32(x) | y
			case uint64:
				v = uint64(x) | y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x | uint64(y)
			case int64:
				v = x | uint64(y)
			case uint32:
				v = x | uint64(y)
			case uint64:
				v = x | y
			}
		}
		return v
	}
	return nil
}

func (t *Translator) EvalExclusiveOrExpression(ex *cc.ExclusiveOrExpression) Value {
	switch ex.Case {
	case 0: // AndExpression
		return t.EvalAndExpression(ex.AndExpression)
	case 1: // ExclusiveOrExpression '^' AndExpression
		v := t.EvalExclusiveOrExpression(ex.ExclusiveOrExpression)
		v2 := t.EvalAndExpression(ex.AndExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x ^ y
			case int64:
				v = int64(x) ^ y
			case uint32:
				v = uint32(x) ^ y
			case uint64:
				v = uint64(x) ^ y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x ^ int64(y)
			case int64:
				v = x ^ y
			case uint32:
				v = uint64(x) ^ uint64(y)
			case uint64:
				v = uint64(x) ^ y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x ^ uint32(y)
			case int64:
				v = uint64(x) ^ uint64(y)
			case uint32:
				v = x ^ y
			case uint64:
				v = uint64(x) ^ y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x ^ uint64(y)
			case int64:
				v = x ^ uint64(y)
			case uint32:
				v = x ^ uint64(y)
			case uint64:
				v = x ^ y
			}
		}
		return v
	default:
		return nil
	}
}

func (t *Translator) EvalAndExpression(ex *cc.AndExpression) Value {
	switch ex.Case {
	case 0: // EqualityExpression
		return t.EvalEqualityExpression(ex.EqualityExpression)
	case 1: // AndExpression '&' EqualityExpression
		v := t.EvalAndExpression(ex.AndExpression)
		v2 := t.EvalEqualityExpression(ex.EqualityExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x & y
			case int64:
				v = int64(x) & y
			case uint32:
				v = uint32(x) & y
			case uint64:
				v = uint64(x) & y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x & int64(y)
			case int64:
				v = x & y
			case uint32:
				v = uint64(x) & uint64(y)
			case uint64:
				v = uint64(x) & y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x & uint32(y)
			case int64:
				v = uint64(x) & uint64(y)
			case uint32:
				v = x & y
			case uint64:
				v = uint64(x) & y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x & uint64(y)
			case int64:
				v = x & uint64(y)
			case uint32:
				v = x & uint64(y)
			case uint64:
				v = x & y
			}
		}
		return v
	default:
		return nil
	}
}

func (t *Translator) EvalEqualityExpression(ex *cc.EqualityExpression) Value {
	switch ex.Case {
	case 0: // RelationalExpression
		return t.EvalRelationalExpression(ex.RelationalExpression)
	case 1: // EqualityExpression "==" RelationalExpression
		v := t.EvalEqualityExpression(ex.EqualityExpression)
		v2 := t.EvalRelationalExpression(ex.RelationalExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x == y)
			case int64:
				v = mustBool(int64(x) == y)
			case uint32:
				v = mustBool(uint32(x) == y)
			case uint64:
				v = mustBool(uint64(x) == y)
			case float32:
				v = mustBool(float32(x) == y)
			case float64:
				v = mustBool(float64(x) == y)
			default:
				return mustBool(false)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x == int64(y))
			case int64:
				v = mustBool(x == y)
			case uint32:
				v = mustBool(uint64(x) == uint64(y))
			case uint64:
				v = mustBool(uint64(x) == y)
			case float32:
				v = mustBool(float32(x) == y)
			case float64:
				v = mustBool(float64(x) == y)
			default:
				return mustBool(false)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x == uint32(y))
			case int64:
				v = mustBool(uint64(x) == uint64(y))
			case uint32:
				v = mustBool(x == y)
			case uint64:
				v = mustBool(uint64(x) == y)
			case float32:
				v = mustBool(float32(x) == y)
			case float64:
				v = mustBool(float64(x) == y)
			default:
				return mustBool(false)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x == uint64(y))
			case int64:
				v = mustBool(x == uint64(y))
			case uint32:
				v = mustBool(x == uint64(y))
			case uint64:
				v = mustBool(x == y)
			case float32:
				v = mustBool(float32(x) == y)
			case float64:
				v = mustBool(float64(x) == y)
			default:
				return mustBool(false)
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x == float32(y))
			case int64:
				v = mustBool(x == float32(y))
			case uint32:
				v = mustBool(x == float32(y))
			case uint64:
				v = mustBool(x == float32(y))
			case float32:
				v = mustBool(x == y)
			case float64:
				v = mustBool(float64(x) == y)
			default:
				return mustBool(false)
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x == float64(y))
			case int64:
				v = mustBool(x == float64(y))
			case uint32:
				v = mustBool(x == float64(y))
			case uint64:
				v = mustBool(x == float64(y))
			case float32:
				v = mustBool(x == float64(y))
			case float64:
				v = mustBool(x == y)
			default:
				return mustBool(false)
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = mustBool(x == y)
			default:
				return mustBool(false)
			}
		default:
			return mustBool(false)
		}
	case 2: // EqualityExpression "!=" RelationalExpression
		v := t.EvalEqualityExpression(ex.EqualityExpression)
		v2 := t.EvalRelationalExpression(ex.RelationalExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x != y)
			case int64:
				v = mustBool(int64(x) != y)
			case uint32:
				v = mustBool(uint32(x) != y)
			case uint64:
				v = mustBool(uint64(x) != y)
			case float32:
				v = mustBool(float32(x) != y)
			case float64:
				v = mustBool(float64(x) != y)
			default:
				return mustBool(false)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x != int64(y))
			case int64:
				v = mustBool(x != y)
			case uint32:
				v = mustBool(uint64(x) != uint64(y))
			case uint64:
				v = mustBool(uint64(x) != y)
			case float32:
				v = mustBool(float32(x) != y)
			case float64:
				v = mustBool(float64(x) != y)
			default:
				return mustBool(false)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x != uint32(y))
			case int64:
				v = mustBool(uint64(x) != uint64(y))
			case uint32:
				v = mustBool(x != y)
			case uint64:
				v = mustBool(uint64(x) != y)
			case float32:
				v = mustBool(float32(x) != y)
			case float64:
				v = mustBool(float64(x) != y)
			default:
				return mustBool(false)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x != uint64(y))
			case int64:
				v = mustBool(x != uint64(y))
			case uint32:
				v = mustBool(x != uint64(y))
			case uint64:
				v = mustBool(x != y)
			case float32:
				v = mustBool(float32(x) != y)
			case float64:
				v = mustBool(float64(x) != y)
			default:
				return mustBool(false)
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x != float32(y))
			case int64:
				v = mustBool(x != float32(y))
			case uint32:
				v = mustBool(x != float32(y))
			case uint64:
				v = mustBool(x != float32(y))
			case float32:
				v = mustBool(x != y)
			case float64:
				v = mustBool(float64(x) != y)
			default:
				return mustBool(false)
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x != float64(y))
			case int64:
				v = mustBool(x != float64(y))
			case uint32:
				v = mustBool(x != float64(y))
			case uint64:
				v = mustBool(x != float64(y))
			case float32:
				v = mustBool(x != float64(y))
			case float64:
				v = mustBool(x != y)
			default:
				return mustBool(false)
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = mustBool(x != y)
			default:
				return mustBool(false)
			}
		default:
			return mustBool(false)
		}
	}
	return nil
}

func (t *Translator) EvalRelationalExpression(ex *cc.RelationalExpression) Value {
	switch ex.Case {
	case 0: // ShiftExpression
		return t.EvalShiftExpression(ex.ShiftExpression)
	case 1: // RelationalExpression '<' ShiftExpression
		v := t.EvalRelationalExpression(ex.RelationalExpression)
		v2 := t.EvalShiftExpression(ex.ShiftExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x < y)
			case int64:
				v = mustBool(int64(x) < y)
			case uint32:
				v = mustBool(uint32(x) < y)
			case uint64:
				v = mustBool(uint64(x) < y)
			case float32:
				v = mustBool(float32(x) < y)
			case float64:
				v = mustBool(float64(x) < y)
			default:
				return mustBool(false)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x < int64(y))
			case int64:
				v = mustBool(x < y)
			case uint32:
				v = mustBool(uint64(x) < uint64(y))
			case uint64:
				v = mustBool(uint64(x) < y)
			case float32:
				v = mustBool(float32(x) < y)
			case float64:
				v = mustBool(float64(x) < y)
			default:
				return mustBool(false)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x < uint32(y))
			case int64:
				v = mustBool(uint64(x) < uint64(y))
			case uint32:
				v = mustBool(x < y)
			case uint64:
				v = mustBool(uint64(x) < y)
			case float32:
				v = mustBool(float32(x) < y)
			case float64:
				v = mustBool(float64(x) < y)
			default:
				return mustBool(false)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x < uint64(y))
			case int64:
				v = mustBool(x < uint64(y))
			case uint32:
				v = mustBool(x < uint64(y))
			case uint64:
				v = mustBool(x < y)
			case float32:
				v = mustBool(float32(x) < y)
			case float64:
				v = mustBool(float64(x) < y)
			default:
				return mustBool(false)
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x < float32(y))
			case int64:
				v = mustBool(x < float32(y))
			case uint32:
				v = mustBool(x < float32(y))
			case uint64:
				v = mustBool(x < float32(y))
			case float32:
				v = mustBool(x < y)
			case float64:
				v = mustBool(float64(x) < y)
			default:
				return mustBool(false)
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x < float64(y))
			case int64:
				v = mustBool(x < float64(y))
			case uint32:
				v = mustBool(x < float64(y))
			case uint64:
				v = mustBool(x < float64(y))
			case float32:
				v = mustBool(x < float64(y))
			case float64:
				v = mustBool(x < y)
			default:
				return mustBool(false)
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = mustBool(x < y)
			default:
				return mustBool(false)
			}
		default:
			return mustBool(false)
		}
	case 2: // RelationalExpression '>' ShiftExpression
		v := t.EvalRelationalExpression(ex.RelationalExpression)
		v2 := t.EvalShiftExpression(ex.ShiftExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x > y)
			case int64:
				v = mustBool(int64(x) > y)
			case uint32:
				v = mustBool(uint32(x) > y)
			case uint64:
				v = mustBool(uint64(x) > y)
			case float32:
				v = mustBool(float32(x) > y)
			case float64:
				v = mustBool(float64(x) > y)
			default:
				return mustBool(false)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x > int64(y))
			case int64:
				v = mustBool(x > y)
			case uint32:
				v = mustBool(uint64(x) > uint64(y))
			case uint64:
				v = mustBool(uint64(x) > y)
			case float32:
				v = mustBool(float32(x) > y)
			case float64:
				v = mustBool(float64(x) > y)
			default:
				return mustBool(false)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x > uint32(y))
			case int64:
				v = mustBool(uint64(x) > uint64(y))
			case uint32:
				v = mustBool(x > y)
			case uint64:
				v = mustBool(uint64(x) > y)
			case float32:
				v = mustBool(float32(x) > y)
			case float64:
				v = mustBool(float64(x) > y)
			default:
				return mustBool(false)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x > uint64(y))
			case int64:
				v = mustBool(x > uint64(y))
			case uint32:
				v = mustBool(x > uint64(y))
			case uint64:
				v = mustBool(x > y)
			case float32:
				v = mustBool(float32(x) > y)
			case float64:
				v = mustBool(float64(x) > y)
			default:
				return mustBool(false)
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x > float32(y))
			case int64:
				v = mustBool(x > float32(y))
			case uint32:
				v = mustBool(x > float32(y))
			case uint64:
				v = mustBool(x > float32(y))
			case float32:
				v = mustBool(x > y)
			case float64:
				v = mustBool(float64(x) > y)
			default:
				return mustBool(false)
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x > float64(y))
			case int64:
				v = mustBool(x > float64(y))
			case uint32:
				v = mustBool(x > float64(y))
			case uint64:
				v = mustBool(x > float64(y))
			case float32:
				v = mustBool(x > float64(y))
			case float64:
				v = mustBool(x > y)
			default:
				return mustBool(false)
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = mustBool(x > y)
			default:
				return mustBool(false)
			}
		default:
			return mustBool(false)
		}
	case 3: // RelationalExpression "<=" ShiftExpression
		v := t.EvalRelationalExpression(ex.RelationalExpression)
		v2 := t.EvalShiftExpression(ex.ShiftExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x <= y)
			case int64:
				v = mustBool(int64(x) <= y)
			case uint32:
				v = mustBool(uint32(x) <= y)
			case uint64:
				v = mustBool(uint64(x) <= y)
			case float32:
				v = mustBool(float32(x) <= y)
			case float64:
				v = mustBool(float64(x) <= y)
			default:
				return mustBool(false)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x <= int64(y))
			case int64:
				v = mustBool(x <= y)
			case uint32:
				v = mustBool(uint64(x) <= uint64(y))
			case uint64:
				v = mustBool(uint64(x) <= y)
			case float32:
				v = mustBool(float32(x) <= y)
			case float64:
				v = mustBool(float64(x) <= y)
			default:
				return mustBool(false)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x <= uint32(y))
			case int64:
				v = mustBool(uint64(x) <= uint64(y))
			case uint32:
				v = mustBool(x <= y)
			case uint64:
				v = mustBool(uint64(x) <= y)
			case float32:
				v = mustBool(float32(x) <= y)
			case float64:
				v = mustBool(float64(x) <= y)
			default:
				return mustBool(false)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x <= uint64(y))
			case int64:
				v = mustBool(x <= uint64(y))
			case uint32:
				v = mustBool(x <= uint64(y))
			case uint64:
				v = mustBool(x <= y)
			case float32:
				v = mustBool(float32(x) <= y)
			case float64:
				v = mustBool(float64(x) <= y)
			default:
				return mustBool(false)
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x <= float32(y))
			case int64:
				v = mustBool(x <= float32(y))
			case uint32:
				v = mustBool(x <= float32(y))
			case uint64:
				v = mustBool(x <= float32(y))
			case float32:
				v = mustBool(x <= y)
			case float64:
				v = mustBool(float64(x) <= y)
			default:
				return mustBool(false)
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x <= float64(y))
			case int64:
				v = mustBool(x <= float64(y))
			case uint32:
				v = mustBool(x <= float64(y))
			case uint64:
				v = mustBool(x <= float64(y))
			case float32:
				v = mustBool(x <= float64(y))
			case float64:
				v = mustBool(x <= y)
			default:
				return mustBool(false)
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = mustBool(x <= y)
			default:
				return mustBool(false)
			}
		default:
			return mustBool(false)
		}
	case 4: // RelationalExpression ">=" ShiftExpression
		v := t.EvalRelationalExpression(ex.RelationalExpression)
		v2 := t.EvalShiftExpression(ex.ShiftExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x >= y)
			case int64:
				v = mustBool(int64(x) >= y)
			case uint32:
				v = mustBool(uint32(x) >= y)
			case uint64:
				v = mustBool(uint64(x) >= y)
			case float32:
				v = mustBool(float32(x) >= y)
			case float64:
				v = mustBool(float64(x) >= y)
			default:
				return mustBool(false)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x >= int64(y))
			case int64:
				v = mustBool(x >= y)
			case uint32:
				v = mustBool(uint64(x) >= uint64(y))
			case uint64:
				v = mustBool(uint64(x) >= y)
			case float32:
				v = mustBool(float32(x) >= y)
			case float64:
				v = mustBool(float64(x) >= y)
			default:
				return mustBool(false)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x >= uint32(y))
			case int64:
				v = mustBool(uint64(x) >= uint64(y))
			case uint32:
				v = mustBool(x >= y)
			case uint64:
				v = mustBool(uint64(x) >= y)
			case float32:
				v = mustBool(float32(x) >= y)
			case float64:
				v = mustBool(float64(x) >= y)
			default:
				return mustBool(false)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x >= uint64(y))
			case int64:
				v = mustBool(x >= uint64(y))
			case uint32:
				v = mustBool(x >= uint64(y))
			case uint64:
				v = mustBool(x >= y)
			case float32:
				v = mustBool(float32(x) >= y)
			case float64:
				v = mustBool(float64(x) >= y)
			default:
				return mustBool(false)
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x >= float32(y))
			case int64:
				v = mustBool(x >= float32(y))
			case uint32:
				v = mustBool(x >= float32(y))
			case uint64:
				v = mustBool(x >= float32(y))
			case float32:
				v = mustBool(x >= y)
			case float64:
				v = mustBool(float64(x) >= y)
			default:
				return mustBool(false)
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = mustBool(x >= float64(y))
			case int64:
				v = mustBool(x >= float64(y))
			case uint32:
				v = mustBool(x >= float64(y))
			case uint64:
				v = mustBool(x >= float64(y))
			case float32:
				v = mustBool(x >= float64(y))
			case float64:
				v = mustBool(x >= y)
			default:
				return mustBool(false)
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = mustBool(x >= y)
			default:
				return mustBool(false)
			}
		default:
			return mustBool(false)
		}
	}
	return nil
}

func (t *Translator) EvalShiftExpression(ex *cc.ShiftExpression) Value {
	switch ex.Case {
	case 0: // AdditiveExpression
		return t.EvalAdditiveExpression(ex.AdditiveExpression)
	case 1: // ShiftExpression "<<" AdditiveExpression
		v := t.EvalShiftExpression(ex.ShiftExpression)
		v2 := t.EvalAdditiveExpression(ex.AdditiveExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		}
		return v
	case 2: // ShiftExpression ">>" AdditiveExpression
		v := t.EvalShiftExpression(ex.ShiftExpression)
		v2 := t.EvalAdditiveExpression(ex.AdditiveExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		}
		return v
	default:
		return nil
	}
}

func (t *Translator) EvalAdditiveExpression(ex *cc.AdditiveExpression) Value {
	switch ex.Case {
	case 0: // MultiplicativeExpression
		return t.EvalMultiplicativeExpression(ex.MultiplicativeExpression)
	case 1: // AdditiveExpression '+' MultiplicativeExpression
		v := t.EvalAdditiveExpression(ex.AdditiveExpression)
		v2 := t.EvalMultiplicativeExpression(ex.MultiplicativeExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x + y
			case int64:
				v = int64(x) + y
			case uint32:
				v = uint32(x) + y
			case uint64:
				v = uint64(x) + y
			case float32:
				v = float32(x) + y
			case float64:
				v = float64(x) + y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x + int64(y)
			case int64:
				v = x + y
			case uint32:
				v = uint64(x) + uint64(y)
			case uint64:
				v = uint64(x) + y
			case float32:
				v = float32(x) + y
			case float64:
				v = float64(x) + y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x + uint32(y)
			case int64:
				v = uint64(x) + uint64(y)
			case uint32:
				v = x + y
			case uint64:
				v = uint64(x) + y
			case float32:
				v = float32(x) + y
			case float64:
				v = float64(x) + y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x + uint64(y)
			case int64:
				v = x + uint64(y)
			case uint32:
				v = x + uint64(y)
			case uint64:
				v = x + y
			case float32:
				v = float32(x) + y
			case float64:
				v = float64(x) + y
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = x + float32(y)
			case int64:
				v = x + float32(y)
			case uint32:
				v = x + float32(y)
			case uint64:
				v = x + float32(y)
			case float32:
				v = x + y
			case float64:
				v = float64(x) + y
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = x + float64(y)
			case int64:
				v = x + float64(y)
			case uint32:
				v = x + float64(y)
			case uint64:
				v = x + float64(y)
			case float32:
				v = x + float64(y)
			case float64:
				v = x + y
			}
		case string:
			switch y := v2.(type) {
			case string:
				v = x + y
			}
		}
		return v
	case 2: // AdditiveExpression '-' MultiplicativeExpression
		v := t.EvalAdditiveExpression(ex.AdditiveExpression)
		v2 := t.EvalMultiplicativeExpression(ex.MultiplicativeExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x - y
			case int64:
				v = int64(x) - y
			case uint32:
				v = uint32(x) - y
			case uint64:
				v = uint64(x) - y
			case float32:
				v = float32(x) - y
			case float64:
				v = float64(x) - y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x - int64(y)
			case int64:
				v = x - y
			case uint32:
				v = uint64(x) - uint64(y)
			case uint64:
				v = uint64(x) - y
			case float32:
				v = float32(x) - y
			case float64:
				v = float64(x) - y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x - uint32(y)
			case int64:
				v = uint64(x) - uint64(y)
			case uint32:
				v = x - y
			case uint64:
				v = uint64(x) - y
			case float32:
				v = float32(x) - y
			case float64:
				v = float64(x) - y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x - uint64(y)
			case int64:
				v = x - uint64(y)
			case uint32:
				v = x - uint64(y)
			case uint64:
				v = x - y
			case float32:
				v = float32(x) - y
			case float64:
				v = float64(x) - y
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = x - float32(y)
			case int64:
				v = x - float32(y)
			case uint32:
				v = x - float32(y)
			case uint64:
				v = x - float32(y)
			case float32:
				v = x - y
			case float64:
				v = float64(x) - y
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = x - float64(y)
			case int64:
				v = x - float64(y)
			case uint32:
				v = x - float64(y)
			case uint64:
				v = x - float64(y)
			case float32:
				v = x - float64(y)
			case float64:
				v = x - y
			}
		}
		return v
	default:
		return nil
	}
}

func (t *Translator) EvalMultiplicativeExpression(ex *cc.MultiplicativeExpression) Value {
	switch ex.Case {
	case 0: // CastExpression
		return t.EvalCastExpression(ex.CastExpression)
	case 1: // MultiplicativeExpression '*' CastExpression
		v := t.EvalMultiplicativeExpression(ex.MultiplicativeExpression)
		v2 := t.EvalCastExpression(ex.CastExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x * y
			case int64:
				v = int64(x) * y
			case uint32:
				v = uint32(x) * y
			case uint64:
				v = uint64(x) * y
			case float32:
				v = float32(x) * y
			case float64:
				v = float64(x) * y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x * int64(y)
			case int64:
				v = x * y
			case uint32:
				v = uint64(x) * uint64(y)
			case uint64:
				v = uint64(x) * y
			case float32:
				v = float32(x) * y
			case float64:
				v = float64(x) * y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x * uint32(y)
			case int64:
				v = uint64(x) * uint64(y)
			case uint32:
				v = x * y
			case uint64:
				v = uint64(x) * y
			case float32:
				v = float32(x) * y
			case float64:
				v = float64(x) * y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x * uint64(y)
			case int64:
				v = x * uint64(y)
			case uint32:
				v = x * uint64(y)
			case uint64:
				v = x * y
			case float32:
				v = float32(x) * y
			case float64:
				v = float64(x) * y
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				v = x * float32(y)
			case int64:
				v = x * float32(y)
			case uint32:
				v = x * float32(y)
			case uint64:
				v = x * float32(y)
			case float32:
				v = x * y
			case float64:
				v = float64(x) * y
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				v = x * float64(y)
			case int64:
				v = x * float64(y)
			case uint32:
				v = x * float64(y)
			case uint64:
				v = x * float64(y)
			case float32:
				v = x * float64(y)
			case float64:
				v = x * y
			}
		}
		return v
	case 2: // MultiplicativeExpression '/' CastExpression
		v := t.EvalMultiplicativeExpression(ex.MultiplicativeExpression)
		v2 := t.EvalCastExpression(ex.CastExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = int32(0)
					break
				}
				v = x / y
			case int64:
				if y == 0 {
					v = int64(0)
					break
				}
				v = int64(x) / y
			case uint32:
				if y == 0 {
					v = uint32(0)
					break
				}
				v = uint32(x) / y
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) / y
			case float32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = float32(x) / y
			case float64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = float64(x) / y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = int64(0)
					break
				}
				v = x / int64(y)
			case int64:
				if y == 0 {
					v = int64(0)
					break
				}
				v = x / y
			case uint32:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) / uint64(y)
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) / y
			case float32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = float32(x) / y
			case float64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = float64(x) / y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = uint32(0)
					break
				}
				v = x / uint32(y)
			case int64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) / uint64(y)
			case uint32:
				if y == 0 {
					v = uint32(0)
					break
				}
				v = x / y
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) / y
			case float32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = float32(x) / y
			case float64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = float64(x) / y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x / uint64(y)
			case int64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x / uint64(y)
			case uint32:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x / uint64(y)
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x / y
			case float32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = float32(x) / y
			case float64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = float64(x) / y
			}
		case float32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = x / float32(y)
			case int64:
				if y == 0 {
					v = float32(0)
					break
				}
				v = x / float32(y)
			case uint32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = x / float32(y)
			case uint64:
				if y == 0 {
					v = float32(0)
					break
				}
				v = x / float32(y)
			case float32:
				if y == 0 {
					v = float32(0)
					break
				}
				v = x / y
			case float64:
				if y == 0 {
					v = float32(0)
					break
				}
				v = float64(x) / y
			}
		case float64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = float64(0)
					break
				}
				v = x / float64(y)
			case int64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = x / float64(y)
			case uint32:
				if y == 0 {
					v = float64(0)
					break
				}
				v = x / float64(y)
			case uint64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = x / float64(y)
			case float32:
				if y == 0 {
					v = float64(0)
					break
				}
				v = x / float64(y)
			case float64:
				if y == 0 {
					v = float64(0)
					break
				}
				v = x / y
			}
		}
		return v
	case 3: // MultiplicativeExpression '%' CastExpression
		v := t.EvalMultiplicativeExpression(ex.MultiplicativeExpression)
		v2 := t.EvalCastExpression(ex.CastExpression)
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = int32(0)
					break
				}
				v = x % y
			case int64:
				if y == 0 {
					v = int64(0)
					break
				}
				v = int64(x) % y
			case uint32:
				if y == 0 {
					v = uint32(0)
					break
				}
				v = uint32(x) % y
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) % y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = int64(0)
					break
				}
				v = x % int64(y)
			case int64:
				if y == 0 {
					v = int64(0)
					break
				}
				v = x % y
			case uint32:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) % uint64(y)
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) % y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = uint32(0)
					break
				}
				v = x % uint32(y)
			case int64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) % uint64(y)
			case uint32:
				if y == 0 {
					v = uint32(0)
					break
				}
				v = x % y
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = uint64(x) % y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x % uint64(y)
			case int64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x % uint64(y)
			case uint32:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x % uint64(y)
			case uint64:
				if y == 0 {
					v = uint64(0)
					break
				}
				v = x % y
			}
		}
		return v
	default:
		return nil
	}
}

func (t *Translator) EvalCastExpression(ex *cc.CastExpression) Value {
	switch ex.Case {
	case 0: // UnaryExpression
		return t.EvalUnaryExpression(ex.UnaryExpression)
	case 1: // '(' TypeName ')' CastExpression
		return t.EvalCastExpression(ex.CastExpression)
	default:
		return nil
	}
}

func (t *Translator) EvalUnaryExpression(ex *cc.UnaryExpression) Value {
	switch ex.Case {
	case 0: // PostfixExpression
		return t.EvalPostfixExpression(ex.PostfixExpression)
	case 1: // "++" UnaryExpression
		x := t.EvalUnaryExpression(ex.UnaryExpression)
		switch x := x.(type) {
		case int32:
			return x + 1
		case int64:
			return x + 1
		case uint32:
			return x + 1
		case uint64:
			return x + 1
		default:
			return x
		}
	case 2: // "--" UnaryExpression
		x := t.EvalUnaryExpression(ex.UnaryExpression)
		switch x := x.(type) {
		case int32:
			return x - 1
		case int64:
			return x - 1
		case uint32:
			return x - 1
		case uint64:
			return x - 1
		default:
			return x
		}
	case 3: // UnaryOperator CastExpression
		x := t.EvalCastExpression(ex.CastExpression)
		switch ex.UnaryOperator.Case {
		case 2: // '+'
			switch x := x.(type) {
			case int32:
				return +x
			case int64:
				return +x
			case uint32:
				return +x
			case uint64:
				return +x
			case float32:
				return +x
			case float64:
				return +x
			default:
				return x
			}
		case 3: // '-'
			switch x := x.(type) {
			case int32:
				return -x
			case int64:
				return -x
			case uint32:
				return -x
			case uint64:
				return -x
			case float32:
				return -x
			case float64:
				return -x
			default:
				return x
			}
		case 4: // '~'
			switch x := x.(type) {
			case int32:
				return ^x
			case int64:
				return ^x
			case uint32:
				return ^x
			case uint64:
				return ^x
			default:
				return x
			}
		case 5: // '!'
			return mustBool(isZero(x))
		default:
			return x
		}
	case 4, // "sizeof" UnaryExpression
		5: // "sizeof" '(' TypeName ')'
		// sizeof will be always 0, why not?
		return mustInt(0)
	case 6, // "defined" IDENTIFIER
		7: // "defined" '(' IDENTIFIER ')'
		if _, ok := cc.Macros[xc.Dict.ID(ex.Token.S())]; ok {
			return mustBool(true)
		}
		return mustBool(false)
	default:
		return nil
	}
}

func (t *Translator) EvalPostfixExpression(ex *cc.PostfixExpression) Value {
	switch ex.Case {
	case 0: // PrimaryExpression
		return t.EvalPrimaryExpression(ex.PrimaryExpression)
	case 5, // PostfixExpression "++"
		6: // PostfixExpression "--"
		return t.EvalPostfixExpression(ex.PostfixExpression)
	default:
		return nil
	}
}

func (t *Translator) EvalPrimaryExpression(ex *cc.PrimaryExpression) Value {
	switch ex.Case {
	case 0: // IDENTIFIER
		return t.valueMap[string(ex.Token.S())]
	case 1: // Constant
		return (*constantValue)(ex.Constant).Eval()
	case 2:
		// get only the first one from expression list
		return t.EvalExpressionList(ex.ExpressionList)
	default:
		return nil
	}
}

func (t *Translator) EvalExpressionList(ex *cc.ExpressionList) Value {
	switch ex.Case {
	case 0, // AssignmentExpression
		1: // ExpressionList ',' AssignmentExpression
		// get only the first one from the list
		return t.EvalAssignmentExpression(ex.AssignmentExpression)
	default:
		return nil
	}
}

func (t *Translator) EvalArgumentExpressionList(ex *cc.ArgumentExpressionList) Value {
	switch ex.Case {
	case 0, // AssignmentExpression
		1: // ExpressionList ',' AssignmentExpression
		// get only the first one from the list
		return t.EvalAssignmentExpression(ex.AssignmentExpression)
	default:
		return nil
	}
}

func (cv *constantValue) Eval() Value {
	switch cv.Case {
	case 0, // CHARCONST
		3: // LONGCHARCONST
		return charConst(cv.Token.S())
	case 1: // FLOATCONST
		return floatConst(cv.Token.S())
	case 2: // INTCONST
		return intConst(cv.Token.S())
	case 4, // LONGSTRINGLITERAL
		5: // STRINGLITERAL
		return stringConst(cv.Token.S())
	default:
		return nil
	}
}

func mustInt(v interface{}) int32 {
	switch x := v.(type) {
	case int:
		return int32(x)
	case int32:
		return x
	case int64:
		return int32(x)
	default:
		panic(fmt.Sprintf("%T can't be casted to int", v))
	}
}

func mustUint(v interface{}) uint32 {
	switch x := v.(type) {
	case uint32:
		return x
	case uint64:
		return uint32(x)
	default:
		panic(fmt.Sprintf("%T can't be casted to unsigned int", v))
	}
}

func mustLong(v interface{}) int64 {
	switch x := v.(type) {
	case int64:
		return x
	case int:
		return int64(x)
	default:
		panic(fmt.Sprintf("%T can't be casted to long", v))
	}
}

func mustUlong(v interface{}) uint64 {
	switch x := v.(type) {
	case uint64:
		return x
	default:
		panic(fmt.Sprintf("%T can't be cased to unsigned long", v))
	}
}

func mustBool(v bool) interface{} {
	if v {
		return mustInt(1)
	}
	return mustInt(0)
}

func isZero(v interface{}) bool {
	switch x := v.(type) {
	case int32:
		return x == 0
	case int64:
		return x == 0
	case uint32:
		return x == 0
	case uint64:
		return x == 0
	}
	panic(fmt.Sprintf("isZero can't be used with %T", v))
}

func floatConst(b []byte) float64 {
	f, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return 0
	}
	return f
}

func stringConst(b []byte) string {
	if len(b) < 2 {
		return ""
	}
	return string(b[1 : len(b)-1])
}

func intConst(b []byte) interface{} {
	if len(b) == 0 {
		return mustInt(0)
	}
	const (
		l = 1 << iota
		ll
		u
	)
	k := 0
	i := len(b) - 1
more:
	switch c := b[i]; c {
	case 'u', 'U':
		k |= u
		i--
		goto more
	case 'l', 'L':
		if i > 0 && (b[i-1] == 'l' || b[i-1] == 'L') {
			k |= ll
			i -= 2
			goto more
		}

		k |= l
		i--
		goto more
	}
	n, err := strconv.ParseUint(string(b[:i+1]), 0, 64)
	if err != nil {
		return mustInt(0)
	}

	switch {
	case k == 0:
		return mustInt(int(n))
	case k == l, k == ll:
		return mustLong(int64(n))
	case k == u:
		return mustUint(n)
	case k == u|l, k == u|ll:
		return mustUlong(n)
	default:
		return mustInt(0)
	}
}

func charConst(b []byte) interface{} {
	if len(b) == 0 {
		return mustInt(0)
	}
	var long bool
	if b[0] == 'L' {
		long = true
		b = b[1:]
	}
	if len(b) == 3 { // 'A'
		if long {
			return mustLong(int(b[1]))
		}
		return mustInt(int(b[1]))
	}
	switch b[2] {
	case '0', '1', '2', '3', '4', '5', '6', '7':
		v := 0
		i := 0
		for _, c := range b[2:] {
			if c < '0' || c > '7' {
				break
			}
			v = 8*v + int(c) - '0'
			i++
			if i == 3 {
				break
			}
		}
		if long {
			return mustLong(v)
		}
		return mustInt(int(v))
	case '\'':
		return mustInt('\'')
	case '"':
		return mustInt('"')
	case '?':
		return mustInt('?')
	case '\\':
		return mustInt('\\')
	case 'a':
		return mustInt('\a')
	case 'b':
		return mustInt('\b')
	case 'f':
		return mustInt('\f')
	case 'n':
		return mustInt('\n')
	case 'r':
		return mustInt('\r')
	case 't':
		return mustInt('\t')
	case 'v':
		return mustInt('\v')
	default:
		return mustInt(0)
	}
}
