package main

import "fmt"

type Player struct {
	nickname string
}

func (p *Player) update(itemName string) {
	fmt.Printf("Player %s can get a new update %s\n", p.nickname, itemName)

}

func (p *Player) getNickname() string {
	return p.nickname
}
