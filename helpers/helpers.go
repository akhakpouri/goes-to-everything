package helpers

func Multiply(x, y int) int {
	return x * y
}

func Split(num int) (x, y int) {
	x = num * 4
	y = (num / x) + 32
	return
}
