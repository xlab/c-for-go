package translator

type Rules map[RuleTarget][]RuleSpec
type ConstRules map[ConstScope]ConstRule
type PtrTips map[TipScope][]TipSpec
type MemTips []TipSpec

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
	NoTarget         RuleTarget = ""
	TargetGlobal     RuleTarget = "global"
	TargetPostGlobal RuleTarget = "post-global"
	//
	TargetConst    RuleTarget = "const"
	TargetType     RuleTarget = "type"
	TargetFunction RuleTarget = "function"
	//
	TargetPublic  RuleTarget = "public"
	TargetPrivate RuleTarget = "private"
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
	ConstEnum ConstScope = "enum"
	ConstDecl ConstScope = "decl"
)

type Tip string

const (
	TipPtrSRef Tip = "sref"
	TipPtrRef  Tip = "ref"
	TipPtrArr  Tip = "arr"
	TipMemRaw  Tip = "raw"
	NoTip      Tip = ""
)

func (t Tip) IsValid() bool {
	switch t {
	case TipPtrArr, TipPtrRef, TipPtrSRef, TipMemRaw:
		return true
	default:
		return false
	}
}

type TipSpec struct {
	Target  string
	Tips    Tips
	Self    Tip
	Default Tip
}

type TipScope string

const (
	TipScopeAny      TipScope = "any"
	TipScopeStruct   TipScope = "struct"
	TipScopeFunction TipScope = "function"
)

type Tips []Tip

var builtinRules = map[string]RuleSpec{
	"snakecase": RuleSpec{From: "_([^_]+)", To: "$1", Action: ActionReplace, Transform: TransformTitle},
}
