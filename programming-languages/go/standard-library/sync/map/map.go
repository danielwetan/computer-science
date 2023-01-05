package main

import (
	"fmt"
	"sync"
)

// Map is like a Go map[interface{}]interface{} but is safe for concurrent use by multiple goroutines
// without additional locking or coordination.
// Loads, stores, and deletes run in amortized constant time.

// The Map type is specialized.
// Most code should use a plain Go map instead, with separate locking or coordination,
// for better type safety and to make it easier to maintain other invariants along with the map content.

// References
// https://pkg.go.dev/sync@go1.19.4#Map

func main() {
	syncMapGeneralUsage()
}

func syncMapGeneralUsage() {
	var syncMap sync.Map

	// Store sets the value for a key
	syncMap.Store("name", "Daniel")
	syncMap.Store("temp", "temporary value")
	syncMap.Store("temp2", "temporary value 2")

	// LoadOrStore returns the existing value for the key if present
	// Otherwise, it stores and returns the given value
	loadOrStore, _ := syncMap.LoadOrStore("location", "Jakarta")
	fmt.Println(loadOrStore)

	// Delete deletes the value for the key
	syncMap.Delete("temp")

	// Load returns the value stored in map for a key, or nil if no value is presenet
	name, _ := syncMap.Load("name")
	temp, _ := syncMap.Load("temp")
	fmt.Println(name)
	fmt.Println(temp)

	// LoadAndDelete the value for a key, returning the previous value if any
	temp2, _ := syncMap.LoadAndDelete("temp2")
	temp2deleted, _ := syncMap.Load("temp2")
	fmt.Println(temp2)
	fmt.Println(temp2deleted)

	fmt.Println("---------")
	// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
	syncMap.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
