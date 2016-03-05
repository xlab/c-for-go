package translator

import (
	"bytes"
	"fmt"
	"go/token"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/cznic/xc"
)

var (
	qualConst       = []byte("const")
	specStruct      = []byte("struct")
	specUnion       = []byte("union")
	specUnsigned    = []byte("unsigned")
	specSigned      = []byte("signed")
	specLong        = []byte("long")
	specShort       = []byte("short")
	ptrStr          = []byte("*")
	sliceStr        = []byte("[]")
	spaceStr        = []byte(" ")
	skipStr         = []byte("_")
	emptyStr        = []byte{}
	restrictedNames = bytes.Join([][]byte{
		qualConst, specStruct, specUnion, specUnsigned, specSigned, specShort,
	}, spaceStr)
)

type bytesSlice [][]byte

func (s bytesSlice) Len() int           { return len(s) }
func (s bytesSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bytesSlice) Less(i, j int) bool { return bytes.Compare(s[i], s[j]) < 0 }

func bytesJoin(buf1, buf2 []byte, sep string) []byte {
	if len(buf1) == 0 {
		return buf2
	}
	if len(buf2) == 0 {
		return buf1
	}
	// both are not empty
	return bytes.Join([][]byte{buf1, buf2}, []byte(sep))
}

func bytesWrap(buf []byte, w1, w2 string) []byte {
	if len(buf) == 0 {
		return nil
	}
	res := make([]byte, 0, len(buf)+2)
	res = append(res, []byte(w1)...)
	res = append(res, buf...)
	res = append(res, []byte(w2)...)
	return res
}

func isRestrictedBase(b []byte) bool {
	return bytes.Contains(restrictedNames, b)
}

// narrowPath reduces full path to file name and parent dir only.
func narrowPath(fp string) string {
	if !filepath.IsAbs(fp) {
		if abs, err := filepath.Abs(fp); err != nil {
			// seems to be reduced already
			return fp
		} else {
			fp = abs
		}
	}
	return filepath.Join(filepath.Base(filepath.Dir(fp)), filepath.Base(fp))
}

func alterBytesPart(buf []byte, idx []int, altFn func([]byte) []byte) []byte {
	copy(buf[idx[0]:idx[1]], altFn(buf[idx[0]:idx[1]]))
	return buf // for chaining

	// a copying version:
	//
	// altered := make([]byte, len(buf))
	// copy(altered[:idx[0]], buf[:idx[0]])
	// copy(altered[idx[0]:idx[1]], altFn(buf[idx[0]:idx[1]]))
	// copy(altered[idx[1]:], buf[idx[1]:])
	// return altered
}

func replaceBytes(buf []byte, idx []int, piece []byte) []byte {
	a, b := idx[0], idx[1]
	altered := make([]byte, len(buf)-(b-a)+len(piece))
	copy(altered[:a], buf[:a])
	pLen := copy(altered[a:], piece)
	copy(altered[a+pLen:], buf[b:])
	return altered
}

var srcReferenceRx = regexp.MustCompile(`(?P<path>[^;]+);(?P<file>[^;]+);(?P<line>[^;]+);(?P<name>[^;]+);`)

func (t *Translator) SrcLocation(docTarget RuleTarget, name string, p token.Pos) string {
	pos := xc.FileSet.Position(p)
	filename := filepath.Base(pos.Filename)
	defaultLocation := func() string {
		return fmt.Sprintf("%s:%d", narrowPath(pos.Filename), pos.Line)
	}
	rxs, ok := t.compiledRxs[ActionDocument][docTarget]
	if !ok {
		global, ok1 := t.compiledRxs[ActionDocument][TargetGlobal]
		post, ok2 := t.compiledRxs[ActionDocument][TargetPostGlobal]
		// other targets won't be supported
		switch {
		case ok1:
			rxs = global
		case ok2:
			rxs = post
		default:
			return defaultLocation()
		}
	}

	var template string
	for i := range rxs {
		if rxs[i].From.MatchString(name) {
			template = string(rxs[i].To)
			break
		}
	}
	if len(template) == 0 {
		return defaultLocation()
	}

	values := fmt.Sprintf("%s;%s;%d;%s;", narrowPath(pos.Filename), filename, pos.Line, name)
	// indices := srcReferenceRx.FindAllStringSubmatchIndex(values, -1)
	location := srcReferenceRx.ReplaceAllString(values, template)
	// for i := len(indices); i > 0; i-- {
	// 	idx := indices[i-1]
	// 	location := srcReferenceRx.ExpandString(nil, template, src, idx)
	// }
	return location
}

func incVal(v Value) Value {
	switch v := v.(type) {
	case int32:
		return v + 1
	case int64:
		return v + 1
	case uint32:
		return v + 1
	case uint64:
		return v + 1
	default:
		return v
	}
}

