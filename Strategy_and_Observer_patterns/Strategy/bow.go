package main

import "fmt"

type Bow struct {
}

func (b *Bow) attack(e *Entity) {
	fmt.Println("Attacking with bow!!!")
}
