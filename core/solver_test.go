package core

import "testing"

func TestSudoku_Solve(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		isSolve bool
	}{
		{
			name: "Solvable Partial Board",
			input: "0,7,0,0,0,0,0,8,0\n" +
				"0,3,0,7,6,2,0,0,1\n" +
				"0,0,1,9,8,0,0,0,0\n" +
				"1,0,0,0,0,0,0,0,0\n" +
				"8,0,3,0,0,0,0,0,2\n" +
				"0,0,6,0,0,0,0,0,8\n" +
				"0,0,0,0,3,1,6,0,0\n" +
				"5,0,0,2,4,9,0,1,0\n" +
				"0,1,0,0,0,0,0,9,0",
			isSolve: true,
		},
		{
			name: "Unsolvable Board (Immediate Conflict)",
			input: "1,1,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0\n" +
				"0,0,0,0,0,0,0,0,0",
			isSolve: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, err := NewSudokuFromString(tt.input)
			if err != nil {
				t.Fatalf("Failed to parse board: %v", err)
			}

			got := board.Solve()
			if got != tt.isSolve {
				t.Errorf("Solve() = %v, want %v", got, tt.isSolve)
			}

			if got && !board.Validate(true) {
				t.Error("Solve() returned true but Validate(true) failed")
			}
		})
	}
}
