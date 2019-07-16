package main

import (
	"fmt"
	t "test1"
)

type Person struct {
	t.Human
	Age int
}

func main() {
	p := &Person{Age: 10, Eye: 2, Ear: 2}
	fmt.Println(p)
}
