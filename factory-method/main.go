package main

import (
	"fmt"
	"github.com/akhakpouri/goes-to-everything/models"
)

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g models.iGun) {
	fmt.Printf("Gun: %s", g.Name)
	fmt.Println()
	fmt.Printf("Power: %d", g.Power)
	fmt.Println()
}
