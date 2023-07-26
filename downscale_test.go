package pxtl_test

import (
	"os"
	"testing"

	"github.com/zamsyt/pxtl"
)

func TestDownscale(t *testing.T) {
	want := getImg("test-images/c-nn.png")
	downscaled := pxtl.AutoDownscale(getImg("test-images/c_x10.png"), 0)
	if !imagesEq(want, downscaled) {
		path := "test-output/c_x10_test-downscale.png"
		t.Errorf("Downscale failed. Saving output at %v", path)
		os.MkdirAll("test-output", 0755)
		savePng(downscaled, path)
		return
	}

	// Visual
	img := getImg("test-images/23-1_ss-cropped.png")
	downscaled = pxtl.AutoDownscale(img, 30)
	savePng(downscaled, "test-output/23-1_ss-cropped_downscaled.png")

	img = getImg("test-images/23-2_ss.png")
	downscaled = pxtl.AutoDownscale(img, 30)
	savePng(downscaled, "test-output/23-2_ss_downscaled.png")
}

/*
mario

no white border
13, 16

outer
14, 19

full pixels
12, 17
*/
