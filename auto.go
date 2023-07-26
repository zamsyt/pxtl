package pxtl

import (
	"image"
	"image/color"
)

func DetectLines(img image.Image, tolerance uint8) (columns, rows []line) {
	edgeImg := EdgeFilter(img, 1, 0)
	b := edgeImg.Bounds()
	columns = detectLines1d(vLine(b.Min.X, b), 1, 0, edgeImg, tolerance)

	edgeImg = EdgeFilter(img, 0, 1)
	b = edgeImg.Bounds()
	rows = detectLines1d(hLine(b.Min.Y, b), 0, 1, edgeImg, tolerance)
	return
}

func detectLines1d(l0 image.Rectangle, oX, oY int, img *image.Gray, tolerance uint8) []line {
	var lines []line
	b := img.Bounds()
	o := image.Pt(oX, oY)
	var pos, w int
	for l := l0; l.In(b); l = l.Add(o) {
		// not edge
		if isBlack(img, l, tolerance) {
			w++
			continue
		}
		// edge
		if w > 0 {
			lines = append(lines, line{pos, w + 1})
		}
		pos += w + 1
		w = 0
	}
	if w > 0 {
		lines = append(lines, line{pos, w + 1})
	}
	return lines
}

type line struct {
	pos, w int
}

func tileAt(column, row line) image.Rectangle {

	return image.Rect(column.pos, row.pos,
		column.pos+column.w, row.pos+row.w)
}

func isBlack(img *image.Gray, r image.Rectangle, t uint8) bool {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if img.At(x, y).(color.Gray).Y > t {
				return false
			}
		}
	}
	return true
}

func vLine(x int, bounds image.Rectangle) image.Rectangle {
	return image.Rect(x, bounds.Min.Y, x+1, bounds.Max.Y)
}
func hLine(y int, bounds image.Rectangle) image.Rectangle {
	return image.Rect(bounds.Min.X, y, bounds.Max.X, y+1)
}
