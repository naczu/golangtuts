package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go count("sheep", ch)

	// channel receiving and sending blocks the execution
	msg := <-ch

	fmt.Println(msg)

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// channel receiving and sending blocks the execution
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}
