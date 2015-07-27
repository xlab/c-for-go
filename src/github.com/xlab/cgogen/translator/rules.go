package translator

type Rules map[RuleTarget][]RuleSpec

type RuleSpec struct {
	From, To string
	Action   RuleAction
}

type RuleAction string

const (
	ActionAccept  RuleAction = "accept"
	ActionReplace            = "replace"
	ActionAbbr               = "abbr"
	ActionLower              = "lower"
	ActionUpper              = "upper"
	ActionIgnore             = "ignore"
)

type RuleTarget string

const (
	TargetGlobal   RuleTarget = "global"
	TargetType                = "type"
	TargetConst               = "const"
	TargetFunction            = "function"
)
