package translator

import (
	"fmt"
	"go/token"
	"log"
	"path/filepath"

	"github.com/cznic/c/internal/xc"
)

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

func unmanagedCaseWarn(c int, p token.Pos) {
	log.Printf("unmanaged ParameterDeclaration case %d at %s\n", c, srcLocation(p))
}
