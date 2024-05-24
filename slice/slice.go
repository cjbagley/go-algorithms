package slice

import (
	"fmt"
	"math/rand"
)

func MakeRandomInts(numItems, max int) []int {
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

func Print(slice []int, items int) {
	if len(slice) <= items {
		fmt.Println(slice)
		return
	}

	fmt.Println(slice[:items])
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
