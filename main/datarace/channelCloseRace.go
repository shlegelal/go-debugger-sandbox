package main

import "time"

func main() {
	ch := make(chan int)

	go func() { ch <- 17 }()

	close(ch)

	time.Sleep(500 * time.Millisecond)
}
