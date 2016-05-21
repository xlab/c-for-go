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
	location := srcReferenceRx.ReplaceAllString(values, template)
	return location
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
	long := false
	unsigned := false
	result := make([]rune, 0, len(v))

	wrap := func() string {
		switch {
		case unsigned && long:
			return "uint64(" + string(result) + ")"
		case unsigned:
			return "uint32(" + string(result) + ")"
		case long:
			return "int64(" + string(result) + ")"
		default:
			return string(result)
		}
	}

	for _, r := range v {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result = append(result, r)
		case '-', '+', '.', 'x':
			result = append(result, r)
		case 'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f':
			result = append(result, r)
		// c'mon, get some roman numerals here
		case 'L', 'l':
			long = true
			continue
		case 'U', 'u':
			unsigned = true
			continue
		default:
			return wrap()
		}
	}
	return wrap()
}
