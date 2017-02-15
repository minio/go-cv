package main

import (
	"testing"
	"github.com/lazywei/go-opencv/opencv"

)

const Resolution = 2048

func SimdSetup() (View, View) {

	a, b := View{}, View{}
	a.Recreate(Resolution, Resolution, GRAY8)
	b.Recreate(Resolution, Resolution, GRAY8)

	return a, b
}

/////////////////////////////////////////////////////////////////////
//
// S i m d
//
/////////////////////////////////////////////////////////////////////

func BenchmarkSimdGaussian(b *testing.B) {

	src, dst := SimdSetup()

	for i := 0; i < b.N; i++ {
		GaussianBlur3x3(src, dst)
	}
}

func BenchmarkSimdBlur(b *testing.B) {

	src, dst := SimdSetup()

	for i := 0; i < b.N; i++ {
		MeanFilter3x3(src, dst)
	}
}

func BenchmarkSimdMedian3x3(b *testing.B) {

	src, dst := SimdSetup()

	for i := 0; i < b.N; i++ {
		MedianFilterSquare3x3(src, dst)
	}
}

func BenchmarkSimdMedian5x5(b *testing.B) {

	src, dst := SimdSetup()

	for i := 0; i < b.N; i++ {
		MedianFilterSquare5x5(src, dst)
	}
}

func BenchmarkSimdSobel(b *testing.B) {

	src, _ := SimdSetup()
	dst := View{}
	dst.Recreate(Resolution, Resolution, INT16)

	for i := 0; i < b.N; i++ {
		SobelDx(src, dst)
		SobelDy(src, dst)
	}
}

/////////////////////////////////////////////////////////////////////
//
// O p e n C V
//
/////////////////////////////////////////////////////////////////////

func OpenCVSetup() (*opencv.IplImage, *opencv.IplImage) {

	a := opencv.CreateImage(Resolution, Resolution, 1, 1)
	b := opencv.CreateImage(Resolution, Resolution, 1, 1)

	return a, b
}

func BenchmarkOpenCVGaussian(b *testing.B) {

	src, dst := OpenCVSetup()

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_GAUSSIAN, 3, 3, 0, 0)
	}
}

func BenchmarkOpenCVBlur(b *testing.B) {

	src, dst := OpenCVSetup()

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_BLUR, 3, 3, 0, 0)
	}
}

func BenchmarkOpenCVMedian3x3(b *testing.B) {

	src, dst := OpenCVSetup()

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_MEDIAN, 3, 3, 0, 0)
	}
}

func BenchmarkOpenCVMedian5x5(b *testing.B) {

	src, dst := OpenCVSetup()

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_MEDIAN, 5, 5, 0, 0)
	}
}

func BenchmarkOpenCVSobel(b *testing.B) {

	src, dst := OpenCVSetup()

	for i := 0; i < b.N; i++ {
		opencv.Sobel(src, dst, 1, 1, 3)
	}
}
