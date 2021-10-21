package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doTask(id int) {
	defer wg.Done()

	fmt.Printf("Goroutine %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Goroutine %d done\n", id)
}

func main() {
	wg.Add(5)

	for i := 0; i < 5; i++ {
		doTask(i)
	}

	// wg.Wait()
}
