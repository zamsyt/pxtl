package pxtl

import (
	"image"
	"image/draw"
	"log"
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

func Downscale(factor, sampleOffset int, img image.Image) image.Image {
	b := img.Bounds()

	if factor <= 1 {
		log.Fatalf("Invalid downscale factor %v", factor)
	}
	if sampleOffset < 0 {
		sampleOffset = factor / 2
	}
	if sampleOffset >= factor {
		log.Fatalf("sampleOffset must be less than downscale factor (%v >= %v)", sampleOffset, factor)
	}

	newImg := image.NewRGBA(
		image.Rect(b.Min.X/factor, b.Min.Y/factor, b.Max.X/factor, b.Max.Y/factor),
	)

	nB := newImg.Bounds()

	for y := nB.Min.Y; y < nB.Max.Y; y++ {
		for x := nB.Min.X; x < nB.Max.X; x++ {
			newImg.Set(x, y, img.At(factor*x+sampleOffset, factor*y+sampleOffset))
		}
	}

	return newImg
}

func Crop(img image.Image, x1, y1, x2, y2 int) image.Image {
	cropped := image.NewRGBA(image.Rect(x1, y1, x2+1, y2+1))
	draw.Draw(cropped, cropped.Bounds(), img, image.Point{x1, y1}, draw.Src)
	return cropped
}
