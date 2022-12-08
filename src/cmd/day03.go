package cmd

import (
	"github.com/AbirRazzak/2022-advent-of-code/day03"
	"github.com/AbirRazzak/2022-advent-of-code/inputreader"
	"github.com/spf13/cobra"
)

var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Advent Of Code Day 03",
	Long:  `https://adventofcode.com/2022/day/3`,
	Run: func(cmd *cobra.Command, args []string) {
		err := day03.CommandRunner{
			Logger:      GetDefaultStartCommandLogger(),
			InputReader: &inputreader.HTTPInputReader{},
		}.Run(part2Flag)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	startCmd.AddCommand(day03Cmd)
}
