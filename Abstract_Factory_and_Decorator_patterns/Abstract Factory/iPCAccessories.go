package main

import "fmt"

type IPCAccessories interface {
	makeMouse() IMouse
	makeKeyboard() IKeyboard
}

func GetPCAccessories(brand string) (IPCAccessories, error) {
	if brand == "logitech" {
		return &Logitech{}, nil
	}

	if brand == "razer" {
		return &Razer{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
