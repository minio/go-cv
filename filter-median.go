package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup median_filter

func MeanFilter3x3(src, dst View) {
	C.SimdMeanFilter3x3((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(ChannelCount(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func MedianFilterRhomb3x3(src, dst View) {
	C.SimdMedianFilterRhomb3x3((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(ChannelCount(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func MedianFilterRhomb5x5(src, dst View) {
	C.SimdMedianFilterRhomb5x5((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(ChannelCount(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func MedianFilterSquare3x3(src, dst View) {
	C.SimdMedianFilterSquare3x3((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(ChannelCount(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func MedianFilterSquare5x5(src, dst View) {
	C.SimdMedianFilterSquare5x5((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(ChannelCount(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}
