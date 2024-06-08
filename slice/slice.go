package slice

import (
	"math/rand"
)

func MakeRandom(numItems, max int) []int {
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

func IsSorted(slice []int) bool {
	if len(slice) == 0 {
		return true
	}

	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}

	return true
}

func Pop(slice []int, index int) []int {
	length := len(slice)
	if length == 0 {
		return slice
	}

	if index > length || index < 0 {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func GetLowest(slice []int) (lowestIndex int) {
	lowestValue := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < lowestValue {
			lowestValue = slice[i]
			lowestIndex = i
		}
	}

	return lowestIndex
}
