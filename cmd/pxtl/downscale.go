package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

// TODO: Downscale command
var downscaleCmd = &cobra.Command{
	Use:   "downscale <image_path>",
	Short: "Downscale image",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		// Read image (arg 0)
		// Get downscale arguments (/flags)
		// - scale
		// - global offset
		//fmt.Println("downscale called")
		inputImg := getImg(args[0])
		scaledImg := inputImg.Downscale(scaleFactor, int(sampleOffset))
		savePng(scaledImg, outPath)
	},
}

// Custom integer type
type intValue int

// Implement pflag.Value
func (p *intValue) Type() string { return "int" }
func (p *intValue) String() string {
	if *p < 0 {
		return ""
	}
	return strconv.Itoa(int(*p))
}
func (p *intValue) Set(s string) error {
	n, err := strconv.Atoi(s)
	*p = intValue(n)
	return err
}

var sampleOffset intValue = -1

func init() {
	downscaleCmd.Flags().IntVarP(&scaleFactor, "factor", "f", 0, "scaling factor")
	downscaleCmd.Flags().Var(&sampleOffset, "sample-offset", "offset from the corner of tile where to sample its color (default: factor/2)")

	rootCmd.AddCommand(downscaleCmd)
}
