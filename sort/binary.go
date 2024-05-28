package sort

func Binary(s []int, needle int) (bool, int) {
	if len(s) == 0 {
		return false, -1
	}

	lowest := 0
	highest := len(s)
	for lowest < highest {
		mid := lowest + ((highest - lowest) / 2)
		i := s[mid]

		if i == needle {
			return true, mid
		}

		if needle > i {
			lowest = mid + 1
		} else {
			highest = mid
		}
	}

	return false, -1
}
