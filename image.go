package pxtl

import (
	"image"
	"image/color"
)

// AKA CompositeImage?
type Image struct {
	Layers []image.Image
}

func NewImage(img image.Image) Image {
	return Image{[]image.Image{img}}
}

func (img *Image) Flat() image.Image /* image.RGBA */ {
	return img.Layers[0] // TODO: Add actual implementation
}

// Implement image.Image
func (img Image) Bounds() image.Rectangle { return img.Flat().Bounds() }
func (img Image) At(x, y int) color.Color { return img.Flat().At(x, y) }
func (img Image) ColorModel() color.Model { return img.Flat().ColorModel() }

type RGBA struct {
	image.RGBA
}

func newRGBA(r image.Rectangle) *RGBA {
	return &RGBA{*image.NewRGBA(r)}
}

func (img *RGBA) Fill(r image.Rectangle, c color.Color) {
	b := img.Bounds()
	if !r.In(b) {
		panic("Rectangle not In image")
	}
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			img.Set(x, y, c)
		}
	}
}
