package main

import (
	"fmt"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(200 * time.Millisecond))
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

// Reference
// https://go.dev/talks/2012/concurrency.slide#46
func main() {
	start := time.Now()

	results := Google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}
