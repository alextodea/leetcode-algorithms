package arrays

import "testing"

func Test1(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expectedResult := 49
	result := maxArea(input)

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}

func Test2(t *testing.T) {
	input := []int{0}
	expectedResult := 0
	result := maxArea(input)

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}

func Test3(t *testing.T) {
	input := []int{1, 1}
	expectedResult := 1
	result := maxArea(input)

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}

func Test4(t *testing.T) {
	input := []int{4, 3, 2, 1, 4}
	expectedResult := 16
	result := maxArea(input)

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}

func Test5(t *testing.T) {
	input := []int{1, 2, 1}
	expectedResult := 2
	result := maxArea(input)

	if result != expectedResult {
		t.Error("result should be", expectedResult)
	}
}
