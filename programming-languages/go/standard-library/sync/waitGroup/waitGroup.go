package main

import (
	"fmt"
	"sync"
)

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

// A WaitGroup must not be copied after first use.
// In the terminology of the Go memory model, a call to Done “synchronizes before” the return of any Wait call that it unblocks.

// Reference
// https://pkg.go.dev/sync@go1.19.4#WaitGroup

func main() {
	var wg sync.WaitGroup

	doAction := func() {
		fmt.Println("do action")
		wg.Done()
	}

	wg.Add(3)
	go doAction()
	go doAction()
	go doAction()

	wg.Wait()
}
