package main

import (
	"fmt"

	"log"

	helpers "github.com/akhakpouri/go-helpers/math"
)

func main() {
	x, y := 0, 5
	num, err := helpers.Multiply(x, y)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(num)
}
