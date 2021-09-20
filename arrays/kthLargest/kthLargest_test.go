package arrays

import (
	"testing"
)

var arr1 = []int{3, 2, 1, 5, 6, 4}
var arr2 = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}

func TestWhenInputNull(t *testing.T) {
	result := findKthLargest(nil, 3)

	if result != -1 {
		t.Error("reult should be -1")
	}
}

func TestWithArray1(t *testing.T) {
	result := findKthLargest(arr1, 2)

	if result != 5 {
		t.Error("reult should be 5")
	}
}

func TestWithArray2(t *testing.T) {
	result := findKthLargest(arr2, 4)

	if result != 4 {
		t.Error("reult should be 4")
	}
}

func TestWithOneElementArray(t *testing.T) {
	result := findKthLargest([]int{1}, 1)

	if result != 1 {
		t.Error("reult should be 1")
	}
}

func TestWithArray3(t *testing.T) {
	result := findKthLargest([]int{99,99}, 1)

	if result != 99 {
		t.Error("reult should be 99")
	}
}