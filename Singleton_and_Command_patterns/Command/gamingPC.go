package main

import "fmt"

type GamingPC struct {
	isTurnOn bool
	isGameOn bool
}

func (p *GamingPC) turnOnPC() {
	p.isTurnOn = true
	fmt.Println("The PC turned on.")
}

func (p *GamingPC) turnOffPC() {
	p.isTurnOn = false
	fmt.Println("The PC turned off.")
}

func (p *GamingPC) openGame() {
	p.isGameOn = true
	fmt.Println("The game has opened.")
}

func (p *GamingPC) closeGame() {
	p.isGameOn = false
	fmt.Println("The game has closed.")
}
