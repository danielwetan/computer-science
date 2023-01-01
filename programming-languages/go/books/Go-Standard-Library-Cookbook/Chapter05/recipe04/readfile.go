package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// Reading the file into a string
// The reading from the file is simple because the File type implements both the Reader and Writer interfaces.
// This way, all functions and approaches applicable to the Reader interface are applicable to the File type.
// The preceding example shows how to read the file with the use of Scanner and write the content to the bytes buffer (which is more performant than string concatenation).
// This way, you are able to control the volume of content read from a file.

func main() {
	fmt.Println("### Read as reader ###")
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text() + " edited\n")
	}

	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// for smaller files
	fContent, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fContent))

}
