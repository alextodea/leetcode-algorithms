package recursion

import "testing"

func Test1(t *testing.T) {
	input := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	result := searchMatrix(input, 5)
	if result != true {
		t.Error("res should be true")
	}
}

func Test2(t *testing.T) {
	input := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	result := searchMatrix(input, 20)
	if result != false {
		t.Error("res should be false")
	}
}

func Test3(t *testing.T) {
	input := [][]int{{1,1}}
	result := searchMatrix(input, 2)
	if result != false {
		t.Error("res should be false")
	}
}

func Test4(t *testing.T) {
	input := [][]int{{-1,3}}
	result := searchMatrix(input, 3)
	if result != true {
		t.Error("res should be true")
	}
}