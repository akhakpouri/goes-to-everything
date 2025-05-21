package learning

const (
	englishPrefix = "hello, "
	spanishPrefix = "holla, "
	frenchPrefix  = "Bonjour, "
	french        = "french"
	spanish       = "spanish"
)

func SayHello(universe, language string) string {
	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}

	if universe == "" {
		universe = "world"
	}

	return prefix + universe
}
