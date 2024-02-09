package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	query := args[0]
	rx := regexp.MustCompile(`[^a-z]+`)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}

	fmt.Printf("The input %q does not exist\n", query)
}
