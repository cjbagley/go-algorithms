package sort

import (
	"github.com/cjbagley/go-algorithms/slice"
	"testing"
)

func Test_selection(t *testing.T) {
	tests := []struct {
		name string
		size int
		max  int
	}{
		{name: "Array size 3", size: 3, max: 20},
		{name: "Array size 5", size: 5, max: 50},
		{name: "Array size 20", size: 20, max: 100},
		{name: "Array size 50", size: 50, max: 200},
		{name: "Array size 200", size: 200, max: 1000},
		{name: "Array size 2000", size: 2000, max: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := slice.Random(tt.size, tt.max)
			s = selection(s)

			if !slice.IsSorted(s) {
				// Could really print out a portion of the slice
				// but for now, YAGNI and KISS
				t.Errorf("%v not sorted correctly", tt.name)
			}

			onlyHasZeros := true
			for i := 0; i < len(s); i++ {
				if s[i] != 0 {
					onlyHasZeros = false
					break
				}
			}

			if onlyHasZeros {
				t.Errorf("%v only contains zeros", tt.name)
			}
		})
	}
}
