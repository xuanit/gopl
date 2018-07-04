package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	MAX = 4
)

const (
	API       = "https://xkcd.com"
	FILE_NAME = "info.0.json"
)

const (
	PATH      = "data"
	ITEM_FILE = "item"
)

type Comic struct {
	Transcript string
}

type Item struct {
	URL        string
	Transcript string
}

func crawlData(url string) (Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Comic{}, fmt.Errorf("getting data at %s: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return Comic{}, fmt.Errorf("getting data at %s: %v", url, resp.Status)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		return Comic{}, fmt.Errorf("Decoding data: %v", err)
	}
	resp.Body.Close()

	return comic, nil
}

func saveComic(comic Comic, url string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {

		return fmt.Errorf("Creating item file: %v", err)
	}
	writer := bufio.NewWriter(file)
	itemData, err := json.Marshal(Item{URL: url, Transcript: comic.Transcript})
	if err != nil {
		return fmt.Errorf("Mashalling data: %v", err)
	}

	fmt.Printf("item Data %s", string(itemData))
	_, err = writer.WriteString(string(itemData))
	if err != nil {
		return fmt.Errorf("Writing file: %v", err)
	}
	writer.Flush()
	file.Close()
	return nil
}

func updateIndex(comic Comic, i int, index map[string][]int) {
	words := strings.Fields(comic.Transcript)
	for _, v := range words {
		index[v] = append(index[v], i)
	}
}

func saveIndex(index map[string][]int) error {
	file, err := os.Create(PATH + "/" + "index.txt")
	if err != nil {
		return fmt.Errorf("Creating index file: %v", err)
	}
	writer := bufio.NewWriter(file)
	indexData, err := json.Marshal(index)
	if err != nil {
		return fmt.Errorf("Mashalling index: %v", err)
	}
	_, err = writer.WriteString(string(indexData))
	if err != nil {
		return fmt.Errorf("Writing index file: %v", err)
	}
	writer.Flush()
	file.Close()
	return nil
}

func buildIndex() error {
	var index = make(map[string][]int)

	for i := 1; i < MAX; i++ {
		url := API + "/" + strconv.Itoa(i) + "/" + FILE_NAME
		comic, err := crawlData(url)
		if err != nil {
			return fmt.Errorf("building index: %v", err)
		}

		fileName := PATH + "/" + ITEM_FILE + "_" + strconv.Itoa(i)
		err = saveComic(comic, url, fileName)
		if err != nil {
			return fmt.Errorf("building index: %v", err)
		}

		updateIndex(comic, i, index)
	}
	err := saveIndex(index)
	if err != nil {
		return fmt.Errorf("builiding index: %v", err)
	}

	return nil
}

func main() {
	err := buildIndex()
	if err != nil {
		log.Fatalf("xkcd: %v", err)
	}
}
