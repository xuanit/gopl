package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(resp *http.Response) ([]string, error) {

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing in as HTML")
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

const SEPARATOR = "/"

func isSlash(r rune) bool {
	return r == '/'
}

const DATA = "./data"

func savePage(resp *http.Response) {
	canonicalURL := strings.TrimRightFunc(resp.Request.URL.Path, isSlash)

	if resp.Request.Host == "golang.org" {
		fmt.Printf("path %v\n", canonicalURL)
		dir := DATA + filepath.Dir(canonicalURL)

		fmt.Printf("dir %s\n", dir)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("creating dir2 %s: %v", dir, err)
		}

		file, err := os.Create(DATA + canonicalURL)
		if err != nil {
			log.Printf("creating file %s: %v", DATA+canonicalURL, err)
		}

		var body []byte
		out := bufio.NewWriter(file)
		_, err = resp.Body.Read(body)
		if err != nil {
			log.Printf("reading response: %v", err)
		}
		_, err = out.Write(body)
		if err != nil {
			log.Printf("writing data: %v", err)
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Printf("getting %s: %s", url, resp.Status)
	}

	savePage(resp)

	list, err := Extract(resp)
	if err != nil {
		resp.Body.Close()
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
