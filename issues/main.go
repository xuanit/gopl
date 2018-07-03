package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

const (
	LESS_THAN_A_MONTH_OLD = "less than a month old"
	LESS_THAN_A_YEAR_OLD  = "less than a year old"
	MORE_THAN_A_YEAR_OLD  = "more than a year old"
)

const aMonth = 30 * 24 * time.Hour
const aYear = 365 * 24 * time.Hour

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	resultByCategory := make(map[string][]*Issue)
	for _, v := range result.Items {
		duration := now.Sub(v.CreatedAt)

		if duration < aMonth {
			resultByCategory[LESS_THAN_A_MONTH_OLD] = append(resultByCategory[LESS_THAN_A_MONTH_OLD], v)
			continue
		}

		if duration < aYear {
			resultByCategory[LESS_THAN_A_YEAR_OLD] = append(resultByCategory[LESS_THAN_A_YEAR_OLD], v)
			continue
		}

		resultByCategory[MORE_THAN_A_YEAR_OLD] = append(resultByCategory[MORE_THAN_A_YEAR_OLD], v)

	}
	fmt.Printf("%d issus:\n", result.TotalCount)
	for category, issues := range resultByCategory {
		fmt.Printf("category %s: ", category)
		for _, item := range issues {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.CreatedAt)
		}
	}
}
