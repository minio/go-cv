package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import (
	"fmt"
	"io/ioutil"
	"unsafe"
	"github.com/lazywei/go-opencv/opencv"
)

/* DecodeImageMemM decodes an image from an in memory byte buffer. */
func DecodeImageMemM(data []byte) *opencv.Mat {
	buf := opencv.CreateMatHeader(1, len(data), opencv.CV_8U)
	buf.SetData(unsafe.Pointer(&data[0]), opencv.CV_AUTOSTEP)
	defer buf.Release()

	return opencv.DecodeImageM(unsafe.Pointer(buf), opencv.CV_LOAD_IMAGE_UNCHANGED)
}

func LoadDetect() string {

    buf, err :=  ioutil.ReadFile("/home/ec2-user/work/src/github.com/lazywei/go-opencv/images/lena.jpg")
    if err != nil {
        panic("error loading from file")
    }
    pmat := DecodeImageMemM(buf)
    fmt.Println("pmat:", pmat)

    detect := DetectInitialize("/home/ec2-user/c_apps/Simd/data/cascade/haar_face_0.xml")

    return DetectObjects(pmat, detect)
}

func DetectInitialize(cascade string) unsafe.Pointer {

    var detect unsafe.Pointer
    detect = C.SimdDetectInitialize(C.CString(cascade))

    return detect
}

func DetectObjects(pmat *opencv.Mat, detect unsafe.Pointer) string {

    return C.GoString(C.SimdDetectObjects(unsafe.Pointer(pmat), detect))
}

func main() {
	fmt.Println(LoadDetect())
}
