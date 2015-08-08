package vpx

const (
	// CodecABIVersion as defined in vpx/vpx_codec.h:86
	CodecABIVersion = 6
	// CodecCapDecoder as defined in vpx/vpx_codec.h:154
	CodecCapDecoder = 0x1
	// CodecCapEncoder as defined in vpx/vpx_codec.h:155
	CodecCapEncoder = 0x2
	// ImageABIVersion as defined in vpx/vpx_image.h:31
	ImageABIVersion = 3
	// ImageFormatPlanar as defined in vpx/vpx_image.h:34
	ImageFormatPlanar = 0x100
	// ImageFormatUvFlip as defined in vpx/vpx_image.h:35
	ImageFormatUvFlip = 0x200
	// ImageFormatHasAlpha as defined in vpx/vpx_image.h:36
	ImageFormatHasAlpha = 0x400
	// ImageFormatHighbitdepth as defined in vpx/vpx_image.h:37
	ImageFormatHighbitdepth = 0x800
	// PlanePacked as defined in vpx/vpx_image.h:100
	PlanePacked = 0
	// PlaneY as defined in vpx/vpx_image.h:101
	PlaneY = 0
	// PlaneU as defined in vpx/vpx_image.h:102
	PlaneU = 1
	// PlaneV as defined in vpx/vpx_image.h:103
	PlaneV = 2
	// PlaneAlpha as defined in vpx/vpx_image.h:104
	PlaneAlpha = 3
)

// RefFrameType as declared in vpx/vp8.h:99
type RefFrameType uint32

// RefFrameType enumeration from vpx/vp8.h:99
const (
	LastFrame RefFrameType = 1
	GoldFrame              = 2
	AltrFrame              = 4
)

// CodecErr as declared in vpx/vpx_codec.h:142
type CodecErr uint32

// CodecErr enumeration from vpx/vpx_codec.h:142
const (
	CodecOk             CodecErr = 0
	CodecError                   = 1
	CodecMemError                = 2
	CodecABIMismatch             = 3
	CodecIncapable               = 4
	CodecUnsupBitstream          = 5
	CodecUnsupFeature            = 6
	CodecCorruptFrame            = 7
	CodecInvalidParam            = 8
	CodecListEnd                 = 9
)

// BitDepth as declared in vpx/vpx_codec.h:223
type BitDepth uint32

// BitDepth enumeration from vpx/vpx_codec.h:223
const (
	Bits8  BitDepth = 8
	Bits10          = 10
	Bits12          = 12
)

// ImageFormat as declared in vpx/vpx_image.h:67
type ImageFormat uint32

// ImageFormat enumeration from vpx/vpx_image.h:67
const (
	ImageFormatNone     ImageFormat = 0
	ImageFormatRgb24                = 1
	ImageFormatRgb32                = 2
	ImageFormatRgb565               = 3
	ImageFormatRgb555               = 4
	ImageFormatUyvy                 = 5
	ImageFormatYuy2                 = 6
	ImageFormatYvyu                 = 7
	ImageFormatBgr24                = 8
	ImageFormatRgb32Le              = 9
	ImageFormatArgb                 = 10
	ImageFormatArgbLe               = 11
	ImageFormatRgb565Le             = 12
	ImageFormatRgb555Le             = 13
	ImageFormatYv12                 = 0x100 | 0x200 | 1
	ImageFormatI420                 = 0x100 | 2
	ImageFormatVpxyv12              = 0x100 | 0x200 | 3
	ImageFormatVpxi420              = 0x100 | 4
	ImageFormatI422                 = 0x100 | 5
	ImageFormatI444                 = 0x100 | 6
	ImageFormatI440                 = 0x100 | 7
	ImageFormat444a                 = 0x100 | 0x400 | 6
	ImageFormatI42016               = ImageFormatI420 | 0x800
	ImageFormatI42216               = ImageFormatI422 | 0x800
	ImageFormatI44416               = ImageFormatI444 | 0x800
	ImageFormatI44016               = ImageFormatI440 | 0x800
)

// ColorSpace as declared in vpx/vpx_image.h:79
type ColorSpace uint32

// ColorSpace enumeration from vpx/vpx_image.h:79
const (
	ColorSpaceUnknown  ColorSpace = 0
	ColorSpaceBt601               = 1
	ColorSpaceBt709               = 2
	ColorSpaceSmpte170            = 3
	ColorSpaceSmpte240            = 4
	ColorSpaceBt2020              = 5
	ColorSpaceReserved            = 6
	ColorSpaceSrgb                = 7
)
