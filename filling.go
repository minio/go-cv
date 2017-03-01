package gocv

// #cgo pkg-config: Simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup filling

func Fill(dst View, val uint8) {
	C.SimdFill((*C.uint8_t)(dst.data), C.size_t(dst.stride), C.size_t(dst.width), C.size_t(dst.height), C.size_t(PixelSize(dst.format)), (C.uint8_t)(val))
}

//SimdFillFrame(uint8_t * dst, size_t stride, size_t width, size_t height, size_t pixelSize, size_t frameLeft, size_t frameTop, size_t frameRight, size_t frameBottom, uint8_t value)
//SimdFillBgr(uint8_t * dst, size_t stride, size_t width, size_t height, uint8_t blue, uint8_t green, uint8_t red)
//SimdFillBgra(uint8_t * dst, size_t stride, size_t width, size_t height, uint8_t blue, uint8_t green, uint8_t red, uint8_t alpha)
