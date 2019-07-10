package main

import (
	"errors"
	"fmt"
	"time"
)

func client(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("time out")
	}
}

func server(ch chan string) {
	for {
		data := <-ch
		fmt.Println("server receiverd", data)
		time.Sleep(time.Second * 3)
		ch <- "roger"
	}
}

func main() {
	ch := make(chan string)
	go server(ch)

	recv, err := client(ch, "hello world")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(recv)
	}
}
