package main

import (
	"fmt"
	"io/ioutil"

	"github.com/minio/go-cv"
)

func main() {

	buf, err := ioutil.ReadFile("./data/images/lena.jpg")
	if err != nil {
		panic("error loading from file")
	}
	rgba, err := gocv.DecodeImageMem(buf)
	if err != nil {
		panic(err)
	}

	detect := gocv.DetectInitialize("./data/cascades/haar_face_0.xml")

	json := gocv.DetectObjects(rgba, detect)

	fmt.Println(json)
}
