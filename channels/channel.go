package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go count("sheep", ch)

	//
	for {
		// When we close channel we can get second parameter as channel is opened or closed
		msg, open := <-ch
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// channel receiving and sending blocks the execution
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	// We use close
	close(c)
}
