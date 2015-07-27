package translator

type Rules map[RuleTarget][]RuleSpec

type RuleSpec struct {
	From, To  string
	Action    RuleAction
	Transform RuleTransform
}

type RuleAction string

const (
	ActionNone    RuleAction = ""
	ActionAccept             = "accept"
	ActionIgnore             = "ignore"
	ActionReplace            = "replace"
)

var ruleActions = []RuleAction{
	ActionAccept, ActionIgnore, ActionReplace,
}

type RuleTransform string

const (
	TransformLower RuleTransform = "lower"
	TransformTitle               = "title"
	TransformUpper               = "upper"
)

type RuleTarget string

const (
	TargetGlobal   RuleTarget = "global"
	TargetType                = "type"
	TargetConst               = "const"
	TargetFunction            = "function"
)
