package main

import (
	"flag"
	"fmt"
	"github.com/tokiie/fetch/utils"
)

func main() {
	showMetadata := flag.Bool("metadata", false, "print metadata for downloaded url")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("No url provided, exiting...")
	}

	for _, url := range urls {
		processedURL, err := utils.ProcessURL(url)
		if err != nil {
			fmt.Printf("%s %s\n", url, err)
			return
		}
		switch *showMetadata {
		case true:
			m, err := utils.GetMetadata(processedURL)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
			utils.PrintMetadata(m)
		case false:
			err := utils.GetHtml(processedURL)
			if err != nil {
				fmt.Printf("%s %s\n", url, err)
			}
		}
	}
}
