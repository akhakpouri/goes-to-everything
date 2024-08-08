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

	x := helpers.Multiply(1, 5)

	fmt.Println(message)
	fmt.Println(x)
}
