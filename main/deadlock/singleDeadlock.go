package main

import "fmt"

/*
 * The program will get stuck on the channel send operation waiting forever for someone to read the value.
 * Go is able to detect situations like this at runtime.
 */

// Single go self deadlock
func main() {
	ch := make(chan int)
	ch <- 789
	num := <-ch
	fmt.Println(num)
}

/*
 * At first glance, you may feel that there is no problem with the above paragraph,
 * but if you look closely, you will find that this ch is a bufferless channel.
 * When 789 writes to the buffer, the reader is not ready.
 * Therefore, the writer will block and the subsequent code will no longer run.
 *
 * Currently, Go only detects when the program as a whole freezes, not when a subset of goroutines get stuck.
 */

/*
 * To avoid this situation we use a default case in the select statement.

select {
	case <-ch:
	default:
		fmt.Println("!.. Default case..!")
}
*/
