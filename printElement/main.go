package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	Prettier(os.Stdin, os.Stdout)
}

func Prettier(in io.Reader, out io.Writer) {
	n, err := html.Parse(in)
	if err != nil {
		log.Fatalln(err)
	}

	visit(n, startELement(out), endElement(out))
}

var depth int

func startELement(out io.Writer) func(*html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attrBuf bytes.Buffer
			for _, attr := range n.Attr {
				attrBuf.WriteString(attr.Key)
				attrBuf.WriteString("=")
				attrBuf.WriteString(attr.Val)
				attrBuf.WriteString(" ")
			}

			if n.FirstChild == nil {
				fmt.Fprintf(out, "%*s<%s %s", depth*2, " ", n.Data, attrBuf.String())
			} else {
				fmt.Fprintf(out, "%*s<%s %s>\n", depth*2, " ", n.Data, attrBuf.String())
				depth++
			}
			return
		}

		if n.Type == html.TextNode {
			fmt.Fprintf(out, "%*s%s\n", depth*2, " ", n.Data)
			return
		}
	}

}

func endElement(out io.Writer) func(*html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.FirstChild == nil {
				fmt.Fprintf(out, " />\n")
			} else {
				depth--
				fmt.Fprintf(out, "%*s</%s>\n", 2*depth, " ", n.Data)
			}
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
