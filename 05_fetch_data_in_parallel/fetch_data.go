package main

import (
	"fmt"
	"net/http"
	"time"
)

// Function to fetch data from a URL
func fetchData(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	ch <- fmt.Sprintf("Fetched %s in %v seconds with status code %d", url, elapsed.Seconds(), resp.StatusCode)
}
