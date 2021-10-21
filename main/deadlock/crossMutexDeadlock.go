package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex1 sync.Mutex
var mutex2 sync.Mutex

func main() {
	go func() {
		mutex1.Lock()
		mutex2.Lock()

		fmt.Println("Doing stuff here...")

		mutex2.Unlock()
		mutex1.Unlock()
	}()

	mutex2.Lock()
	time.Sleep(300 * time.Millisecond)
	mutex1.Lock()

	fmt.Println("Doing stuff here...")

	mutex2.Unlock()
	mutex1.Unlock()
}
