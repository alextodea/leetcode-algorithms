package recursion

type IslandsMap struct {
	nrIslands int
	grid      [][]int
}

func maxAreaOfIsland(grid [][]int) int {
	islandsMap := IslandsMap{nrIslands: 0, grid: grid}

	for m := 0; m < len(islandsMap.grid); m++ {
		for n := 0; n < len(islandsMap.grid[m]); n++ {
			if islandsMap.grid[m][n] == 1 {
				area := islandsMap.getArea(m, n)
				if area > islandsMap.nrIslands {
					islandsMap.nrIslands = area
				}
			}
		}
	}
	return islandsMap.nrIslands
}

func (im *IslandsMap) getArea(m, n int) int {
	if m < 0 || n < 0 || m >= len(im.grid) || n >= len(im.grid[0]) || im.grid[m][n] == 0 {
		return 0
	}

	im.grid[m][n] = 0

	return 1 + im.getArea(m+1, n) + im.getArea(m, n+1) + im.getArea(m-1, n) + im.getArea(m, n-1)
}
