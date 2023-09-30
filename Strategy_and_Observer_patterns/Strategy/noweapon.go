package main

import "fmt"

type NoWeapon struct {
}

func (n *NoWeapon) attack(e *Entity) {
	fmt.Println("Attacking without weapon, just ARMS!!!")
}
