package custom

import "testing"

func Test1(t *testing.T) {
	var A = []int{1, 2, 3, 3}
	var B = []int{2, 3, 1, 4}
	N := 4
	result := Solution(A, B, N)
	expectedResult := 4

	if result != expectedResult {
		t.Error("sasa")
	}
}

func Test2(t *testing.T) {
	var A = []int{1, 2, 4, 5}
	var B = []int{2, 3, 5, 6}
	N := 6
	result := Solution(A, B, N)
	expectedResult := 2

	if result != expectedResult {
		t.Error("sasa")
	}
}
