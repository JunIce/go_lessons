package main

import (
	"fmt"
)

func main() {
	var ch = make(chan string)
	go func() {
		fmt.Println("go goroutine\n")
		ch <- "jack"
		fmt.Println("send data ok\n")
	}()
	fmt.Println("start goroutine\n")
	c := <-ch
	fmt.Println("get go routine data", c)
	fmt.Println("done\n")

}
