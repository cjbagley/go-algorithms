package recursion

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{name: "Test 1", slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, want: 9},
		{name: "Test 2", slice: []int{100, 69, 25, 89, 60, 291}, want: 6},
		{name: "Test 3", slice: []int{20, -29, 17, -55}, want: 4},
		{name: "Test handles empty", slice: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.slice); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHighest(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{name: "Test single digits", slice: []int{0, 1, 2, 3, 8, 5, 6, 7, 4}, want: 8},
		{name: "Test multi digits", slice: []int{100, 69, 25, 89, 60, 291, 89, 8}, want: 291},
		{name: "Test negatives", slice: []int{20, -29, 17, -55}, want: 20},
		{name: "Test handles empty", slice: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Highest(tt.slice); got != tt.want {
				t.Errorf("Highest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{name: "Test single digits", slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, want: 36},
		{name: "Test multi digits", slice: []int{100, 69, 25, 89, 60, 291, 89, 8}, want: 731},
		{name: "Test negatives", slice: []int{20, -29, 17, -55}, want: -47},
		{name: "Test handles empty", slice: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.slice); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
