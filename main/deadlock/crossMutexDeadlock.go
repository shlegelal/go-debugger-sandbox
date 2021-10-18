package main

import (
	"sync"
	"time"
)

var mutex1 sync.Mutex
var mutex2 sync.Mutex

// Cross mutex deadlock
func main() {
	go func() {
		mutex1.Lock()
		mutex2.Lock()

		// code

		mutex2.Unlock()
		mutex1.Unlock()
	}()

	mutex2.Lock()

	time.Sleep(time.Millisecond * 300)

	mutex1.Lock()

	// code

	mutex2.Unlock()
	mutex1.Unlock()
}
