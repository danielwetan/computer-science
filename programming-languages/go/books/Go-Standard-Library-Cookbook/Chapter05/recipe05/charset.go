package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	// Write the string ecoded to Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, err := encoder.String("This is sample text with runes Š")
	if err != nil {
		panic(err)
	}

	os.WriteFile("example.txt", []byte(s), os.ModePerm)

	// Decode to UTF-8
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
