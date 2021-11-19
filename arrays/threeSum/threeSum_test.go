package arrays

import "testing"

func Test1(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(input)
	expectedResult := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	for i := 0; i < len(expectedResult); i++ {
		subArr := expectedResult[i]
		for j:=0; j < len(subArr);j++ {
			val := subArr[j]
			if val != result[i][j] {
				t.Error("values should be equal")
			}
		}
	}
}
