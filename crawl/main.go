package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
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

func savePage(resp *http.Response) {
	if resp.Request.Host == "golang.org" {
		parts := strings.Split(url, SEPARATOR)
		fmt.Printf("%v parts %v\n", len(parts), parts)
		if len(parts) == 3 {
			parts = append(parts, "index")
		}
		path := strings.Join(parts[3:len(parts)-1], SEPARATOR)
		err := os.MkdirAll(path, os.ModeDir)
		if err != nil {
			log.Printf("creating dir %s: %v", path, err)
		}
		file, err := os.Create(path + "/" + parts[len(parts)-1])
		if err != nil {
			log.Printf("creating file %s: %v", parts[len(parts)-1], err)
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

	if resp.Request.Host == "golang.org" {
		parts := strings.Split(url, SEPARATOR)
		fmt.Printf("%v parts %v\n", len(parts), parts)
		if len(parts) == 3 {
			parts = append(parts, "index")
		}
		path := strings.Join(parts[3:len(parts)-1], SEPARATOR)
		err := os.MkdirAll(path, os.ModeDir)
		if err != nil {
			log.Printf("creating dir %s: %v", path, err)
		}
		file, err := os.Create(path + "/" + parts[len(parts)-1])
		if err != nil {
			log.Printf("creating file %s: %v", parts[len(parts)-1], err)
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
