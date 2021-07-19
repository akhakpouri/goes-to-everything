package main

import (
	"fmt"

	"github.com/akhakpouri/goes-to-everything/factory/models"
)

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g models.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}

func getGun(gynType string) (models.IGun, error) {
	if gynType == "ak47" {
		return models.NewAk47(), nil
	}
	if gynType == "musket" {
		return models.NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type was passed")
}
