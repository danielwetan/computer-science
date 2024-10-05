package main

import (
	"fmt"
	"time"
)

// Reference
// https://go.dev/talks/2012/concurrency.slide#20
func main() {
	c := make(chan string)
	go greet("hello", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("finish!")
}

func greet(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- msg
		time.Sleep(200 * time.Millisecond)
	}
}
