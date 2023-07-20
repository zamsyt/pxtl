package pxtl_test

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func imagesEq(img0, img1 image.Image) bool {
	b0 := img0.Bounds()
	b1 := img1.Bounds()
	if !b0.Eq(b1) {
		return false
	}
	for y := b0.Min.Y; y < b0.Max.Y; y++ {
		for x := b0.Min.X; x < b0.Max.X; x++ {
			if !colorEq(img0.At(x, y), img1.At(x, y)) {
				return false
			}
		}
	}
	return true
}

func colorEq(a, b color.Color) bool {
	aR, aG, aB, aA := a.RGBA()
	bR, bG, bB, bA := b.RGBA()

	return (aR == bR &&
		aG == bG &&
		aB == bB &&
		aA == bA)
}

func getImg(path string) image.Image {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	check(err)
	return img
}

func savePng(img image.Image, path string) {
	f, err := os.Create(path)
	check(err)
	png.Encode(f, img)
	f.Close()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
