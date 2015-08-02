package translator

import (
	"bytes"
	"fmt"
	"go/token"
	"path/filepath"
	"runtime"

	"github.com/cznic/c/internal/xc"
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
	emptyStr        = []byte{}
	restrictedNames = bytes.Join([][]byte{
		qualConst, specStruct, specUnion, specUnsigned, specSigned, specShort,
	}, spaceStr)
)

type Evaler interface {
	Eval() []byte
}

func evalAndJoin(ex1, ex2 Evaler, sep string) []byte {
	buf1 := ex1.Eval()
	if len(buf1) == 0 {
		return ex2.Eval()
	}
	buf2 := ex2.Eval()
	if len(buf2) == 0 {
		return buf1
	}
	// both are not empty
	return bytes.Join([][]byte{buf1, buf2}, []byte(sep))
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
	altered := make([]byte, 2*len(buf)+len(piece)-a-b)
	copy(altered[:a], buf[:a])
	pLen := copy(altered[a:], piece)
	copy(altered[a+pLen:], buf[b:])
	return altered
}

func srcLocation(p token.Pos) string {
	pos := xc.FileSet.Position(p)
	return fmt.Sprintf("%s:%d", narrowPath(pos.Filename), pos.Line)
}

func unresolvedIdentifierWarn(name string, p token.Pos) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("WARN: %s:%d unresolved identifier %s at %s\n", narrowPath(file), line, name, srcLocation(p))
}

func unmanagedCaseWarn(c int, p token.Pos) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("WARN: %s:%d unmanaged case %d at %s\n", narrowPath(file), line, c, srcLocation(p))
}
