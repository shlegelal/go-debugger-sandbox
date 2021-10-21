package main

import (
	"fmt"
	"sync"
	"time"
)

func doTask(id int) {
	fmt.Printf("Goroutine %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Goroutine %d done\n", id)
}

func main() {
	i := 0
	var wg sync.WaitGroup
	wg.Add(5)

	for i < 5 {
		go func(id int) {
			defer wg.Done()
			doTask(id)
		}(i)
		i++
	}
	//wg.Wait()
}
