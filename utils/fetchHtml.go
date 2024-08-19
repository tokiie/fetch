package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetDirPathFromUrl(rawURL string) string {
	rawURL = filepath.Base(rawURL)

	return fmt.Sprintf("%s.html", rawURL)
}

func ProcessURL(rawURL string) (string, error) {
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

func GetHtml(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // Always close the response body

	return SaveHtml(resp.Body, url)
}

func SaveHtml(r io.Reader, url string) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	targetDir := GetDirPathFromUrl(url)
	if err := os.MkdirAll(targetDir, 0777); err != nil { // MkdirAll to handle nested dirs
		return err
	}

	filePath := fmt.Sprintf("%s/index.html", targetDir)
	return writeToFile(filePath, body)
}

func writeToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}
