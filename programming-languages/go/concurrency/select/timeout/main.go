package main

import (
	"fmt"
	"time"
)

// Main function
func main() {
	c := greet("first")

	for {
		select {
		case s, ok := <-c:
			if !ok {
				fmt.Println("Channel closed!")
				return
			}
			fmt.Println(s)

		// timeout for each message
		// change to 100 milisecond to simulate the timeout
		case <-time.After(500 * time.Millisecond): // Timeout if no message is received within 500ms
			fmt.Println("timeout!")
			return
		}
	}
}

// greet function generates messages and sends them to the channel
func greet(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ { // Send a limited number of messages
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(200 * time.Millisecond)
		}
		close(c) // Close the channel after sending all messages
	}()

	return c
}
