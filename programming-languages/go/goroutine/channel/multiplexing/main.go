package main

import (
	"fmt"
	"time"
)

// Reference
// https://go.dev/talks/2012/concurrency.slide#27
func main() {
	c := fanIn(greet("first"), greet("second"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("finish!")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
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
