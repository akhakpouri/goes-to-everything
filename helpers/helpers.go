package helpers

import "fmt"

func Multiply(x, y int) int {
	return x * y
}

func Split(num int) (x, y int) {
	x = num * 4
	y = (num / x) + 32
	return
}

func PrintSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
