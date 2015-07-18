package sudoku

import "testing"

func TestKiller(t *testing.T) {
	t.Parallel()
	tests := []struct {
		start, solution []int
		solveable       bool
		chars           int
		constraints     []Constraint
	}{
		{make([]int, 81),
			[]int{
				6, 2, 9, 1, 4, 3, 8, 5, 7,
				3, 5, 7, 9, 6, 8, 4, 2, 1,
				1, 8, 4, 7, 5, 2, 6, 9, 3,
				2, 9, 5, 3, 1, 4, 7, 8, 6,
				7, 3, 6, 2, 8, 9, 1, 4, 5,
				8, 4, 1, 6, 7, 5, 2, 3, 9,
				9, 1, 8, 4, 3, 7, 5, 6, 2,
				4, 7, 3, 5, 2, 6, 9, 1, 8,
				5, 6, 2, 8, 9, 1, 3, 7, 4,
			},
			true,
			9,
			append(s9,
				UniqueSum{[]int{0, 1, 2, 11}, 24},
				UniqueSum{[]int{3, 4, 5}, 8},
				UniqueSum{[]int{6, 7, 8}, 20},
				UniqueSum{[]int{9, 10, 18}, 9},
				UniqueSum{[]int{12, 21}, 16},
				UniqueSum{[]int{13, 14, 15}, 18},
				UniqueSum{[]int{16, 17}, 3},
				UniqueSum{[]int{19, 20, 28, 29}, 26},
				UniqueSum{[]int{22, 30, 31}, 9},
				UniqueSum{[]int{23, 32, 40, 41}, 23},
				UniqueSum{[]int{24, 25, 33}, 22},
				UniqueSum{[]int{26, 34, 35}, 17},
				UniqueSum{[]int{27, 36, 45, 54}, 26},
				UniqueSum{[]int{37, 46, 55}, 8},
				UniqueSum{[]int{38, 39}, 8},
				UniqueSum{[]int{42, 43, 51, 60}, 12},
				UniqueSum{[]int{44, 53}, 14},
				UniqueSum{[]int{47, 48, 49, 50}, 19},
				UniqueSum{[]int{52, 61, 62}, 11},
				UniqueSum{[]int{56, 57}, 12},
				UniqueSum{[]int{58, 67, 75, 76}, 22},
				UniqueSum{[]int{59, 68, 69}, 22},
				UniqueSum{[]int{63, 72}, 9},
				UniqueSum{[]int{64, 73}, 13},
				UniqueSum{[]int{65, 66, 74}, 10},
				UniqueSum{[]int{70, 79}, 8},
				UniqueSum{[]int{71, 80}, 12},
				UniqueSum{[]int{77, 78}, 4},
			)}}
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
