package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	p := newParser()

	for in.Scan() {
		p.lines++
		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		p = update(p, parsed)
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.stats[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
