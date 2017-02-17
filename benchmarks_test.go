package main

import (
	"os"
	"fmt"
	"image"
	"image/draw"
	"testing"
	"github.com/lazywei/go-opencv/opencv"
	"github.com/anthonynsimon/bild/noise"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
)

const Resolution = 2048

/////////////////////////////////////////////////////////////////////
//
// S i m d
//
/////////////////////////////////////////////////////////////////////

func SimdSetup(f Format) (View, View) {

	a, b := View{}, View{}
	a.Recreate(Resolution, Resolution, f)
	b.Recreate(Resolution, Resolution, f)

	return a, b
}

func benchmarkSimdGaussian(b *testing.B, f Format) {

	src, dst := SimdSetup(f)

	for i := 0; i < b.N; i++ {
		GaussianBlur3x3(src, dst)
	}
}

func BenchmarkSimdGaussian(b *testing.B) {

	benchmarkSimdGaussian(b, GRAY8)
}

func BenchmarkSimdGaussianRGB(b *testing.B) {

	benchmarkSimdGaussian(b, BGR24)
}

func benchmarkSimdBlur(b *testing.B, f Format) {

	src, dst := SimdSetup(f)

	for i := 0; i < b.N; i++ {
		MeanFilter3x3(src, dst)
	}
}

func BenchmarkSimdBlur(b *testing.B) {

	benchmarkSimdBlur(b, GRAY8)
}

func BenchmarkSimdBlurRGB(b *testing.B) {

	benchmarkSimdBlur(b, BGR24)
}

func benchmarkSimdMedian3x3(b *testing.B, f Format) {

	src, dst := SimdSetup(f)

	for i := 0; i < b.N; i++ {
		MedianFilterSquare3x3(src, dst)
	}
}

func BenchmarkSimdMedian3x3(b *testing.B) {

	benchmarkSimdMedian3x3(b, GRAY8)
}

func BenchmarkSimdMedian3x3RGB(b *testing.B) {

	benchmarkSimdMedian3x3(b, BGR24)
}

func benchmarkSimdMedian5x5(b *testing.B, f Format) {

	src, dst := SimdSetup(f)

	for i := 0; i < b.N; i++ {
		MedianFilterSquare5x5(src, dst)
	}
}

func BenchmarkSimdMedian5x5(b *testing.B) {

	benchmarkSimdMedian5x5(b, GRAY8)
}

func BenchmarkSimdMedian5x5RGB(b *testing.B) {

	benchmarkSimdMedian5x5(b, BGR24)
}

//func benchmarkSimdLaplace(b *testing.B, f Format) {
//
//	src, _ := SimdSetup(f)
//	dst := View{}
//	dst.Recreate(Resolution, Resolution, INT16)
//
//	for i := 0; i < b.N; i++ {
//		Laplace(src, dst)
//	}
//}
//
//func BenchmarkSimdLaplace(b *testing.B) {
//
//	benchmarkSimdLaplace(b, GRAY8)
//}
//
//func BenchmarkSimdLaplaceRGB(b *testing.B) {
//
//	benchmarkSimdLaplace(b, BGR24)
//}

func BenchmarkSimdBGRtoGray(b *testing.B) {

	src, _ := SimdSetup(BGR24)
	_, dst := SimdSetup(GRAY8)

	for i := 0; i < b.N; i++ {
		BgrToGray(src, dst)
	}
}

func BenchmarkSimdBGRtoHsl(b *testing.B) {

	src, _ := SimdSetup(BGR24)
	_, dst := SimdSetup(HSL24)

	for i := 0; i < b.N; i++ {
		BgrToHsl(src, dst)
	}
}

func BenchmarkSimdBGRtoHsv(b *testing.B) {

	src, _ := SimdSetup(BGR24)
	_, dst := SimdSetup(HSV24)

	for i := 0; i < b.N; i++ {
		BgrToHsl(src, dst)
	}
}

// AsRGBA returns an RGBA copy of the supplied image.
func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

