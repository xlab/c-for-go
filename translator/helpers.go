package translator

import (
	"bytes"
	"fmt"
	"go/token"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/cznic/c/xc"
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

func SrcLocation(p token.Pos) string {
	pos := xc.FileSet.Position(p)
	return fmt.Sprintf("%s:%d", narrowPath(pos.Filename), pos.Line)
}

func unresolvedIdentifierWarn(name string, p token.Pos) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("WARN: %s:%d unresolved identifier %s at %s\n", narrowPath(file), line, name, SrcLocation(p))
}

func unmanagedCaseWarn(c int, p token.Pos) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("WARN: %s:%d unmanaged case %d at %s\n", narrowPath(file), line, c, SrcLocation(p))
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

func getVarArrayCount(arraySizes []uint32) (count uint8) {
	for i := range arraySizes {
		if arraySizes[i] == 0 {
			count++
		}
	}
	return
}

func tagAnonymousMembers(decl CDecl) {
	if decl.Kind() != StructKind {
		return
	}
	structSpec := decl.Spec.(*CStructSpec)
	for i := range structSpec.Members {
		switch spec := structSpec.Members[i].Spec.(type) {
		case *CStructSpec:
			if len(spec.Tag) == 0 {
				spec.Tag = fmt.Sprintf("%s_%s", decl.Name, structSpec.Members[i].Name)
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
