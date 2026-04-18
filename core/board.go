package core

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Sudoku [9][9]uint8

// Each cell is seperated by a comma char
// Each row is seperated by a newline char
func NewSudokuFromString(raw string) (Sudoku, error) {
	if raw == "" {
		return Sudoku{}, fmt.Errorf("Empty input string")
	}

	var board Sudoku
	for row, line := range strings.Split(raw, "\n") {
		for col, cell := range strings.Split(line, ",") {
			if row >= 9 || col >= 9 {
				return board, fmt.Errorf("Only supports 9x9 Sudoku inputs")
			}

			if cell == "" {
				cell = "0"
			}

			if len(cell) != 1 || cell[0] < '0' || cell[0] > '9' {
				return board, fmt.Errorf("Invalid value for cell @ [%v, %v], got: %v", row, col, cell)
			}

			board[row][col] = uint8(cell[0] - '0')
		}
	}

	return board, nil
}

// Read from a file
func NewSudokuFromCSV(path string) (Sudoku, error) {
	f, err := os.Open(path)
	if err != nil {
		return Sudoku{}, fmt.Errorf("Failed to open sudoku input file %s: %w", path, err)
	}

	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return Sudoku{}, fmt.Errorf("Error reading file: %w", err)
	}

	return NewSudokuFromString(string(raw))
}

// Return board as a csv string
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
