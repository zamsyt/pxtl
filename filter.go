package pxtl

import (
	"image"
	"image/color"
)

func EdgeFilter(img image.Image, oX, oY int) *image.Gray {
	o := image.Pt(oX, oY)
	b := img.Bounds()
	nb := b.Intersect(b.Sub(o)) // b.Max=10, o=1 -> nb.Max=9
	edgeImg := image.NewGray(nb)
	for y := nb.Min.Y; y < nb.Max.Y; y++ {
		for x := nb.Min.X; x < nb.Max.X; x++ {
			d := simpleColorDiff(img.At(x, y), img.At(x+o.X, y+o.Y))
			edgeImg.Set(x, y, color.Gray{d})
		}
	}
	return edgeImg
}

func simpleColorDiff(a, b color.Color) uint8 {
	na := color.NRGBAModel.Convert(a).(color.NRGBA)
	nb := color.NRGBAModel.Convert(b).(color.NRGBA)
	return max(diff(na.R, nb.R), diff(na.G, nb.G), diff(na.B, nb.B), diff(na.A, nb.A))
}

type ordered interface {
	uint8 | uint32 | int
}

func diff[T ordered](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}
func max[T ordered](nums ...T) T {
	var largest T = 0
	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest
}
