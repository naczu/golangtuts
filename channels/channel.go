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
		fmt.Println(<-ch1) // we are waiting for slow one
		fmt.Println(<-ch2) // this is blocking the queue for 2 seconds. We need to solve this with select
	}
}
