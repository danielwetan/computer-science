package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	// Create a file to store the trace
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start tracing
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Your application logic here
	sum := 1 + 2
	fmt.Println(sum)
}
