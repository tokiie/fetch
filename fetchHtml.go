package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getDirPathFromUrl(rawURL string) string {
	splitOn := "://"
	index := strings.Index(rawURL, splitOn)

	if index != -1 {
		rawURL = rawURL[index+len(splitOn):]
	}

	rawURL = strings.ReplaceAll(rawURL, "/", "_")

	return fmt.Sprintf("%s.html", rawURL)
}

func processURL(rawURL string) (string, error) {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	if parsedURL.Host == "" {
		return "", fmt.Errorf("invalid URL, missing host: %s", rawURL)
	}

	return parsedURL.String(), nil
}

func getHtml(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	targetDir := getDirPathFromUrl(url)
	if err := os.Mkdir(targetDir, 0777); err != nil {
		return err
	}
	if err := os.WriteFile(fmt.Sprintf("%s/index.html", targetDir), body, 0666); err != nil {
		return err
	}
	return nil
}
