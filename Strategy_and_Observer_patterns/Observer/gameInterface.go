package main

type GameInterface interface {
	register(player PlayerInterface)
	deregister(player PlayerInterface)
	notifyAll()
}
