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
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
}
