// Go의 2가지 컨셉

package main // 1. package : 작성할 패키지의 이름을 작성해주는 곳.

import (
	"fmt"
	"prgmr99/learngo/something"
)

// 2. function : 아래의 방법이 Go에서 function을 만드는 방법
func main() {
	fmt.Println("Hello world!")
	something.SayBye()   // exported function
	something.sayHello() // private function
}
