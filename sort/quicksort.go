package sort

import (
	"github.com/cjbagley/go-algorithms/slice"
)

func Quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := (len(s) - 1) / 2
	val := s[mid]
	s = slice.Pop(s, mid)

	var less []int
	var more []int
	for i := 0; i < len(s); i++ {
		if s[i] < val {
			less = append(less, s[i])
		} else {
			more = append(more, s[i])
		}
	}

	less = append(Quicksort(less), val)
	more = Quicksort(more)
	return append(less, more...)
}
