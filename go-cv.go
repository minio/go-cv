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

func main() {
	fmt.Println("Simd version:", SimdVersion())
	fmt.Println("Alignment   :", SimdAlignment())
    fmt.Println("Crc32c      :", SimdCrc32c("aap"))
    fmt.Println("AbsDiffSum  :", SimdAbsDifferenceSum())
}
