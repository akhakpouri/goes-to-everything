package learning

import "fmt"

func SayHello(universe string) string {
	return "hello, " + universe
}

func main() {
	fmt.Println(SayHello("world"))
	fmt.Println(SayHello("earth"))
}
