package learning

const (
	englishPrefix = "hello, "
	spanishPrefix = "holla, "
	frenchPrefix  = "Bonjour, "
	french        = "french"
	spanish       = "spanish"
)

func SayHello(universe, language string) string {
	if universe == "" {
		universe = "world"
	}
	return greeting(language) + universe
}

func greeting(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
