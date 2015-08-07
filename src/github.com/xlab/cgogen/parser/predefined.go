package parser

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
