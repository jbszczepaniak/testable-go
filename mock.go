package main

import (
	"net/http"
)

// START OMIT
func getPage(url string) int {
	resp, _ := http.Get(url)
	return resp.StatusCode
}

func crawl(urls ...string) {
	for _, url := range urls {
		getPage(url)
	}
}

func main() {
	// getPage = func(url string) int { return 200 }
	crawl("http://www.facebook.com", "http://matters.com")
}

// END OMIT
