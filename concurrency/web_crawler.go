package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited map[string]bool, mu *sync.Mutex, donech chan bool) {
	fmt.Println("crawl:", url, " depth: ", depth)
	if depth <= 0 {
		donech <- true
		return
	}
	// Don't fetch the same URL twice.
	//mu.Lock()
	if _, ok := visited[url]; ok {
		fmt.Println("already visited:", url)
		donech <- true
		return
	}
	visited[url] = true
	//mu.Unlock()
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		donech <- true
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	children_are_done := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, visited, mu, children_are_done)
	}
	for _, _ = range urls {
		<-children_are_done
	}
	donech <- true
	return
}

func main() {
	visited := make(map[string]bool)
	var mu sync.Mutex
	donech := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, visited, &mu, donech)
	<-donech
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
