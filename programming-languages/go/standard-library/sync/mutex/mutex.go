package main

import (
	"fmt"
	"sync"
)

type Container interface {
	incr()
	toString() string
}

type MutexContainer struct {
	mu       sync.Mutex
	counters int
}

type NormalContainer struct {
	counters int
}

func (c *MutexContainer) incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters++
}

func (c *MutexContainer) toString() string {
	return fmt.Sprintf("%v", c.counters)
}

func (c *NormalContainer) incr() {
	c.counters++
}

func (c *NormalContainer) toString() string {
	return fmt.Sprintf("%v", c.counters)
}

func NewMutexContainer() Container {
	return &MutexContainer{counters: 0}
}

func NewNormalContainer() Container {
	return &NormalContainer{counters: 0}
}

func main() {
	var wg sync.WaitGroup

	normalContainer := NewNormalContainer()
	mutexContainer := NewMutexContainer()

	doIncrement := func(target Container, n int) {
		for i := 0; i < n; i++ {
			target.incr()
		}
		wg.Done()
	}

	wg.Add(4)
	go doIncrement(normalContainer, 100000)
	go doIncrement(normalContainer, 100000)
	go doIncrement(mutexContainer, 100000)
	go doIncrement(mutexContainer, 100000)

	wg.Wait()

	fmt.Println(normalContainer.toString())
	fmt.Println(mutexContainer.toString())
}
