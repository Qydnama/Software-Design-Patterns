package main

import "fmt"

func getChips(chipsType string) (IChips, error) {
	if chipsType == "lays" {
		return newLays(), nil
	}
	if chipsType == "hrustim" {
		return newHrustim(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
