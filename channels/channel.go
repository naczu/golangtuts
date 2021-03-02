package main

import "fmt"

func main() {
	// We just a create fibonacci function which is shows nth fibonacci value
	// The problem how to show fibonacci series with goroutines and channels
	fmt.Println(fib(10))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
