package main

import (
	"fmt"
	"runtime"
)

const (
	limit = 10000
)

func main() {
	n := runtime.GOMAXPROCS(0)

	ch := make(chan int)

	for i := 0; i < n; i++ {
		go func(i int, r chan<- int) {
			sum := 0
			start := (limit / n) * i
			end := start + (limit / n)

			for j := start; j < end; j++ {
				sum += j
			}

			r <- sum
		}(i, ch)
	}

	total := 0

	for i := 0; i < n; i++ {
		total += <-ch
	}

	fmt.Println(total)
}
