package main

import (
	"fmt"
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
		fmt.Println("downscale called")
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
	downscaleCmd.Flags().Var(&sampleOffset, "sample-offset", "offset from the corner of tile where to sample its color (default: factor/2)")

	rootCmd.AddCommand(downscaleCmd)
}
