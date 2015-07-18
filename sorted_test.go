package sudoku

import "testing"

func TestFutoshiki(t *testing.T) {
	t.Parallel()
	g := []Constraint{
		Unique{0, 1, 2, 3, 4},
		Unique{5, 6, 7, 8, 9},
		Unique{10, 11, 12, 13, 14},
		Unique{15, 16, 17, 18, 19},
		Unique{20, 21, 22, 23, 24},
		Unique{0, 5, 10, 15, 20},
		Unique{1, 6, 11, 16, 21},
		Unique{2, 7, 12, 17, 22},
		Unique{3, 8, 13, 18, 23},
		Unique{4, 9, 14, 19, 24},
	}
	tests := []struct {
		start, solution []int
		solveable       bool
		chars           int
		constraints     []Constraint
	}{
		{
			make([]int, 25),
			[]int{
				4, 2, 3, 1, 5,
				5, 3, 1, 2, 4,
				1, 5, 4, 3, 2,
				2, 1, 5, 4, 3,
				3, 4, 2, 5, 1,
			},
			true,
			5,
			append(g,
				Sorted{1, 0, 5},
				Sorted{3, 8, 9},
				Sorted{13, 18, 17},
				Sorted{16, 15, 20},
			),
		},
	}
	for i, test := range tests {
		solved := Solve(test.start, test.chars, test.constraints)
		if test.solveable {
			if solved {
				for j, num := range test.start {
					if num != test.solution[j] {
						t.Errorf("solution found does not match solution given for test %d", i+1)
						break
					}
				}
			} else {
				t.Errorf("didn't find solution in puzzle %d when solution expected", i+1)
			}
		} else {

		}
	}
}
