package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func ping(url string) bool {
	time.Sleep(2 * time.Second)
	if strings.Contains(url, "face") {
		return false
	}
	return true
}

// check status of these sites and respond to traffic
func main() {
	var wg sync.WaitGroup
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
	}

	for _, val := range links {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println(url, ping(url))
		}(val)
	}

	wg.Wait()
}
