package pxtl

import "image"

// Quantity with width and height dimensions
type dim struct {
	w, h int
}

type grid struct {
	// number of tiles horizontally and vertically
	tiles dim
	// area where the tiles are contained
	rect image.Rectangle
	// Total width of gridlines in the image. Distributed as evenly as possible.
	// There's one gridline between each row/column of tiles,
	// as well as one before the first, and one after the last.
	// For a uniform width line, set to `lineWidth*(tiles.w + 1)`
	// (assuming tile sizes also fit uniformly).
	lines dim
}
