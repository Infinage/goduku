package main

import (
	"flag"
	"fmt"
	"goduku/core"
	"io"
	"net/http"
	"os"
	"strings"

	webview "github.com/webview/webview_go"
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
		startGUI()
		return
	}

	switch flag.Arg(0) {
	case "serve":
		startWebServer(flag.Args()[1:])

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

// Run as:
// WEBKIT_DISABLE_COMPOSITING_MODE=1 ./goduku
func startGUI() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Goduku")
	w.SetSize(480, 320, webview.Hint(webview.HintNone))

	// --- Bindings ---
	w.Bind("generate", func(string) map[string]any {
		board, err := core.GenerateSudoku()
		return core.Response(board.String(false), err)
	})

	w.Bind("validate", func(csv string) map[string]any {
		board, err := core.NewSudokuFromString(csv)
		if err != nil {
			return core.Response("", err)
		}

		var results []any
		errIndices := board.Validate(false)
		for _, idx := range errIndices {
			results = append(results, map[string]any{"row": idx.Row, "col": idx.Col})
		}

		return core.Response(results, nil)
	})

	w.Bind("solve", func(csv string) map[string]any {
		board, err := core.NewSudokuFromString(csv)
		if err != nil {
			return core.Response("", err)
		}

		if board.Solve() {
			return core.Response(board.String(false), nil)
		} else {
			return core.Response("", fmt.Errorf("No solution found"))
		}
	})

	htmlRaw, err := os.ReadFile("./ui/index.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read index.html asset\n")
		os.Exit(1)
	}

	var html string = strings.Replace(string(htmlRaw), "const isDesktop = false", "const isDesktop = true", 1)
	w.SetHtml(html)

	w.Run()
}

func startWebServer(args []string) {
	subCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	port := subCmd.String("p", "8080", "Port to serve on")
	subCmd.Parse(args)

	// Ensure the web directory exists
	dir := "./ui"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: directory %s not found.\n", dir)
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Printf("Goduku Web UI available at http://localhost:%s\n", *port)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
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
	if subCmd.NArg() >= 1 {
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
	if subCmd.NArg() >= 1 {
		filePath = subCmd.Arg(0)
	}

	var board = handleSudokuRead(filePath)
	if len(board.Validate(*checkSolved)) > 0 {
		status := "invalid"
		if *checkSolved {
			status = "incomplete or invalid"
		}
		fmt.Fprintf(os.Stderr, "The board is %s.\n", status)
		os.Exit(1)
	}
	fmt.Println("Board is valid")
}
