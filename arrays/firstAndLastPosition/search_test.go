package arrays

import "testing"

func Test1(t *testing.T) {
	result := searchRange([]int{}, 1)
	for _, val := range result {
		if val != -1 {
			t.Error("val should be -1")
		}
	}
}

func Test2(t *testing.T) {
	result := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	if result[0] != 3 || result[1] != 4 {
		t.Error("res incorrect")
	}
}

func Test3(t *testing.T) {
	result := searchRange([]int{2, 2}, 3)
	if result[0] != -1 || result[1] != -1 {
		t.Error("res incorrect")
	}
}

func Test4(t *testing.T) {
	result := searchRange([]int{1}, 1)
	if result[0] != 0 || result[1] != 0 {
		t.Error("res incorrect")
	}
}

func Test5(t *testing.T) {
	result := searchRange([]int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 10}, 8)
	if result[0] != 3 || result[1] != 14 {
		t.Error("res incorrect")
	}
}

func Test6(t *testing.T) {
	result := searchRange([]int{1,2,3}, 2)
	if result[0] != 1 || result[1] != 1 {
		t.Error("res incorrect")
	}
}
