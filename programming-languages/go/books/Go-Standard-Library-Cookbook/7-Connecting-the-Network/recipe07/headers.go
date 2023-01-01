package main 

import (
	"fmt"
	"net/http"
)

func main() {
	header := http.Header{}

	// Using the header as slice
	header.Set("Auth-X", "abcdef1234")
	header.Add("Auth-X", "defghijkl")
	fmt.Println(header)

	// retrieving slice of values in header
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// gte the first value 
	resFirst := header.Get("Auth-x")
	fmt.Println(resFirst)

	// replace all existing value with this one 
	header.Set("Auth-X", "newValue")
	fmt.Println(header)

	// Remove header 
	header.Del("Auth-X")
	fmt.Println(header)
}
