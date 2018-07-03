package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const APIURL = "https://api.github.com/repos/xuanit/gopl/issues"

const TOKEN = ""

type GetIssueResult struct {
	Items []*Issue
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

func getIssues() ([]*Issue, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", APIURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "token "+TOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getIssues: %v", resp.Status)
	}

	var result []*Issue
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func createIssue(title string) error {
	client := &http.Client{}

	body := Issue{Title: title}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return err
	}

	fmt.Printf("json %v\n", body)

	req, err := http.NewRequest("POST", APIURL, &buf)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "token "+TOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("request %v\n", req.Method)

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return fmt.Errorf("createIssue: %v", resp.Status)
	}
	resp.Body.Close()
	return nil
}

func main() {
	err := createIssue("my first issue from go")
	if err != nil {
		fmt.Printf("create issue : %v", err)
	}

	issues, err := getIssues()
	if err != nil {
		fmt.Printf("getIssues : %v", err)
	}
	for _, v := range issues {
		fmt.Println(v.Title)
	}
}
