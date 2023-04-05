package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
	results := make(map[string]string)

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
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for key, value := range results {
		fmt.Println(key, value)
	}
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("hit: ", url)
	res, err := http.Get(url)
	status := "OK"

	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- requestResult{url: url, status: status}
}
