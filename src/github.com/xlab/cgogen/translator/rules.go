package translator

type Rules map[RuleTarget][]RuleSpec
type ConstRules map[ConstScope]ConstRule
type PointerLayouts map[PointerScope][]PointerLayoutSpec

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
	TransformLower    RuleTransform = "lower"
	TransformTitle    RuleTransform = "title"
	TransformExport   RuleTransform = "export"
	TransformUnexport RuleTransform = "unexport"
	TransformUpper    RuleTransform = "upper"
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
	ConstCGOAlias   ConstRule = "cgo"
	ConstExpand     ConstRule = "expand"
	ConstExpandFull ConstRule = "expandFull"
	ConstEval       ConstRule = "eval"
	ConstEvalFull   ConstRule = "evalFull"
)

type ConstScope string

const (
	ConstEnum    ConstScope = "enums"
	ConstDeclare ConstScope = "declares"
)

type PointerSpec string

const (
	PointerRef PointerSpec = "ref"
	PointerArr PointerSpec = "arr"
)

type PointerLayoutSpec struct {
	Name    string
	Layout  []PointerSpec
	Default PointerSpec
}

type PointerScope string

const (
	PointerScopeAny      PointerScope = "any"
	PointerScopeStruct   PointerScope = "struct"
	PointerScopeFunction PointerScope = "function"
)

var builtinRules = map[string]RuleSpec{
	"snakecase": RuleSpec{From: "_([^_]+)", To: "$1", Action: ActionReplace, Transform: TransformTitle},
}
