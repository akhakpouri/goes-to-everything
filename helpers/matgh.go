package helpers

import (
	"errors"
	"math/rand/v2"
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

func GetRandNum() int {
	numbers := []int{}

	for i := 0; i <= 2500; i++ {
		if i%2 == 0 || i%3 == 0 {
			numbers = append(numbers, i)
		}
		if i%5 == 0 || i%8 == 0 {
			continue
		}
	}

	len, min := len(numbers), 1
	return rand.IntN(len-min) + min
}
