package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zamsyt/pxtl"
)

var rootCmd = &cobra.Command{
	Use:   "pxtl",
	Short: "SHORT DESCRIPTION",
	Long:  `PXTL - pixel art tools`,
}

var ScaleFactor int
var OutPath string

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&OutPath, "out", "output.png", "image output path")

	var upscaleCmd = &cobra.Command{
		Use:   "upscale <image_path>",
		Short: "Upscale pixel art",
		//Long:  "",
		//Args: cobra.MinimumNArgs(1),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			savePng(pxtl.Upscale(ScaleFactor, getImg(args[0])), OutPath)
		},
	}

	upscaleCmd.Flags().IntVarP(&ScaleFactor, "factor", "f", 0, "scaling factor")
	upscaleCmd.MarkFlagRequired("factor")

	var sampleOffset int
	var downscaleCmd = &cobra.Command{
		Use:   "downscale <image_path>",
		Short: "Downscale image",
		Long:  `Downscale image. By default, attempts to revert upscaled pixel art to 1:1 scale.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			savePng(pxtl.Downscale(ScaleFactor, sampleOffset, getImg(args[0])), OutPath)
		},
	}
	downscaleCmd.Flags().IntVarP(&ScaleFactor, "factor", "f", 0, "scaling factor")
	downscaleCmd.MarkFlagRequired("factor")
	downscaleCmd.Flags().IntVar(&sampleOffset, "sample-offset", -1, "offset within tile where to pick color (default: factor/2)") // FIXME: hide "(default -1)"

	var cmds = [](*cobra.Command){
		{
			Use:   "version",
			Short: "Print version information",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("pxtl dev")
			},
		},
		upscaleCmd,
		downscaleCmd,
		cropCmd,
	}

	rootCmd.AddCommand(cmds...)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
