package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pxtl",
	Short: "SHORT DESCRIPTION",
	Long:  `PXTL - pixel art tools`,
}

var OutPath string

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&OutPath, "out", "output.png", "image output path")

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
