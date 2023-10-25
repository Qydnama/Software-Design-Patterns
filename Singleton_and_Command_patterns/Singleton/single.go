package main

import (
	"fmt"
	"sync"
)

type account struct {
	nickname string
	password int
}

var (
	singleAccount *account
	once          sync.Once
)

func getAccount() *account {
	if singleAccount == nil {
		once.Do(func() {
			singleAccount = &account{
				nickname: "bobir328",
				password: 87654321,
			}
			fmt.Println("Account created.")
		})
	} else {
		fmt.Println("Account already exists.")
	}

	return singleAccount
}
