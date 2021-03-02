package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	// We created another goroutine and send value to channel in that gourutine.
	// So the execution will not stop
	go anotherRoutine(ch)
	msg := <-ch // this will block the execution until it receives the value from channel
	fmt.Println(msg)
}

func anotherRoutine(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "hello"
}
