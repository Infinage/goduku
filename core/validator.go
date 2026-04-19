package core

import (
	"iter"
)

func (board *Sudoku) Grid(r, c int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		startRow, startCol := (r/3)*3, (c/3)*3
		for row := startRow; row < startRow+3; row++ {
			for col := startCol; col < startCol+3; col++ {
				if !yield(row, col) {
					return
				}
			}
		}
	}
}

// Validate entry for a single cell
// Returns false for zeroed and cells with invalid no entries
func (board *Sudoku) validate(row, col int) bool {
	if board[row][col] < 1 || board[row][col] > 9 {
		return false
	}

	// Check for repeats across row
	for x := range 9 {
		if x == row {
			continue
		}
		if board[x][col] == board[row][col] {
			return false
		}
	}

	// Check for repeats across col
	for y := range 9 {
		if y == col {
			continue
		}
		if board[row][y] == board[row][col] {
			return false
		}
	}

	// Check for repeats across grid
	for x, y := range board.Grid(row, col) {
		if x == row && y == col {
			continue
		}
		if board[x][y] == board[row][col] {
			return false
		}
	}

	return true
}

// Validate that the entire board is 'valid' by checking constraints
// Returns a list of cells that are unfilled (and unsolved if `checkSolved` is set)
func (board *Sudoku) Validate(checkSolved bool) []Index {
	var errors []Index
	for row := range 9 {
		for col := range 9 {
			if board[row][col] == 0 && checkSolved {
				errors = append(errors, Index{uint8(row), uint8(col)})
			} else if board[row][col] != 0 && !board.validate(row, col) {
				errors = append(errors, Index{uint8(row), uint8(col)})
			}
		}
	}
	return errors
}
