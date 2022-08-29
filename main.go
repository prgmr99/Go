// Go의 2가지 컨셉

package main // 1. package : 작성할 패키지의 이름을 작성해주는 곳.

import (
	"fmt"
	"prgmr99/learngo/something"
	"strings"
)

// 2. function : 아래의 방법이 Go에서 function을 만드는 방법
func main() {
	fmt.Println("Hello world!")
	something.SayBye() // exported function
	// something.sayHello() // private function
}

// 중요한 것은 기본적으로 package가 어떻게 동작하는 지
// 그리고 왜 Println은 대문자로 시작하는 지 알고 넘어가야한다.

//---------------------------------------------------------------------------------------------------
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

//---------------------------------------------------------------------------------------------------
// return 값의 type도 명시해줘야한다.
// 어떤 종류의 value를 받는지, arguments, return에 명시해줘야한다.

func multiply(a int, b int) int { //(a,b int)도 가능
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) { // arguments를 무제한으로 받는 방법.
	fmt.Println(words)
}

func main3() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}

// totalLength만 반환 받으려 했으면 에러가 발생한다.
// lenAndUpper은 두 개의 값을 return하기 때문이다.
// totalLength, _ := lenAndUpper("nico") -> 이렇게 작성하면 value값을 무시할 수 있다.

// 'defer'은 function이 값을 return하고 나면 실행된다.
// 그렇다고 나중에 출력되는 것은 아니고 아래 문장 우선 출력하고 그 이후 return값 출력한다.
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // length, uppercase를 여기에 적지 않아도 동작한다. 하지만, 명시하는 것이 편하면 하는 것이 좋다.
}

//---------------------------------------------------------------------------------------------------

// Go에서 loop은 오로지 for을 사용하는 것만으로 가능하다.
// range는 array에 loop을 적용할 수 있도록 한다.
// 아래와 같이 작성하면 0부터 시작한다. range는 index를 주기 때문이다.
// 그래서 index를 추가했다.

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAdd3(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main4() {
	superAdd(1, 2, 3, 4, 5, 6)
	superAdd2(1, 2, 3, 4, 5, 6)
	result := superAdd3(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

//---------------------------------------------------------------------------------------------------

// 놀라운 점은 if 속에서 변수를 생성할 수 있다는 것이다.

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // if-else 조건에서만 사용하는 변수라는 것을 알 수 있다.
		return false
	} else {
		return true
	}

	// else를 없애고 바로 return true를 적어도 가능.
	// if age < 18 {
	//	return false
	// }
	// return true  -> 이렇게!
}

func main5() {
	fmt.Println(canIDrink(16))
}

//---------------------------------------------------------------------------------------------------

// switch 활용

func canIDrink1(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIDrink2(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	}
	return false
}

func canIDrink3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main6() {
	fmt.Println(canIDrink1(16))
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink3(16))

}

//---------------------------------------------------------------------------------------------------

// Pointers
// Go 언어는 높은 수준의 코드로 Low-level Programming을 하도록 해준다. -> 니꼬 형님의 생각

/*func main() {
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
}*/

/*func main1() {
	a := 2
	b := 5
	fmt.Println(&a, &b) // a,b의 메모리 주소를 확인할 수 있다.
}*/

/*func main() {
	a := 2
	b := &a		// 메모리 주소를 저장한다.
	fmt.Println(&a, b)	// 같은 값이 나온다.
}*/

// * 은 메모리 주소를 살펴볼 수 있게 한다. 살펴보거나 훑어본다는 느낌으로 사용된다.

// 값을 단순히 복사시키는 것이 아니라
// 메모리에 저장된 object를 서로 똑같이 가지고 싶도록 할때 사용한다.
func main7() {
	a := 2
	b := &a
	*b = 20 // b는 a와 연결되어 있기 때문에 b를 통해서 a를 변경시킬 수 있다.
	fmt.Println(*b)
}

//---------------------------------------------------------------------------------------------------

// Go의 array는 다른 언어들과 차이가 있다.

func main8() {
	names := [5]string{"nico", "lynn", "dal"} // 배열의 크기와 타입, 요소들을 명시해줘야한다.
	names[3] = "alala"
	names[4] = "alala"
	// names[5] = "alala" -> 인덱스 error 발생

	// 배열의 크기를 정하고 싶지 않을 때는 slice를 사용한다. -> [] 이렇게만 사용하면 된다.
	names1 := []string{"j", "k", "q"}
	names1 = append(names1, "a") // 위의 배열에서 값을 추가하고 싶을 때 -> argument: slice, elem
	// names1을 매개변수로 받고 요소를 추가하여 새로운 slice를 return 하는 구조이다. (javascript/python과 다른 구조)

	fmt.Println(names)
	fmt.Println(names1)
}
