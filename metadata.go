package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/html"
)

type Metadata struct {
	LastFetch time.Time
	NumImage  int
	NumLink   int
	Site      string
}

func getMetadata(url string) (Metadata, error) {
	metadata := Metadata{Site: url}

	dir := getDirPathFromUrl(url)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist for URL: %s. Downloading HTML...\n", url)
		err = getHtml(url)
		if err != nil {
			return metadata, fmt.Errorf("failed to download HTML for %s: %v", url, err)
		}
	}

	file, err := os.Open(fmt.Sprintf("%s/index.html", dir))
	if err != nil {
		return metadata, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return metadata, err
	}
	metadata.LastFetch = info.ModTime()

	tokens := html.NewTokenizer(file)
	for {
		switch tokens.Next() {
		case html.ErrorToken:
			return metadata, nil
		case html.StartTagToken:
			token := tokens.Token()
			if token.Data == "a" {
				metadata.NumLink++
			}
			if token.Data == "img" {
				metadata.NumImage++
			}
		}
	}
}

func printMetadata(meta Metadata) {
	fmt.Printf("Site: %s\n", meta.Site)
	fmt.Printf("Number of Links: %d\n", meta.NumLink)
	fmt.Printf("Number of Images: %d\n", meta.NumImage)
	fmt.Printf("Last Fetch: %s\n", meta.LastFetch.Format(time.UnixDate))
}
