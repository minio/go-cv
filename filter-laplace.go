package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

func Laplace(a, b View) {
	C.SimdLaplace((*C.uint8_t)(a.data), C.size_t(a.stride), C.size_t(a.width), C.size_t(a.height), (*C.uint8_t)(b.data), C.size_t(b.stride))
}

//SimdLaplaceAbs(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t * dst, size_t dstStride)
