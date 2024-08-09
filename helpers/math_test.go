package helpers

import (
	"testing"
)

func TestMultiplyPass(t *testing.T) {
	x := 1
	y := 5

	answer, err := Multiply(x, y)

	if err != nil && answer == 0 {
		t.Fatalf("multiply(%v, %v) got %v, error", x, y, err)
	}
}

func TestMultiplyFail(t *testing.T) {
	x := 0
	y := 5

	answer, err := Multiply(x, y)

	if err != nil && answer == 0 {
		t.Fatalf("multiply(%v, %v) got '%v' error", x, y, err)
	}
}

func TestRandNum(t *testing.T) {
	num := GetRandNum()
	t.Logf("number is: %v", num)
}
