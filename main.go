package main

import (
	"flag"
	"fmt"
	"goduku/core"
)

// Command line usage:
// `./goduku` 	  	   => Start the GUI
// `./goduku serve`    => Start the static webserver
// `./goduku generate` => Generate a sudoku puzzle
// `./goduku validate` => Validate the input board
// `./goduku solve`    => Solve the input puzzle
func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, `Usage: goduku [command]

A Sudoku toolkit with built-in Web UI.

Commands:
  (no command)  Launch the desktop GUI
  serve         Start the static webserver for the Web UI
  generate      Create a new puzzle and output to stdout
  validate      Check if a board is valid and calculate difficulty
  solve         Find a solution for the provided board
		`)
	}

	flag.Parse()
	var cmd string
	if flag.NArg() == 1 {
		cmd = flag.Args()[0]
	}

	switch cmd {
	case "":
		fmt.Println("Starting the GUI")
	case "serve":
		fmt.Println("Starting the Web UI")
	case "generate":
		fmt.Println("Generating a puzzle")
	case "solve":
		fmt.Println("Solving the puzzle")
	}

	board, _ := core.NewSudokuFromString(
		"1,3,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0\n" +
			"0,0,0,0,0,0,0,0,0",
	)
	if !board.Solve() {
		fmt.Println("Something is wrong")
		return
	}

	fmt.Println(board.String(true))
}
