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

	counts := visit(map[string]int{}, node)

	fmt.Printf("node\tcount\n")
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func visit(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counts = visit(counts, c)
	}
	return counts
}
