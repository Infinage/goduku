package core

import (
	"errors"
	"math/rand/v2"
)

// Generate a random sudoku puzzle
func generateSolvedBoard() Sudoku {
	var board Sudoku
	var ctx = newBtContext(board, true, 1)
	backtrack(&board, 0, 0, &ctx)
	board = ctx.solutions[0]
	return board
}

// Generates a random puzzle by using backtracking solver and punching holes
func GenerateSudoku() (Sudoku, error) {
	// Prepare a randomly filled board
	board := generateSolvedBoard()

	// Get a list of indices to punch out
	var indices [81]Index
	for i := range 81 {
		indices[i] = Index{
			Row: uint8(i / 9),
			Col: uint8(i % 9),
		}
	}

	// Shuffle the indices
	rand.Shuffle(81, func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	for _, idx := range indices {
		var prev = board[idx.Row][idx.Col]
		board[idx.Row][idx.Col] = 0
		var ctx = newBtContext(board, false, 2)
		if !backtrack(&board, 0, 0, &ctx) || len(ctx.solutions) == 0 {
			return board, errors.New("Something went wrong with backtracking logic")
		}

		// If solution is not unique, fill the hole back in and try another
		if len(ctx.solutions) > 1 {
			board[idx.Row][idx.Col] = prev
		}
	}

	if len(board.Validate(false)) > 0 {
		return board, errors.New("Something went wrong, validation failed on returned puzzle")
	}

	return board, nil
}
