package recursion

type Chessboard struct {
	n                        int
	solutions                int
	colsUnderAttack          map[int]bool
	diagonalsUnderAttack     map[int]bool
	antidiagonalsUnderAttack map[int]bool
}

func (c *Chessboard) isCellUnderAttack(col, diagonal, antidiagonal int) bool {
	if c.colsUnderAttack[col] || c.diagonalsUnderAttack[diagonal] || c.antidiagonalsUnderAttack[antidiagonal] {
		return true
	}
	return false
}

func (c *Chessboard) placeQueen(col, diagonal, antidiagonal int) {
	c.colsUnderAttack[col], c.diagonalsUnderAttack[diagonal], c.antidiagonalsUnderAttack[antidiagonal] = true, true, true
}

func (c *Chessboard) removeQueen(col, diagonal, antidiagonal int) {
	delete(c.colsUnderAttack, col)
	delete(c.diagonalsUnderAttack, diagonal)
	delete(c.antidiagonalsUnderAttack, antidiagonal)
}

func (c *Chessboard) backtrack(row int) int {
	if row == c.n {
		return 1
	}

	solutions := 0
	for col := 0; col < c.n; col++ {
		if !c.isCellUnderAttack(col, row-col, row+col) {
			c.placeQueen(col, row-col, row+col)
			solutions += c.backtrack(row + 1)
			c.removeQueen(col, row-col, row+col)
		}
	}

	return solutions
}

func totalNQueens(n int) int {
	if n <= 1 {
		return n
	}

	chessboard := Chessboard{n: n, colsUnderAttack: make(map[int]bool), diagonalsUnderAttack: make(map[int]bool), antidiagonalsUnderAttack: make(map[int]bool)}
	return chessboard.backtrack(0)
}
