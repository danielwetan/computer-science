package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/test", testHandler)

	wrappedMux := loggingMiddleware(mux)
	log.Println("Starting http server on :8000...")
	log.Fatal(http.ListenAndServe(":8000", wrappedMux))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test")
}
