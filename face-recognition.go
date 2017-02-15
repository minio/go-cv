package main

// ingroup face_recognition

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

func HogDirectionHistograms(src View) float64 {
	cellX, cellY := 0, 0
	quantization := 0
	histograms := C.float(0.0)
	C.SimdHogDirectionHistograms((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), C.size_t(cellX), C.size_t(cellY), C.size_t(quantization), &histograms)
	return float64(histograms)
}
