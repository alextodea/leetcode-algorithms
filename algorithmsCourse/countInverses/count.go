package countinverses

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var filePath = "data.txt"

func main() {
	arr, err := getArrayOfIntegersFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	countNumberOfInversions(arr)
}

func countNumberOfInversions(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	}

	middle := len(arr) / 2

	left, nrLeftInversions := countNumberOfInversions(arr[:middle])
	right, nrRightInversions := countNumberOfInversions(arr[middle:])
	split, nrSplitInversions := mergeAndCountInversions(left, right)

	return split, nrLeftInversions + nrRightInversions + nrSplitInversions
}

func mergeAndCountInversions(left, right []int) ([]int, int) {
	output := make([]int, len(left)+len(right))
	var nrInversions int

	var i int
	var j int
	var k int

	for k < len(output) && i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			output[k] = left[i]
			i++
		} else {
			output[k] = right[j]
			j++
			nrInversions += len(left[i:])
		}
		k++
	}

	for i < len(left) {
		output[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		output[k] = right[j]
		j++
		k++
	}

	return output, nrInversions
}

func getArrayOfIntegersFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return []int{}, err
	}

	defer file.Close()

	processedFile, err := processFile(file)

	if err != nil {
		return []int{}, err
	}

	return processedFile, nil
}

func processFile(file *os.File) ([]int, error) {
	scanner := bufio.NewScanner(file)
	outputArr := make([]int, 0)

	for scanner.Scan() {
		processedLine, err := processLine(scanner.Text())

		if err != nil {
			return []int{}, err
		}

		outputArr = append(outputArr, processedLine)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	fmt.Println("succesfully created array of integers from file")
	return outputArr, nil
}

func processLine(line string) (int, error) {
	val, err := strconv.Atoi(line)
	if err != nil {
		return -1, err
	}
	return val, nil
}
