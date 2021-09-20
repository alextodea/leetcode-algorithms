package quicksort

import (
	"fmt"
	"testing"
)

var null_arr = input(nil)
var unsorted_arr = input([]int{3, 2, 5, 1, 4, 6})
var inputFilePath = "./input.txt"
var inputFilePath2 = "./input2.txt"

func TestWhenInputNull(t *testing.T) {
	null_arr.sort(0, null_arr.get_length())

	if null_arr != nil {
		t.Error("result should be of type null")
	}
}

func TestWithUnsortedArray1(t *testing.T) {
	unsorted_arr.sort(0, unsorted_arr.get_length()-1)

	expected_result := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(expected_result); i++ {
		if unsorted_arr[i] != expected_result[i] {
			t.Error("array values should be identical")
		}
	}

	fmt.Println(totalNumberComparisons)
}

func TestWithInputFromFile(t *testing.T) {
	inputArr, err := parseInputFile(inputFilePath)
	expectedResult := [10000]int{}
	var j int

	for i := 0; i < 10000; i++ {
		j = i + 1
		expectedResult[i] = j
	}

	if err != nil {
		t.Error(err)
	}

	result := input(inputArr)
	result.sort(0, result.get_length()-1)

	for i := 0; i < len(expectedResult); i++ {
		if result[i] != expectedResult[i] {
			t.Error("array values should be identical")
		}
	}

	fmt.Println(totalNumberComparisons)
}

func TestWithInputFromFile2(t *testing.T) {
	inputArr, err := parseInputFile(inputFilePath2)
	expectedResult := []int{504, 609, 2148, 3153, 5469, 6324, 7017, 7628, 7742, 9058}

	if err != nil {
		t.Error(err)
	}

	result := input(inputArr)
	result.sort(0, result.get_length()-1)

	for i := 0; i < len(expectedResult); i++ {
		if result[i] != expectedResult[i] {
			t.Error("array values should be identical")
		}
	}

	fmt.Println(totalNumberComparisons)
}
