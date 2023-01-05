package main

import (
	"fmt"
	"sync"
)

// Once is an object that will perform exactly one action.
// A Once must not be copied after first use.
// In the terminology of the Go memory model, the return from f “synchronizes before” the return from any call of once.Do(f).

// Reference
// https://pkg.go.dev/sync@go1.19.4#Once

func main() {
	var once sync.Once
	getMessage := func() {
		fmt.Println("only once")
	}

	for i := 0; i < 100; i++ {
		once.Do(getMessage)
	}
}
