package day03

import (
	"fmt"
	"unicode"
)

type RuckSackCollection struct {
	Sacks []ItemFinder
}

func NewRuckSackCollection() *RuckSackCollection {
	return &RuckSackCollection{
		Sacks: make([]ItemFinder, 0),
	}
}

func (c *RuckSackCollection) AddRuckSack(s ItemFinder) error {
	c.Sacks = append(c.Sacks, s)
	return nil
}

func (c *RuckSackCollection) SumPriorities() (int, error) {
	if len(c.Sacks) == 0 {
		return 0, fmt.Errorf("no sacks in the collection")
	}

	sum := 0

	for _, sack := range c.Sacks {
		dupItem, err := sack.FindDuplicateItem()
		if err != nil {
			return 0, err
		}

		priority, err := c.CalculateItemPriority(dupItem)
		sum += priority
	}

	return sum, nil
}

func (c *RuckSackCollection) CalculateItemPriority(item string) (int, error) {
	if len(item) != 1 {
		return 0, fmt.Errorf("item length must be 1: %s", item)
	}

	for _, itemUnicode := range item {
		if unicode.IsUpper(itemUnicode) {
			return int(itemUnicode) - 38, nil
		}
		if unicode.IsLower(itemUnicode) {
			return int(itemUnicode) - 96, nil
		}
	}

	return 0, fmt.Errorf("could not calculate item priority")
}
