package sort

import (
	"testing"
)

func TestBinary(t *testing.T) {
	s1 := []int{1, 2, 4, 5, 8, 10, 11, 25, 44, 57, 68, 79, 82, 99, 102, 105}
	t.Run("Slice one should return index", func(t *testing.T) {
		needle := 5
		want := 3
		found, index := Binary(s1, needle)
		if !found {
			t.Errorf("Binary(s1, %v) did not find item when it exists in the slice", needle)
		}
		if index != want {
			t.Errorf("Binary(s1, %v) expected index %v, got %v", needle, want, index)
		}
	})
	t.Run("Slice one should not find needle", func(t *testing.T) {
		needle := 45
		found, index := Binary(s1, needle)
		if found {
			t.Errorf("Binary(s1, %v) found item when it does not exist in the slice", needle)
		}
		if index != -1 {
			t.Errorf("Binary(s1, %v) expected index -1, got %v", needle, index)
		}
	})

	s2 := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 5, 6, 7, 8, 9, 10}
	t.Run("Slice two should return index", func(t *testing.T) {
		needle := 2
		want := 4
		found, index := Binary(s2, needle)
		if !found {
			t.Errorf("Binary(s2, %v) did not find item when it exists in the slice", needle)
		}
		if index != want {
			t.Errorf("Binary(s2, %v) expected index %v, got %v", needle, want, index)
		}
	})
	t.Run("Slice two should not find needle", func(t *testing.T) {
		needle := 45
		found, index := Binary(s2, needle)
		if found {
			t.Errorf("Binary(s2, %v) found item when it does not exist in the slice", needle)
		}
		if index != -1 {
			t.Errorf("Binary(s2, %v) expected index -1, got %v", needle, index)
		}
	})

	s3 := []int{0, 2, 8, 16}
	t.Run("Slice three should return index", func(t *testing.T) {
		needle := 16
		want := 3
		found, index := Binary(s3, needle)
		if !found {
			t.Errorf("Binary(s3, %v) did not find item when it exists in the slice", needle)
		}
		if index != want {
			t.Errorf("Binary(s3, %v) expected index %v, got %v", needle, want, index)
		}
	})
	t.Run("Slice three should not find needle", func(t *testing.T) {
		needle := 9999
		found, index := Binary(s3, needle)
		if found {
			t.Errorf("Binary(s3, %v) found item when it does not exist in the slice", needle)
		}
		if index != -1 {
			t.Errorf("Binary(s3, %v) expected index -1, got %v", needle, index)
		}
	})
}
