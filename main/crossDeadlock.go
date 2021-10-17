package main

// Multi go, multi channel cross deadlock
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		num := <-ch1
		ch2 <- num
	}()

	num := <-ch2
	ch1 <- num
}

/*
 * When reading (writing) at one end of the channel,
 * it is necessary to ensure that the other end of the channel has a chance
 * to execute the write (read) operation at the other end.
 */
