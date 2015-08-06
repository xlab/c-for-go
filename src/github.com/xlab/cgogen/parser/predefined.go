package parser

import "github.com/cznic/c/internal/cc"

type TargetArch int

const (
	Arch32 TargetArch = 32
	Arch64 TargetArch = 64
)

var predefinedBase = `
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1
#define __signed signed
#define __GNUC__ 0
#define __asm__(x)
#define __inline
#define __attribute__(x)
`

var predefines = map[TargetArch]string{
	Arch32: predefinedBase + `#define __i386__ 1`,
	Arch64: predefinedBase + `#define __x86_64__ 1`,
}

var archs = map[string]TargetArch{
	"386":    Arch32,
	"arm":    Arch32,
	"armbe":  Arch32,
	"mips":   Arch32,
	"mipsle": Arch32,
	"sparc":  Arch32,
	//
	"amd64":       Arch64,
	"amd64p32":    Arch64,
	"arm64":       Arch64,
	"arm64be":     Arch64,
	"ppc64":       Arch64,
	"ppc64le":     Arch64,
	"mips64":      Arch64,
	"mips64le":    Arch64,
	"mips64p32":   Arch64,
	"mips64p32le": Arch64,
	"sparc64":     Arch64,
}

var models = map[TargetArch]cc.Model{
	Arch32: cc.Model{
		cc.Ptr:       {Size: 4, Align: 4, More: "__TODO_PTR"},
		cc.Void:      {Size: 0, Align: 1, More: "__TODO_VOID"},
		cc.Char:      {Size: 1, Align: 1, More: "int8"},
		cc.UChar:     {Size: 1, Align: 1, More: "byte"},
		cc.Short:     {Size: 2, Align: 2, More: "int16"},
		cc.UShort:    {Size: 2, Align: 2, More: "uint16"},
		cc.Int:       {Size: 4, Align: 4, More: "int32"},
		cc.UInt:      {Size: 4, Align: 4, More: "uint32"},
		cc.Long:      {Size: 4, Align: 4, More: "int32"},
		cc.ULong:     {Size: 4, Align: 4, More: "uint32"},
		cc.LongLong:  {Size: 8, Align: 8, More: "int64"},
		cc.ULongLong: {Size: 8, Align: 8, More: "uint64"},
		cc.Float:     {Size: 4, Align: 4, More: "float32"},
		cc.Double:    {Size: 8, Align: 8, More: "float64"},
		cc.Bool:      {Size: 1, Align: 1, More: "bool"},
		cc.Complex:   {Size: 8, Align: 8, More: "complex128"},
	},
	Arch64: cc.Model{
		cc.Ptr:       {Size: 8, Align: 8, More: "__TODO_PTR"},
		cc.Void:      {Size: 0, Align: 1, More: "__TODO_VOID"},
		cc.Char:      {Size: 1, Align: 1, More: "int8"},
		cc.UChar:     {Size: 1, Align: 1, More: "byte"},
		cc.Short:     {Size: 2, Align: 2, More: "int16"},
		cc.UShort:    {Size: 2, Align: 2, More: "uint16"},
		cc.Int:       {Size: 4, Align: 4, More: "int32"},
		cc.UInt:      {Size: 4, Align: 4, More: "uint32"},
		cc.Long:      {Size: 8, Align: 8, More: "int64"},
		cc.ULong:     {Size: 8, Align: 8, More: "uint64"},
		cc.LongLong:  {Size: 8, Align: 8, More: "int64"},
		cc.ULongLong: {Size: 8, Align: 8, More: "uint64"},
		cc.Float:     {Size: 4, Align: 4, More: "float32"},
		cc.Double:    {Size: 8, Align: 8, More: "float64"},
		cc.Bool:      {Size: 1, Align: 1, More: "bool"},
		cc.Complex:   {Size: 8, Align: 8, More: "complex128"},
	},
}
