package main

import (
	"fmt"
	"time"
)

func main() {
	// we create two channels and send strings to values in goroutine
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ch2 <- "every 2 second"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// We have solved the step13 problem here with select.
		// select applies the execution when the value received
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
