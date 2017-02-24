package main

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import (
	"os"
	"fmt"
	"time"
	"image"
	"github.com/anthonynsimon/bild/imgio"
)

func Version() string {
	return C.GoString(C.SimdVersion())
}

func LoadImage() {

	reader, err := os.Open("/Users/frankw/c_apps/opencv-3.2.0/samples/data/lena.jpg")
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	img := AsRGBA(m)

	src := View{}
	src.Recreate(img.Bounds().Size().X, img.Bounds().Size().Y, BGRA32)
	src.CopyFrom(img)

	dst := View{}
	dst.Recreate(src.width, src.height, src.format)

	Copy(src, dst)

	dst.CopyTo(img)
	if err := imgio.Save("./filename", img, imgio.JPEG); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Simd version:", Version())
	fmt.Println("Alignment   :", Alignment())
	fmt.Println("Crc32c      :", Crc32c("aap"))
	fmt.Println("AbsDiffSum  :", AbsDifferenceSum(View{}, View{}))
}
