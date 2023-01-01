package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Extracting data from an incomplete JSON array

// Besides the Unmarshall function, the json package also contains the Decoder API. 
// With NewDecoder, the Decoder could be created. 
// By calling the Token method on the decoder, the underlying Reader is read and returns the Token interface. 
// This could hold multiple values.

const js = `
	[
		{
			"name":"Muhammad",
			"lastname":"Adam"
		},
		{
			"name":"Ahmad",
			"lastname":"Idris"
		},
		{
			"name":"Invalid",
			"lastname":"Name"
`

type User struct {
	Name string `json:"name"`
	LastName string `json:"lastname"`
}

func main() {

	userSlice := make([]User, 0)
	r := strings.NewReader(js)
	dec := json.NewDecoder(r)
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
			case json.Delim:
				str := tp.String()
				if str == "[" || str == "{" {
					for dec.More() {
						u := User{}
						err := dec.Decode(&u)
						if err == nil {
							userSlice = append(userSlice, u)
						} else {
							break
						}
					}
				}
			}
		}

		fmt.Println(userSlice)
	}