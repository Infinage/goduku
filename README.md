# Goduku

A fast, lightweight Sudoku toolkit with CLI, Web (WASM), and Desktop support — all powered by a single Go engine.

🌐 Live: http://goduku.infinage.space

---

## Features

* 🎲 Generate valid Sudoku puzzles (unique solutions)
* ✅ Validate board state in real-time
* 🧠 Solve any valid Sudoku instantly
* 🖥️ Desktop GUI (via webview)
* 🌍 Browser support using WebAssembly
* ⚡ Minimal dependencies, high performance

---

## Installation

```bash
git clone https://github.com/yourusername/goduku
cd goduku
go build -o goduku
```

---

## Usage

### Launch Desktop App

```bash
./goduku
```

### Run Web Server

```bash
./goduku serve -p 8080
```

Then open: http://localhost:8080

---

### Generate Puzzle

```bash
./goduku generate
```

### Solve Puzzle

```bash
./goduku solve < file.txt
```

### Validate Puzzle

```bash
./goduku validate < file.txt
```

---

## Input Format

* Accepts CSV or raw digits
* Use `0` for empty cells
* Must contain exactly 81 cells

Example:

```
530070000600195000098000060...
```

---

## Architecture

* `core/` → Sudoku engine (solver, generator, validator)
* `ui/` → Web UI + WASM bindings
* `main.go` → CLI + Desktop entrypoint

Single engine reused across:

* CLI
* WebAssembly
* Desktop GUI

---

## Build WASM (already part of version control)

```bash
cp $(go env GOROOT)/lib/wasm/wasm_exec.js ui/
GOOS=js GOARCH=wasm go build -o ui/main.wasm ui/main.go
```

---

## Tech Stack

* Go (core logic)
* WebAssembly (browser support)
* Webview (desktop UI)
* Vanilla JS (frontend)

---

## Notes

* Puzzles are generated with guaranteed unique solutions
* Solver uses backtracking with pruning
* Designed for speed and portability

---

## License

MIT
