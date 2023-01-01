package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "This is example")
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		_, err := io.Copy(os.Stdout, pReader)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
