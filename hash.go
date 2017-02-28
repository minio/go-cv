package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import "unsafe"

// ingroup hash
//SimdCrc32c(const void * src, size_t size)

func Crc32c(src string) int {
	p := unsafe.Pointer(C.CBytes(src))
	defer C.free(p)
	return int(C.SimdCrc32c(p, C.size_t(len(src))))
}
