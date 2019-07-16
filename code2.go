package main

import (
	"fmt"
)

func removeDuplicates(s []int) int {
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
		fmt.Println(s)
	}

	return i + 1
}

func main() {
	s := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	c := removeDuplicates(s)
	fmt.Println(c)
}
