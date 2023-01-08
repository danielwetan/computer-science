package main

import (
	"fmt"
	"time"
)

// Converting dates to epoch and vice versa

func main() {

	// Set the epoch from int64
	t := time.Unix(0, 0)
	fmt.Println(t)

	// Get the epoch
	// from Time instance
	epoch := t.Unix()
	fmt.Println(epoch)

	// Current epoch time
	apochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d\n", apochNow)

	apochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", apochNano)

}