package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// var limit int = 50
var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?search_area=main&search_done=y&search_optional_item=n&searchType=search&searchword={}&recruitPage={}&recruitSort=relation&recruitPageCount=100"

func main() {
	getPages()
}

func getPage(page int) {

}

func getPages() int {
	req, rErr := http.NewRequest("GET", baseURL, nil)
	checkErr(rErr)

	// 프록시로 호출
	fmt.Println("called Proxy")
	purl, err := url.Parse(baseURL)
	checkErr(err)

	fmt.Println("create Http Client")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	res, err := client.Do(req)
	checkErr(err)
	checkStatus(res)

	fmt.Println("called goquery")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("a").Length())
	})

	return 0
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("caused error")
		log.Fatalln(err)
	}
}

func checkStatus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
