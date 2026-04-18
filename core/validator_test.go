package core

import "testing"

func TestSudoku_Validate(t *testing.T) {
	t.Run("Valid Partial Board", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[0][8] = 9
		board[8][0] = 1
		// No conflicts
		if !board.Validate(false) {
			t.Error("Validate(false) failed on a valid partial board")
		}
	})

	t.Run("Row Conflict", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[0][5] = 5 // Conflict in row 0
		if board.Validate(false) {
			t.Error("Validate() failed to detect row conflict")
		}
	})

	t.Run("Column Conflict", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[5][0] = 5 // Conflict in column 0
		if board.Validate(false) {
			t.Error("Validate() failed to detect column conflict")
		}
	})

	t.Run("Grid Conflict", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[1][1] = 5 // Conflict in top-left 3x3 grid
		if board.Validate(false) {
			t.Error("Validate() failed to detect grid conflict")
		}
	})

	t.Run("CheckSolved with zeros", func(t *testing.T) {
		var board Sudoku
		// Board is empty (all zeros)
		if board.Validate(true) {
			t.Error("Validate(true) should return false for empty board")
		}
	})
}
