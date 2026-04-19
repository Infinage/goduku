package core

import "testing"

func TestSudoku_Validate(t *testing.T) {
	t.Run("Valid Partial Board", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[0][8] = 9
		board[8][0] = 1
		// No conflicts
		if len(board.Validate(false)) > 0 {
			t.Error("Validate(false) failed on a valid partial board")
		}
	})

	t.Run("Row Conflict", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[0][5] = 5 // Conflict in row 0
		if len(board.Validate(false)) == 0 {
			t.Error("Validate() failed to detect row conflict")
		}
	})

	t.Run("Column Conflict", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[5][0] = 5 // Conflict in column 0
		if len(board.Validate(false)) == 0 {
			t.Error("Validate() failed to detect column conflict")
		}
	})

	t.Run("Grid Conflict", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[1][1] = 5 // Conflict in top-left 3x3 grid
		if len(board.Validate(false)) == 0 {
			t.Error("Validate() failed to detect grid conflict")
		}
	})

	t.Run("CheckSolved with zeros", func(t *testing.T) {
		var board Sudoku
		// Board is empty (all zeros)
		if len(board.Validate(true)) == 0 {
			t.Error("Validate(true) should return false for empty board")
		}
	})
}

func TestSudoku_Validate_Diagnostic(t *testing.T) {
	t.Run("Identify Conflicting Cells", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 5
		board[0][5] = 5

		errors := board.Validate(false)

		// We expect at least two errors: both [0,0] and [0,5] are invalid
		if len(errors) < 2 {
			t.Errorf("Expected at least 2 conflicting cells, got %d", len(errors))
		}

		// Verify the coordinates are exactly what we expect
		expected := map[Index]bool{
			{0, 0}: true,
			{0, 5}: true,
		}

		for _, errIdx := range errors {
			if !expected[errIdx] {
				t.Errorf("Unexpected error index reported: %+v", errIdx)
			}
		}
	})

	t.Run("Valid Full Board", func(t *testing.T) {
		// You would need a known-solved board here
		board := generateSolvedBoard()
		if len(board.Validate(true)) != 0 {
			t.Error("Validate(true) failed on a solved board")
		}
	})

	t.Run("Sparse Board with no conflicts", func(t *testing.T) {
		var board Sudoku
		board[0][0] = 1
		board[1][1] = 2
		// Should be valid (checkSolved=false)
		if len(board.Validate(false)) != 0 {
			t.Error("Validate(false) should pass sparse valid board")
		}
		// Should fail (checkSolved=true) because 79 cells are empty
		if len(board.Validate(true)) == 0 {
			t.Error("Validate(true) should fail sparse board due to empty cells")
		}
	})
}
