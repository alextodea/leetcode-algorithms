package hashMaps

import "testing"

func TestRepeatedNTimes1(t *testing.T) {
	inputArr := []int{1, 2, 3, 3}
	result := repeatedNTimes(inputArr)

	if result != 3 {
		t.Error("result should be 3")
	}
}

func TestRepeatedNTimes2(t *testing.T) {
	inputArr := []int{2, 1, 2, 5, 3, 2}
	result := repeatedNTimes(inputArr)

	if result != 2 {
		t.Error("result should be 2")
	}
}

func TestRepeatedNTimes3(t *testing.T) {
	inputArr := []int{5, 1, 5, 2, 5, 3, 5, 4}
	result := repeatedNTimes(inputArr)

	if result != 5 {
		t.Error("result should be 5")
	}
}
