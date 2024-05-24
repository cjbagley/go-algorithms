package sort

func Bubble(s []int) []int {
	sorted := false
	length := len(s)
	for sorted == false {
		sorted = true
		for i := 1; i < length; i++ {
			if s[i-1] > s[i] {
				sorted = false
				s[i], s[i-1] = s[i-1], s[i]
			}
		}
	}
	return s
}
