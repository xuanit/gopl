package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const APIURL = "https://api.github.com/repos/xuanit/gopl/issues"

const TOKEN = "4e003ad1c8f9f6bfe9902d67ac6491093ffcd3ba"

type GetIssueResult struct {
	Items []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string `json:"state"`
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

func getIssue(number string) (*Issue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", APIURL+"/"+number, nil)
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

	var result Issue
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func createIssue(title string) error {
	client := &http.Client{}

	body := Issue{Title: title}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", APIURL, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token "+TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return fmt.Errorf("createIssue: %v", resp.Status)
	}
	resp.Body.Close()
	return nil
}

func updateIssue(number string, title, state string) error {
	client := &http.Client{}

	body := Issue{Title: title, State: state}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", APIURL+"/"+number, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token "+TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resErr, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return fmt.Errorf("createIssue: %s", string(resErr))
	}
	resp.Body.Close()
	return nil
}

func deleteIssue(number string) error {
	issue, err := getIssue(number)
	if err != nil {
		return fmt.Errorf("getting issue %v", err)
	}
	return updateIssue(number, issue.Title, "closed")
}

const (
	FILE = "message.txt"
)

func getMessage() (string, error) {
	cmd := exec.Command("vim", FILE)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("starting vim %v\n", err)
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("waiting for vim %v\n", err)
	}

	file, err := os.Open(FILE)
	if err != nil {
		return "", fmt.Errorf("opening file %v\n", err)
	}
	msgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("reading file %v\n", err)
	}
	file.Close()
	return string(msgBytes), nil
}

const (
	CREATE = "create"
	UPDATE = "update"
	DELETE = "delete"
	LIST   = "list"
	VIEW   = "view"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case CREATE:
		{
			msg, err := getMessage()
			if err != nil {
				log.Fatalf("getting message: %v", err)
			}
			err = createIssue(msg)
			if err != nil {
				fmt.Printf("create issue : %v", err)
			}
		}
	case LIST:
		{
			issues, err := getIssues()
			if err != nil {
				fmt.Printf("getIssues : %v", err)
			}
			fmt.Printf("#\ttitle\n")
			for _, v := range issues {
				fmt.Printf("%d\t%s\n", v.Number, v.Title)
			}
		}
	case DELETE:
		{
			err := deleteIssue(os.Args[2])
			if err != nil {
				log.Fatalf("deleting issue: %v", err)
			}
		}
	case UPDATE:
		{
			msg, err := getMessage()
			if err != nil {
				log.Fatalf("getting message: %v", err)
			}
			err = updateIssue(os.Args[2], msg, "")
			if err != nil {
				fmt.Printf("updating issue issue : %v", err)
			}
		}
	case VIEW:
		{
			issue, err := getIssue(os.Args[2])
			if err != nil {
				log.Fatalf("deleting issue: %v", err)
			}
			jsonString, err := json.Marshal(issue)

			if err != nil {
				log.Fatalf("mashaling data: %v", err)
			}
			fmt.Printf("issue details:\n%s", jsonString)
		}
	}

}
