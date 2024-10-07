package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *SafeCounter, iAmDone chan bool) {
	if depth <= 0 {
		iAmDone <- true
		return
	}
	// Don't fetch the same URL twice.
	if count := c.Value(url); count > 0 {
		fmt.Println("already visited:", url)
		iAmDone <- true
		return
	}
	c.Inc(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		iAmDone <- true
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	childrenAreDone := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, c, childrenAreDone)
	}
	for _, _ = range urls {
		<-childrenAreDone
	}
	iAmDone <- true
	return
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	childrenAreDone := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, &c, childrenAreDone)
	<-childrenAreDone
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
