package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import "fmt"
import "unsafe"

func SimdVersion() string {
	return C.GoString(C.SimdVersion())
}

func SimdAlignment() int {
    return int(C.SimdAlignment())
}

func SimdAbsDifferenceSum() int64 {
    a := unsafe.Pointer(C.CBytes("aaaa"))
    defer C.free(a)
    aStride := 4
    b := unsafe.Pointer(C.CBytes("cccc"))
    defer C.free(b)
    bStride := 4
    width, height := 4, 1
    sum := C.uint64_t(0)
    C.SimdAbsDifferenceSum((*C.uint8_t)(a), C.size_t(aStride), (*C.uint8_t)(b), C.size_t(bStride), C.size_t(width), C.size_t(height), &sum)
    return int64(sum)
}

// ingroup difference_estimation
//SimdAddFeatureDifference(const uint8_t * value, size_t valueStride, size_t width, size_t height, const uint8_t * lo, size_t loStride, const uint8_t * hi, size_t hiStride,  uint16_t weight, uint8_t * difference, size_t differenceStride)

// ingroup drawing
//SimdAlphaBlending(const uint8_t * src, size_t srcStride, size_t width, size_t height, size_t channelCount, const uint8_t * alpha, size_t alphaStride, uint8_t * dst, size_t dstStride)

// ingroup integral
//SimdIntegral(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t * sum, size_t sumStride, uint8_t * sqsum, size_t sqsumStride, uint8_t * tilted, size_t tiltedStride, SimdPixelFormatType sumFormat, SimdPixelFormatType sqsumFormat)

// ingroup shifting
//SimdShiftBilinear(const uint8_t * src, size_t srcStride, size_t width, size_t height, size_t channelCount, const uint8_t * bkg, size_t bkgStride, const double * shiftX, const double * shiftY, size_t cropLeft, size_t cropTop, size_t cropRight, size_t cropBottom, uint8_t * dst, size_t dstStride)

// ingroup svm
//SimdSvmSumLinear(const float * x, const float * svs, const float * weights, size_t length, size_t count, float * sum)

func main() {
	fmt.Println("Simd version:", SimdVersion())
	fmt.Println("Alignment   :", SimdAlignment())
    fmt.Println("Crc32c      :", SimdCrc32c("aap"))
    fmt.Println("AbsDiffSum  :", SimdAbsDifferenceSum())
}
