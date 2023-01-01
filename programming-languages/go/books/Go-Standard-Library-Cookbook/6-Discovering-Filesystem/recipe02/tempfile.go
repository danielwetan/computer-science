package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	tFile, err := ioutil.TempFile("", "gostdcookbook")
	if err != nil {
		panic(err)
	}

	// The called is responsible for handling clean up
	defer os.Remove(tFile.Name())

	fmt.Println(tFile.Name())


	// TempDir returns the path in string
	tDir, err := ioutil.TempDir("", "gostdcookbookdir")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tDir)

	fmt.Println(tDir)
}

