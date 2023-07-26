package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:     "pxtl",
	Short:   "pixel art tools",
	Long:    `PXTL - Pixel art tools`,
	Version: "0.1.0",
}

func Execute() {
	check(rootCmd.Execute())
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
