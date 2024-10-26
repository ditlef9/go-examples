package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fetching data from multiple sources in parallel...")

	// List of URLs to fetch data from
	urls := []string{
		"https://www.example.com",
		"https://www.example.org",
		"https://www.example.net",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
	}

	// Channel to collect results
	// A channel in Go is a built-in data structure that allows goroutines (lightweight threads)
	// to communicate with each other and synchronize their execution.
	// Channels provide a way to send and receive values between goroutines, enabling concurrent programming.
	ch := make(chan string)

	// Start fetching data in parallel
	for _, url := range urls {
		go fetchData(url, ch) // Launching a goroutine for each fetch operation
	}

	// Collect results
	for range urls {
		fmt.Println(<-ch) // Read from the channel and print the result
	}
}
