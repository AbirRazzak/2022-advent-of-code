package day03

import "fmt"

type ItemFinder interface {
	FindDuplicateItem() (string, error)
}

type RuckSack struct {
	Contents string
}

func NewRuckSack(contents string) *RuckSack {
	return &RuckSack{
		Contents: contents,
	}
}

// FindDuplicateItem looks at the items in Compartment1 and Compartment2 and returns the item that appears in both
func (r *RuckSack) FindDuplicateItem() (string, error) {
	compartment1, compartment2, err := r.GetCompartments()
	if err != nil {
		return "", err
	}

	if compartment1 == "" || compartment2 == "" {
		return "", fmt.Errorf("missing compartments")
	}

	for _, item1 := range compartment1 {
		for _, item2 := range compartment2 {
			if item1 == item2 {
				return string(item1), nil
			}
		}
	}

	return "", fmt.Errorf("could not find matching items between compartments %s %s", compartment1, compartment2)
}

func (r *RuckSack) GetCompartments() (string, string, error) {
	if r.Contents == "" {
		return "", "", fmt.Errorf("missing contents in RuckSack")
	}

	// Split Contents in half
	length := len(r.Contents)
	compartment1 := r.Contents[0 : length/2]
	compartment2 := r.Contents[length/2:]

	return compartment1, compartment2, nil
}
