package gocv

// #cgo pkg-config: Simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"
import (
	"unsafe"
)

// ingroup object_detection

func DetectionLoadA(path string) unsafe.Pointer {

	return C.SimdDetectionLoadA(C.CString(path))
}

func DetectionInfo(data unsafe.Pointer) (w, h, f int) {
	width, height := C.size_t(0), C.size_t(0)
	flags := C.SimdDetectionInfoFlags(0)
	C.SimdDetectionInfo(data, &width, &height, &flags)
	return int(width), int(height), int(flags)
}

func DetectionInit(data unsafe.Pointer, sum, sqsum, tilted View, throughColumn, int16 int) unsafe.Pointer {

	return C.SimdDetectionInit(data, (*C.uint8_t)(sum.data), C.size_t(sum.stride), C.size_t(sum.width), C.size_t(sum.height), (*C.uint8_t)(sqsum.data), C.size_t(sqsum.stride), (*C.uint8_t)(tilted.data), C.size_t(tilted.stride), C.int(throughColumn), C.int(int16))
}

func DetectionPrepare(hid unsafe.Pointer) {

	C.SimdDetectionPrepare(hid)
}

func DetectionHaarDetect32fp(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionHaarDetect32fp(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionHaarDetect32fi(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionHaarDetect32fi(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect32fp(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect32fp(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect32fi(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect32fi(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect16ip(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect16ip(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect16ii(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect16ii(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionFree(data unsafe.Pointer) {

	C.SimdDetectionFree(data)
}
