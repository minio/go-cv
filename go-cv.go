package gocv

// #cgo pkg-config: Simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
//
// #cgo LDFLAGS: -lstdc++
import "C"

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/jpeg"
	"io/ioutil"
	"unsafe"
)

func DecodeImageMem(buf []byte) (*image.RGBA, error) {
	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	// Extract pixel information.
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)
	return rgba, nil
}

func LoadDetect() string {
	buf, err := ioutil.ReadFile("images/lena.jpg")
	if err != nil {
		panic(err)
	}

	img, err := DecodeImageMem(buf)
	if err != nil {
		panic(err)
	}

	detect := DetectInitialize("cascade/haar_face_0.xml")
	return DetectObjects(img, detect)
}

func DetectObjects(img *image.RGBA, detect unsafe.Pointer) string {
	return C.GoString(C.SimdDetectObjectsRaw(C.int(img.Rect.Max.X),
		C.int(img.Rect.Max.Y),
		C.int(img.Stride),
		unsafe.Pointer(&img.Pix[0]),
		detect))
}

func DetectInitialize(cascade string) unsafe.Pointer {
	return C.SimdDetectInitialize(C.CString(cascade))
}
