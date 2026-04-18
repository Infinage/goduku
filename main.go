package main

import (
	"flag"
	"fmt"
	"goduku/core"
	"io"
	"os"
)

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
	if flag.NArg() < 1 {
		fmt.Println("Starting GUI (No command provided)")
		return
	}

	switch flag.Arg(0) {
	case "serve":
		fmt.Println("Starting the Web UI")

	case "generate":
		handleGenerate(flag.Args()[1:])

	case "validate":
		handleValidate(flag.Args()[1:])

	case "solve":
		handleSolve(flag.Args()[1:])

	default:
		flag.Usage()
		os.Exit(1)
	}
}

func handleGenerate(args []string) {
	subCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	pretty := subCmd.Bool("p", false, "Pretty print output")
	subCmd.Parse(args)

	board, err := core.GenerateSudoku()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating puzzle: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(board.String(*pretty))
}

func handleSudokuRead(filePath string) core.Sudoku {
	var input []byte
	var err error

	// Read string as CSV string from stdin / filepath
	if filePath == "" {
		input, err = io.ReadAll(os.Stdin)
	} else {
		input, err = os.ReadFile(filePath)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin/file: %v\n", err)
		os.Exit(1)
	}

	var board core.Sudoku
	board, err = core.NewSudokuFromString(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid board: %v\n", err)
		os.Exit(1)
	}

	return board
}

func handleSolve(args []string) {
	subCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	pretty := subCmd.Bool("p", false, "Pretty print output")
	subCmd.Parse(args)

	var filePath string
	if subCmd.NArg() > 1 {
		filePath = subCmd.Arg(0)
	}

	var board = handleSudokuRead(filePath)
	if board.Solve() {
		fmt.Println(board.String(*pretty))
	} else {
		fmt.Fprintln(os.Stderr, "No solution exists for this board.")
		os.Exit(1)
	}
}

func handleValidate(args []string) {
	subCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	checkSolved := subCmd.Bool("solved", false, "Require all cells to be filled")
	subCmd.Parse(args)

	var filePath string
	if subCmd.NArg() > 1 {
		filePath = subCmd.Arg(0)
	}

	var board = handleSudokuRead(filePath)
	if !board.Validate(*checkSolved) {
		status := "invalid"
		if *checkSolved {
			status = "incomplete or invalid"
		}
		fmt.Fprintf(os.Stderr, "The board is %s.\n", status)
		os.Exit(1)
	}
	fmt.Println("Board is valid")
}
