package recursion

import (
	"strconv"
)

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	return getRowIndex(rowIndex)
}

var memo = make(map[string]int)

func getRowIndex(rowIndex int) []int {
	var output []int

	for colIndex := 0; colIndex <= rowIndex; colIndex++ {
		output = append(output, getIthIndex(rowIndex, colIndex))
	}

	return output
}

func getIthIndex(row, col int) int {
	if col == 0 || col == row {
		return 1
	}

	key := strconv.Itoa(row) + strconv.Itoa(col)

	if _, ithIndexExistsInMemo := memo[key]; !ithIndexExistsInMemo {
		memo[key] = getIthIndex(row-1, col-1) + getIthIndex(row-1, col)
	}

	return memo[key]
}
