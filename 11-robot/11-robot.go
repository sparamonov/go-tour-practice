package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var same = make(map[string]bool)
var mux sync.Mutex
var ch = make(chan string)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	
	if _, ok := same[url]; !ok {
		body, urls, err := fetcher.Fetch(url)
		
		mux.Lock()
		same[url] = true
		mux.Unlock()
		
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		
		ch <- fmt.Sprintf("found: %s %q", url, body)
		
		for _, u := range urls {
			go Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	go Crawl("http://golang.org/", 4, fetcher)
	timeout := time.After(1 * time.Second)
	
	for {
		select {
		case r := <- ch:
			fmt.Println(r)
		case <- timeout:
			return
		}
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
