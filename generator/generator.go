package generator

import (
	"errors"
	"io"
	"math/rand"
	"sort"

	tl "github.com/xlab/c-for-go/translator"
)

type Generator struct {
	pkg string
	cfg *Config
	tr  *tl.Translator
	//
	closed        bool
	closeC, doneC chan struct{}
	helpersChan   chan *Helper
	rand          *rand.Rand
	noTimestamps  bool
	maxMem        MemSpec
}

func (g *Generator) DisableTimestamps() {
	g.noTimestamps = true
}

type TraitFlagGroup struct {
	Name   string   `yaml:"name"`
	Traits []string `yaml:"traits"`
	Flags  []string `yaml:"flags"`
}

type Config struct {
	PackageName        string           `yaml:"PackageName"`
	PackageDescription string           `yaml:"PackageDescription"`
	PackageLicense     string           `yaml:"PackageLicense"`
	PkgConfigOpts      []string         `yaml:"PkgConfigOpts"`
	FlagGroups         []TraitFlagGroup `yaml:"FlagGroups"`
	SysIncludes        []string         `yaml:"SysIncludes"`
	Includes           []string         `yaml:"Includes"`
	Options            GenOptions       `yaml:"Options"`
}

type GenOptions struct {
	SafeStrings     bool `yaml:"SafeStrings"`
	StructAccessors bool `yaml:"StructAccessors"`
	KeepAlive       bool `yaml:"KeepAlive"`
}

func New(pkg string, cfg *Config, tr *tl.Translator) (*Generator, error) {
	if cfg == nil || len(cfg.PackageName) == 0 {
		return nil, errors.New("no package name provided")
	} else if tr == nil {
		return nil, errors.New("no translator provided")
	}
	gen := &Generator{
		pkg: pkg,
		cfg: cfg,
		tr:  tr,
		//
		helpersChan: make(chan *Helper, 1),
		closeC:      make(chan struct{}),
		doneC:       make(chan struct{}),
		rand:        rand.New(rand.NewSource(+79269965690)),
		maxMem:      MemSpecDefault,
	}
	return gen, nil
}

type declList []*tl.CDecl

func (s declList) Len() int      { return len(s) }
func (s declList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s declList) Less(i, j int) bool {
	if s[i].Position.Filename != s[j].Position.Filename {
		return s[i].Position.Filename < s[j].Position.Filename
	} else if s[i].Position.Offset != s[j].Position.Offset {
		return s[i].Position.Offset < s[j].Position.Offset
	} else {
		return s[i].Name < s[j].Name
	}
}

func (gen *Generator) WriteConst(wr io.Writer) int {
	var count int
	if defines := gen.tr.Defines(); len(defines) > 0 {
		n := gen.writeDefinesGroup(wr, defines)
		count = count + n
	}
	writeSpace(wr, 1)
	tagsSeen := make(map[string]bool)
	namesSeen := make(map[string]bool)

	gen.submitHelper(cgoGenTag)
	expandEnum := func(decl *tl.CDecl) bool {
		if tag := decl.Spec.GetTag(); len(tag) == 0 {
			gen.expandEnumAnonymous(wr, decl, namesSeen)
			return true
		} else if tagsSeen[tag] {
			return false
		} else {
			gen.expandEnum(wr, decl, namesSeen)
			if decl.Spec.IsComplete() {
				tagsSeen[tag] = true
			}
			return true
		}
	}

	enumList := make([]*tl.CDecl, 0, len(gen.tr.TagMap()))
	for tag, decl := range gen.tr.TagMap() {
		if decl.Spec.Kind() != tl.EnumKind {
			continue
		} else if !decl.Spec.IsComplete() {
			continue
		}
		if !gen.tr.IsAcceptableName(tl.TargetType, tag) {
			continue
		}
		enumList = append(enumList, decl)
	}
	sort.Sort(declList(enumList))
	for i := range enumList {
		if expandEnum(enumList[i]) {
			count++
		}
	}

	for _, decl := range gen.tr.Typedefs() {
		if decl.Spec.Kind() != tl.EnumKind {
			continue
		} else if !decl.Spec.IsComplete() {
			continue
		}
		if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
			continue
		}
		if expandEnum(decl) {
			count++
		}
	}
	for _, decl := range gen.tr.Declares() {
		if decl.Spec.Kind() == tl.EnumKind {
			if !decl.Spec.IsComplete() {
				continue
			}
			if len(decl.Name) > 0 {
				if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
					continue
				}
			}
			if expandEnum(decl) {
				count++
			}
			continue
		}
		if decl.IsStatic {
			continue
		}
		if len(decl.Name) == 0 {
			continue
		}
		if decl.Spec.Kind() != tl.TypeKind {
			continue
		} else if !decl.Spec.IsConst() {
			continue
		}
		if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
			continue
		}
		gen.writeConstDeclaration(wr, decl)
		writeSpace(wr, 1)
		count++
	}
	return count
}

