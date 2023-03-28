package pxtl

import (
	"image"
	"image/draw"
)

func Upscale(s int, img image.Image) image.Image {
	b := img.Bounds()

	newImg := image.NewRGBA(
		image.Rect(s*b.Min.X, s*b.Min.Y, s*b.Max.X, s*b.Max.Y),
	)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		sY := s * y
		for i := 0; i < s; i++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				sX := s * x
				for j := 0; j < s; j++ {
					newImg.Set(sX+j, sY+i, img.At(x, y))
				}
			}
		}
	}

	return newImg
}

func Crop(img image.Image, x1, y1, x2, y2 int) image.Image {
	cropped := image.NewRGBA(image.Rect(x1, y1, x2+1, y2+1))
	draw.Draw(cropped, cropped.Bounds(), img, image.Point{x1, y1}, draw.Src)
	return cropped
}
