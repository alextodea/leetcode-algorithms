package recursion

import "testing"

func Test1(t *testing.T) {
	result := totalNQueens(1)

	if result != 1 {
		t.Error("result should be 1")
	}
}

func Test2(t *testing.T) {
	result := totalNQueens(4)

	if result != 2 {
		t.Error("result should be 2")
	}
}

func Test3(t *testing.T) {
	result := totalNQueens(5)

	if result != 10 {
		t.Error("result should be 10")
	}
}
