package custom

import "testing"

func TestWhenInputNull(t *testing.T) {
	result := Solution(nil)

	if result != false {
		t.Error("res should be nil")
	}
}

func TestWhenArrayLenIs1(t *testing.T) {
	result := Solution([]int{1})

	if result != false {
		t.Error("res should be nil")
	}
}

func TestWhenArrayContainsTwoDistinctValues(t *testing.T) {
	result := Solution([]int{1, 2})

	if result != false {
		t.Error("res should be nil")
	}
}

func TestWhenArrayContainsTwoPairs(t *testing.T) {
	result := Solution([]int{1, 2, 2, 1})

	if result != true {
		t.Error("res should be true")
	}
}

func TestWhenArrayContainsThreeOfAKind(t *testing.T) {
	result := Solution([]int{7, 7, 7})

	if result != false {
		t.Error("res should be nil")
	}
}

func TestWhenArrayContainsOnlyOnePair(t *testing.T) {
	result := Solution([]int{1, 2, 2, 3})

	if result != false {
		t.Error("res should be nil")
	}
}
