package pxtl

import (
	"image"
	"image/color"
)

func AutoDownscale(img image.Image, tolerance uint8) image.Image {
	cols, rows := DetectLines(img, tolerance)
	downscaled := newDrawImageMatch(img, image.Rect(0, 0, len(cols), len(rows)))
	for y, r := range rows {
		for x, c := range cols {
			tile := tileAt(c, r)
			//log.Println(x, y, c, r, tile)
			downscaled.Set(x, y, sample(img, tile))
		}
	}
	return downscaled
}

func sample(img image.Image, tile image.Rectangle) color.Color {
	centerX := tile.Min.X + tile.Dx()/2
	centerY := tile.Min.Y + tile.Dy()/2
	return img.At(centerX, centerY)
}
