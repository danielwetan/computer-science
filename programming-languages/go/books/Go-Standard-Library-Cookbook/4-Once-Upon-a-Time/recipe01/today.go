package main

import (
	"fmt"
	"time"
)

// Finding today's date
// Reference
// https://pkg.go.dev/time

func main() {
	today := time.Now()
	fmt.Println(today)
}
