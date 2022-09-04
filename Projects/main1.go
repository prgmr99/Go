package main

import (
	"fmt"

	accounts "github.com/prgmr99/learngo/banking"
)

func main() {
	account := accounts.NewAccount("john")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
