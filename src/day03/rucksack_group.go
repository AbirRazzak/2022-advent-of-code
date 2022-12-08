package day03

import "fmt"

type RuckSackGroup struct {
	group []*RuckSack
}

func NewRuckSackGroup(r1 *RuckSack, r2 *RuckSack, r3 *RuckSack) *RuckSackGroup {
	return &RuckSackGroup{
		group: []*RuckSack{r1, r2, r3},
	}
}

// FindDuplicateItem returns the item that is repreated between all 3 RuckSack's
func (g *RuckSackGroup) FindDuplicateItem() (string, error) {
	for _, item1 := range g.group[0].Contents {
		for _, item2 := range g.group[1].Contents {
			if item1 == item2 {
				for _, item3 := range g.group[2].Contents {
					if item2 == item3 {
						return string(item1), nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("could not find badge in group")
}

type RuckSackGroupFactory struct {
	tempRuckSacks []*RuckSack
}

func NewRuckSackGroupFactory() *RuckSackGroupFactory {
	return &RuckSackGroupFactory{}
}

// ProcessRuckSack will take in a RuckSack and return a RuckSackGroup if 3 RuckSack's are processed successfully
func (f *RuckSackGroupFactory) ProcessRuckSack(s *RuckSack) *RuckSackGroup {
	f.tempRuckSacks = append(f.tempRuckSacks, s)

	if len(f.tempRuckSacks) == 3 {
		group := NewRuckSackGroup(f.tempRuckSacks[0], f.tempRuckSacks[1], f.tempRuckSacks[2])
		f.tempRuckSacks = []*RuckSack{}

		return group
	}

	return nil
}
