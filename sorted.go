package sudoku

type Sorted []int

func (s Sorted) Constrain(su *Sudoku, pos int, marked []bool) bool {
	sp := slicePos([]int(s), pos)
	if sp == -1 {
		return true
	}
	min := 1
	nums := make([]int, len(s))
	for n, p := range s {
		nums[n] = su.data[p]
	}
	f := false
	for i := 0; i < sp; i++ {
		if mp := su.data[s[i]]; mp == 0 {
			min++
		} else if mp+1 > min {
			min = mp + 1
		} else if f {
			return false
		}
		f = true
	}
	max := su.chars
	f = false
	for i := len(s) - 1; i > sp; i-- {
		if mp := su.data[s[i]]; mp == 0 {
			max--
		} else if mp-1 < max {
			max = mp - 1
		} else if f {
			return false
		}
		f = true
	}
	if max-min < 0 {
		return false
	}
	for i := 1; i < min; i++ {
		marked[i] = true
	}
	for i := su.chars; i > max; i-- {
		marked[i] = true
	}
	return true
}
