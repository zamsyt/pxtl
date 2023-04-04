package main

import (
	"image"
	"image/png"
	"os"

	"github.com/zamsyt/pxtl"
)

func getImg(path string) pxtl.Image {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	check(err)
	return pxtl.NewImage(img)
}

func savePng(img pxtl.Image, path string) {
	f, err := os.Create(path)
	check(err)
	png.Encode(f, img)
	f.Close()
}
