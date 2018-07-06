package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestParseOut(t *testing.T) {
	var result bytes.Buffer
	page := `<html><body><div>header</div><img src="http://google.com" /></body>`
	Prettier(strings.NewReader(page), &result)

	if _, err := html.Parse(&result); err != nil {
		t.Errorf("can't parse output %v, error %v", result.String(), err)
	}
}
