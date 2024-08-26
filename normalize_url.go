package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")
	return fullPath, nil
}
