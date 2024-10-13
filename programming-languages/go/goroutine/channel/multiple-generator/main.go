package main

import (
	"fmt"
	"time"
)

// Reference
// https://go.dev/talks/2012/concurrency.slide#25
func main() {
	first := greet("hello")
	second := greet("hello")

	for i := 0; i < 5; i++ {
		fmt.Printf("First say: %q\n", <-first)
		fmt.Printf("Second say: %q\n", <-second)
	}

	fmt.Println("finish!")
}

func greet(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- msg
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return c
}