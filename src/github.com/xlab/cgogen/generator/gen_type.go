package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func writeType(wr io.Writer, name []byte, spec tl.GoTypeSpec) {
	fmt.Fprintf(wr, "type %s %s\n", name, spec)
}
