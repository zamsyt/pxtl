package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zamsyt/pxtl"
)

type coords []int

func (cs *coords) String() string {
	if *cs == nil {
		return ""
	}
	return fmt.Sprint(*cs)
}

func (cs *coords) Set(s string) error {
	strs := strings.Split(s, " ")
	for _, v := range strs {
		num, err := strconv.Atoi(v)
		check(err)
		*cs = append(*cs, num)
	}
	return nil
}

func (cs *coords) Type() string { return "numbers" }

var cropCmd = &cobra.Command{
	Use:   "crop <image_path>",
	Short: "Crop image",
	Run: func(cmd *cobra.Command, args []string) {
		savePng(pxtl.Crop(getImg(args[0]), bounds[0], bounds[1], bounds[2], bounds[3]), OutPath)
	},
}

var bounds coords

func init() {
	cropCmd.Flags().VarP(&bounds, "bounds", "b", "The values to use for cropping")
}
