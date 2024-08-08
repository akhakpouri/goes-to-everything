package main

import (
	"example/greetings"
	"fmt"
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
}
