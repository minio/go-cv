package gocv

// #cgo pkg-config: Simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"
import "unsafe"

// Allocate allocates an aligned memory block.
// The memory allocated by this function is must be deleted by function Free
// [in] size - a size of memory block.
// [in] align - a required alignment of memory block.
// return a pointer to allocated memory.
func Allocate(size, align int) unsafe.Pointer {
	return C.SimdAllocate(C.size_t(size), C.size_t(align))
}

// Free frees aligned memory block.
// This function frees a memory allocated by function Allocate.
// [in] ptr - a pointer to the memory to be deleted.
func Free(ptr unsafe.Pointer) {
	C.SimdFree(ptr)
}

// Align gets aligned size.
// [in] size - an original size.
// [in] align - a required alignment.
// return an aligned size.
func Align(size, align int) int {
	return int(C.SimdAlign(C.size_t(size), C.size_t(align)))
}

// Alignment gets alignment required for the most productive work of the Simd Library.
// a required alignment.
func Alignment() int {
	return int(C.SimdAlignment())
}
