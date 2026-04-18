package core

// Solves and exits after first solution
// Returns back to original state if no solution found
func backtrack(board *Sudoku, row, col int) bool {
	if row >= 9 {
		return true
	}

	if col >= 9 {
		return backtrack(board, row+1, 0)
	}

	if (*board)[row][col] != 0 {
		return backtrack(board, row, col+1)
	}

	for n := range 9 {
		(*board)[row][col] = uint8(n + 1)
		if board.validate(row, col) {
			if backtrack(board, row, col+1) {
				return true
			}
		}
	}

	(*board)[row][col] = 0
	return false
}

// Solve the sudoku puzzle board, returns true is solved else false
func (board *Sudoku) Solve() bool {
	return backtrack(board, 0, 0)
}
