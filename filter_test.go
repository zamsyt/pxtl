package pxtl_test

import (
	"testing"

	"github.com/zamsyt/pxtl"
)

func TestFilter(t *testing.T) {
	// Visual test
	filtered := pxtl.EdgeFilter(getImg("test-images/c_x10.png"), 1, 1)
	savePng(filtered, "test-output/c_x10_filtered.png")

	filtered = pxtl.EdgeFilter(getImg("test-images/c_x10.png"), 0, 1)
	savePng(filtered, "test-output/c_x10_filtered2.png")

	filtered = pxtl.EdgeFilter(getImg("test-images/c_x10.png"), 1, 0)
	savePng(filtered, "test-output/c_x10_filtered3.png")
}
