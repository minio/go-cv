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

func SimdCrc32c(src string) int {
    p := unsafe.Pointer(C.CBytes(src))
    defer C.free(p)
    return int(C.SimdCrc32c(p, C.size_t(len(src))))
}

func main() {
	fmt.Println("Simd version:", SimdVersion())
	fmt.Println("Alignment   :", SimdAlignment())
    fmt.Println("Crc32c      :", SimdCrc32c("aap"))
}
