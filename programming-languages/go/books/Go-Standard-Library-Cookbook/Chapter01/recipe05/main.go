package main 

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Notes:
	// If the code is run by the command `go run`, the actual executable is located in a temporary directory
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Path to executable file
	fmt.Println(ex)

	// Resolve the directory of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path: " + exPath)

	// Use EvalSymlinks to get the real path
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated: " + realPath)
}