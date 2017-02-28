package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup integral
func Integral(src, sum, sqsum, tilted View) {
	C.SimdIntegral((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(sum.data), C.size_t(sum.stride), (*C.uint8_t)(sqsum.data), C.size_t(sqsum.stride), (*C.uint8_t)(tilted.data), C.size_t(tilted.stride), C.SimdPixelFormatType(sum.format), C.SimdPixelFormatType(sqsum.format))
}
