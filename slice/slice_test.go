package slice_test

import (
	"github.com/cjbagley/go-algorithms/slice"
	"reflect"
	"testing"
)

func TestRandom(t *testing.T) {
	tests := []struct {
		name   string
		length int
		max    int
	}{
		{name: "one", length: 6, max: 10},
		{name: "two", length: 88, max: 300},
		{name: "three", length: 51, max: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.Random(tt.length, tt.max)
			length := len(got)
			if length != tt.length {
				t.Errorf("Random() contains %v items, want %v", length, tt.length)
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
					t.Errorf("Random() value %v exceeds max %v", v, tt.max)
				}
			}

			if duplicate == tt.length {
				t.Errorf("Random() contains only the same int: %v", got)
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

func TestPop(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  []int
	}{
		{name: "Pop first index", index: 0, want: []int{8, 2, 1, 8, 2, 0, 10}},
		{name: "Pop an inner index", index: 3, want: []int{8, 8, 2, 8, 2, 0, 10}},
		{name: "Pop last index", index: 7, want: []int{8, 8, 2, 1, 8, 2, 0}},
		{name: "Handle invalid index", index: 10, want: []int{8, 8, 2, 1, 8, 2, 0, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := []int{8, 8, 2, 1, 8, 2, 0, 10}
			if got := slice.Pop(test, tt.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLowest(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{name: "Lowest is first index", slice: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 0},
		{name: "Lowest is inner index", slice: []int{9, 7, 10, 4, 1, 6, 5, 8}, want: 4},
		{name: "Lowest is last index", slice: []int{8, 7, 6, 5, 4, 3, 2, 1, 0}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := slice.GetLowest(tt.slice); gotIndex != tt.want {
				t.Errorf("GetLowest() = %v, want %v", gotIndex, tt.want)
			}
		})
	}
}
