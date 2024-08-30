package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("failed to parse base url: %v", err)
	}

	scheme := baseURL.Scheme

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse current url: %v", err)
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normalizedRawCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize current url: %v", err)
		return
	}

	if _, visited := pages[normalizedRawCurrentURL]; visited {
		pages[normalizedRawCurrentURL]++
		return
	}

	pages[normalizedRawCurrentURL] = 1

	fmt.Printf("Crawling: %v\n", normalizedRawCurrentURL)

	fullPath := scheme + "://" + normalizedRawCurrentURL
	returnedHTML, err := getHTML(fullPath)
	if err != nil {
		fmt.Printf("failed to get HTML for URL %s: %v\n", normalizedRawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(returnedHTML, fullPath)
	if err != nil {
		fmt.Printf("failed to get URLs from HTML for URL %s: %v\n", normalizedRawCurrentURL, err)
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
