package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zamsyt/pxtl/pxtl"
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

	var upscaleCmd = &cobra.Command{
		Use:   "upscale <image_path>",
		Short: "Upscale some pixel art",
		//Long:  "",
		//Args: cobra.MinimumNArgs(1),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			savePng(pxtl.Upscale(ScaleFactor, getImg(args[0])), OutPath)
		},
	}

	upscaleCmd.Flags().IntVarP(&ScaleFactor, "factor", "f", 1, "scaling factor")
	upscaleCmd.MarkFlagRequired("factor")
	upscaleCmd.Flags().StringVar(&OutPath, "out", "output.png", "where to save scaled image")

	var cmds = [](*cobra.Command){
		{
			Use:   "version",
			Short: "Print version information",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("pxtl dev")
			},
		},
		upscaleCmd,
	}

	rootCmd.AddCommand(cmds...)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
