package quicksort

import (
	"fmt"
	"testing"
)

var null_arr = input(nil)
var unsorted_arr = input([]int{3, 8, 2, 5, 1, 4, 7, 6})
var inputFilePath = "./input.txt"

func TestWhenInputNull(t *testing.T) {
	null_arr.sort(0, null_arr.get_length())

	if null_arr != nil {
		t.Error("result should be of type null")
	}
}

func TestWithUnsortedArray1(t *testing.T) {
	unsorted_arr.sort(0, unsorted_arr.get_length())

	expected_result := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(expected_result); i++ {
		if unsorted_arr[i] != expected_result[i] {
			t.Error("array values should be identical")
		}
	}

	fmt.Println(totalNumberComparisons)
}

func TestWithInputFromFile(t *testing.T) {
	inputArr, err := parseInputFile(inputFilePath)

	if err != nil {
		t.Error(err)
	}

	result := input(inputArr)
	result.sort(0, result.get_length())
	fmt.Println(totalNumberComparisons)
}
