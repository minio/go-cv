package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup sobel_filter

func SobelDx(src, dst View) {
	C.SimdSobelDx((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func SobelDxAbs(src, dst View) {
	C.SimdSobelDxAbs((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func SobelDxAbsSum(src View) uint64 {
	sum := C.uint64_t(0)
	C.SimdSobelDxAbsSum((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint64_t)(&sum))
	return uint64(sum)
}

func SobelDy(src, dst View) {
	C.SimdSobelDy((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func SobelDyAbs(src, dst View) {
	C.SimdSobelDyAbs((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func SobelDyAbsSum(src View) uint64 {
	sum := C.uint64_t(0)
	C.SimdSobelDyAbsSum((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint64_t)(&sum))
	return uint64(sum)
}
