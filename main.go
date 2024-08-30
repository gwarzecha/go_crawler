package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Println("no website provided")
		return
	} else if len(arg) > 1 {
		fmt.Println("too many arguments provided")
		return
	}

	BASE_URL := arg[0]

	const maxConcurrency = 10
	cfg, err := configure(BASE_URL, maxConcurrency)
	if err != nil {
		fmt.Printf("error - configure: %v", err)
		return
	}

	fmt.Printf("beginning crawl of: %s\n", BASE_URL)

	// Kicks off the crawling process with initial goroutine and url
	cfg.wg.Add(1)
	go cfg.crawlPage(BASE_URL)
	// Ensures the program doesn't exit before all goroutines are done
	cfg.wg.Wait()

	for normalizedRawCurrentURL, count := range cfg.pages {
		fmt.Printf("%s: %d\n", normalizedRawCurrentURL, count)
	}
}
