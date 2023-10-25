package main

import "fmt"

func main() {
	gamingPc := &GamingPC{}

	turnOnCommand := &TurnOnCommand{gamingPc}
	turnOffCommand := &TurnOffCommand{gamingPc}
	openGameCommand := &OpenGameCommand{gamingPc}
	closeGameCommand := &CloseGameCommand{gamingPc}

	turnOnButton := &Button{turnOnCommand}
	turnOffButton := &Button{turnOffCommand}
	openGameButton := &Button{openGameCommand}
	closeGameButton := &Button{closeGameCommand}

	turnOnButton.press()
	openGameButton.press()
	fmt.Println("You played minecraft with your bros")
	closeGameButton.press()
	turnOffButton.press()

	//The PC turned on.
	//The game has opened.
	//You played minecraft with your bros
	//The game has closed.
	//The PC turned off.

}
