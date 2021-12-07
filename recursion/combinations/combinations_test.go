package recursion

import "testing"

func Test1(t *testing.T) {
	k, n := 2, 4
	result := combine(n, k)
	expected := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}

	if len(result) != len(expected) {
		t.Error("array lengths are not the same")
	}

	for j, subArr := range result {
		for i := 0; i < len(subArr); i++ {
			if result[j][i] != subArr[i] {
				t.Error("combination element is not the same")
			}
		}
	}
}

func Test2(t *testing.T) {
	k, n := 1, 1
	result := combine(n, k)
	expected := [][]int{
		{1},
	}

	if len(result) != len(expected) {
		t.Error("array lengths are not the same")
	}
}

func Test3(t *testing.T) {
	k, n := 1, 2
	result := combine(n, k)
	expected := [][]int{
		{1},
		{2},
	}

	if len(result) != len(expected) {
		t.Error("array lengths are not the same")
	}
}
