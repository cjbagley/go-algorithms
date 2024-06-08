package slice

import (
	"math/rand"
)

func Random(length, max int) []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

func IsSorted(s []int) bool {
	if len(s) == 0 {
		return true
	}

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}

	return true
}

func Pop(s []int, index int) []int {
	length := len(s)
	if length == 0 {
		return s
	}

	if index > length || index < 0 {
		return s
	}

	return append(s[:index], s[index+1:]...)
}

func GetLowest(s []int) (lowestIndex int) {
	lowestValue := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < lowestValue {
			lowestValue = s[i]
			lowestIndex = i
		}
	}

	return lowestIndex
}
