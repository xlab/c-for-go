package utils

import (
	"fmt"
	"strings"
)

func CTypeString(name string) string {
	if strings.HasPrefix(name, "enum") {
		return "uint32"
	}
	return fmt.Sprintf("C.%s", name)
}
