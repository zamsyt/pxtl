package pxtl_test

import (
	"os"
	"testing"

	"github.com/zamsyt/pxtl"
)

func init() {
	os.MkdirAll("test-output", 0755)
}

func TestUpscale(t *testing.T) {
	want := getImg("test-images/c_x10.png")
	upscaled := pxtl.Upscale(getImg("test-images/c-nn.png"), 10)
	if !imagesEq(want, upscaled) {
		path := "test-output/c_test-upscale-x10.png"
		t.Errorf("Upscale failed. Saving output at %v", path)
		savePng(upscaled, path)
		return
	}
}