func (gen *Generator) MemTipOf(decl *tl.CDecl) tl.Tip {
	var memTip tl.Tip
	if name := decl.Spec.CGoName(); len(name) > 0 {
		if memTipRx, ok := gen.tr.MemTipRx(name); ok {
			memTip = memTipRx.Self()
		}
	}
	return memTip
}

func (gen *Generator) WriteTypedefs(wr io.Writer) int {
	var count int
	typedefs := gen.tr.Typedefs()
	seenStructTags := make(map[string]bool, len(typedefs))
	seenUnionTags := make(map[string]bool, len(typedefs))
	seenTypeNames := make(map[string]bool, len(typedefs))
	seenStructNames := make(map[string]bool, len(typedefs))
	seenUnionNames := make(map[string]bool, len(typedefs))
	seenFunctionNames := make(map[string]bool, len(typedefs))
	for _, decl := range typedefs {
		if !gen.tr.IsAcceptableName(tl.TargetType, decl.Name) {
			continue
		}
		switch decl.Spec.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			if tag := decl.Spec.GetTag(); len(tag) > 0 {
				if len(decl.Name) == 0 || decl.Name == tag {
					if seenStructTags[tag] {
						continue
					}
				}
				seenStructTags[tag] = true
			}
			memTip := gen.MemTipOf(decl)
			gen.writeStructTypedef(wr, decl, memTip == tl.TipMemRaw, seenStructNames)
		case tl.UnionKind:
			if len(decl.Name) > 0 {
				if seenUnionNames[decl.Name] {
					continue
				}
				seenUnionNames[decl.Name] = true
			}
			if tag := decl.Spec.GetTag(); len(tag) > 0 {
				if len(decl.Name) == 0 || decl.Name == tag {
					if seenUnionTags[tag] {
						continue
					}
				}
				seenUnionTags[tag] = true
			}
			gen.writeUnionTypedef(wr, decl)
		case tl.EnumKind:
			if !decl.Spec.IsComplete() {
				gen.writeEnumTypedef(wr, decl)
			}
		case tl.TypeKind:
			gen.writeTypeTypedef(wr, decl, seenTypeNames)
		case tl.FunctionKind:
			gen.writeFunctionTypedef(wr, decl, seenFunctionNames)
		}
		writeSpace(wr, 1)
		count++
	}

	tagDefs := sortedTagDefs(gen.tr.TagMap())
	for _, def := range tagDefs {
		decl := def.tagDecl
		tag := def.tagName
		switch decl.Spec.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			if seenStructTags[tag] {
				continue
			}
			if !gen.tr.IsAcceptableName(tl.TargetPublic, tag) {
				continue
			} else if !gen.tr.IsAcceptableName(tl.TargetType, tag) {
				continue
			}
			if memTipRx, ok := gen.tr.MemTipRx(tag); ok {
				gen.writeStructTypedef(wr, decl, memTipRx.Self() == tl.TipMemRaw, seenStructNames)
			} else {
				gen.writeStructTypedef(wr, decl, false, seenStructNames)
			}
			writeSpace(wr, 1)
			count++
		case tl.UnionKind:
			if seenUnionTags[tag] {
				continue
			}
			if !gen.tr.IsAcceptableName(tl.TargetPublic, tag) {
				continue
			} else if !gen.tr.IsAcceptableName(tl.TargetType, tag) {
				continue
			}
			gen.writeUnionTypedef(wr, decl)
			writeSpace(wr, 1)
			count++
		}
	}
	return count
}

func (gen *Generator) WriteDeclares(wr io.Writer) int {
	var count int
	declares := gen.tr.Declares()
	seenStructs := make(map[string]bool, len(declares))
	seenUnions := make(map[string]bool, len(declares))
	seenEnums := make(map[string]bool, len(declares))
	seenFunctions := make(map[string]bool, len(declares))
	for _, decl := range declares {
		const public = true
		switch decl.Spec.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			if len(decl.Name) == 0 {
				continue
			} else if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			} else if seenStructs[decl.Name] {
				continue
			} else {
				seenStructs[decl.Name] = true
			}
			gen.writeStructDeclaration(wr, decl, tl.NoTip, tl.NoTip, public)
		case tl.UnionKind:
			if len(decl.Name) == 0 {
				continue
			} else if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			} else if seenUnions[decl.Name] {
				continue
			} else {
				seenUnions[decl.Name] = true
			}
			gen.writeUnionDeclaration(wr, decl, tl.NoTip, tl.NoTip, public)
		case tl.EnumKind:
			if !decl.Spec.IsComplete() {
				if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
					continue
				} else if seenEnums[decl.Name] {
					continue
				} else {
					seenEnums[decl.Name] = true
				}
				gen.writeEnumDeclaration(wr, decl, tl.NoTip, tl.NoTip, public)
			}
		case tl.FunctionKind:
			if !gen.tr.IsAcceptableName(tl.TargetFunction, decl.Name) {
				continue
			} else if seenFunctions[decl.Name] {
				continue
			} else {
				seenFunctions[decl.Name] = true
			}
			// defaults to ref for the returns
			ptrTip := tl.TipPtrRef
			if ptrTipRx, ok := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name); ok {
				if tip := ptrTipRx.Self(); tip.IsValid() {
					ptrTip = tip
				}
			}
			typeTip := tl.TipTypeNamed
			if typeTipRx, ok := gen.tr.TypeTipRx(tl.TipScopeFunction, decl.Name); ok {
				if tip := typeTipRx.Self(); tip.IsValid() {
					typeTip = tip
				}
			}
			gen.writeFunctionDeclaration(wr, decl, ptrTip, typeTip, public)
		}
		writeSpace(wr, 1)
		count++
	}
	return count
}

