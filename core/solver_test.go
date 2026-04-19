package core

import "testing"

func TestSudoku_Solve(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		isSolve bool
	}{
		{
			name:    "Solvable Partial",
			isSolve: true,
			input: "0,7,0,0,0,0,0,8,0\n" +
				"0,3,0,7,6,2,0,0,1\n" +
				"0,0,1,9,8,0,0,0,0\n" +
				"1,0,0,0,0,0,0,0,0\n" +
				"8,0,3,0,0,0,0,0,2\n" +
				"0,0,6,0,0,0,0,0,8\n" +
				"0,0,0,0,3,1,6,0,0\n" +
				"5,0,0,2,4,9,0,1,0\n" +
				"0,1,0,0,0,0,0,9,0",
		},
		{
			name:    "Immediate Conflict",
			isSolve: false,
			input: "1,1,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0",
		},
		{
			name:    "Already Solved",
			isSolve: true,
			input: "3,6,9,2,1,5,4,8,7\n" +
				"2,7,1,4,8,3,6,9,5\n" +
				"4,8,5,7,6,9,2,3,1\n" +
				"8,1,7,9,2,6,3,5,4\n" +
				"9,2,4,5,3,1,7,6,8\n" +
				"5,3,6,8,7,4,1,2,9\n" +
				"6,4,3,1,9,8,5,7,2\n" +
				"7,5,8,6,4,2,9,1,3\n" +
				"1,9,2,3,5,7,8,4,6",
		},
		{
			name:    "In-Place Persistence",
			isSolve: true,
			input: "0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := NewSudokuFromString(tt.input)
			if got := board.Solve(); got != tt.isSolve {
				t.Fatalf("Solve() = %v, want %v", got, tt.isSolve)
			}

			if tt.isSolve && len(board.Validate(true)) > 0 {
				t.Error("Board marked solved but is logically invalid/incomplete")
			}
		})
	}
}

func TestBacktrackContext_ExitAfter(t *testing.T) {
	for _, n := range []uint8{1, 2} {
		t.Run(string(n+'0')+" Solutions", func(t *testing.T) {
			var board Sudoku
			ctx := backtrackContext{exitAfter: n}
			backtrack(&board, 0, 0, &ctx)
			if len(ctx.solutions) != int(n) {
				t.Errorf("Expected %d solutions, got %d", n, len(ctx.solutions))
			}
		})
	}
}
