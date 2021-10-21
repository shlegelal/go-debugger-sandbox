package main

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
