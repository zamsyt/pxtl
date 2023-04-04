package main

import (
	"github.com/spf13/cobra"
)

// Upscale command
var upscaleCmd = &cobra.Command{
	Use:   "upscale <image_path>",
	Short: "Upscale pixel art",
	Run: func(cmd *cobra.Command, args []string) {
		inputImg := getImg(args[0])
		scaledImg := inputImg.Upscale(upscaleFactor)
		savePng(scaledImg, outPath)
	},
}

var upscaleFactor int

func init() {
	upscaleCmd.Flags().IntVarP(&upscaleFactor, "factor", "f", 0, "scaling factor")
	upscaleCmd.MarkFlagRequired("factor")

	rootCmd.AddCommand(upscaleCmd)
}
