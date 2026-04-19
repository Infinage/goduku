package core

// backtrackContext maintains search state across recursive calls.
type backtrackContext struct {
	shuffle   bool     // Randomizes digit attempt order (useful for generation)
	exitAfter uint8    // Exits after specified count of solutions are found
	solutions []Sudoku // Stores full board snapshots of found solutions
}

// `backtrack` performs a depth-first search to fill empty cells.
// It returns true if at least one solution was found in the current branch.
// Solutions can be obtained from the context, board will be reset on return
func backtrack(board *Sudoku, row, col int, ctx *backtrackContext) bool {
	// Base case: row 9 means the entire board was traversed successfully
	if row == 9 {
		ctx.solutions = append(ctx.solutions, *board)
		return true
	}

	// Move to the next row once the current row is exhausted
	if col == 9 {
		return backtrack(board, row+1, 0, ctx)
	}

	// Skip immutable cells and proceed to the next column
	if (*board)[row][col] != 0 {
		return backtrack(board, row, col+1, ctx)
	}

	var solved bool = false
	for n := range sequence(1, 9, ctx.shuffle) {
		(*board)[row][col] = n

		// Prune branch if digit violates Sudoku constraints
		if board.validate(row, col) {
			solved = backtrack(board, row, col+1, ctx) || solved
			if solved && len(ctx.solutions) >= int(ctx.exitAfter) {
				break
			}
		}
	}

	// Revert state for parent calls
	(*board)[row][col] = 0
	return solved
}

// Solve attempts to find a solution for the current board state.
// Returns false if the initial state is invalid or no solution exists.
// Updates the board in place
func (board *Sudoku) Solve() bool {
	if len(board.Validate(false)) > 0 {
		return false
	}

	var ctx = backtrackContext{exitAfter: 1}
	if !backtrack(board, 0, 0, &ctx) {
		return false
	}

	*board = ctx.solutions[0]
	return true
}
