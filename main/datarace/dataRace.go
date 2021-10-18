package main

import (
	"fmt"
	"sync"
)

/*
 * Data races are among the most common and hardest to debug types of bugs in concurrent systems.
 * A data race occurs when two goroutines access the same variable concurrently
 * and at least one of the accesses is a write.
 */

func main() {
	n := 0

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// First conflicting access.
		for i := 0; i < 1000000; i++ {
			n++
		}

		wg.Done()
	}()

	// Second conflicting access.
	for i := 0; i < 1000000; i++ {
		n++
	}

	wg.Wait()

	fmt.Println(n)
}

/*
 * To help diagnose such bugs, Go includes a built-in data race detector.
 * To use it, add the -race flag to the go command:
	$ go run -race dataRace.go
*/
