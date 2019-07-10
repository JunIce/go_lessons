package main

import (
	"fmt"
	"runtime"
	"time"
)

func running() {
	for {
		fmt.Println(runtime.NumCPU())
		time.Sleep(time.Second)
	}
}

func main() {
	go running()
	var input string
	fmt.Scanln(&input)
}
