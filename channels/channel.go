package main

import (
	"fmt"
)

func main() {
	// When we use size parameter for chan type this will create new goroutines.
	ch := make(chan string, 2)

	ch <- "hello" // This will not block the execution because we gave 2 gourutines in make for chan type
	msg := <-ch
	fmt.Println(msg)

	ch <- "world"
	msg = <-ch
	fmt.Println(msg)
}
