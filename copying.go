package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup copying

func Copy(src, dst View) {

	C.SimdCopy((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(PixelSize(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

//SimdCopyFrame(const uint8_t * src, size_t srcStride, size_t width, size_t height, size_t pixelSize, size_t frameLeft, size_t frameTop, size_t frameRight, size_t frameBottom, uint8_t * dst, size_t dstStride)
//SimdDeinterleaveUv(const uint8_t * uv, size_t uvStride, size_t width, size_t height, uint8_t * u, size_t uStride, uint8_t * v, size_t vStride)
