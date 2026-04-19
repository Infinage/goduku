// $ cp $(go env GOROOT)/lib/wasm/wasm_exec.js web/
// $ GOOS=js GOARCH=wasm go build -o web/main.wasm cmd/wasm.go

package main

import (
	"fmt"
	"goduku/core"
	"syscall/js"
)

func main() {
	wrapResponse := func(data any, err error) any {
		res := map[string]any{
			"success": err == nil,
			"data":    data,
			"error":   "",
		}
		if err != nil {
			res["error"] = err.Error()
		}
		return res
	}

	// Expose Generate
	js.Global().Set("generate", js.FuncOf(func(this js.Value, args []js.Value) any {
		board, err := core.GenerateSudoku()
		return wrapResponse(board.String(false), err)
	}))

	// Expose Validate
	js.Global().Set("validate", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return wrapResponse("", fmt.Errorf("Missing board input"))
		}

		input := args[0].String()
		board, err := core.NewSudokuFromString(input)
		if err != nil {
			return wrapResponse("", err)
		}

		var results []any
		errIndices := board.Validate(false)
		for _, idx := range errIndices {
			results = append(results, map[string]any{"row": idx.Row, "col": idx.Col})
		}

		return wrapResponse(results, nil)
	}))

	// Expose Solve
	js.Global().Set("solve", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return wrapResponse("", fmt.Errorf("Missing board input"))
		}

		input := args[0].String()
		board, err := core.NewSudokuFromString(input)
		if err != nil {
			return wrapResponse("", err)
		}

		if board.Solve() {
			return wrapResponse(board.String(false), nil)
		} else {
			return wrapResponse("", fmt.Errorf("No solution found"))
		}

	}))

	// Keep alive
	select {}
}
