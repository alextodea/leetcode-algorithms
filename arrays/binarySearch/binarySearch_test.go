package arrays

import "testing"

func TestSearchWhenInputIsNull(t *testing.T) {
	result := search(nil, 2)
	if result != -1 {
		t.Error("res should be -1")
	}
}

func TestSearchWhenInputIsEmpty(t *testing.T) {
	result := search([]int{}, 2)
	if result != -1 {
		t.Error("res should be -1")
	}
}

func TestSearchWhenInputDoesNotContainTarget(t *testing.T) {
	result := search([]int{1, 2, 3}, 4)
	if result != -1 {
		t.Error("res should be -1")
	}
}

func TestSearchWhenInputContainsTarget(t *testing.T) {
	result := search([]int{1, 2, 3, 4, 5, 6, 7}, 7)
	if result != -1 {
		t.Error("res should be 6")
	}
}

func TestSearchWhenTargetIsNegative(t *testing.T) {
	result := search([]int{-1, 0, 5}, -1)
	if result != 0 {
		t.Error("res should be 6")
	}
}
