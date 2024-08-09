package main

import (
	"example/greetings"
	"fmt"
	"helpers"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.SayHello("Ali")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	x, err := helpers.Multiply(1, 5)

	if x == 0 || err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)
}
