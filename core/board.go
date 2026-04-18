package core

import (
	"fmt"
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
	var validBoard bool

	raw = strings.Trim(raw, "\n")
	for row, line := range strings.Split(raw, "\n") {
		for col, cell := range strings.Split(line, ",") {
			if row == 8 && col == 8 {
				validBoard = true
			}

			if row >= 9 || col >= 9 {
				return board, fmt.Errorf("Only 9x9 Sudoku inputs supported")
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

	if !validBoard {
		return board, fmt.Errorf("Only 9x9 Sudoku inputs supported")
	}

	return board, nil
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
