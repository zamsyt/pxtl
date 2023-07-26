package pxtl_test

import (
	"fmt"
	"testing"

	"github.com/zamsyt/pxtl"
)

func TestAuto(t *testing.T) {
	img := getImg("test-images/c_x10.png")
	cols, rows := pxtl.DetectLines(img, 0)
	fmt.Printf("Detected %v columns, %v rows\n", len(cols), len(rows))
	fmt.Println(cols)
	fmt.Println(rows)

	img = getImg("test-images/23-1_ss-cropped.png")
	cols, rows = pxtl.DetectLines(img, 30)
	fmt.Printf("Detected %v columns, %v rows\n", len(cols), len(rows))

	img = getImg("test-images/23-2_ss.png")
	cols, rows = pxtl.DetectLines(img, 40)
	fmt.Printf("Detected %v columns, %v rows\n", len(cols), len(rows))
}
