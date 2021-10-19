package main

import (
	"fmt"
	"time"
)

type Person struct {
	name, surname string
}

func (receiver *Person) fill(name, surname string) {
	receiver.name = name
	time.Sleep(500 * time.Millisecond)
	receiver.surname = surname
}

func (receiver Person) print() {
	fmt.Printf("Name: %s\nSurname: %s\n", receiver.name, receiver.surname)
}

var person Person

func main() {
	go person.fill("Pavel", "Durov")
	time.Sleep(100 * time.Millisecond)
	go person.fill("Elon", "Musk")

	time.Sleep(410 * time.Millisecond)

	person.print()
}
