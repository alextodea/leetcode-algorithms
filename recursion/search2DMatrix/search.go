package recursion

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix) - 1
	col := 0

	for col < len(matrix[0]) && row > -1 {
		if matrix[row][col] == target {
			return true
		} else if target < matrix[row][col] {
			row--
		} else {
			col++
		}
	}
	return false
}
