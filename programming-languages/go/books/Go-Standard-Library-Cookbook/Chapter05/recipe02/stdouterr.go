package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Simply write string
	io.WriteString(os.Stdout, "This is string to standard output.\n")

	io.WriteString(os.Stderr, "This is string to standard error output.\n")

	// Stdout/err implements writer interface
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 10; i++ {
		if _, err := os.Stdout.Write(buf); err != nil {
			panic(err)
		}
	}

	// The fmt package could be used too
	fmt.Fprintln(os.Stdout, "\n")
}
