package main

import (
	"fmt"
	"strings"
)

func handleFunc(list []string, chain []func(string) string) []string {

	for index, str := range list {
		result := str
		for _, props := range chain {
			result = props(result)
		}
		list[index] = result
	}
	return list
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	chain := []func(string) string{
		strings.TrimSpace,
		strings.ToUpper,
	}

	for _, str := range handleFunc(list, chain) {
		fmt.Println(str)
	}
}
