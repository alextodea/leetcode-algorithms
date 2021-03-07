package hashMaps

import "testing"

func TestNumIdenticalPairs1(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	result := numIdenticalPairs(nums)
	if result != 4 {
		t.Error("result should be 4")
	}
}

func TestNumIdenticalPairs2(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	result := numIdenticalPairs(nums)
	if result != 6 {
		t.Error("result should be 6")
	}
}

func TestNumIdenticalPairs3(t *testing.T) {
	nums := []int{1, 2, 3}
	result := numIdenticalPairs(nums)
	if result != 0 {
		t.Error("result should be 0")
	}
}