type TypeCache struct {
	mux   sync.RWMutex
	cache map[string]struct{}
}

func (t *TypeCache) Get(id string) bool {
	t.mux.RLock()
	defer t.mux.RUnlock()
	_, ok := t.cache[id]
	if ok {
		return true
	}
	return false
}

func (t *TypeCache) Set(id string) {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.cache == nil {
		t.cache = make(map[string]struct{})
	}
	t.cache[id] = struct{}{}
}

func (t *TypeCache) Delete(id string) {
	t.mux.Lock()
	defer t.mux.Unlock()
	delete(t.cache, id)
}

type CachedNameTransform struct {
	Target     RuleTarget
	Visibility RuleTarget
	Name       string
}

type NameTransformCache struct {
	mux   sync.RWMutex
	cache map[CachedNameTransform][]byte
}

func (n *NameTransformCache) Get(target, visibility RuleTarget, name string) ([]byte, bool) {
	n.mux.RLock()
	entry := CachedNameTransform{
		Target:     target,
		Visibility: visibility,
		Name:       name,
	}
	if cached, ok := n.cache[entry]; ok {
		n.mux.RUnlock()
		return cached, true
	}
	n.mux.RUnlock()
	return nil, false
}

func (n *NameTransformCache) Set(target, visibility RuleTarget, name string, result []byte) {
	n.mux.Lock()
	if n.cache == nil {
		n.cache = make(map[CachedNameTransform][]byte)
	}
	entry := CachedNameTransform{
		Target:     target,
		Visibility: visibility,
		Name:       name,
	}
	n.cache[entry] = result
	n.mux.Unlock()
}

type TipCache struct {
	mux   sync.RWMutex
	cache map[struct {
		TipScope
		string
	}]TipSpecRx
}

func (n *TipCache) Get(scope TipScope, name string) (TipSpecRx, bool) {
	n.mux.RLock()
	if cached, ok := n.cache[struct {
		TipScope
		string
	}{scope, name}]; ok {
		n.mux.RUnlock()
		return cached, true
	}
	n.mux.RUnlock()
	return TipSpecRx{}, false
}

func (n *TipCache) Set(scope TipScope, name string, result TipSpecRx) {
	n.mux.Lock()
	if n.cache == nil {
		n.cache = make(map[struct {
			TipScope
			string
		}]TipSpecRx)
	}
	n.cache[struct {
		TipScope
		string
	}{scope, name}] = result
	n.mux.Unlock()
}

// func getVarArrayCount(arraySizes []uint6432) (count uint8) {
// 	for i := range arraySizes {
// 		if arraySizes[i] == 0 {
// 			count++
// 		}
// 	}
// 	return
// }

func tagAnonymousMembers(decl CDecl) {
	if decl.Spec.Kind() != StructKind {
		return
	}
	structSpec := decl.Spec.(*CStructSpec)
	for _, m := range structSpec.Members {
		switch spec := m.Spec.(type) {
		case *CStructSpec:
			if len(spec.Tag) == 0 {
				spec.Tag = fmt.Sprintf("%s_%s", decl.Name, m.Name)
			}
		}
	}
}

var v = struct{}{}

var builtinNames = map[string]struct{}{
	"break": v, "default": v, "func": v, "interface": v, "select": v,
	"case": v, "defer": v, "go": v, "map": v, "struct": v,
	"chan": v, "else": v, "goto": v, "package": v, "switch": v,
	"const": v, "fallthrough": v, "if": v, "range": v, "type": v,
	"continue": v, "for": v, "import": v, "return": v, "var": v,
}

// blessName transforms the name to be a valid name in Go and not a keyword.
func blessName(name []byte) string {
	if _, ok := builtinNames[string(name)]; ok {
		return "_" + string(name)
	}
	return string(name)
}

func isBuiltinName(name []byte) bool {
	if _, ok := builtinNames[string(name)]; ok {
		return true
	}
	return false
}

var nameRewrites = map[string][]byte{
	"type": []byte("kind"),
}

func rewriteName(name []byte) []byte {
	if n, ok := nameRewrites[string(name)]; ok {
		return n
	}
	// e.g. type -> _type
	return append(skipStr, name...)
}

// readNumeric checks if the value looks like a numeric value, i.e. -1000.f
func isNumeric(v []rune) bool {
	for _, r := range v {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true
		case '-', '+', '.':
			continue
		default:
			return false
		}
	}
	return false
}

// readNumeric reads numeric part of the value, ex: -1000.f being read as -1000.
func readNumeric(v []rune) string {
	result := make([]rune, 0, len(v))
	for _, r := range v {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result = append(result, r)
		case '-', '+', '.', 'x', 'b':
			result = append(result, r)
		case 'A', 'B', 'C', 'D', 'E', 'F':
			result = append(result, r)
			// c'mon, get some roman numerals here
		default:
			return string(result)
		}
	}
	return string(result)
}
