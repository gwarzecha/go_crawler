package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse current url: %v", err)
	}

	scheme := currentURL.Scheme

	// skip sites that are not on the same domain as the base url
	if currentURL.Host != cfg.baseURL.Host {
		return
	}

	normalizedRawCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize current url: %v", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedRawCurrentURL)
	if !isFirst {
		return
	}

	fmt.Printf("Crawling: %v\n", normalizedRawCurrentURL)

	fullPath := scheme + "://" + normalizedRawCurrentURL
	returnedHTML, err := getHTML(fullPath)
	if err != nil {
		fmt.Printf("failed to get HTML for URL %s: %v\n", normalizedRawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(returnedHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("failed to get URLs from HTML for URL %s: %v\n", normalizedRawCurrentURL, err)
	}

	// Continues crawling the page for new URLs and spawns new goroutines to crawl each discovered URL
	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
