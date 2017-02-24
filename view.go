package main

import (
	"C"
	"fmt"
	"image"
	"unsafe"
)

type Format uint8

const (
	NONE Format = iota
	GRAY8
	UV16
	BGR24
	BGRA32
	INT16
	INT32
	INT64
	FLOAT
	DOUBLE
	BAYERGRBG
	BAYERGBRG
	BAYERRGGB
	BAYERBGGR
	HSV24
	HSL24
)

func PixelSize(f Format) int {
	switch f {
	case NONE:
		return 0
	case GRAY8:
		return 1
	case UV16:
		return 2
	case BGR24:
		return 3
	case BGRA32:
		return 4
	case INT16:
		return 2
	case INT32:
		return 4
	case INT64:
		return 8
	case FLOAT:
		return 4
	case DOUBLE:
		return 8
	case BAYERGRBG:
		return 1
	case BAYERGBRG:
		return 1
	case BAYERRGGB:
		return 1
	case BAYERBGGR:
		return 1
	case HSV24:
		return 3
	case HSL24:
		return 3
	default:
		return 0
	}
}

func ChannelCount(f Format) int {
	switch f {
	case NONE:
		return 0
	case GRAY8:
		return 1
	case UV16:
		return 2
	case BGR24:
		return 3
	case BGRA32:
		return 4
	case INT16:
		return 1
	case INT32:
		return 1
	case INT64:
		return 1
	case FLOAT:
		return 1
	case DOUBLE:
		return 1
	case BAYERGRBG:
		return 1
	case BAYERGBRG:
		return 1
	case BAYERRGGB:
		return 1
	case BAYERBGGR:
		return 1
	case HSV24:
		return 3
	case HSL24:
		return 3
	default:
		return 0
	}
}

type View struct {
	width, height int
	format        Format
	stride        int
	owner         bool
	data          unsafe.Pointer
}

// Recreate
func (v *View) Recreate(w, h int, f Format) {

	if v.owner && v.data != nil {
		Free(v.data)
		v.data = nil
		v.owner = false
	}
	v.width = w
	v.height = h
	v.format = f
	v.stride = Align(v.width*PixelSize(v.format), Alignment())
	v.data = Allocate(v.height*v.stride, Alignment())
	fmt.Println(*v)
}

// Load
func (v *View) Load(path string) error {
	fmt.Println(path)

	return nil
}

func (v *View) CopyFrom(img *image.RGBA) error {

	for y := 0; y < img.Bounds().Size().Y; y++ {
		start := y*img.Bounds().Size().X*4
		psrcstart := unsafe.Pointer(uintptr(v.data)+uintptr(start))
		for x := 0; x < img.Bounds().Size().X; x++ {
			psrcdata := (*C.uint)(unsafe.Pointer(uintptr(psrcstart)+uintptr(x*4)))
			*psrcdata = C.uint(uint32(img.Pix[start+x*4+0]) << 16 +
			                   uint32(img.Pix[start+x*4+1]) << 8 +
			                   uint32(img.Pix[start+x*4+2]) << 0 +
			                   uint32(img.Pix[start+x*4+3]) << 24)
		}
	}

	return nil
}

func (v *View) CopyTo(img *image.RGBA) error {

	for y := 0; y < img.Bounds().Size().Y; y++ {
		start := y*img.Bounds().Size().X*4
		psrcstart := unsafe.Pointer(uintptr(v.data)+uintptr(start))
		for x := 0; x < img.Bounds().Size().X; x++ {
			psrcdata := (*C.uint)(unsafe.Pointer(uintptr(psrcstart)+uintptr(x*4)))
			img.Pix[start+x*4+0] = uint8((*psrcdata >> 16) & 0xff)
			img.Pix[start+x*4+1] = uint8((*psrcdata >> 8) & 0xff)
			img.Pix[start+x*4+2] = uint8((*psrcdata >> 0) & 0xff)
			img.Pix[start+x*4+3] = uint8((*psrcdata >> 24) & 0xff)
		}
	}


	return nil
}