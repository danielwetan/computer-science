package main 

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title string `xml:"title"`
	Author string `xml:"author"`
}

func main() {
	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)

	// Read the book one by one
	books := make([]Book, 0)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "book" {
				// Decode the element to struct
				var book Book 
				decoder.DecodeElement(&book, &tokenType)
				books = append(books, book)
			}
		}
	}

	fmt.Println(books)
}