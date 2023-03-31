package pxtl

import (
	"image"
	"image/color"
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

func diff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func max(nums ...uint32) uint32 {
	var largest uint32 = 0
	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func simpleColorDiff(a, b color.Color) uint32 {
	aR, aG, aB, aA := a.RGBA()
	bR, bG, bB, bA := b.RGBA()
	return max(diff(aR, bR), diff(aG, bG), diff(aB, bB), diff(aA, bA))
}

func colorMatch(a, b color.Color, tolerance float32) bool {
	t := uint32(tolerance * 0xffff)
	// TODO: consider smarter algorithm
	return simpleColorDiff(a, b) <= t
}

func pxsMatch(img image.Image, r image.Rectangle, c color.Color, tolerance float32) bool {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if !colorMatch(c, img.At(x, y), tolerance) {
				return false
			}
		}
	}
	return true
}

func vLine(x, y0, y1 int) image.Rectangle {
	return image.Rect(x, y0, x+1, y1)
}
func hLine(y, x0, x1 int) image.Rectangle {
	return image.Rect(x0, y, x1, y+1)
}

type direction int

const (
	up direction = iota
	right
	down
	left
	topLeft
	topRight
	bottomRight
	bottomLeft
)

func matchExpand(d direction, tile image.Rectangle, img image.Image, c color.Color, tolerance float32) (image.Rectangle, bool) {
	if d > 3 {
		panic("Invalid direction")
	}
	var edge image.Rectangle
	switch d {
	case up:
		edge = hLine(tile.Min.Y-1, tile.Min.X, tile.Max.X)
	case right:
		edge = vLine(tile.Max.X, tile.Min.Y, tile.Max.Y)
	case down:
		edge = hLine(tile.Max.Y, tile.Min.X, tile.Max.X)
	case left:
		edge = vLine(tile.Min.X-1, tile.Min.Y, tile.Max.Y)
	}
	if edge.In(img.Bounds()) && pxsMatch(img, edge, c, tolerance) {
		return tile.Union(edge), true
	}
	return tile, false
}

// d must be a diagonal direction
func matchExpandSquare(d direction, tile image.Rectangle, img image.Image, c color.Color, tolerance float32) (image.Rectangle, bool) {
	var newTile image.Rectangle
	var changed bool
	switch d {
	case topLeft, topRight:
		newTile, changed = matchExpand(up, tile, img, c, tolerance)
	case bottomLeft, bottomRight:
		newTile, changed = matchExpand(down, tile, img, c, tolerance)
	default:
		panic("Invalid direction")
	}
	if !changed {
		return tile, false
	}
	switch d {
	case topLeft, bottomLeft:
		newTile, changed = matchExpand(left, newTile, img, c, tolerance)
	case topRight, bottomRight:
		newTile, changed = matchExpand(right, newTile, img, c, tolerance)
	}
	if !changed {
		return tile, false
	}
	return newTile, true
}

func findTileAt(x, y int, img image.Image, tolerance float32) image.Rectangle {
	tile := image.Rect(x, y, x+1, y+1)
	c := img.At(x, y)
	for d := topLeft; d <= bottomLeft; d++ {
		changed := true
		for changed {
			tile, changed = matchExpandSquare(d, tile, img, c, tolerance)
		}
	}
	return tile
}

func DetectFactor(img image.Image, tolerance float32) int {
	b := img.Bounds()
	maxW := findTileAt(0, 0, img, tolerance).Dx()
MAXW:
	for maxW > 1 {
		//fmt.Println("---")
		co := maxW / 2 // center offset
		for i := 0; i < b.Dy()/maxW; i++ {
			for j := 0; j < b.Dx()/maxW; j++ {
				tile := findTileAt(j*maxW+co, i*maxW+co, img, tolerance)
				tileW := tile.Dx()
				if tileW < maxW {
					maxW = tileW
					continue MAXW
				}
				//fmt.Println(j, i, tile)
			}
		}
		//log.Fatal("MAXW:", maxW)
		return maxW
	}
	return 0
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
