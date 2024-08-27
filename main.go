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

	if len(arg) == 1 {
		BASE_URL := arg[0]
		fmt.Printf("starting crawl of: %v", BASE_URL)
	}
}
