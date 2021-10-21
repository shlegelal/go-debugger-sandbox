package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 322
	num := <-ch
	fmt.Println(num)
}
