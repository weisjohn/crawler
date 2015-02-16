/*
Package crawler creates a very simple web crawler which generates sha1 hashes
*/

package main

import (
	"crypto/sha1"
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

// this is a simple map of urls to sha1 hashes
var visited = make(map[string]string)

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

	// read the body of the page into a string
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body", url, err)
		return
	}
	defer resp.Body.Close()

	// get sha1 of that content
	hash := sha1.New()
	hash.Write(body)
	bodysha := hash.Sum(nil)

	// once we have the sha1, put it into the map
	// this helps us not refetch, and also we'll persist this later (in the db)
	visited[url] = fmt.Sprintf("%x", bodysha)

	// TODO: fork https://github.com/JackDanger/collectlinks/ into something else
	// TODO: find links, js, css, img, picture
	// TODO: don't follow links that don't match this domain, or something like that...

}

func main() {

	// preflight checks
	args := os.Args
	if len(args) < 2 {
		bolt("You must pass a valid URL")
	}

	// grab the URL they sent in
	first := args[1]

	_, err := url.ParseRequestURI(first)
	if err != nil {
		bolt("You must pass a valid URL")
	}

	Crawl(first)
	fmt.Println(visited)
}
