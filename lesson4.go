package main

import (
	"fmt"
)

func printer(c chan int) {
	for {
		data := <-c
		fmt.Printf("data is %d\n", data)
		if data == 0 {
			break
		}
	}
	c <- 0
}

func main() {
	var c = make(chan int)

	go printer(c)

	for i := 1; i < 10; i++ {
		c <- i
	}

	c <- 0
	z := <-c
	fmt.Println("z is", z)
}
