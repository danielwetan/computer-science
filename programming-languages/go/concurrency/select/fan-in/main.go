package main

import (
	"fmt"
	"time"
)

// Select
// The select statement provides another way to handle multiple channels.
// It's like a switch, but each case is a communication:
// - All channels are evaluated.
// - Selection blocks until one communication can proceed, which then does.
// - If multiple can proceed, select chooses pseudo-randomly.
// - A default clause, if present, executes immediately if no channel is ready.

// Reference
// https://go.dev/talks/2012/concurrency.slide#33
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
			select {
			case s := <-input1:
				c <- s
			case w := <-input2:
				c <- w
			}
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
