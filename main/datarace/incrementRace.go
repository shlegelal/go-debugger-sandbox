package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func increment(counter *int, times int) {
	for i := 0; i < times; i++ {
		*counter++
	}

	wg.Done()
}

func main() {
	n := 0
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go increment(&n, 1000000)
	}

	// wg.Wait()
	fmt.Println(n)
}
