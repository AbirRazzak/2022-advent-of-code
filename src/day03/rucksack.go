package day03

import "fmt"

type RuckSack struct {
	Contents     string
	Compartment1 string
	Compartment2 string
}

func NewRuckSack(contents string) *RuckSack {
	return &RuckSack{
		Contents:     contents,
		Compartment1: "",
		Compartment2: "",
	}
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

// FindDuplicateItem looks at the items in Compartment1 and Compartment2 and returns the item that appears in both
func (r *RuckSack) FindDuplicateItem() (string, error) {
	err := error(nil)
	r.Compartment1, r.Compartment2, err = r.GetCompartments()
	if err != nil {
		return "", err
	}

	if r.Compartment1 == "" || r.Compartment2 == "" {
		return "", fmt.Errorf("missing compartments")
	}

	for _, item1 := range r.Compartment1 {
		for _, item2 := range r.Compartment2 {
			if item1 == item2 {
				return string(item1), nil
			}
		}
	}

	return "", fmt.Errorf("could not find matching items between compartments %s %s", r.Compartment1, r.Compartment2)
}
