package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	n, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	visit(n, startELement, endElement)

}

var depth int

func startELement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attrBuf bytes.Buffer
		for _, attr := range n.Attr {
			attrBuf.WriteString(attr.Key)
			attrBuf.WriteString("=")
			attrBuf.WriteString(attr.Val)
			attrBuf.WriteString(" ")
		}

		if n.FirstChild == nil {
			fmt.Printf("%*s<%s %s", depth*2, " ", n.Data, attrBuf.String())
		} else {
			fmt.Printf("%*s<%s %s>\n", depth*2, " ", n.Data, attrBuf.String())
			depth++
		}
		return
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, " ", n.Data)
		return
	}

}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf(" />\n")
		} else {
			depth--
			fmt.Printf("%*s</%s>\n", 2*depth, " ", n.Data)
		}
	}
}

func visit(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pre, post)
	}

	if post != nil {
		post(n)
	}

}
