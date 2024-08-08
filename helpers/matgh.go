package helpers

import (
	"errors"
	"fmt"
)

func Multiply(x int, y int) (int, error) {
	if x == 0 || y == 0 {
		return 0, errors.New("both x & y must be greater than 0")
	}
	return x * y, nil
}

func Split(num int) (x, y int) {
	x = num * 4
	y = (num / x) + 32
	return
}

func PrintSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
