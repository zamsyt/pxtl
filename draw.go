package pxtl

import (
	"image"
	"image/color"
	"image/draw"
)

func fill(img draw.Image, r image.Rectangle, c color.Color) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			img.Set(x, y, c)
		}
	}
}
