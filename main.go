package main

import (
	"fmt"
	"net/http"
	"time"
)

type URLStatus struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
}

func checkURL(url string, ch chan<- URLStatus) {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	status := URLStatus{
		URL:          url,
		StatusCode:   0,
		ResponseTime: duration,
	}

	if err != nil {
		fmt.Printf("Error checking %s: %v\n", url, err)
	} else {
		status.StatusCode = resp.StatusCode
		defer resp.Body.Close()
	}

	ch <- status
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
	}

	ch := make(chan URLStatus)
	for _, url := range urls {
		go checkURL(url, ch)
	}

	for range urls {
		status := <-ch
		fmt.Printf("URL: %s, Status Code: %d, Response Time: %s\n",
			status.URL, status.StatusCode, status.ResponseTime)
	}
}

