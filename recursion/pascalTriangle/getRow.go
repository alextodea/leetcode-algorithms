package recursion

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	var output []int

	for j := 0; j <= rowIndex; j++ {
		output = append(output, getIndex(rowIndex, j))
	}

	return output
}

func getIndex(i, j int) int {
	if i == 0 || j==0 || j == i {
		return 1
	}

	return getIndex(i-1, j-1) + getIndex(i-1, j)
}
