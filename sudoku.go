// Package sudoku is a generic solver for sudoku type puzzles. It can solve a puzzle of any shape and size.
package sudoku

import "fmt"

var (
	s4 = []Constraint{
		Unique{0, 1, 2, 3},
		Unique{4, 5, 6, 7},
		Unique{8, 9, 10, 11},
		Unique{12, 13, 14, 15},
		Unique{0, 4, 8, 12},
		Unique{1, 5, 9, 13},
		Unique{2, 6, 10, 14},
		Unique{3, 7, 11, 15},
		Unique{0, 1, 4, 5},
		Unique{2, 3, 6, 7},
		Unique{8, 9, 12, 13},
		Unique{10, 11, 14, 15},
	}
	c4 = []int{1, 2, 3, 4}
	s9 = []Constraint{
		Unique{0, 1, 2, 3, 4, 5, 6, 7, 8},
		Unique{9, 10, 11, 12, 13, 14, 15, 16, 17},
		Unique{18, 19, 20, 21, 22, 23, 24, 25, 26},
		Unique{27, 28, 29, 30, 31, 32, 33, 34, 35},
		Unique{36, 37, 38, 39, 40, 41, 42, 43, 44},
		Unique{45, 46, 47, 48, 49, 50, 51, 52, 53},
		Unique{54, 55, 56, 57, 58, 59, 60, 61, 62},
		Unique{63, 64, 65, 66, 67, 68, 69, 70, 71},
		Unique{72, 73, 74, 75, 76, 77, 78, 79, 80},
		Unique{0, 9, 18, 27, 36, 45, 54, 63, 72},
		Unique{1, 10, 19, 28, 37, 46, 55, 64, 73},
		Unique{2, 11, 20, 29, 38, 47, 56, 65, 74},
		Unique{3, 12, 21, 30, 39, 48, 57, 66, 75},
		Unique{4, 13, 22, 31, 40, 49, 58, 67, 76},
		Unique{5, 14, 23, 32, 41, 50, 59, 68, 77},
		Unique{6, 15, 24, 33, 42, 51, 60, 69, 78},
		Unique{7, 16, 25, 34, 43, 52, 61, 70, 79},
		Unique{8, 17, 26, 35, 44, 53, 62, 71, 80},
		Unique{0, 1, 2, 9, 10, 11, 18, 19, 20},
		Unique{3, 4, 5, 12, 13, 14, 21, 22, 23},
		Unique{6, 7, 8, 15, 16, 17, 24, 25, 26},
		Unique{27, 28, 29, 36, 37, 38, 45, 46, 47},
		Unique{30, 31, 32, 39, 40, 41, 48, 49, 50},
		Unique{33, 34, 35, 42, 43, 44, 51, 52, 53},
		Unique{54, 55, 56, 63, 64, 65, 72, 73, 74},
		Unique{57, 58, 59, 66, 67, 68, 75, 76, 77},
		Unique{60, 61, 62, 69, 70, 71, 78, 79, 80},
	}
	c9 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

// MakeBox is a helper function to make it easier to create the sections in standard rectangular puzzles
func MakeBox(gridWidth, gridHeight, boxWidth, boxHeight int) []Constraint {
	toRet := make([]Constraint, 0, gridWidth/boxWidth*gridHeight/boxHeight)
	for j := 0; j < gridHeight; j += boxHeight {
		for i := 0; i < gridWidth; i += boxWidth {
			thisGrid := make(Unique, 0, boxWidth*boxHeight)
			for y := 0; y < boxHeight; y++ {
				for x := 0; x < boxWidth; x++ {
					thisGrid = append(thisGrid, i+x+(j+y)*gridWidth)
				}
			}
			toRet = append(toRet, thisGrid)
		}
	}
	return toRet
}

// Solve9 is a sudoku puzzle of the standard 9x9 format
func Solve9(data []int) bool {
	if len(data) != 81 {
		return false
	}
	s := Sudoku{
		data:        data,
		chars:       c9,
		constraints: s9,
	}
	return s.Solve()
}

// Solve4 is a sudoku puzzle of the 4x4 format
func Solve4(data []int) bool {
	if len(data) != 16 {
		return false
	}
	s := Sudoku{
		data:        data,
		chars:       c4,
		constraints: s4,
	}
	return s.Solve()
}

// Solve allows the creation of a non-standard Sudoku puzzle and solves it.
//
// data is the puzzle information, layed out left to right, then top to bottom
//
// chars is any valid 'character' the puzzle uses - 0 is used for an unfilled space
//
// structure is a slice of sections, each of which is a slice of positions, len(chars)
// in length, which describes the rows, columns, boxes or other shapes in which there
// can only be one of each character
//
// Will return true if puzzle is solveable and the solution will be stored in the data slice.
// Upon a failure, will return false and the data slice will be as original.
func Solve(data, chars []int, constraints []Constraint) bool {
	s := Sudoku{data, chars, constraints}
	return s.Solve()
}

func slicePos(s []int, p int) int {
	for n, sp := range s {
		if sp == p {
			return n
		}
	}
	return -1
}

type Constraint interface {
	Constrain(*Sudoku, int, []bool) bool
}

type Unique []int

func (u Unique) Constrain(s *Sudoku, pos int, marked []bool) bool {
	if slicePos([]int(u), pos) == -1 {
		return true
	}
	myMark := make([]int, 0, len(u))
	for _, p := range u {
		if mp := s.data[p]; mp != 0 {
			if slicePos(myMark, mp) != -1 {
				fmt.Println(u, pos, myMark, mp)
				return false
			}
			myMark = append(myMark, mp)
			marked[slicePos(s.chars, mp)] = true
		}
	}
	return true
}

/*type Sorted []int

func (s Sorted) Constrain(su *Sudoku, pos int, marked []bool) bool {
	if slicePos([]int(s), pos) == -1 {
		return true
	}
	num := 0
	for _, p := range s {
		if mp := su.data[p]; mp != 0 {
			if mp > num {
				num = mp
			} else {
				return false
			}
		} else if mp == pos {

		}
	}
	return true
}

type UniqueSum struct {
	positions []int
	total     int
}

func (u UniqueSum) Constrain(s *Sudoku, pos int, marked []bool) bool {
	if slicePos(u.positions, pos) == -1 {
		return true
	}
	total := 0
	myMark := make([]int, len(u.positions))
	unmarked := 0
	for _, p := range u.positions {
		if mp := s.data[p]; mp == 0 {
			unmarked++
		} else {
			if slicePos(myMark, mp) != -1 {
				return false
			}
			myMark = append(myMark, mp)
			total += mp
		}
	}
	left := u.total - total
	if left < 0 {
		return false
	}

	return true
}*/

// Sudoku puzzle information
type Sudoku struct {
	data        []int
	chars       []int
	constraints []Constraint
}

func (s *Sudoku) possible(pos int) []int {
	if pos < 0 || pos > len(s.data) || s.data[pos] != 0 {
		return nil
	}
	marked := make([]bool, len(s.chars))
	for _, c := range s.constraints {
		any := c.Constrain(s, pos, marked)
		if !any {
			return []int{}
		}
	}
	toRet := make([]int, 0, len(s.chars))
	for p, m := range marked {
		if !m {
			toRet = append(toRet, s.chars[p])
		}
	}
	return toRet
}

// Solve will solve any solveable puzzle and return whether is was sucessful.
func (s *Sudoku) Solve() bool {
	l := len(s.data)
	possibilities := make([][]int, l)
	var pos int
	for ; pos < l; pos++ {
		if p := s.possible(pos); p != nil {
			possibilities[pos] = p
			break
		}
	}
	for pos < l {
		if len(possibilities[pos]) == 0 {
			s.data[pos] = 0
			for pos--; pos >= 0 && len(possibilities[pos]) == 0; pos-- {
				if possibilities[pos] != nil {
					s.data[pos] = 0
				}
			}
			if pos < 0 {
				return false
			}
			continue
		}
		s.data[pos] = possibilities[pos][0]
		possibilities[pos] = possibilities[pos][1:]
		for pos++; pos < l; pos++ {
			if p := s.possible(pos); p != nil {
				possibilities[pos] = p
				break
			}
		}
	}
	return true
}
