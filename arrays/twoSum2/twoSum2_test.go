package arrays

import "testing"

func Test1(t *testing.T) {
	result := twoSum2([]int{2, 7, 11, 15}, 9)
	expected := []int{1, 2}

	for i, val := range expected {
		if val != result[i] {
			t.Error("expected does not equal actual result")
		}
	}
}

func Test2(t *testing.T) {
	result := twoSum2([]int{2, 3, 4}, 6)
	expected := []int{1, 3}

	for i, val := range expected {
		if val != result[i] {
			t.Error("expected does not equal actual result")
		}
	}
}

func Test3(t *testing.T) {
	result := twoSum2([]int{-1, 0}, -1)
	expected := []int{1, 2}

	for i, val := range expected {
		if val != result[i] {
			t.Error("expected does not equal actual result")
		}
	}
}
