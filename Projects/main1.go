package main

import (
	"fmt"

	accounts "github.com/prgmr99/learngo/banking"
)

func main() {
	account := accounts.NewAccount("john")
	fmt.Println(account)
}
