package cmd

import (
	"github.com/AbirRazzak/2022-advent-of-code/src/day02"
	"github.com/AbirRazzak/2022-advent-of-code/src/inputreader"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runner := day02.CommandRunner{
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
	startCmd.AddCommand(day02Cmd)
}
