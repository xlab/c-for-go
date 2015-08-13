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
	ActionAccept  RuleAction = "accept"
	ActionIgnore  RuleAction = "ignore"
	ActionReplace RuleAction = "replace"
)

var ruleActions = []RuleAction{
	ActionAccept, ActionIgnore, ActionReplace,
}

type RuleTransform string

const (
	TransformLower RuleTransform = "lower"
	TransformTitle RuleTransform = "title"
	TransformUpper RuleTransform = "upper"
)

type RuleTarget string

const (
	TargetGlobal     RuleTarget = "global"
	TargetPostGlobal RuleTarget = "post-global"
	//
	TargetConst    RuleTarget = "const"
	TargetType     RuleTarget = "type"
	TargetPublic   RuleTarget = "public"
	TargetPrivate  RuleTarget = "private"
	TargetFunction RuleTarget = "function"
)

type ConstRule string

const (
	ConstCGOAlias   ConstRule = "cgo_alias"
	ConstExpand     ConstRule = "expand"
	ConstExpandFull ConstRule = "expand_full"
	ConstEval       ConstRule = "eval"
	ConstEvalFull   ConstRule = "eval_full"
)

type ConstScope string

const (
	ConstEnum    ConstScope = "enum"
	ConstDeclare ConstScope = "declare"
)

var builtinRules = map[string]RuleSpec{
	"snakecase": RuleSpec{From: "_([^_]+)", To: "$1", Action: ActionReplace, Transform: TransformTitle},
}
