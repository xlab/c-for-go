package generator

import (
	"errors"
	"io"
	"math/rand"
	"sort"

	tl "github.com/xlab/cgogen/translator"
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
	}
	return gen, nil
}

type declList []*tl.CDecl

func (s declList) Len() int      { return len(s) }
func (s declList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s declList) Less(i, j int) bool {
	if s[i].Pos != s[j].Pos {
		return s[i].Pos < s[j].Pos
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
			gen.expandEnum(wr, decl)
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
	seenStructTags := make(map[string]bool)
	seenUnionTags := make(map[string]bool)
	for _, decl := range gen.tr.Typedefs() {
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
			gen.writeStructTypedef(wr, decl, memTip == tl.TipMemRaw)
		case tl.UnionKind:
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
			gen.writeTypeTypedef(wr, decl)
		case tl.FunctionKind:
			gen.writeFunctionTypedef(wr, decl)
		}
		writeSpace(wr, 1)
		count++
	}
	for tag, decl := range gen.tr.TagMap() {
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
				gen.writeStructTypedef(wr, decl, memTipRx.Self() == tl.TipMemRaw)
			} else {
				gen.writeStructTypedef(wr, decl, false)
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
	for _, decl := range gen.tr.Declares() {
		if decl.IsStatic {
			continue
		}
		const public = true
		switch decl.Spec.Kind() {
		case tl.StructKind, tl.OpaqueStructKind:
			if len(decl.Name) == 0 {
				continue
			} else if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			}
			gen.writeStructDeclaration(wr, decl, tl.NoTip, tl.NoTip, public)
		case tl.UnionKind:
			if len(decl.Name) == 0 {
				continue
			} else if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
				continue
			}
			gen.writeUnionDeclaration(wr, decl, tl.NoTip, tl.NoTip, public)
		case tl.EnumKind:
			if !decl.Spec.IsComplete() {
				if !gen.tr.IsAcceptableName(tl.TargetPublic, decl.Name) {
					continue
				}
				gen.writeEnumDeclaration(wr, decl, tl.NoTip, tl.NoTip, public)
			}
		case tl.FunctionKind:
			if !gen.tr.IsAcceptableName(tl.TargetFunction, decl.Name) {
				continue
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
