package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		go getAccount()
	}
	time.Sleep(time.Second * 4)
}
