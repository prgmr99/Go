package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("request failed")

// URL checker
func main11() {
	var results = make(map[string]string) // 이런 식으로 하지 않으면 map은 nil이 되어버린다.
	// nil인 map에는 어떤 값을 넣을 수가 없게된다. 중요한 내용!!
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	// // 초기화되지 않은 map에 어떤 값을 넣을 수 없다. -> 넣으면 에러발생(어디서 발생한지 몰라서 알려주지 않음.)
	// 비어있는 empty map을 만들기 위해서는 끝에 {}를 해주면 된다.
	// 이 방법(var 변수 선언)말고 할 수 있는 방법은 make()를 하는 것.
	// make()는 map을 만들어주는 함수.

	for _, url := range urls { // _, 을 안써주면 value를 얻을 수 없다. url -> index
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

// 함수를 만들어보자
// 웹사이트로 접속(hit, 인테넷 웹 서버의 파일 1개에 접속하는 것)하고
// 그 결과를 알려주는 함수.

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}

// Go에서 작업을 최적화하는 방법은 동시에 작업을 처리하는 것.
// Goroutines이란 기본적으로 다른 함수와 동시에 실행시키는 함수.

func main12() {
	// sexyCount("nico")
	// sexyCount("fynn")
	// 위와 같이 하면 20초가 걸린다.

	go sexyCount("nico") // 이렇게 하면 동시에 실행된다. -> 10 소요. go만 추가하면 된다.
	sexyCount("fynn")    // 두 번째 sexycount에도 go를 추가하면 아무것도 출력되지 않는다.
	// 프로그램이 종료되었기 때문에 그렇다. -> main 함수가 작업을 마쳤기 때문이다.
	// Goroutines는 프로그램이 작동하는 동안만 유효하다.
	// 어떤 작업을 동시에 진행하려고하면 메인 함수는 다른 goroutines를 기다려주지 않는다.
	// 메인 함수가 끝나면 goroutines도 소멸된다.
	// 위의 코드가 실행되는 이유는 메인 함수가 fynn을 카운팅하고 있기 때문이다.
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

// channel: goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법.
// 채널은 파이프 같은 것 - 메세지를 보내거나 받을 수 있다.
// 아래의 코드는 기본적으로 이 person이 sexy한지 확인하는 것

func main13() {
	// c := make(chan bool)                 // c: channel name 아무거나 해도 괜찮다.
	c := make(chan string)
	people := [2]string{"nico", "flynn"} // 배열
	for _, person := range people {
		go isSexy(person, c) // result := go isSexy(person) 같이 작성할 수 없다. 이 isSexy함수의 작업이 끝나면
		// 아래의 함수의 true값을 channel을 통해 보내고 싶은 것.
	}
	// result := <-c  // 채널에서 보낸 메세지를 result에 저장한다.
	// <-c, 채널로부터 메세지를 가져오는 것. 메세지를 받는 것 = blocking operation.

	// loop 활용하여 작성해보자.
	//resultOne := <-c
	//resultTwo := <-c
	//fmt.Println("Received this message", resultOne)
	//fmt.Println("Received this message", resultTwo)

	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for messages", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) { // c: argument name, chan: channel type, bool: return type
	time.Sleep(time.Second * 10)
	// fmt.Println(person)
	// c <- true // c에 true를 보내준다.
	c <- person + " is sexy"
}

// 채널을 만들었고 그 채널은 두 개의 함수(isSexy nico, isSexyflynn로 보내진다.
// 이 두 함수는 5초뒤에 나에게 true를 보내준다.

// Rules of Channel
// 1. 먼저 해야하는 것을 먼저 해치운다.
// channel과 goroutines가 우선이다.
// 만약, 메인함수가 끝이나면 나의 goroutines는 무의미해진다. 끝이났던 끝이나지 않았던.
// 2. 내가 받을 데이터와 보낼 데이터에 대해서
// 어떤 타입을 받을 것인지를 구체적으로 지정해줘야한다는 것.
// 3. 메세지를 채널로 보내는 방법은 메세지에 화살표를 붙여서 채널로 항햐게 작성하면된다.
// 메세지를 받을 수 있는 곳이 없더라도 메세지를 보낼 수 있다.
// blocking operation: 프로그램(이 경우에는 메인함수)이 뭔가를 받기 전까지 동작을 멈춘다는 것.

// URL checker

type requestResult struct {
	url    string
	status string
}

func main15() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL1(url, c) // hitURL로 채널을 보낸다. ",c"를 추가함으로써
	}
	// 이렇게 작성하고 바로 실행해보면 바로 종료된다. 왜? -> 아무도 메세지를 기다리지 않기 때문이다.
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL1(url string, c chan<- requestResult) { // chan<- : send only
	// c <- result{} 이런 식으로 작성하면 채널로 메세지를 보낼 수 있다.
	// fmt.Println(<-c) 이런 식으로 작성하면 채널에서 메세지를 받을 수 있다.
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} else {
		c <- requestResult{url: url, status: status}
	}
}
