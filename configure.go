package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

// Method on the config struct, can modify the config instance it's called on
func (cfg *config) addPageVisit(normalizedRawCurrentURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedRawCurrentURL]; visited {
		cfg.pages[normalizedRawCurrentURL]++
		return false
	}

	cfg.pages[normalizedRawCurrentURL] = 1
	return true
}

// Initializes and returns a new config struct, used to manage crawler's state and control concurrency
func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("failed to parse base url: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
