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

// ingroup other_filter
func GaussianBlur3x3(src, dst View) {
	C.SimdGaussianBlur3x3((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(ChannelCount(src.format)), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

//SimdAbsGradientSaturatedSum(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t * dst, size_t dstStride)
//SimdLbpEstimate(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t * dst, size_t dstStride)

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
