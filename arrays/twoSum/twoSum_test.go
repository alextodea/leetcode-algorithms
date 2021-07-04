package arrays

import "testing"

func Test1(t *testing.T) {
	input := []int{3,2,3}
	result := twoSum(input,6)
	if result != nil{
		t.Error("s")
	}

}

func Test2(t *testing.T) {
	input := []int{2,7,11,15}
	result := twoSum(input,9)
	if result != nil{
		t.Error("s")
	}

}