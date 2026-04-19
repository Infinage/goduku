package core

import (
	"testing"
)

func TestGenerateSudoku(t *testing.T) {
	t.Run("Generate and Solve Flow", func(t *testing.T) {
		board, err := GenerateSudoku()
		if err != nil {
			t.Fatalf("Failed to generate: %v", err)
		}

		hasHoles := false
		for r := range 9 {
			for c := range 9 {
				if board[r][c] == 0 {
					hasHoles = true
					break
				}
			}
		}
		if !hasHoles {
			t.Error("Generator returned a completely full board")
		}

		if !board.Solve() {
			t.Error("Generated puzzle is reported as unsolvable")
		}

		if len(board.Validate(true)) > 0 {
			t.Error("The solution provided by Solve() is invalid")
		}
	})

	t.Run("Uniqueness Check", func(t *testing.T) {
		board, _ := GenerateSudoku()

		// Run backtrack with exitAfter: 2 to verify uniqueness
		ctx := backtrackContext{exitAfter: 2, shuffle: false}
		backtrack(&board, 0, 0, &ctx)

		if len(ctx.solutions) != 1 {
			t.Errorf("Generator produced a puzzle with %d solutions; want 1", len(ctx.solutions))
		}
	})
}
