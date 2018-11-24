package strategy

import "testing"

func TestAgeComparator(t *testing.T) {
	h1 := NewHuman(24, 173, 75)
	h2 := NewHuman(20, 173, 63)
	comparator := &AgeComparator{}

	result := comparator.Compare(h1, h2)
	expect := 1

	if result != expect {
		t.Errorf("Expect result to equal %d, but %d.", expect, result)
	}
}

func TestHeightComparator(t *testing.T) {
	h1 := NewHuman(24, 173, 75)
	h2 := NewHuman(20, 174, 63)
	comparator := &HeightComparator{}

	result := comparator.Compare(h1, h2)
	expect := -1

	if result != expect {
		t.Errorf("Expect result to equal %d, but %d.", expect, result)
	}
}

func TestWeightComparator(t *testing.T) {
	h1 := NewHuman(24, 173, 75)
	h2 := NewHuman(20, 173, 63)
	comparator := &WeightComparator{}

	result := comparator.Compare(h1, h2)
	expect := 1

	if result != expect {
		t.Errorf("Expect result to equal %d, but %d.", expect, result)
	}
}
