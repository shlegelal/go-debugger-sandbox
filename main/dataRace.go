package main

import "fmt"

/*
 * Data races are among the most common and hardest to debug types of bugs in concurrent systems.
 * A data race occurs when two goroutines access the same variable concurrently
 * and at least one of the accesses is a write.
 */

/*
 * Here is an example of a data race that can lead to crashes and memory corruption:
 */

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
 * To help diagnose such bugs, Go includes a built-in data race detector.
 * To use it, add the -race flag to the go command:
	$ go run -race dataRace.go
*/
