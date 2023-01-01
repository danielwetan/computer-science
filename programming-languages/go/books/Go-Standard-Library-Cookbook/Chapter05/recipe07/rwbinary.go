package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Reading and writing binary data

func main() {
	// Writing binary values
	buf := bytes.NewBuffer([]byte{})

	err := binary.Write(buf, binary.BigEndian, 1.004)
	if err != nil {
		panic(err)
	}

	err = binary.Write(buf, binary.BigEndian, []byte("Hello"))
	if err != nil {
		panic(err)
	}

	err = binary.Write(buf, binary.BigEndian, []byte("Testing"))
	if err != nil {
		panic(err)
	}

	// Reading the written values
	var num float64
	err = binary.Read(buf, binary.BigEndian, &num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)

	greeting := make([]byte, 5)
	err = binary.Read(buf, binary.BigEndian, &greeting)
	if err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", string(greeting))

	testingMsg := make([]byte, 7)
	err = binary.Read(buf, binary.BigEndian, &testingMsg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", string(testingMsg))
}
