package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go count("sheep", ch)

	// we wrapped with for to receive all sended items to channels.
	// But this will give an error which is "fatal error: all goroutines are asleep - deadlock!"
	// because msq will still wait for new values from ch because of infinitive for
	for {
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
