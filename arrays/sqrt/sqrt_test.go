package arrays

import "testing"

func TestSearchWhenInputIsNull(t *testing.T) {
	result := mySqrt(2)
	if result != 1 {
		t.Error("res should be 1")
	}
}

func TestSearchWhenInputIs4(t *testing.T) {
	result := mySqrt(4)
	if result != 2 {
		t.Error("res should be 2")
	}
}

func TestSearchWhenInputIs16(t *testing.T) {
	result := mySqrt(16)
	if result != 4 {
		t.Error("res should be 4")
	}
}

// func TestSearchWhenInputDoesNotContainTarget(t *testing.T) {
// 	result := search([]int{1, 2, 3}, 4)
// 	if result != -1 {
// 		t.Error("res should be -1")
// 	}
// }

// func TestSearchWhenInputContainsTarget(t *testing.T) {
// 	result := search([]int{1, 2, 3, 4, 5, 6, 7}, 7)
// 	if result != -1 {
// 		t.Error("res should be 6")
// 	}
// }

// func TestSearchWhenTargetIsNegative(t *testing.T) {
// 	result := search([]int{-1, 0, 5}, -1)
// 	if result != 0 {
// 		t.Error("res should be 6")
// 	}
// }
