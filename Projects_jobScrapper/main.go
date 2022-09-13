package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// job scrapper

// 1. 페이지 수를 알아야한다.
// 2. 각각의 페이지를 방문한 후, 그 페이지에서 각각의 job들을 추출한다.
// 3. 추출한 job들을 엑셀에 집어넣을 것.

// goquery를 사용할 때마다 에러체크를 해줘야한다.

var baseURL string = "https://kr.indeed.com/jobs?q=%ED%8C%8C%EC%9D%B4%EC%8D%AC&l&vjk=058526105f75dc1d"

func main() {
	getPages()
}

func getPages() int {
	resp, err := http.Get(baseURL)
	checkErr(err)
	checkCode(resp)

	defer resp.Body.Close() // 이걸 닫아줘야한다. 그래야지 메모리가 세어나가는 것을 막을 수 있다.

	doc, err := goquery.NewDocumentFromReader(resp.Body) // Body는 기본적으로 I/O이다.
	checkErr(err)

	//fmt.Println(doc)
	doc.Find(".pagination")
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", resp.StatusCode)
	}
}

/*func getHttp(url string) *http.Response {
	fmt.Println("[GET]Request :", url)
	req, rErr := http.NewRequest("GET", url, nil)
	checkErr(rErr)
	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	res, err := client.Do(req)
	checkErr(err)
	checkCode(res)

	return res
}*/
