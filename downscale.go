package pxtl

import (
	"image"
	"log"
)

// o: sample offset
func (img *Image) downscale(s int, o image.Point) Image {
	b := img.Bounds()
	nb := image.Rect(b.Min.X/s, b.Min.Y/s, b.Max.X/s, b.Max.Y/s)
	newImg := newRGBA(nb)
	for y := nb.Min.Y; y < nb.Max.Y; y++ {
		for x := nb.Min.X; x < nb.Max.X; x++ {
			newImg.Set(x, y, img.At(x*s+o.X, y*s+o.Y))
		}
	}
	return NewImage(newImg)
}

func (img *Image) Downscale(factor, offset int) Image {
	if offset >= factor {
		log.Fatal("Offset must be smaller than factor")
	} else if offset < 0 {
		offset = factor / 2
	}
	return img.downscale(factor, image.Pt(offset, offset))
}
