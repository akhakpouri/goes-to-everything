package iterator

import "strings"

const counter = 5

func Repeat(character string, c int) string {

	if c == 0 {
		c = counter
	}

	var repeat strings.Builder

	for range c {
		repeat.WriteString(character)
	}
	return repeat.String()
}
