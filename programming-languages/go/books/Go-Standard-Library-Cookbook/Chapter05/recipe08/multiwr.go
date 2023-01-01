package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Writing to multiple writers at once

func main() {
	buf := bytes.NewBuffer([]byte{})
	anotherBuf := bytes.NewBuffer([]byte{})

	f, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	wr := io.MultiWriter(buf, f, anotherBuf)
	_, err = io.WriteString(wr, "Go is awesome")
	if err != nil {
		panic(err)
	}

	fmt.Println("Content of buffer: " + buf.String())
	io.WriteString(os.Stdout, anotherBuf.String())
}
