package main

import "fmt"

func main() {
	ch := make(chan string)
	// Here we blocked the execution in same goroutine. So this will not work.
	// We cannot send value to channels in same goroutine.
	fmt.Println("Now execution will be blocked")
	ch <- "hello"
	fmt.Println("This will never fire")
	msg := <-ch
	fmt.Println(msg)

}
