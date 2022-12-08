package cmd

import (
	"github.com/AbirRazzak/2022-advent-of-code/day01"
	"github.com/AbirRazzak/2022-advent-of-code/inputreader"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runner := day01.CommandRunner{
			Logger:      log.New(os.Stderr, "", log.LstdFlags),
			InputReader: &inputreader.HTTPInputReader{},
		}

		err := runner.Run(part2Flag)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	startCmd.AddCommand(day01Cmd)
}
