package main

import (
	"fmt"

	"github.com/prgmr99/learngo/mydict"
)

// banking
/*import (
	"fmt"

	accounts "github.com/prgmr99/learngo/banking"
)

func main() {
	account := accounts.NewAccount("john")
	account.Deposit(10)
	//fmt.Println(account.Balance())

	// Go에서 error를 handling하는 법.
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err) // Println()을 호출하고 프로그램을 종료시킨다. fmt.Println(err)로 작성하면 can't withdraw만 바로 나온다.
	}
	//fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}*/

// dictionary
/*func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//fmt.Println(dictionary["first"]) 	// 이렇게 할 수도 있다.
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else { // else는 if 대괄호에 붙여서 써야한다.
		fmt.Println(definition)
	}
}*/

// Add
/*func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}*/

// Update
/*func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")            // 먼저 baseword를 first로 정의하고
	err := dictionary.Update(baseword, "Second") // 그 후 baseword를 second로 update하고 싶다.
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word) // 그리고 word를 출력했을 때 Second가 출력되어야한다.

}*/

// Delete
func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
