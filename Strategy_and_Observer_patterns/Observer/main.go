package main

func main() {
	newVersion := newVersion("version 2.3.1")
	firstPlayer := &Player{nickname: "Qydnama"}
	secondPlayer := &Player{nickname: "Zooom_15"}

	newVersion.register(firstPlayer)
	newVersion.register(secondPlayer)

	newVersion.updateAvailability()
}
