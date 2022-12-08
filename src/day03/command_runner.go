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

	r.Logger.Println("Getting input...")

	input, err := r.InputReader.GetInputForDay(3)
	if err != nil {
		return err
	}

	r.Logger.Println("Input retrieved.")
	r.Logger.Println("Running part 1...")

	coll := NewRuckSackCollection()
	for _, line := range strings.Split(input, "\n") {
		err := coll.AddRuckSack(NewRuckSack(line))
		if err != nil {
			return err
		}
	}

	sum, err := coll.SumPriorities()
	if err != nil {
		return err
	}

	r.Logger.Println("Part 1 = ", sum)

	if part2Flag == false {
		return nil
	}

	r.Logger.Println("Running part 2...")
	factory := NewRuckSackGroupFactory()
	coll2 := NewRuckSackCollection()

	for _, line := range strings.Split(input, "\n") {
		group := factory.ProcessRuckSack(NewRuckSack(line))
		if group != nil {
			err := coll2.AddRuckSack(group)
			if err != nil {
				return err
			}
		}
	}

	sum2, err := coll2.SumPriorities()
	if err != nil {
		return err
	}

	r.Logger.Println("Part 2 = ", sum2)

	return nil
}
