package main

type TurnOffCommand struct {
	pc Desktop
}

func (o *TurnOffCommand) execute() {
	o.pc.turnOffPC()
}
