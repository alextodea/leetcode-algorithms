package quicksort

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type PivotType int

const (
	FirstArrayElement PivotType = iota
	LastArrayElement
	MedianArrayElement
)

type ListOf3 [3]int

func (a ListOf3) Len() int           { return len(a) }
func (a ListOf3) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ListOf3) Less(i, j int) bool { return a[i] < a[j] }

type input []int

var totalNumberComparisons int

func (arr input) get_length() int {
	return len(arr)
}

func (arr input) sort(left, right int) {
	if right <= left {
		return
	}

	if (right - left) < len(arr) {
		totalNumberComparisons += right - left
	}

	// pivotIndex := arr.chooseIndexAsPivotIndex(left, right, FirstArrayElement)
	// pivotIndex := arr.chooseIndexAsPivotIndex(left, right, LastArrayElement)
	pivotIndex := arr.chooseIndexAsPivotIndex(left, right, MedianArrayElement)

	arr.swap(left, pivotIndex)

	limit := arr.partition(left, right)

	arr.sort(left, limit-1)
	arr.sort(limit+1, right)
}

func (arr input) chooseIndexAsPivotIndex(left, right int, pivotIndex PivotType) int {
	if pivotIndex == FirstArrayElement {
		return left
	} else if pivotIndex == MedianArrayElement {
		return arr.calculateMedianOfThree(left, right)
	}
	return right
}

func (arr input) calculateMedianOfThree(left, right int) int {
	arrMiddle := (left + right) / 2
	values := [3]int{arr[left], arr[arrMiddle], arr[right]}
	mapOfIndices := map[int]int{
		arr[left]:      left,
		arr[arrMiddle]: arrMiddle,
		arr[right]:   right,
	}

	sort.Ints(values[:])

	return mapOfIndices[values[1]]
}

func (arr input) partition(left, right int) int {
	i := left + 1
	j := left + 1
	p := left

	for j <= right {
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
