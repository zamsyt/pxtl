package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

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
