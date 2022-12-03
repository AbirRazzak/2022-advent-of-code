package day01

import "strconv"

type Elf struct {
	Calories int
}

func NewElf(food []string) *Elf {
	totalCalories := 0

	for _, c := range food {
		// convert the string to an int
		current, _ := strconv.Atoi(c)
		totalCalories += current
	}

	return &Elf{
		Calories: totalCalories,
	}
}

type ElfGroup struct {
	Elves []*Elf
}

func (g *ElfGroup) FindMaxCalorieElf(removeElf bool) *Elf {
	var max int
	var maxElf *Elf
	var maxElfIndex int

	for i, e := range g.Elves {
		if e.Calories > max {
			max = e.Calories
			maxElf = e
			maxElfIndex = i
		}
	}

	if removeElf {
		g.Elves = append(g.Elves[:maxElfIndex], g.Elves[maxElfIndex+1:]...)
	}

	return maxElf
}

func (g *ElfGroup) FindTop3CalorieElves() []*Elf {
	var maxElves []*Elf

	maxElves = append(maxElves, g.FindMaxCalorieElf(true))
	maxElves = append(maxElves, g.FindMaxCalorieElf(true))
	maxElves = append(maxElves, g.FindMaxCalorieElf(true))

	return maxElves
}
