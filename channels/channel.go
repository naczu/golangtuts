package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go count("sheep", ch)

	// we stopped the execution with limit our for loop
	// so we will not get deadlock error
	for i := 1; i <= 5; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// channel receiving and sending blocks the execution
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}
