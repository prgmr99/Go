package main

import (
	"fmt"
	"log"
)

func main() {
	account := accounts.NewAccount("john") //accounts.NewAccount("john")
	account.Deposit(10)
	fmt.Println(account.Balance())

	// Go에서 error를 handling하는 법.
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err) // Println()을 호출하고 프로그램을 종료시킨다. fmt.Println(err)로 작성하면 can't withdraw만 바로 나온다.
	}
	fmt.Println(account.Balance())
}
