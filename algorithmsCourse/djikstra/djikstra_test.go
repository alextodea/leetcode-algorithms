package djikstra

import (
	"testing"
)

func Test1(t *testing.T) {
	result := []int{}
	input := []int{2, 3, 4, 5, 6, 7, 8}

	for _, val := range input {
		result = append(result, main("test1", val))
	}

	expected_arr := []int{1, 2, 3, 4, 4, 3, 2}

	for i, exp_val := range expected_arr {
		if result[i] != exp_val {
			t.Error("elements should be equal")
		}
	}
}

func Test2(t *testing.T) {
	result := []int{}
	input := []int{2,3,4,5,6,7,8,9,10,11}

	for _, val := range input {
		result = append(result, main("test2", val))
	}

	expected_arr := []int{3,5,8,5,7,11,4,6,10,10}

	for i, exp_val := range expected_arr {
		if result[i] != exp_val {
			t.Error("elements should be equal")
		}
	}
}

func Test3(t *testing.T) {
	result := main("test3", 7)

	if result != 9 {
		t.Error("res should be 7")
	}
}

func TestMain(t *testing.T) {
	result := []int{}
	input := []int{7,37,59,82,99,115,133,165,188,197}

	for _, val := range input {
		result = append(result, main("datasetMain", val))
	}

	if result[0] != 3 {
		t.Error("result should be 3")
	}
}
