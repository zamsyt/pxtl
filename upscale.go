package pxtl

import "image"

func upscaleRect(r image.Rectangle, s int) image.Rectangle {
	return image.Rect(r.Min.X*s, r.Min.Y*s, r.Max.X*s, r.Max.Y*s)
}

func Upscale(img image.Image, s int) image.Image {
	b := img.Bounds()
	upscaledImg := newDrawImageMatch(img, upscaleRect(b, s))
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			fill(upscaledImg, image.Rect(s*x, s*y, s*x+s, s*y+s), img.At(x, y))
		}
	}
	return upscaledImg
}
