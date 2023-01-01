package main

import (
	"fmt"
	"io"
	"os"
)

// Opening a file by name

func main() {
	// The os package offers a simple way of opening file, the function Open() opens the file by path just in read-only mode
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	c, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("## File content ###\n%s\n", string(c))
	f.Close()

	// Openfile() is the more powerful one and consumer the path to the file, flags, and permissions
	f, err = os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	io.WriteString(f, "Test string")
	f.Close()
}