func benchmarkSimdSobel(b *testing.B, f Format) {

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

func OpenCVSetup(channels int) (*opencv.IplImage, *opencv.IplImage) {

	a := opencv.CreateImage(Resolution, Resolution, 8, channels)
	b := opencv.CreateImage(Resolution, Resolution, 8, channels)

	return a, b
}

func benchmarkOpenCVGaussian(b *testing.B, channels int) {

	src, dst := OpenCVSetup(channels)

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_GAUSSIAN, 3, 3, 0, 0)
	}
}

func BenchmarkOpenCVGaussian(b *testing.B) {

	benchmarkOpenCVGaussian(b, 1)
}

func BenchmarkOpenCVGaussianRGB(b *testing.B) {

	benchmarkOpenCVGaussian(b, 3)
}

func benchmarkOpenCVBlur(b *testing.B, channels int) {

	src, dst := OpenCVSetup(channels)

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_BLUR, 3, 3, 0, 0)
	}
}

func BenchmarkOpenCVBlur(b *testing.B) {

	benchmarkOpenCVBlur(b, 1)
}

func BenchmarkOpenCVBlurRGB(b *testing.B) {

	benchmarkOpenCVBlur(b, 3)
}

func benchmarkOpenCVMedian3x3(b *testing.B, channels int) {

	src, dst := OpenCVSetup(channels)

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_MEDIAN, 3, 3, 0, 0)
	}
}

func BenchmarkOpenCVMedian3x3(b *testing.B) {

	benchmarkOpenCVMedian3x3(b, 1)
}

func BenchmarkOpenCVMedian3x3RGB(b *testing.B) {

	benchmarkOpenCVMedian3x3(b, 3)
}

func benchmarkOpenCVMedian5x5(b *testing.B, channels int) {

	src, dst := OpenCVSetup(channels)

	for i := 0; i < b.N; i++ {
		opencv.Smooth(src, dst, opencv.CV_MEDIAN, 5, 5, 0, 0)
	}
}

func BenchmarkOpenCVMedian5x5(b *testing.B) {

	benchmarkOpenCVMedian5x5(b, 1)
}

func BenchmarkOpenCVMedian5x5RGB(b *testing.B) {

	benchmarkOpenCVMedian5x5(b, 3)
}

func BenchmarkOpenCVBGRtoGray(b *testing.B) {

	src, _ := OpenCVSetup(3)
	_, dst := OpenCVSetup(1)

	for i := 0; i < b.N; i++ {
		opencv.CvtColor(src, dst, opencv.CV_BGR2GRAY)
	}
}

func BenchmarkOpenCVBGRtoHsv(b *testing.B) {

	src, _ := OpenCVSetup(3)
	_, dst := OpenCVSetup(3)

	for i := 0; i < b.N; i++ {
		opencv.CvtColor(src, dst, 40) // from /* CV_BGR2HSV = 40 from opencv-2.4.10/modules/imgproc/include/opencv2/imgproc/imgproc.hpp */
	}
}

func BenchmarkOpenCVBGRtoHsl(b *testing.B) {

	src, _ := OpenCVSetup(3)
	_, dst := OpenCVSetup(3)

	for i := 0; i < b.N; i++ {
		opencv.CvtColor(src, dst, 52) // from /* COLOR_BGR2HLS = 52 from opencv-2.4.10/modules/imgproc/include/opencv2/imgproc/imgproc.hpp */
	}
}

func benchmarkOpenCVSobel(b *testing.B, channels int) {

	src, dst := OpenCVSetup(channels)

	for i := 0; i < b.N; i++ {
		opencv.Sobel(src, dst, 1, 0, 3)
	}
}

func BenchmarkOpenCVSobel(b *testing.B) {

	benchmarkOpenCVSobel(b, 1)
}

func BenchmarkOpenCVSobelRGB(b *testing.B) {

	benchmarkOpenCVSobel(b, 3)
}

/////////////////////////////////////////////////////////////////////
//
// B i l d
//
/////////////////////////////////////////////////////////////////////

func BenchmarkBildBlurRGBA(b *testing.B) {

	img := noise.Generate(Resolution, Resolution, &noise.Options{Monochrome: true, NoiseFn: noise.Gaussian})

	for i := 0; i < b.N; i++ {
		blur.Box(img, 1.0)
	}
}
