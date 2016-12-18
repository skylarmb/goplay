package main

import (
  "fmt"
  "time"
  s "strings"
)

func main() {
  board := &Board { w: 40, h: 40 }
  populate(board)
  for true {
    printBoard (board)
    evolve (board)
    time.Sleep(100 * time.Millisecond)
  }
}

// our game board
type Board struct {
  w, h int
  cells [][]bool
}

// populate our board with some initial state
func populate (board *Board) {
  cells := make([][]bool, board.h)
  for i := range cells {
    cells[i] = make([]bool, board.w)
  }
  // hardcoded for now
  cells[3][3] = true
  cells[3][4] = true
  cells[4][2] = true
  cells[4][3] = true
  cells[5][3] = true
  cells[5][4] = true
  board.cells = cells
}

// create a clone of a set of cells
func clone (cells [][]bool) [][]bool {
  clone := make([][]bool, len(cells))
  for row := range clone {
    clone[row] = make([]bool, len(cells[row]))
    for col := range clone[row] {
      clone[row][col] = cells[row][col]
    }
  }
  return clone
}

// evolves the board
func evolve (board *Board) {
  // copy current board cells in order
  // to calculate the next generation of cells
  evolved := clone(board.cells)
  for row := range evolved {
    for col := range evolved[row] {
      evolved[row][col] = next(board, row, col)
    }
  }
  board.cells = evolved
}

//determine if a cell is alive or dead in the next generation
func next(board *Board, x, y int) bool {
  neighbors := 0
  for i := -1; i <= 1; i++ {
    for j := -1; j <= 1; j++ {
      if i == 0 && j == 0 {
        continue
      } else if alive(board, x+i, y+j) {
        neighbors ++
      }
    }
  }
  // if 3 neighbors, its alive
  // if 2 neighbors, it maintains it current state
  // otherwise its dead
  return neighbors == 3 || neighbors == 2 && alive(board, x, y)
}

// determine if a cell at a given coord is alive or dead
// if the x or y coordinates are outside the game board
// they are wrapped toroidally.
func alive (board *Board, x, y int) bool {
  x += board.w
  x %= board.w
  y += board.h
  y %= board.h
  return board.cells[x][y]
}

func printBoard(board *Board) {
  // overwrite existing terminal content
  fmt.Printf("\033[0;0H")
  // top board edge
  fmt.Print(s.Repeat("\u2581", board.w + 2))
  fmt.Println()
  for x := range board.cells {
    // left board edge
    fmt.Print("\u258F")
    for y := range board.cells[x] {
      // anonymous functions ftw
      cell := func() string {
        if alive(board, x, y) {
          return "\u25EF"
        } else {
          return " "
        }
      }()
      fmt.Print(cell)
    }
    // right board edge
    fmt.Print("\u2595")
    fmt.Println()
  }
  // bottom board edge
  fmt.Print(s.Repeat("\u2594", board.w + 2))
  fmt.Println()
}
