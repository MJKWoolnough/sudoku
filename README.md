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

#### type Sudoku

```go
type Sudoku struct {
}
```

Sudoku puzzle information

#### func  NewSudoku

```go
func NewSudoku(data, chars []int, structure [][]int) *Sudoku
```
NewSudoku creates a custom puzzle solver.

data is layed out left to right, then top to bottom

chars is any valid 'character' the puzzle uses - 0 is used for an empty space

structure is a slice of sections, each of which is a slice of positions,
len(chars) in length, which describes the rows, columns, boxes or other shapes
in which there can only be one of each character

#### func  NewSudoku4

```go
func NewSudoku4(data []int) *Sudoku
```
A sudoku puzzle of the 4x4 format

#### func  NewSudoku9

```go
func NewSudoku9(data []int) *Sudoku
```
A sudoku puzzle of the standard 9x9 format

#### func (*Sudoku) Solve

```go
func (s *Sudoku) Solve() bool
```
Solve will solve any solveable puzzle and return whether is was sucessful.

The solution is stored in the data slice initially given to any NewSudoku*
function.
