# sudoku
--
    import "github.com/MJKWoolnough/sudoku"

Sudoku is a generic solver for sudoku type puzzles. It can solve a puzzle of any shape and size.

## Usage

#### func  MakeBox

```go
func MakeBox(gridWidth, gridHeight, boxWidth, boxHeight int) [][]int
```
A helper function to make it easier to create the sections in standard
rectangular puzzles

#### func  Solve

```go
func Solve(data, chars []int, structure [][]int) bool
```
Solve allows the creation of a non-standard Sudoku puzzle and solves it.

data is layed out left to right, then top to bottom

chars is any valid 'character' the puzzle uses - 0 is used for an empty space

structure is a slice of sections, each of which is a slice of positions,
len(chars) in length, which describes the rows, columns, boxes or other shapes
in which there can only be one of each character

Will return true if puzzle is solveable and the solution will be stored in the
data slice. Upon a failure, will return false and the data slice will be as
original.

#### func  Solve4

```go
func Solve4(data []int) bool
```
A sudoku puzzle of the 4x4 format

#### func  Solve9

```go
func Solve9(data []int) bool
```
A sudoku puzzle of the standard 9x9 format
