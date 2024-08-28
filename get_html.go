package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed with: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed with: %w", err)
	}
	res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("failed with a status code of: %v", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("content type is: %v, expected text/html", res.Header.Get("Content-Type"))
	}

	return string(body), nil
}
