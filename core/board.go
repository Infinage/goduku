package core

import (
	"fmt"
	"strings"
)

type Sudoku [9][9]uint8

// Parses a raw string into a Sudoku board.
// It supports CSV format with optional newlines
// It ignores commas and newlines, expecting exactly 81 numeric digits.
func NewSudokuFromString(raw string) (Sudoku, error) {
	if raw == "" {
		return Sudoku{}, fmt.Errorf("Empty input string")
	}

	var board Sudoku
	var validBoard bool

	// Normalize input by removing formatting characters
	raw = strings.ReplaceAll(raw, "\n", "")
	raw = strings.ReplaceAll(raw, ",", "")
	raw = strings.TrimSpace(raw)

	for idx, cell := range raw {
		row, col := idx/9, idx%9

		if idx >= 81 {
			return board, fmt.Errorf("Only 9x9 Sudoku inputs supported")
		}

		if cell < '0' || cell > '9' {
			return board, fmt.Errorf("Invalid value for cell @ [%v, %v], got: %v", row, col, cell)
		}

		board[row][col] = uint8(cell - '0')

		if idx == 80 {
			validBoard = true
		}

	}

	if !validBoard {
		return board, fmt.Errorf("Only 9x9 Sudoku inputs supported")
	}

	return board, nil
}

// String returns the board as a string.
// If pretty is true, it returns a human-readable grid with 3x3 block dividers.
// If pretty is false, it returns a compact CSV format suitable for storage or WASM.
func (board Sudoku) String(pretty bool) string {
	var buffer []string
	for row, line := range board {
		for col, cell := range line {
			cellRaw := string(cell + '0')
			buffer = append(buffer, cellRaw)

			if pretty {
				if col == 2 || col == 5 {
					buffer = append(buffer, " |")
				}
			}

			if col < 8 {
				if pretty {
					buffer = append(buffer, " ")
				} else {
					buffer = append(buffer, ",")
				}
			}

		}

		if row < 8 {
			buffer = append(buffer, "\n")
			if pretty && (row == 2 || row == 5) {
				buffer = append(buffer, "------+-------+------\n")
			}
		}
	}
	return strings.Join(buffer, "")
}
