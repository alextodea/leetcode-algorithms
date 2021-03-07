package custom

import "testing"

func Test1(t *testing.T) {
	arr := []int32{7, 1, 2}
	result := howManySwaps(arr)
	if result != 2 {
		t.Error("result should be 4")
	}
}
