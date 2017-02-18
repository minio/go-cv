package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"
import "unsafe"

// ingroup object_detection

func DetectionLoadA(path string) unsafe.Pointer {

	return C.SimdDetectionLoadA(C.CString(path))
}

//SimdDetectionInfo(const void * data, size_t * width, size_t * height, SimdDetectionInfoFlags * flags)

func DetectionInit(data, sum, sqsum, tilted View, throughColumn, int16 int) unsafe.Pointer {

	return C.SimdDetectionInit(data.data, (*C.uint8_t)(sum.data), C.size_t(sum.stride), C.size_t(sum.width), C.size_t(sum.height), (*C.uint8_t)(sqsum.data), C.size_t(sqsum.stride), (*C.uint8_t)(tilted.data), C.size_t(tilted.stride), C.int(throughColumn), C.int(int16))
}

//SimdDetectionPrepare(void * hid)
//SimdDetectionHaarDetect32fp(const void * hid, const uint8_t * mask, size_t maskStride, ptrdiff_t left, ptrdiff_t top, ptrdiff_t right, ptrdiff_t bottom, uint8_t * dst, size_t dstStride)
//SimdDetectionHaarDetect32fi(const void * hid, const uint8_t * mask, size_t maskStride, ptrdiff_t left, ptrdiff_t top, ptrdiff_t right, ptrdiff_t bottom, uint8_t * dst, size_t dstStride)
//SimdDetectionLbpDetect32fp(const void * hid, const uint8_t * mask, size_t maskStride, ptrdiff_t left, ptrdiff_t top, ptrdiff_t right, ptrdiff_t bottom, uint8_t * dst, size_t dstStride)
//SimdDetectionLbpDetect32fi(const void * hid, const uint8_t * mask, size_t maskStride, ptrdiff_t left, ptrdiff_t top, ptrdiff_t right, ptrdiff_t bottom, uint8_t * dst, size_t dstStride)
//SimdDetectionLbpDetect16ip(const void * hid, const uint8_t * mask, size_t maskStride, ptrdiff_t left, ptrdiff_t top, ptrdiff_t right, ptrdiff_t bottom, uint8_t * dst, size_t dstStride)
//SimdDetectionLbpDetect16ii(const void * hid, const uint8_t * mask, size_t maskStride, ptrdiff_t left, ptrdiff_t top, ptrdiff_t right, ptrdiff_t bottom, uint8_t * dst, size_t dstStride)

func DetectionFree(data unsafe.Pointer) {

	C.SimdDetectionFree(data)
}
