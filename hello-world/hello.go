package learning

const helloPrefix = "hello, "

func SayHello(universe string) string {
	if universe == "" {
		universe = "world"
	}

	return helloPrefix + universe
}
