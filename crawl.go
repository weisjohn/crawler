/*
Package crawler creates a very simple web crawler
that is made for demo usage. This could evolve into the backend
for my competitor track website thingie
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func bolt(message string) {
	fmt.Println(message)
	os.Exit(1)
}

// this is a simple map of urls to md5 hashes
// to see if content has changed
var visited = make(map[string]int)

func Crawl(url string) {

	// if we have already visited it, bounce
	if _, ok := visited[url]; ok {
		return
	}

	// get the page
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET", url, "failed", err)
		return
	}

	// read the body of the page
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body", url, err)
		return
	}

	fmt.Println("body")
	fmt.Println(string(body))

}

func main() {

	// preflight checks
	args := os.Args
	if len(args) < 2 {
		bolt("You must pass a valid URL")
	}

	first := args[1]

	_, err := url.ParseRequestURI(first)
	if err != nil {
		bolt("You must pass a valid URL")
	}

	Crawl(first)
	fmt.Println(visited)
}
