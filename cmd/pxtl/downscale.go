package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zamsyt/pxtl"
)

// Downscale command
var downscaleCmd = &cobra.Command{
	Use:   "downscale <image_path>",
	Short: "Downscale pixel art to 1:1 scale",
	Run: func(cmd *cobra.Command, args []string) {
		inPath := args[0]
		inputImg := getImg(inPath)
		scaledImg := pxtl.AutoDownscale(inputImg, tolerance)
		outPath := strings.TrimSuffix(filepath.Base(inPath), filepath.Ext(inPath)) + "_auto-downscale.png"
		b := scaledImg.Bounds()
		savePng(scaledImg, outPath)
		fmt.Printf("Downscaled size: %vx%v (%v)\n", b.Dx(), b.Dy(), outPath)
	},
}

var tolerance uint8

func init() {
	downscaleCmd.Flags().Uint8Var(&tolerance, "tolerance", 5, "tolerance for tile detection (0-255)")

	rootCmd.AddCommand(downscaleCmd)
}
