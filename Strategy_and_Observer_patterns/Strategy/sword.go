package main

import "fmt"

type Sword struct {
}

func (s *Sword) attack(e *Entity) {
	fmt.Println("Attacking with sword!!")
}
