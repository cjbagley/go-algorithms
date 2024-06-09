package search

func Binary(s []int, needle int) (bool, int) {
	if len(s) == 0 {
		return false, -1
	}

	lowest := 0
	highest := len(s) - 1
	for lowest <= highest {
		//mid := lowest + ((highest - lowest) / 2)
		mid := (lowest + highest) / 2
		i := s[mid]

		if i == needle {
			return true, mid
		}

		if needle > i {
			lowest = mid + 1
		} else {
			highest = mid - 1
		}
	}

	return false, -1
}

// RecursiveBinary
// Wrapper for the actual recursive binary, just to save the initial high/low
// calculations for the user
func RecursiveBinary(s []int, needle int) (bool, int) {
	return recursiveBinaryInner(s, needle, 0, len(s)-1)
}

func recursiveBinaryInner(s []int, needle, low, high int) (bool, int) {
	// Handle zero
	if len(s) == 0 {
		return false, -1
	}

	// Base case
	if low > high {
		return false, -1
	}

	mid := (low + high) / 2
	value := s[mid]

	if value == needle {
		return true, mid
	}

	if needle > value {
		low = mid + 1
	} else {
		high = mid - 1
	}

	return recursiveBinaryInner(s, needle, low, high)
}
