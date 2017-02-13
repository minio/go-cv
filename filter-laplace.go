package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

func Laplace(src, dst View) {
	C.SimdLaplace((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func LaplaceAbs(src, dst View) {

	C.SimdLaplaceAbs((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}
