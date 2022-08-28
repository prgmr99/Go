// Go의 2가지 컨셉

package main // 1. package : 작성할 패키지의 이름을 작성해주는 곳.

import (
	"fmt"
	"prgmr99/learngo/something"
)

// 2. function : 아래의 방법이 Go에서 function을 만드는 방법
func main() {
	fmt.Println("Hello world!")
	something.SayBye() // exported function
	// something.sayHello() // private function
}

// 중요한 것은 기본적으로 package가 어떻게 동작하는 지
// 그리고 왜 Println은 대문자로 시작하는 지 알고 넘어가야한다.

// constant(상수): 변수지만 값을 바꿀 수 없다.
// variable(변수): 값을 변경할 수 있다.

// Go는 내가 작성한 값의 타입을 찾아내려는 시도를 한다. -> Go = type language
// Go에게 '타입이 무엇이다'라는 것을 알려줘야한다. Java/C처럼

func main2() {
	const name string = "john" // 상수 정의
	// name = "nico" -> error 발생

	var name1 string = "john" // 변수 정의
	name1 = "yeom"            // 변수라서 error 발생 x

	// 매우 중요
	// 위와 같이 var name1 string = "john"으로 사용하는 것 대신
	// func안에서는
	name2 := "johnyeom" // 이렇게 작성할 수도 있다. type은 Go가 찾아서 사용한다. 내가 임의로 변경할 수 없다.

	// 축약형은 오직 func안에서 그리고 변수에만 가능하다.

	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)
}
