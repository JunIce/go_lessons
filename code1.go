package main

import (
	"fmt"
)

func isValid(str string) bool {
	// stack := []byte

	for _, s := range str {
		fmt.Println(s)
	}

	return true
}

func main() {
	s := "()[]{}"
	isValid(s)
}
