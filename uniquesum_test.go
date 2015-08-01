package sudoku

import (
	"fmt"
	"testing"
)

func TestSumdoku(t *testing.T) {
	t.Parallel()
	tests := []struct {
		start, solution []int
		solveable       bool
		chars           int
		constraints     []Constraint
	}{
		{
			make([]int, 81),
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
			),
		},
		{
			make([]int, 81),
			[]int{
				2, 6, 7, 9, 5, 1, 3, 4, 8,
				3, 5, 4, 7, 2, 8, 1, 6, 9,
				8, 1, 9, 6, 3, 4, 2, 5, 7,
				9, 3, 6, 2, 4, 7, 5, 8, 1,
				5, 2, 8, 1, 9, 6, 7, 3, 4,
				4, 7, 1, 3, 8, 5, 9, 2, 6,
				1, 4, 5, 8, 7, 3, 6, 9, 2,
				6, 8, 2, 5, 1, 9, 4, 7, 3,
				7, 9, 3, 4, 6, 2, 8, 1, 5,
			},
			true,
			9,
			append(s9,
				UniqueSum{[]int{0, 9, 18, 27}, 22},
				UniqueSum{[]int{1, 2, 3, 10}, 27},
				UniqueSum{[]int{4, 5, 6}, 9},
				UniqueSum{[]int{7, 16}, 10},
				UniqueSum{[]int{8, 17}, 17},
				UniqueSum{[]int{11, 12, 21}, 17},
				UniqueSum{[]int{13, 22, 31}, 9},
				UniqueSum{[]int{14, 15, 23, 24}, 15},
				UniqueSum{[]int{19, 28}, 4},
				UniqueSum{[]int{20, 29, 38}, 23},
				UniqueSum{[]int{25, 26, 34}, 20},
				UniqueSum{[]int{30, 39}, 3},
				UniqueSum{[]int{32, 33}, 12},
				UniqueSum{[]int{35, 43, 44}, 8},
				UniqueSum{[]int{36, 37, 45, 54}, 12},
				UniqueSum{[]int{40, 41, 49}, 23},
				UniqueSum{[]int{42, 51, 60}, 22},
				UniqueSum{[]int{46, 55}, 11},
				UniqueSum{[]int{47, 48, 56}, 9},
				UniqueSum{[]int{50, 57, 58, 59}, 23},
				UniqueSum{[]int{52, 53, 61, 70}, 24},
				UniqueSum{[]int{62, 71, 79, 80}, 11},
				UniqueSum{[]int{63, 64, 72, 73}, 30},
				UniqueSum{[]int{65, 66, 67}, 8},
				UniqueSum{[]int{68, 69}, 13},
				UniqueSum{[]int{74, 75}, 7},
				UniqueSum{[]int{76, 77, 78}, 16},
			),
		},
		{
			make([]int, 81),
			[]int{
				3, 5, 4, 8, 1, 7, 2, 6, 9,
				8, 9, 2, 6, 3, 4, 7, 1, 5,
				1, 7, 6, 5, 9, 2, 8, 4, 3,
				9, 4, 8, 1, 2, 5, 6, 3, 7,
				7, 6, 5, 4, 8, 3, 1, 9, 2,
				2, 3, 1, 7, 6, 9, 4, 5, 8,
				6, 1, 9, 2, 5, 8, 3, 7, 4,
				4, 8, 3, 9, 7, 6, 5, 2, 1,
				5, 2, 7, 3, 4, 1, 9, 8, 6,
			},
			true,
			9,
			append(s9,
				UniqueSum{[]int{0, 1}, 8},
				UniqueSum{[]int{2, 9, 10, 11}, 23},
				UniqueSum{[]int{3, 12}, 14},
				UniqueSum{[]int{4, 13, 21, 22}, 18},
				UniqueSum{[]int{5, 6}, 9},
				UniqueSum{[]int{7, 8, 17, 26}, 23},
				UniqueSum{[]int{14, 15, 23, 24}, 21},
				UniqueSum{[]int{16, 25, 34}, 8},
				UniqueSum{[]int{18, 19}, 8},
				UniqueSum{[]int{20, 29, 30, 31}, 17},
				UniqueSum{[]int{27, 36}, 16},
				UniqueSum{[]int{28, 37, 38, 47}, 16},
				UniqueSum{[]int{32, 33}, 11},
				UniqueSum{[]int{35, 44}, 9},
				UniqueSum{[]int{39, 48}, 11},
				UniqueSum{[]int{40, 41, 42, 49}, 18},
				UniqueSum{[]int{43, 52, 53}, 22},
				UniqueSum{[]int{45, 54}, 8},
				UniqueSum{[]int{46, 55}, 4},
				UniqueSum{[]int{50, 57, 58, 59}, 24},
				UniqueSum{[]int{51, 60, 69}, 12},
				UniqueSum{[]int{56, 65}, 12},
				UniqueSum{[]int{61, 62}, 11},
				UniqueSum{[]int{63, 64, 72, 73}, 19},
				UniqueSum{[]int{66, 67, 68}, 22},
				UniqueSum{[]int{70, 71}, 3},
				UniqueSum{[]int{74, 75, 76, 77}, 15},
				UniqueSum{[]int{78, 79, 80}, 23},
			),
		},
	}
	for i, test := range tests {
		solved := Solve(test.start, test.chars, test.constraints)
		if test.solveable {
			if solved {
				for j, num := range test.start {
					if num != test.solution[j] {
						fmt.Println(test.start[:9])
						fmt.Println(test.start[9:18])
						fmt.Println(test.start[18:27])
						fmt.Println(test.start[27:36])
						fmt.Println(test.start[36:45])
						fmt.Println(test.start[45:54])
						fmt.Println(test.start[54:63])
						fmt.Println(test.start[63:72])
						fmt.Println(test.start[72:])
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
