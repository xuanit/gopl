package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	node, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	visit(node)

}

func visit(n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		fmt.Printf("%v\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
