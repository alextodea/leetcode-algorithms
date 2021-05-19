package recursion

import "testing"

func Test1(t *testing.T) {
	result := sortArray([]int{5, 2, 3, 1})
	expectedResult := []int{1, 2, 3, 5}

	for i, val := range result {
		if val != expectedResult[i] {
			t.Error("values are not equal")
		}
	}
}

func Test2(t *testing.T) {
	result := sortArray([]int{5, 1, 1, 2, 0, 0})
	expectedResult := []int{0, 0, 1, 1, 2, 5}

	for i, val := range result {
		if val != expectedResult[i] {
			t.Error("values are not equal")
		}
	}
}

func Test3(t *testing.T) {
	result := sortArray([]int{0, 0, 0})
	expectedResult := []int{0, 0, 0}

	for i, val := range result {
		if val != expectedResult[i] {
			t.Error("values are not equal")
		}
	}
}

func Test4(t *testing.T) {
	result := sortArray([]int{-2, 3, -5})
	expectedResult := []int{-5, -2, 3}

	for i, val := range result {
		if val != expectedResult[i] {
			t.Error("values are not equal")
		}
	}
}
