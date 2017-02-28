package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// AbsDifferenceSum gets sum of absolute difference of two gray 8-bit images.
// Both images must have the same width and height.
// [in] a - first image.
// [in] b - second image.
// [out] sum - the result sum of absolute difference of two images.
func AbsDifferenceSum(a, b View) uint64 {
	sum := C.uint64_t(0)
	C.SimdAbsDifferenceSum((*C.uint8_t)(a.data), C.size_t(a.stride), (*C.uint8_t)(b.data), C.size_t(b.stride), C.size_t(a.width), C.size_t(a.height), &sum)
	return uint64(sum)
}

//SimdAbsDifferenceSumMasked(const uint8_t * a, size_t aStride, const uint8_t * b, size_t bStride, const uint8_t * mask, size_t maskStride, uint8_t index, size_t width, size_t height, uint64_t * sum)
//SimdAbsDifferenceSums3x3(const uint8_t * current, size_t currentStride, const uint8_t * background, size_t backgroundStride, size_t width, size_t height, uint64_t * sums)
//SimdAbsDifferenceSums3x3Masked(const uint8_t *current, size_t currentStride, const uint8_t *background, size_t backgroundStride, const uint8_t *mask, size_t maskStride, uint8_t index, size_t width, size_t height, uint64_t * sums)
//SimdSquaredDifferenceSum(const uint8_t * a, size_t aStride, const uint8_t * b, size_t bStride, size_t width, size_t height, uint64_t * sum)
//SimdSquaredDifferenceSumMasked(const uint8_t * a, size_t aStride, const uint8_t * b, size_t bStride, const uint8_t * mask, size_t maskStride, uint8_t index, size_t width, size_t height, uint64_t * sum)
//SimdSquaredDifferenceSum32f(const float * a, const float * b, size_t size, float * sum)
//SimdSquaredDifferenceKahanSum32f(const float * a, const float * b, size_t size, float * sum)
