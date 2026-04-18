package core

// Solves and exits after first solution
// Returns back to original state if no solution found
// shuffle introduces randomness into the order in which entries are attempted
func backtrack(board *Sudoku, row, col int, shuffle bool) bool {
	if row >= 9 {
		return true
	}

	if col >= 9 {
		return backtrack(board, row+1, 0, shuffle)
	}

	if (*board)[row][col] != 0 {
		return backtrack(board, row, col+1, shuffle)
	}

	for n := range sequence(1, 9, shuffle) {
		(*board)[row][col] = n
		if board.validate(row, col) {
			if backtrack(board, row, col+1, shuffle) {
				return true
			}
		}
	}

	(*board)[row][col] = 0
	return false
}

// Solve the sudoku puzzle board, returns true is solved else false
func (board *Sudoku) Solve() bool {
	if !board.Validate(false) {
		return false
	}
	return backtrack(board, 0, 0, false)
}
