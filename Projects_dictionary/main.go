package main

import (
	"fmt"

	"github.com/prgmr99/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//fmt.Println(dictionary["first"]) 	// 이렇게 할 수도 있다.
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else { // else는 if 대괄호에 붙여서 써야한다.
		fmt.Println(definition)
	}
}
