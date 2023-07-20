package pxtl

import (
	"image"
	"image/draw"
	"log"
)

var defaultNewDrawImageFunc = func(r image.Rectangle) draw.Image { return image.NewNRGBA(r) }

// Attempt to create new image matching the type of original.
func newDrawImageMatch(original image.Image, r image.Rectangle) draw.Image {
	switch v := original.(type) {
	case *image.NRGBA:
		return image.NewNRGBA(r)
	case *image.RGBA:
		return image.NewRGBA(r)
	case *image.Paletted:
		return image.NewPaletted(r, v.Palette)
	case draw.Image:
		log.Printf("Unknown draw.Image type %T", v)
	}
	return defaultNewDrawImageFunc(r)
}
