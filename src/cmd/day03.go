package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Advent Of Code Day 03",
	Long:  `https://adventofcode.com/2022/day/3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day03 called")
	},
}

func init() {
	startCmd.AddCommand(day03Cmd)
}
