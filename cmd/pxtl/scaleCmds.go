package main

import (
	"github.com/spf13/cobra"
	"github.com/zamsyt/pxtl"
)

var scaleFactor int

var upscaleCmd = &cobra.Command{
	Use:   "upscale <image_path>",
	Short: "Upscale pixel art",
	//Long:  "",
	//Args: cobra.MinimumNArgs(1),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		savePng(pxtl.Upscale(scaleFactor, getImg(args[0])), OutPath)
	},
}

var sampleOffset int
var downscaleCmd = &cobra.Command{
	Use:   "downscale <image_path>",
	Short: "Downscale image",
	Long:  `Downscale image. By default, attempts to revert upscaled pixel art to 1:1 scale.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		savePng(pxtl.Downscale(scaleFactor, sampleOffset, getImg(args[0])), OutPath)
	},
}

func init() {
	upscaleCmd.Flags().IntVarP(&scaleFactor, "factor", "f", 0, "scaling factor")
	upscaleCmd.MarkFlagRequired("factor")

	downscaleCmd.Flags().IntVarP(&scaleFactor, "factor", "f", 0, "scaling factor")
	downscaleCmd.MarkFlagRequired("factor")
	downscaleCmd.Flags().IntVar(&sampleOffset, "sample-offset", -1, "offset within tile where to pick color (default: factor/2)") // FIXME: hide "(default -1)"
}
