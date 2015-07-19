package sudoku

// UniqueSum is a Constraint that requires all numbers to be unique and total a specific amount.
type UniqueSum struct {
	positions []int
	total     int
}

// Constrain implements the Constraint interface.
func (u UniqueSum) Constrain(s *Sudoku, pos int, marked []bool) bool {
	if slicePos(u.positions, pos) == -1 {
		return true
	}
	total := 0
	myMark := make([]bool, s.chars+1)
	empty := 0
	for _, p := range u.positions {
		if mp := s.data[p]; mp == 0 {
			empty++
		} else {
			if myMark[mp] {
				return false
			}
			myMark[mp] = true
			marked[mp] = true
			total += mp
		}
	}
	remaining := u.total - total
	if remaining < 0 {
		return false
	}
	myNums := make([]int, 0, s.chars)
	for n, b := range myMark[1:] {
		if !b {
			myNums = append(myNums, n+1)
		}
	}
	data := make([]int, 0, empty)
	marks := make([]bool, s.chars+1)
	if !getCombinations(myNums, data, 0, remaining, marks) {
		return false
	}
	r := false
	for n, m := range marks {
		if !m {
			marked[n] = true
		} else if !r && !marked[n] {
			r = true
		}
	}
	return r
}

func getCombinations(nums, data []int, pos, remaining int, marks []bool) bool {
	if len(data) == cap(data) {
		if remaining == 0 {
			for _, n := range data {
				marks[n] = true
			}
			return true
		}
		return false
	}
	toRet := false
	o := data
	for i := pos; i < len(nums); i++ {
		if nums[i] > remaining {
			continue
		}
		data = append(data, nums[i])
		if getCombinations(nums, data, i+1, remaining-nums[i], marks) {
			toRet = true
		}
		data = o
	}
	return toRet
}
