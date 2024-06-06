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
