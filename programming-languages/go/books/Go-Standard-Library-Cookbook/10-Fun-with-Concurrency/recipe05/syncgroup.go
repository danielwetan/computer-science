package main

import (
	"fmt"
	"sync"
)

// Synchronizing goroutines with WaitGroup

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Existing %d\n", idx)
		}(i)
	}

	wg.Wait()
	fmt.Println("All done.")
}
