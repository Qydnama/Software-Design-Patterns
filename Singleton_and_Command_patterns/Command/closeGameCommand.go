package main

type CloseGameCommand struct {
	pc Desktop
}

func (o *CloseGameCommand) execute() {
	o.pc.closeGame()
}
