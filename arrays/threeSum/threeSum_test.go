package arrays

import "testing"

func Test1(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(input)
	expectedResult := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	if len(expectedResult) != len(result) {
		t.Error("lengths should be equal")
	}

	for i := 0; i < len(expectedResult); i++ {
		subArr := expectedResult[i]
		for j := 0; j < len(subArr); j++ {
			val := subArr[j]
			if val != result[i][j] {
				t.Error("values should be equal")
			}
		}
	}
}

func Test2(t *testing.T) {
	input := []int{}
	result := threeSum(input)

	if len(result) != 0 {
		t.Error("should be 0")
	}
}

func Test3(t *testing.T) {
	input := []int{0, 0, 0}
	result := threeSum(input)
	expected := [][]int{{0, 0, 0}}

	if len(expected) != len(result) {
		t.Error("lengths should be equal")
	}

	for i := 0; i < len(expected); i++ {
		subArr := expected[i]
		for j := 0; j < len(subArr); j++ {
			val := subArr[j]
			if val != result[i][j] {
				t.Error("values should be equal")
			}
		}
	}
}
