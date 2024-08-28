package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(arg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	BASE_URL := arg[0]

	crawledHTML, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Printf("failed to crawl: %v", err)
		os.Exit(1)
	}

	fmt.Println(crawledHTML)
}
