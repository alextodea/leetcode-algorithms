package recursion

import "testing"

func Test1(t *testing.T) {
	row1 := []byte{1, 1, 0, 0, 0}
	row3 := []byte{0, 0, 1, 0, 0}
	row4 := []byte{0, 0, 0, 1, 1}
	input := [][]byte{row1, row1, row3, row4}
	result := numIslands(input)

	if result != 3 {
		t.Error("there are 3 islands")
	}
}

func Test2(t *testing.T) {
	row1 := []byte{1, 1, 1, 1, 0}
	row2 := []byte{1, 1, 0, 1, 0}
	row3 := []byte{1, 1, 0, 0, 0}
	row4 := []byte{0, 0, 0, 0, 0}
	input := [][]byte{row1, row2, row3, row4}
	result := numIslands(input)

	if result != 1 {
		t.Error("there are 1 islands")
	}
}

func Test3(t *testing.T) {
	row1 := []byte{1, 1, 0, 0, 0}
	row2 := []byte{1, 1, 0, 0, 0}
	row3 := []byte{0, 0, 1, 0, 0}
	row4 := []byte{0, 0, 0, 1, 1}
	input := [][]byte{row1, row2, row3, row4}
	result := numIslands(input)

	if result != 3 {
		t.Error("there are 3 islands")
	}
}