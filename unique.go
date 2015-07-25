package sudoku

// Unique represents a section of data where each cell needs to be different
type Unique []int

// Constrain implements the Constraint interface
func (u Unique) Constrain(s *Sudoku, pos int, marked []bool) bool {
	if slicePos([]int(u), pos) == -1 {
		return true
	}
	myMark := make([]int, 0, len(u))
	for _, p := range u {
		if mp := s.Pos(p); mp != 0 {
			if slicePos(myMark, mp) != -1 {
				return false
			}
			myMark = append(myMark, mp)
			marked[mp] = true
		}
	}
	return true
}
