package main

import (
	"fmt"
)

type Person struct {
	Name     string
	Age      int
	Password string
	CreateAt int64
}

func main() {
	user := Person{Name: "jack", Age: 64}
	fmt.Println(user.Name)
	fmt.Println(user)
}
