package day01

import (
	"github.com/AbirRazzak/2022-advent-of-code/src/inputreader"
	"log"
	"strings"
)

type CommandRunner struct {
	Logger *log.Logger
}

func (r *CommandRunner) Run(part2 bool) error {
	r.Logger.Println("Welcome to day 01!")

	inputReader := inputreader.InputReader{}

	input, err := inputReader.GetInputForDay(1)
	if err != nil {
		return err
	}

	r.Logger.Println("Running part 1...")

	var elves []*Elf

	caloriesPerPerson := strings.Split(input, "\n\n")
	for _, c := range caloriesPerPerson {
		items := strings.Split(c, "\n")
		elf := NewElf(items)

		elves = append(elves, elf)
	}

	group := ElfGroup{Elves: elves}

	max := group.FindMaxCalorieElf(false)

	r.Logger.Println("Maximum: ", max.Calories)

	if part2 == false {
		return nil
	}

	r.Logger.Println("Running part 2...")

	max3 := group.FindTop3CalorieElves()
	max3Total := max3[0].Calories + max3[1].Calories + max3[2].Calories

	r.Logger.Println("Top 3 Total: ", max3Total)

	return nil
}
