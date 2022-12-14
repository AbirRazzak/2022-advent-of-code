package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var part2Flag bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().BoolVarP(&part2Flag, "part2", "2", false, "Run part 2 of the puzzle instead of only part 1")
}

func GetDefaultStartCommandLogger() *log.Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}
