package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("need to provide 3 arguments")
		fmt.Println("intended usage: crawler <base url> <maxConcurrency> <maxPages>")
		return
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	BASE_URL := os.Args[1]
	maxConcurrencyArg := os.Args[2]
	maxPagesArg := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyArg)
	if err != nil {
		fmt.Printf("error - maxConcurrency: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(maxPagesArg)
	if err != nil {
		fmt.Printf("error - maxPages: %v", err)
		return
	}

	cfg, err := configure(BASE_URL, maxConcurrency, maxPages)
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

	printReport(cfg.pages, BASE_URL)
}
