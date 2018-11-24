package strategy

import "fmt"

type (
	Human struct {
		Age    int
		Height int
		Weight int
	}

	Comparator interface {
		Compare(h1 Human, h2 Human) int
	}

	AgeComparator    struct{}
	HeightComparator struct{}
	WeightComparator struct{}
)

func (c *AgeComparator) Compare(h1 *Human, h2 *Human) int {
	if h1.Age >= h2.Age {
		return 1
	} else {
		return -1
	}
}

func (c *HeightComparator) Compare(h1 *Human, h2 *Human) int {
	if h1.Height >= h2.Height {
		return 1
	} else {
		return -1
	}
}

func (c *WeightComparator) Compare(h1 *Human, h2 *Human) int {
	if h1.Weight >= h2.Weight {
		return 1
	} else {
		return -1
	}
}

func NewHuman(age int, height int, weight int) *Human {
	return &Human{
		Age:    age,
		Height: height,
		Weight: weight,
	}
}

func Run() {
	h1 := NewHuman(24, 173, 75)
	h2 := NewHuman(20, 173, 63)
	comparator := &AgeComparator{}
	result := comparator.Compare(h1, h2)
	fmt.Println(result)
}
