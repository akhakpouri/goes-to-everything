package iterator

func Repeat(character string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat = repeat + character
	}
	return repeat
}
