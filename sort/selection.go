package sort

import (
	"github.com/cjbagley/go-algorithms/slice"
)

func selection(s []int) []int {
	length := len(s)
	if length == 0 {
		return s
	}

	sorted := make([]int, length)
	for i := 0; i < length; i++ {
		lowest := slice.GetLowest(s)
		sorted[i] = s[lowest]
		s = slice.Pop(s, lowest)
	}

	return sorted
}
