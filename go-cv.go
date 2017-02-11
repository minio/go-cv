package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import "fmt"

func Version() string {
	return C.GoString(C.SimdVersion())
}

func main() {
	fmt.Println("Simd version:", Version())
	fmt.Println("Alignment   :", Alignment())
	fmt.Println("Crc32c      :", Crc32c("aap"))
	fmt.Println("AbsDiffSum  :", AbsDifferenceSum(View{}, View{}))
}
