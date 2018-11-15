package main

import (
	"fmt"
	"net/http"
)

func getPage(url string) int {
	resp, _ := http.Get(url)
	return resp.StatusCode
}

// START OMIT
var getPageWrapper = getPage

func crawl(urls ...string) {
	for _, url := range urls {
		getPageWrapper(url)
	}
}

func main() {
	getPageWrapper = func(url string) int {
		fmt.Println("I was called indeed")
		return 200
	}
	crawl("http://www.facebook.com", "http://matters.com")
}

// END OMIT
