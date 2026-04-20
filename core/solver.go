package core

import (
	"math/rand"
)

// backtrackContext maintains search state across recursive calls.
type backtrackContext struct {
	rows [9]uint16 // Stores bitmask to rows
	cols [9]uint16 // Stores bitmask to cols
	grds [9]uint16 // Stores bitmask to grids

	fillOrder [9]uint8 // Order in which backtracker will attempt to fill entries

	shuffle   bool     // Randomizes digit attempt order (useful for generation)
	exitAfter uint8    // Exits after specified count of solutions are found
	solutions []Sudoku // Stores full board snapshots of found solutions
}

func newBtContext(board Sudoku, shuffle bool, exitAfter uint8) backtrackContext {
	var ctx = backtrackContext{shuffle: shuffle, exitAfter: exitAfter}

	// Initialize the fill order sequence
	for i := range uint8(9) {
		ctx.fillOrder[i] = i + 1
	}
	if ctx.shuffle {
		rand.Shuffle(len(ctx.fillOrder), func(i, j int) {
			ctx.fillOrder[i], ctx.fillOrder[j] = ctx.fillOrder[j], ctx.fillOrder[i]
		})
	}

	// Mark the bitmaps
	for row := range 9 {
		for col := range 9 {
			if val := board[row][col]; val != 0 {
				ctx.mark(row, col, val, true)
			}
		}
	}

	return ctx
}

func (ctx *backtrackContext) isSafe(row, col int, n uint8) bool {
	bit := uint16(1 << n)
	grd := (row/3)*3 + col/3
	return ctx.rows[row]&bit == 0 && ctx.cols[col]&bit == 0 && ctx.grds[grd]&bit == 0
}

func (ctx *backtrackContext) mark(row, col int, n uint8, val bool) {
	bit := uint16(1 << n)
	grd := (row/3)*3 + (col / 3)
	if val {
		// OR
		ctx.rows[row] |= bit
		ctx.cols[col] |= bit
		ctx.grds[grd] |= bit
	} else {
		// AND-NOT
		ctx.rows[row] &^= bit
		ctx.cols[col] &^= bit
		ctx.grds[grd] &^= bit
	}
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
	for _, n := range ctx.fillOrder {
		// Prune branch if digit violates Sudoku constraints
		if ctx.isSafe(row, col, n) {
			(*board)[row][col] = n
			ctx.mark(row, col, n, true)
			solved = backtrack(board, row, col+1, ctx) || solved
			ctx.mark(row, col, n, false)
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

	var ctx = newBtContext(*board, false, 1)
	if !backtrack(board, 0, 0, &ctx) {
		return false
	}

	*board = ctx.solutions[0]
	return true
}
