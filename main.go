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

	pages := make(map[string]int)

	crawlPage(BASE_URL, BASE_URL, pages)

	for normalizedRawCurrentURL, count := range pages {
		fmt.Printf("%s: %d\n", normalizedRawCurrentURL, count)
	}
}
