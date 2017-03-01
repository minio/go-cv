package main

import (
	"fmt"
	"github.com/fwessels/go-cv"
	"io/ioutil"
)

func main() {

	buf, err := ioutil.ReadFile("./data/images/lena.jpg")
	if err != nil {
		panic("error loading from file")
	}
	pmat := gocv.DecodeImageMemM(buf)

	detect := gocv.DetectInitialize("./data/cascades/haar_face_0.xml")

	json := gocv.DetectObjects(pmat, detect)

	fmt.Println(json)
}
