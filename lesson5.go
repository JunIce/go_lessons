package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	fmt.Println("start\n")
	time.AfterFunc(time.Second, func() {
		fmt.Println("wait a second")
		c <- 0
	})
	<-c
	fmt.Println("get exit signal")

	fmt.Println("exit")

}
