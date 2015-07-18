package sudoku

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
			marked[mp] = true
			total += mp
		}
	}
	left := u.total - total
	if left < 0 {
		return false
	}
	myNums := make([]int, 0, s.chars)
	for n, m := range marked {
		if !m {
			myNums = append(myNums, n)
		}
	}
	if len(myNums) < unmarked {
		return false
	}
	data := make([]int, 0, unmarked)
	marks := make([]bool, 0, s.chars+1)
	if !getCombinations(myNums, data, left, 0, marks) {
		return false
	}
	r := false
	for n, m := range marks {
		if !m {
			marked[n] = true
			r = true
		} else if !r && !marked[n] {
			r = true
		}
	}
	return r
}
func getCombinations(nums, data []int, addTo, from int, marks []bool) bool {
	if from == cap(data) {
		total := 0
		for _, n := range nums {
			total += n
		}
		if total == addTo {
			for _, n := range nums {
				marks[n] = true // ????
			}
			return true
		}
		return false
	}
	toRet := false
	for i := from; i < cap(data); i++ {
		data = append(data, nums[i])
		if getCombinations(nums, data, addTo, i+1, marks) {
			toRet = true
		}
		data = data[:len(data)-1]
	}
	return toRet
}
