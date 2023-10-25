package main

type TurnOnCommand struct {
	pc Desktop
}

func (o *TurnOnCommand) execute() {
	o.pc.turnOnPC()
}
