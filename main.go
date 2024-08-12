package main

import (
	"flag"
	"fmt"
)

func main() {
	showMetadata := flag.Bool("metadata", false, "print metadata for downloaded url")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("No url provided, exiting...")
	}

	for _, url := range urls {
		processedURL, err := processURL(url)
		if err != nil {
			fmt.Printf("%s %s\n", url, err)
			return
		}
		switch *showMetadata {
		case true:
			m, err := getMetadata(processedURL)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
			printMetadata(m)
		case false:
			err := getHtml(processedURL)
			if err != nil {
				fmt.Printf("%s %s\n", url, err)
			}
		}
	}
}
