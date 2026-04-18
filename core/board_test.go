package core

import (
	"strings"
	"testing"
)

func TestNewSudokuFromString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "Valid Minimal Board",
			expectErr: false,
			input: "1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9\n" +
				"1,2,3,4,5,6,7,8,9",
		},
		{
			name:      "Invalid Character",
			expectErr: true,
			input:     "1,2,A,4,5,6,7,8,9...",
		},
		{
			name:      "Multiple Character Cells",
			expectErr: true,
			input: "1,2,3,4,5,6,7,8,9\n" +
				"12,3,4,5,6,7,8,9,0",
		},
		{
			name:      "Empty Input",
			expectErr: true,
			input:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSudokuFromString(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewSudokuFromString() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestSudoku_String(t *testing.T) {
	var board Sudoku
	board[0][0] = 1
	board[0][1] = 2
	board[8][8] = 9

	t.Run("CSV Format", func(t *testing.T) {
		got := board.String(false)

		if !strings.HasPrefix(got, "1,2") {
			t.Errorf("CSV Prefix mismatch. Got: %q", got)
		}

		if !strings.HasSuffix(got, "9") {
			t.Errorf("CSV Suffix mismatch. Got: %q", got)
		}

		idx := strings.IndexFunc(got, func(r rune) bool {
			return r != ',' && r != '\n' && (r < '0' || r > '9')
		})

		if idx != -1 {
			t.Errorf("Found invalid character %q at index %d, Got: %q", got[idx], idx, got)
		}
	})

	t.Run("Pretty Format", func(t *testing.T) {
		got := board.String(true)
		if !strings.Contains(got, "------+-------+------") {
			t.Errorf("Pretty String missing divider:\n%s", got)
		}

		// Check for the pipe symbol
		if !strings.Contains(got, "|") {
			t.Errorf("Pretty String missing pipe separators")
		}
	})
}
