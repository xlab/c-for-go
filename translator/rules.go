package translator

import (
	"github.com/dlclark/regexp2"
)

type Rules map[RuleTarget][]RuleSpec
type ConstRules map[ConstScope]ConstRule
type PtrTips map[TipScope][]TipSpec
type TypeTips map[TipScope][]TipSpec
type MemTips []TipSpec

type Validations []ValidationSpec

type ValidationSpec struct {
	ValidateFunc string
	Ret          string
	MatchedFunc  string
}

func (v ValidationSpec) MatchFunc(name string) bool {
	reg := regexp2.MustCompile(v.MatchedFunc, 0)
	matched, _ := reg.MatchString(name)
	return matched
}

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
	ConstCGOAlias ConstRule = "cgo"
	ConstExpand   ConstRule = "expand"
	ConstEval     ConstRule = "eval"
)

type ConstScope string

const (
	ConstEnum    ConstScope = "enum"
	ConstDecl    ConstScope = "decl"
	ConstDefines ConstScope = "defines"
)

type Tip string

const (
	TipPtrSRef      Tip = "sref"
	TipPtrRef       Tip = "ref"
	TipPtrArr       Tip = "arr"
	TipPtrInst      Tip = "inst"
	TipMemRaw       Tip = "raw"
	TipTypeNamed    Tip = "named"
	TipTypePlain    Tip = "plain"
	TipTypeString   Tip = "string"
	TipTypeUnsigned Tip = "unsigned"
	NoTip           Tip = ""
)

type TipKind string

const (
	TipKindUnknown TipKind = "unknown"
	TipKindPtr     TipKind = "ptr"
	TipKindType    TipKind = "type"
	TipKindMem     TipKind = "mem"
)

func (t Tip) Kind() TipKind {
	switch t {
	case TipPtrArr, TipPtrRef, TipPtrSRef, TipPtrInst:
		return TipKindPtr
	case TipTypePlain, TipTypeNamed, TipTypeString, TipTypeUnsigned:
		return TipKindType
	case TipMemRaw:
		return TipKindMem
	default:
		return TipKindUnknown
	}
}

func (t Tip) IsValid() bool {
	switch t {
	case TipPtrArr, TipPtrRef, TipPtrSRef, TipPtrInst:
		return true
	case TipTypePlain, TipTypeNamed, TipTypeString:
		return true
	case TipMemRaw:
		return true
	case TipTypeUnsigned:
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
	TipScopeType     TipScope = "type"
	TipScopeEnum     TipScope = "enum"
	TipScopeFunction TipScope = "function"
)

type Tips []Tip

var builtinRules = map[string]RuleSpec{
	"snakecase":  {Action: ActionReplace, From: "_([^_]+)", To: "$1", Transform: TransformTitle},
	"doc.file":   {Action: ActionDocument, To: "$path:$line"},
	"doc.google": {Action: ActionDocument, To: "https://google.com/search?q=$file+$name"},
}
