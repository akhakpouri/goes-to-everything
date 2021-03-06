package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/akhakpouri/goes-to-everything/helpers"
	"github.com/akhakpouri/goes-to-everything/logging"
	"github.com/akhakpouri/goes-to-everything/settings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	logging.Debug(true)

	logging.Log("Start of Application")
	x := 5
	y := 7
	sumReg := x + y
	fmt.Println(sumReg)

	if x > 6 {
		fmt.Println("x is greater than 6")
	}

	a := []int{5, 3, 2, 1, 0}
	a[2] = 7
	a = append(a, 13)
	fmt.Println(a)

	verticies := make(map[string]int)
	verticies["triangle"] = 2
	verticies["square"] = 3
	verticies["dodecagon"] = 12

	delete(verticies, "square")
	fmt.Println(verticies)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	arr := []string{"foo", "bar", "jon", "doe"}
	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
	}

	table := make(map[string]int)
	table["foo"] = 1
	table["bar"] = 2
	table["jon"] = 3
	table["doe"] = 4
	table["yo"] = 5

	for key, value := range table {
		fmt.Println("key: ", key, "value: ", value)
	}
	s := sum(1, 2)
	fmt.Println(s)

	sq, err := sqrt(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sq)
	}

	muliply := helpers.Multiply(6, 8)
	xSplit, ySplit := helpers.Split(17)

	fmt.Println(settings.ReverseRunes("!oG, olleH"))
	fmt.Println(cmp.Diff("Hellow World", "Hello Go"))
	fmt.Println("muliply is: ", muliply)
	fmt.Println("x split is ", xSplit, "y split is", ySplit)

	logging.Log("End of Application")
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined negative type")
	}
	return math.Sqrt(x), nil
}
