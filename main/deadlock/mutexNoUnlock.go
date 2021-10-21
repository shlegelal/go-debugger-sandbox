package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func prettyPrint(str string, wg *sync.WaitGroup) {
	// defer wg.Done()

	mutex.Lock()

	for _, char := range str {
		fmt.Printf("%c ", char)
		time.Sleep(time.Millisecond * 300)
	}

	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	// wg.Add(2)

	go prettyPrint("water", &wg)
	go prettyPrint("punk", &wg)

	time.Sleep(5 * time.Second)
	// wg.Wait()
}
