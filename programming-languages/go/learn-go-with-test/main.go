package main

import (
	dependencyinjection "danielwetan/learn-go-with-test/dependency_injection"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
