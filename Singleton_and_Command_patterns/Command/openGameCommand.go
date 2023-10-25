package main

type OpenGameCommand struct {
	pc Desktop
}

func (o *OpenGameCommand) execute() {
	o.pc.openGame()
}