func (gen *Generator) Close() {
	if gen.closed {
		return
	}
	gen.closed = true
	gen.closeC <- struct{}{}
	<-gen.doneC
}

func (gen *Generator) MonitorAndWriteHelpers(goWr, chWr io.Writer, ccWr io.Writer, initWrFunc ...func() (io.Writer, error)) {
	seenHelperNames := make(map[string]bool)
	var seenGoHelper bool
	var seenCHHelper bool
	var seenCCHelper bool
	for {
		select {
		case <-gen.closeC:
			close(gen.helpersChan)
		case helper, ok := <-gen.helpersChan:
			if !ok {
				close(gen.doneC)
				return
			}
			if seenHelperNames[string(helper.Side)+helper.Name] {
				continue
			}
			seenHelperNames[string(helper.Side)+helper.Name] = true

			var wr io.Writer
			switch helper.Side {
			case NoSide, GoSide:
				if goWr != nil {
					wr = goWr
				} else if len(initWrFunc) < 1 {
					continue
				} else if w, err := initWrFunc[0](); err != nil {
					continue
				} else {
					wr = w
				}
				if !seenGoHelper {
					gen.writeGoHelpersHeader(wr)
					seenGoHelper = true
				}
			case CHSide:
				if chWr != nil {
					wr = chWr
				} else if len(initWrFunc) < 2 {
					continue
				} else if w, err := initWrFunc[1](); err != nil {
					continue
				} else {
					wr = w
				}
				if !seenCHHelper {
					gen.writeCHHelpersHeader(wr)
					seenCHHelper = true
				}
			case CCSide:
				if ccWr != nil {
					wr = ccWr
				} else if len(initWrFunc) < 3 {
					continue
				} else if w, err := initWrFunc[2](); err != nil {
					continue
				} else {
					wr = w
				}
				if !seenCCHelper {
					gen.writeCCHelpersHeader(wr)
					seenCCHelper = true
				}
			default:
				continue
			}

			if len(helper.Description) > 0 {
				writeTextBlock(wr, helper.Description)
			}
			writeSourceBlock(wr, helper.Source)
			writeSpace(wr, 1)
		}
	}
}

// randPostfix generates a simply random 4-byte postfix. Doesn't require a crypto package.
func (gen *Generator) randPostfix() int32 {
	return 0x0f000000 + gen.rand.Int31n(0x00ffffff)
}

type tagDef struct {
	tagName string
	tagDecl *tl.CDecl
}

func sortedTagDefs(m map[string]*tl.CDecl) []tagDef {
	tagDefs := make([]tagDef, 0, len(m))
	for tag, decl := range m {
		tagDefs = append(tagDefs, tagDef{
			tagName: tag,
			tagDecl: decl,
		})
	}
	sort.Slice(tagDefs, func(i, j int) bool {
		return tagDefs[i].tagName < tagDefs[j].tagName
	})
	return tagDefs
}

func NewMemSpec(spec string) MemSpec {
	switch s := MemSpec(spec); s {
	case MemSpec1,
		MemSpec2,
		MemSpec3,
		MemSpec4,
		MemSpec5,
		MemSpec6,
		MemSpec7,
		MemSpec8,
		MemSpec9,
		MemSpecA,
		MemSpecB,
		MemSpecC,
		MemSpecD,
		MemSpecE,
		MemSpecF:
		return s
	default:
		return MemSpecDefault
	}
}

type MemSpec string

const (
	MemSpec1 MemSpec = "0x1fffffff"
	MemSpec2 MemSpec = "0x2fffffff"
	MemSpec3 MemSpec = "0x3fffffff"
	MemSpec4 MemSpec = "0x4fffffff"
	MemSpec5 MemSpec = "0x5fffffff"
	MemSpec6 MemSpec = "0x6fffffff"
	MemSpec7 MemSpec = "0x7fffffff"
	MemSpec8 MemSpec = "0x8fffffff"
	MemSpec9 MemSpec = "0x9fffffff"
	MemSpecA MemSpec = "0xafffffff"
	MemSpecB MemSpec = "0xbfffffff"
	MemSpecC MemSpec = "0xcfffffff"
	MemSpecD MemSpec = "0xdfffffff"
	MemSpecE MemSpec = "0xefffffff"
	MemSpecF MemSpec = "0xffffffff"

	MemSpecDefault MemSpec = "0x7fffffff"
)

func (g *Generator) SetMaxMemory(spec MemSpec) {
	g.maxMem = spec
}
