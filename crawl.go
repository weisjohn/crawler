/*
Package crawler recursively crawls a URL returning a `map[string]string` of with URIs as keys and sha1 hashes as values.
*/

package crawler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/weisjohn/cssrefs"
	"github.com/weisjohn/htmlrefs"
)

func bolt(message string) {
	fmt.Println(message)
	os.Exit(1)
}

// if the contenttype is not known, normazlie the contenttype based on the server-response
func divine(contenttype, header string) string {

	// split the mime-type descript on the /, e.g. text/html == "html"
	header = strings.Split(header, "/")[1]
	header = strings.Split(header, ";")[0]

	// if the contenttype is not explicitly set, determine from header
	// if contenttype != header {
	// 	fmt.Printf("content-type: %q != server-sent: %q\n", contenttype, header)
	// }

	if contenttype != "" {
		return contenttype
	}

	return header
}

// read the body, return a base-64 sha1 of the contents
func hashBody(httpBody io.Reader, url string) string {

	// read the body of the page into a string
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		fmt.Printf("Error reading body", url, err)
		return "foo"
	}

	// get sha1 of that content
	hash := sha1.New()
	hash.Write(body)
	bodysha := hash.Sum(nil)
	return fmt.Sprintf("%x", bodysha)
}

// a simple map of html tokens to content types
var resourcemap = map[string]string{
	"link":   "css",
	"script": "javascript",
	"a":      "html",
}

// resources
type Resource struct{ URI, Type string }

// return a slice of urls to contenttypes
func getRefs(contenttype string, httpBody io.Reader) []Resource {

	resources := make([]Resource, 0)

	// TODO: expand

	switch contenttype {
	case "html":
		refs := htmlrefs.All(httpBody)
		for _, ref := range refs {
			resources = append(resources, Resource{URI: ref.URI, Type: resourcemap[ref.Token]})
		}
	case "css":
		refs := cssrefs.All(httpBody)
		for _, ref := range refs {
			resources = append(resources, Resource{URI: ref.URI, Type: ref.Token})
		}
	}

	return resources
}

// rather than cutting out of the previous slice, we have to normalize all URLs anyways
// so just create a new map of resources and return it
func resolveRefs(refs []Resource, s string) []Resource {

	// use
	u, err := url.Parse(s)
	if err != nil {
		bolt("problem parsing url")
	}

	resources := make([]Resource, 0)

	for _, ref := range refs {

		// all uris must be content-bearing
		uri := ref.URI
		if uri == "" {
			uri = "/"
		}

		r, err := u.Parse(uri)
		if err != nil {
			fmt.Println("problem resolving", ref.URI)
		}

		// for html, hosts must be the same, otherwise add
		if ref.Type != "html" || (ref.Type == "html" && r.Host == u.Host) {
			resources = append(resources, Resource{URI: r.String(), Type: ref.Type})
		}
	}
	return resources
}

// this is a simple map of urls to sha1 hashes
var visited map[string]string

func crawl(url, contenttype string) {

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

	// discover conflict between referenced content-type and server reply
	contenttype = divine(contenttype, resp.Header["Content-Type"][0])

	// grab the refs depending on the type of content
	refs := getRefs(contenttype, resp.Body)

	// resolve to full-urls
	refs = resolveRefs(refs, url)

	// once we have the sha1, put it into the map
	// this helps us not refetch, and also we'll persist this later (in the db)
	visited[url] = hashBody(resp.Body, url)
	defer resp.Body.Close()

	// use a simply bool channel to fan out
	done := make(chan bool)

	// fan out, recursion
	for _, ref := range refs {
		go func(ref Resource) {
			crawl(ref.URI, ref.Type)
			done <- true
		}(ref)
	}

	// sync back up
	for _, _ = range refs {
		<-done
	}

	return

}

// Crawl recusively crawls a URI returning a map of URIs to sha1 hashes.
func Crawl(uri string) map[string]string {
	visited = make(map[string]string)
	crawl(uri, "")
	return visited
}
