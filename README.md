# sudoku
--
    import "github.com/MJKWoolnough/sudoku"

Package sudoku is a generic solver for sudoku type puzzles. It can solve a puzzle of any shape and size.

## Usage

#### func  MakeBox

```go
func MakeBox(gridWidth, gridHeight, boxWidth, boxHeight int) []Constraint
```
MakeBox is a helper function to make it easier to create the sections in
standard rectangular puzzles

#### func  Solve

```go
func Solve(data []int, chars int, constraints []Constraint) bool
```
Solve allows the creation of a non-standard Sudoku puzzle and solves it.

data is the puzzle information, layed out left to right, then top to bottom

chars is any valid 'character' the puzzle uses - 0 is used for an unfilled space

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
Solve4 is a sudoku puzzle of the 4x4 format

#### func  Solve9

```go
func Solve9(data []int) bool
```
Solve9 is a sudoku puzzle of the standard 9x9 format

#### type Constraint

```go
type Constraint interface {
	Constrain(*Sudoku, int, []bool) bool
}
```

Constraint defines the interface through which the character constraints are
processed

#### type Sorted

```go
type Sorted []int
```

Sorted is a Constraint that requires the numbers to be ordered lowest to
highest.

#### func (Sorted) Constrain

```go
func (s Sorted) Constrain(su *Sudoku, pos int, marked []bool) bool
```
Constrain implements the Constraing interface.

#### type Sudoku

```go
type Sudoku struct {
}
```

Sudoku puzzle information

#### func (*Sudoku) Chars

```go
func (s *Sudoku) Chars() int
```
Chars returns the number of different characters used in the puzzle

#### func (*Sudoku) Pos

```go
func (s *Sudoku) Pos(i int) int
```
Pos return the character at the given position

#### func (*Sudoku) Solve

```go
func (s *Sudoku) Solve() bool
```
Solve will solve any solveable puzzle and return whether is was sucessful.

#### type Unique

```go
type Unique []int
```

Unique represents a section of data where each cell needs to be different

#### func (Unique) Constrain

```go
func (u Unique) Constrain(s *Sudoku, pos int, marked []bool) bool
```
Constrain implements the Constraint interface

#### type UniqueSum

```go
type UniqueSum struct {
}
```

UniqueSum is a Constraint that requires all numbers to be unique and total a
specific amount.

#### func (UniqueSum) Constrain

```go
func (u UniqueSum) Constrain(s *Sudoku, pos int, marked []bool) bool
```
Constrain implements the Constraint interface.
