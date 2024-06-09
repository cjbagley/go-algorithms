package recursion

func Sum(s []int) int {
	length := len(s)

	// Handle empty
	if length == 0 {
		return 0
	}

	// Base case
	if length == 1 {
		return s[0]
	}

	// Recursive call
	return s[0] + Sum(s[1:])
}

func Count(s []int) int {
	length := len(s)

	// Handle empty
	if length == 0 {
		return 0
	}

	// Base case
	if length == 1 {
		return 1
	}

	// Recursive call
	return 1 + Count(s[1:])
}

func Highest(s []int) int {
	length := len(s)

	// Handle empty
	if length == 0 {
		return 0
	}

	// Base case
	if length == 1 {
		return s[0]
	}

	// Recursive call
	highest := Highest(s[1:])
	if s[0] > highest {
		return s[0]
	}
	return highest
}
