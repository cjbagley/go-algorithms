package slice_test

import (
	"github.com/cjbagley/go-algorithms/slice"
	"testing"
)

func TestMakeRandomInts(t *testing.T) {
	tests := []struct {
		name     string
		numItems int
		max      int
	}{
		{name: "one", numItems: 6, max: 10},
		{name: "two", numItems: 88, max: 300},
		{name: "three", numItems: 51, max: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.MakeRandomInts(tt.numItems, tt.max)
			length := len(got)
			if length != tt.numItems {
				t.Errorf("MakeRandomInts() contains %v items, want %v", length, tt.numItems)
			}

			first := -1
			duplicate := 0

			for _, v := range got {
				if first < 0 {
					first = v
				}
				if v == first {
					duplicate++
				}

				if v > tt.max {
					t.Errorf("MakeRandomInts() value %v exceeds max %v", v, tt.max)
				}
			}

			if duplicate == tt.numItems {
				t.Errorf("MakeRandomInts() contains only the same int: %v", got)
			}
		})
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  bool
	}{
		{name: "one", slice: []int{8, 8, 2, 1, 8, 2, 0, 10}, want: false},
		{name: "two", slice: []int{0, 1, 2, 3, 4, 4, 5}, want: true},
		{name: "three", slice: []int{0, -1, 2, 3, 4, 5}, want: false},
		{name: "four", slice: []int{1, 1, 1, 1, 1, 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.IsSorted(tt.slice)
			if got != tt.want {
				t.Errorf("IsSorted(%v) returned %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}
