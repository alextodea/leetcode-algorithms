package arrays

import "testing"

func compareLists(t *testing.T, arr1, arr2 []int) {
	if len(arr1) != len(arr2) {
		t.Error("arrays are not equal")
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			t.Error("array values are not identic")
		}
	}
}

func TestReplaceElements1(t *testing.T) {
	result := replaceElements([]int{17, 18, 5, 4, 6, 1})
	expectedResult := []int{18, 6, 6, 6, 1, -1}

	compareLists(t, result, expectedResult)
}
