package print

import "fmt"

func Slice(slice []int, items int) {
	if len(slice) <= items {
		fmt.Println(slice)
		return
	}

	fmt.Println(slice[:items])
}
