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

func (r *RuleSpec) LoadSpec(r2 RuleSpec) {
	if len(r.From) == 0 {
		r.From = r2.From
	}
	if len(r.To) == 0 {
		r.To = r2.To
	}
	if len(r.Action) == 0 {
		r.Action = r2.Action
	}
	if len(r.Transform) == 0 {
		r.Transform = r2.Transform
	}
}

type RuleAction string

const (
	ActionNone     RuleAction = ""
	ActionAccept   RuleAction = "accept"
	ActionIgnore   RuleAction = "ignore"
	ActionReplace  RuleAction = "replace"
	ActionDocument RuleAction = "doc"
)

var ruleActions = []RuleAction{
	ActionAccept, ActionIgnore, ActionReplace, ActionDocument,
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
)

type ConstScope string

const (
	ConstEnum    ConstScope = "enum"
	ConstDecl    ConstScope = "decl"
	ConstDefines ConstScope = "defines"
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
	"snakecase":  RuleSpec{Action: ActionReplace, From: "_([^_]+)", To: "$1", Transform: TransformTitle},
	"doc.file":   RuleSpec{Action: ActionDocument, To: "$path:$line"},
	"doc.google": RuleSpec{Action: ActionDocument, To: "https://google.com/search?q=$file+$name"},
}
