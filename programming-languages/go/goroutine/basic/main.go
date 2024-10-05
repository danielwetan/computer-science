package main

import (
	"fmt"
	"time"
)

// Reference
// https://go.dev/talks/2012/concurrency.slide#16
func main() {
	go greet("hello")
	time.Sleep(1 * time.Second)
}

func greet(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(200 * time.Millisecond)
	}
}
