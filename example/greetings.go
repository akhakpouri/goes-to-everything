package greetings

import (
	"errors"
	"fmt"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return name, errors.New("name is empty")
	}
	message := fmt.Sprintf("hello %s", name)
	return message, nil
}
