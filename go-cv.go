package main

// #cgo pkg-config: simd
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import "fmt"

func SimdVersion() string {
	return C.GoString(C.SimdVersion())
}

func main() {
	fmt.Println("Simd version:", SimdVersion())
}
