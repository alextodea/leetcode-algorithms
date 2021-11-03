package recursion

import "testing"

func Test1(t *testing.T) {
	row1 := []int{1, 1, 0, 0, 0}
	row3 := []int{0, 0, 1, 0, 0}
	row4 := []int{0, 0, 0, 1, 1}
	input := [][]int{row1, row1, row3, row4}
	result := maxAreaOfIsland(input)

	if result != 2 {
		t.Error("max area should be 2")
	}
}

func Test2(t *testing.T) {
	row1 := []int{1, 1, 1, 1, 0}
	row2 := []int{1, 1, 0, 1, 0}
	row3 := []int{1, 1, 0, 0, 0}
	row4 := []int{0, 0, 0, 0, 0}
	input := [][]int{row1, row2, row3, row4}
	result := maxAreaOfIsland(input)

	if result != 9 {
		t.Error("max area should be 9")
	}
}

func Test3(t *testing.T) {
	row1 := []int{1, 1, 0, 0, 0}
	row2 := []int{1, 1, 0, 0, 0}
	row3 := []int{0, 0, 1, 0, 0}
	row4 := []int{0, 0, 0, 1, 1}
	input := [][]int{row1, row2, row3, row4}
	result := maxAreaOfIsland(input)

	if result != 4 {
		t.Error("max area should be 4")
	}
}

func Test4(t *testing.T) {
	row1 := []int{1}
	input := [][]int{row1}
	result := maxAreaOfIsland(input)

	if result != 1 {
		t.Error("max area should be 1")
	}
}
