package day03

import (
	"github.com/AbirRazzak/2022-advent-of-code/inputreader"
	"log"
	"strings"
)

type CommandRunner struct {
	Logger      *log.Logger
	InputReader inputreader.InputReader
}

func (r CommandRunner) Run(part2Flag bool) error {
	r.Logger.Println("Welcome to day 03!")
	r.Logger.Println("Running part 1...")

	input, err := r.InputReader.GetInputForDay(3)
	if err != nil {
		return err
	}

	coll := NewRuckSackCollection()
	for _, line := range strings.Split(input, "\n") {
		coll.AddRuckSack(NewRuckSack(line))
	}

	sum, err := coll.SumPriorities()
	if err != nil {
		return err
	}

	r.Logger.Println("SumPriorities = ", sum)

	return nil
}
