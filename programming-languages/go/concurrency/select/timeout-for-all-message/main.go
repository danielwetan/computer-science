package main

import (
	"fmt"
	"time"
)

// Reference
// https://go.dev/talks/2012/concurrency.slide#36
func main() {
	c := greet("first")
	timeout := time.After(3 * time.Second)

	for {
		select {
		case s, ok := <-c:
			if !ok {
				fmt.Println("Channel closed!")
				return
			}
			fmt.Println(s)

		case <-timeout:
			fmt.Println("timeout!")
			return
		}
	}
}

// greet function generates messages and sends them to the channel
func greet(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 100; i++ { // Send a limited number of messages
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(200 * time.Millisecond)
		}
		close(c) // Close the channel after sending all messages
	}()

	return c
}
