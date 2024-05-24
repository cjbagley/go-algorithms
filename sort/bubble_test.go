package sort_test

import (
	"github.com/cjbagley/go-algorithms/sort"
	"reflect"
	"testing"
)

func Test_bubble(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{name: "one", slice: []int{5, 4, 5, 9, 1, 3, 8}, want: []int{1, 3, 4, 5, 5, 8, 9}},
		{name: "two", slice: []int{0, 1, 2, 3, 4, 5, 6}, want: []int{0, 1, 2, 3, 4, 5, 6}},
		{name: "three", slice: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort.Bubble(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
