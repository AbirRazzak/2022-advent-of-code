package day02

import (
	"github.com/AbirRazzak/2022-advent-of-code/inputreader"
	"log"
	"strings"
)

type CommandRunner struct {
	Logger      *log.Logger
	InputReader inputreader.InputReader
}

func (r *CommandRunner) Run(part2Flag bool) error {
	r.Logger.Println("Welcome to day 02!")

	r.Logger.Println("Running part 1...")
	if part2Flag {
		r.Logger.Println("Also running part 2...")
	}

	input, err := r.InputReader.GetInputForDay(2)
	if err != nil {
		return err
	}

	input = strings.Trim(input, "\n") // remove the trailing new line at the end of the input string

	scoreKeeper := NewScoreKeeper()
	scoreKeeperPart2 := NewScoreKeeper()
	rounds := strings.Split(input, "\n")

	for _, round := range rounds {
		components := strings.Split(round, "")
		column1 := components[0]
		column2 := components[2]

		err := scoreKeeper.Play(column1, column2)
		if err != nil {
			return err
		}

		if part2Flag {
			err = scoreKeeperPart2.PlayPart2(column1, column2)
			if err != nil {
				return err
			}
		}
	}

	r.Logger.Println("Total score for part 1 is: ", scoreKeeper.CurrentScore)

	if part2Flag {
		r.Logger.Println("Total score for part 2 is: ", scoreKeeperPart2.CurrentScore)
	}
	return nil
}
