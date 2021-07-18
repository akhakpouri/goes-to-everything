package main

import (
	"fmt"

	"github.com/akhakpouri/goes-to-everything/models"
)

func getGun(gynType string) (iGun, error) {
	if gynType == "ak47" {
		return models.NewAk47(), nil
	}
	if gynType == "musket" {
		return models.Newmusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type was passed")
}
