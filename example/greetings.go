package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return name, errors.New("name is empty")
	}
	message := fmt.Sprintf(getRandMessage(), name)
	return message, nil
}

func getRandMessage() string {
	repos := []string{
		"Hi %s, welcome!",
		"Great to meet you, %s",
		"Welcome to the jungle, %s",
	}
	return repos[rand.Intn(len(repos))]
}
