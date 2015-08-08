package translator

type Rules map[RuleTarget][]RuleSpec
type ConstRules map[ConstScope]ConstRule

type RuleSpec struct {
	From, To  string
	Action    RuleAction
	Transform RuleTransform
	Load      string
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
	TargetGlobal  RuleTarget = "global"
	TargetDefine             = "define"
	TargetTag                = "tag"
	TargetTypedef            = "typedef"
	TargetDeclare            = "declare"
)

type ConstRule string

const (
	ConstCGOAlias   ConstRule = "cgo_alias"
	ConstExpand               = "expand"
	ConstExpandFull           = "expand_full"
	ConstEval                 = "eval"
	ConstEvalFull             = "eval_full"
)

type ConstScope string

const (
	ConstEnum    = "enum"
	ConstDeclare = "declare"
)

var builtinRules = map[string]RuleSpec{
	"snakecase": RuleSpec{From: "_([^_]+)", To: "$1", Action: ActionReplace, Transform: TransformTitle},
}
