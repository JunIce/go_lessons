package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	stopper := time.NewTimer(time.Second * 5)

	var i int
	for {
		select {
		case <-stopper.C:
			fmt.Println("stop")
			goto Done
		case <-ticker.C:
			i++
			fmt.Println("ticker", i)
		}
	}
Done:
	fmt.Println("stoper is done")
}
