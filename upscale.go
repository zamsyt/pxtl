package pxtl

import (
	"image"
)

func (img *Image) Upscale(s int) Image {
	b := img.Bounds()
	newImg := newRGBA(image.Rect(s*b.Min.X, s*b.Min.Y, s*b.Max.X, s*b.Max.Y))
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			newImg.Fill(image.Rect(s*x, s*y, s*x+s, s*y+s), img.At(x, y))
		}
	}
	return NewImage(newImg)
}
