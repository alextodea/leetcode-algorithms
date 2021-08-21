package quicksort

import (
	"bufio"
	"os"
	"strconv"
)

type input []int

var totalNumberComparisons int

func (arr input) get_length() int {
	return len(arr)
}

func (arr input) sort(left, right int) {
	if right <= left {
		return
	}

	limit := arr.partition(left, right)

	arr.sort(left, limit-1)
	arr.sort(limit+1, right)
}

func (arr input) partition(left, right int) int {
	i := left + 1
	j := i
	p := left

	if (right - left) < len(arr) {
		totalNumberComparisons += right - left
	}

	for j < right {
		if arr[j] < arr[p] {
			arr.swap(j, i)
			i++
		}
		j++
	}

	arr.swap(p, i-1)

	return i - 1
}

func (arr input) swap(x, y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

func parseInputFile(inputFilePath string) (lines []int, err error) {
	if len(inputFilePath) < 1 {
		return nil, err
	}

	file, err := os.Open(inputFilePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineValueTxt := scanner.Text()
		lineValueInt, err := strconv.Atoi(lineValueTxt)

		if err != nil {
			return nil, err
		}

		lines = append(lines, lineValueInt)
	}

	return lines, nil
}
