package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
	// when using fmtScanln it blocks to finishing execution
	fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
