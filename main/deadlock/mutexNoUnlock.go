package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Programs that make heavy use of mutexes can, on the other hand, be notoriously difficult to debug.
 */

// Using the traditional "lock" to complete synchronization -- mutex
var mutex sync.Mutex // Create a mutex (mutex), the new mutex state is 0 > not locked. There is only one lock
func printer(str string) {
	mutex.Lock() // Lock before accessing shared data
	for _, s := range str {
		fmt.Printf("%c ", s)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock() // Shared data access ends, unlock
}

func person1(group *sync.WaitGroup) {
	//defer group.Done()
	printer("hello")
}

func person2(group *sync.WaitGroup) {
	//defer group.Done()
	printer("world")
}

func main() {
	var wg sync.WaitGroup
	//wg.Add(2)
	go person1(&wg)
	go person2(&wg)
	time.Sleep(5 * time.Second)

	//wg.Wait()
}
