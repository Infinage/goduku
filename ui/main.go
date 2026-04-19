// $ cp $(go env GOROOT)/lib/wasm/wasm_exec.js ui/wasm_exec.js
// $ GOOS=js GOARCH=wasm go build -o ui/main.wasm ui/main.go

// $ cp /opt/tinygo/targets/wasm_exec.js ui/wasm_exec.js
// $ /opt/tinygo/bin/tinygo build -o ui/main.wasm -target wasm -no-debug ui/main.go
package main

import (
	"fmt"
	"goduku/core"
	"syscall/js"
)

func main() {
	// Expose Generate
	js.Global().Set("generate", js.FuncOf(func(this js.Value, args []js.Value) any {
		board, err := core.GenerateSudoku()
		return core.Response(board.String(false), err)
	}))

	// Expose Validate
	js.Global().Set("validate", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return core.Response("", fmt.Errorf("Missing board input"))
		}

		input := args[0].String()
		board, err := core.NewSudokuFromString(input)
		if err != nil {
			return core.Response("", err)
		}

		var results []any
		errIndices := board.Validate(false)
		for _, idx := range errIndices {
			results = append(results, map[string]any{"row": idx.Row, "col": idx.Col})
		}

		return core.Response(results, nil)
	}))

	// Expose Solve
	js.Global().Set("solve", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return core.Response("", fmt.Errorf("Missing board input"))
		}

		input := args[0].String()
		board, err := core.NewSudokuFromString(input)
		if err != nil {
			return core.Response("", err)
		}

		if board.Solve() {
			return core.Response(board.String(false), nil)
		} else {
			return core.Response("", fmt.Errorf("No solution found"))
		}

	}))

	// Keep alive
	select {}
}
