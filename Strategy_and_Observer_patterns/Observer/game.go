package main

import "fmt"

type Game struct {
	playerList []PlayerInterface
	name       string
	inList     bool
}

func newVersion(r string) *Game {
	return &Game{name: r}
}

func (i *Game) updateAvailability() {
	fmt.Printf("New version of the game %s is out!\n", i.name)
	i.inList = true
	i.notifyAll()
}

func (i *Game) register(o PlayerInterface) {
	i.playerList = append(i.playerList, o)
}

func (i *Game) deregister(o PlayerInterface) {
	i.playerList = removeFromSlice(i.playerList, o)
}

func (i *Game) notifyAll() {
	for _, player := range i.playerList {
		player.update(i.name)
	}
}

func removeFromSlice(playerList []PlayerInterface, playerToRemove PlayerInterface) []PlayerInterface {
	playerListLength := len(playerList)
	for i, observer := range playerList {
		if playerToRemove.getNickname() == observer.getNickname() {
			playerList[playerListLength-1], playerList[i] = playerList[i], playerList[playerListLength-1]
			return playerList[:playerListLength-1]
		}
	}
	return playerList
}
