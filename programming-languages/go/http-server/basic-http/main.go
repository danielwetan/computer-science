package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/test", testHandler)

	log.Println("starting http server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test")
}
