package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

// Reference
// https://go.dev/talks/2012/concurrency.slide#39
func main() {
	const n = 10000
	leftmost := make(chan int)

	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)

	fmt.Println(<-leftmost)
}
