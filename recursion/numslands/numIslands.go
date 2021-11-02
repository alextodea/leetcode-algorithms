package recursion

type Output struct {
	grid      [][]byte
	nrIslands int
}

type Coordinate struct {
	m int
	n int
}

// bfs using iteration
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	output := Output{grid: grid}
	return output.findNumberOfIslands()
}

func (o *Output) findNumberOfIslands() int {
	for m := 0; m < len(o.grid); m++ {
		for n := 0; n < len(o.grid[0]); n++ {
			if o.grid[m][n] != 0 {
				o.nrIslands++
				o.search(m, n)
			}
		}
	}
	return o.nrIslands
}

func (o *Output) search(m, n int) {
	queue := []Coordinate{}
	queue = append(queue, Coordinate{m,n})

	for len(queue) > 0 {
		currentRow := queue[0].m
		currentCol := queue[0].n
		o.grid[currentRow][currentCol] = 0 // mark as visited

		// check if neighbour is unvisited and add to queue
		if currentRow + 1 < len(o.grid) && o.grid[currentRow + 1][currentCol] == 1 {
			queue = append(queue, Coordinate{currentRow + 1,currentCol})
		}

		// check if neighbour is unvisited and add to queue
		if queue[0].n + 1 < len(o.grid[0]) && o.grid[currentRow][currentCol + 1] == 1 {
			queue = append(queue, Coordinate{currentRow,currentCol + 1})
		}

		queue = queue[1:] // pop first element of queue
	}
}

// dfs using recursion
// func (o *Output) search(row, col int) {
// 	if col < 0 || row < 0 || row >= len(o.grid) || col >= len(o.grid[0]) ||  o.grid[row][col] == 0 {
// 		return
// 	}

// 	o.grid[row][col] = 0

// 	o.search(row+1, col)
// 	o.search(row, col+1)
// 	o.search(row-1, col)
// 	o.search(row, col-1)

// 	return
// }

// func (o *Output) findNumberOfIslands() int {
// 	for m := 0; m < len(o.grid); m++ {
// 		for n := 0; n < len(o.grid[0]); n++ {
// 			if o.grid[m][n] != 0 {
// 				o.nrIslands++
// 				o.search(m, n)
// 			}
// 		}
// 	}
// 	return o.nrIslands
// }

// func numIslands(grid [][]byte) int {
// 	output := Output{grid: grid}
// 	return output.findNumberOfIslands()
// }
