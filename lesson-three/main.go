package main

import (
	"fmt"

	"github.com/akhakpouri/goes-to-everything/helpers"
	"github.com/akhakpouri/goes-to-everything/models"
)

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := models.Vertex{1, 2}
	v.X = 4

	fmt.Println(v)

	a := make([]int, 5)
	helpers.PrintSlices("a", a)

	b := make([]int, 0, 5)
	helpers.PrintSlices("b", b)

	c := b[:2]
	helpers.PrintSlices("c", c)

	d := c[2:5]
	helpers.PrintSlices("d", d)

}
