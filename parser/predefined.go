package parser

import "github.com/cznic/cc"

type TargetArch string

const (
	Arch32    TargetArch = "i386"
	Arch48    TargetArch = "x86_48"
	Arch64    TargetArch = "x86_64"
	ArchArm7A TargetArch = "armv7a"
	ArchArm8A TargetArch = "armv8a"
)

var predefinedBase = `
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1
#define __GNUC__ 4
#define __builtin_va_list void *
#define __asm__(x)
#define __asm(x)
#define __inline inline
#define __inline__ inline
#define __signed signed
#define __signed__ signed
#define __extension__
#define __attribute__(x)
#define __POSIX_C_DEPRECATED(ver)

void __GO__(char*, ...);
`

var predefines = map[TargetArch]string{
	Arch32:    predefinedBase + `#define __i386__ 1`,
	Arch48:    predefinedBase + `#define __x86_64__ 1`,
	Arch64:    predefinedBase + `#define __x86_64__ 1`,
	ArchArm7A: predefinedBase + `#define __ARM_ARCH_7A__ 1` + "\n" + `#define __arm__ 1`,
	ArchArm8A: predefinedBase + `#define __ARM_ARCH_8A__ 1` + "\n" + `#define __aarch64__ 1`,
	// TODO(xlab): https://sourceforge.net/p/predef/wiki/Architectures/
}

var models = map[TargetArch]*cc.Model{
	Arch32:    model32,
	Arch48:    model48,
	Arch64:    model64,
	ArchArm7A: model32,
	ArchArm8A: model64,
}

var arches = map[string]TargetArch{
	"386":         Arch32,
	"arm":         ArchArm7A,
	"armv7a":      ArchArm7A,
	"armv8a":      ArchArm8A,
	"armeabi-v7a": ArchArm7A,
	"armeabi-v8a": ArchArm8A,
	"armbe":       ArchArm7A,
	"mips":        Arch32,
	"mipsle":      Arch32,
	"sparc":       Arch32,
	"amd64":       Arch64,
	"amd64p32":    ArchArm7A,
	"arm64":       ArchArm8A,
	"arm64be":     ArchArm8A,
	"ppc64":       Arch64,
	"ppc64le":     Arch64,
	"mips64":      Arch64,
	"mips64le":    Arch64,
	"mips64p32":   Arch48,
	"mips64p32le": Arch48,
	"sparc64":     Arch64,
}

var model32 = &cc.Model{
	Items: map[cc.Kind]cc.ModelItem{
		cc.Ptr:               {4, 4, 4, "__TODO_PTR"},
		cc.UintPtr:           {4, 4, 4, "__TODO_UINTPTR"},
		cc.Void:              {0, 1, 1, "__TODO_VOID"},
		cc.Char:              {1, 1, 1, "int8"},
		cc.SChar:             {1, 1, 1, "int8"},
		cc.UChar:             {1, 1, 1, "byte"},
		cc.Short:             {2, 2, 2, "int16"},
		cc.UShort:            {2, 2, 2, "uint16"},
		cc.Int:               {4, 4, 4, "int32"},
		cc.UInt:              {4, 4, 4, "uint32"},
		cc.Long:              {4, 4, 4, "int32"},
		cc.ULong:             {4, 4, 4, "uint32"},
		cc.LongLong:          {8, 8, 8, "int64"},
		cc.ULongLong:         {8, 8, 8, "uint64"},
		cc.Float:             {4, 4, 4, "float32"},
		cc.Double:            {8, 8, 8, "float64"},
		cc.LongDouble:        {8, 8, 8, "float64"},
		cc.Bool:              {1, 1, 1, "bool"},
		cc.FloatComplex:      {8, 8, 8, "complex64"},
		cc.DoubleComplex:     {16, 16, 16, "complex128"},
		cc.LongDoubleComplex: {16, 16, 16, "complex128"},
	},
}

var model48 = &cc.Model{
	Items: map[cc.Kind]cc.ModelItem{
		cc.Ptr:               {4, 4, 4, "__TODO_PTR"},
		cc.UintPtr:           {4, 4, 4, "__TODO_UINTPTR"},
		cc.Void:              {0, 1, 1, "__TODO_VOID"},
		cc.Char:              {1, 1, 1, "int8"},
		cc.SChar:             {1, 1, 1, "int8"},
		cc.UChar:             {1, 1, 1, "byte"},
		cc.Short:             {2, 2, 2, "int16"},
		cc.UShort:            {2, 2, 2, "uint16"},
		cc.Int:               {4, 4, 4, "int32"},
		cc.UInt:              {4, 4, 4, "uint32"},
		cc.Long:              {8, 8, 8, "int64"},
		cc.ULong:             {8, 8, 8, "uint64"},
		cc.LongLong:          {8, 8, 8, "int64"},
		cc.ULongLong:         {8, 8, 8, "uint64"},
		cc.Float:             {4, 4, 4, "float32"},
		cc.Double:            {8, 8, 8, "float64"},
		cc.LongDouble:        {8, 8, 8, "float64"},
		cc.Bool:              {1, 1, 1, "bool"},
		cc.FloatComplex:      {8, 8, 8, "complex64"},
		cc.DoubleComplex:     {16, 16, 16, "complex128"},
		cc.LongDoubleComplex: {16, 16, 16, "complex128"},
	},
}

var model64 = &cc.Model{
	Items: map[cc.Kind]cc.ModelItem{
		cc.Ptr:               {8, 8, 8, "__TODO_PTR"},
		cc.UintPtr:           {8, 8, 8, "__TODO_UINTPTR"},
		cc.Void:              {0, 1, 1, "__TODO_VOID"},
		cc.Char:              {1, 1, 1, "int8"},
		cc.SChar:             {1, 1, 1, "int8"},
		cc.UChar:             {1, 1, 1, "byte"},
		cc.Short:             {2, 2, 2, "int16"},
		cc.UShort:            {2, 2, 2, "uint16"},
		cc.Int:               {4, 4, 4, "int32"},
		cc.UInt:              {4, 4, 4, "uint32"},
		cc.Long:              {8, 8, 8, "int64"},
		cc.ULong:             {8, 8, 8, "uint64"},
		cc.LongLong:          {8, 8, 8, "int64"},
		cc.ULongLong:         {8, 8, 8, "uint64"},
		cc.Float:             {4, 4, 4, "float32"},
		cc.Double:            {8, 8, 8, "float64"},
		cc.LongDouble:        {8, 8, 8, "float64"},
		cc.Bool:              {1, 1, 1, "bool"},
		cc.FloatComplex:      {8, 8, 8, "complex64"},
		cc.DoubleComplex:     {16, 16, 16, "complex128"},
		cc.LongDoubleComplex: {16, 16, 16, "complex128"},
	},
}
