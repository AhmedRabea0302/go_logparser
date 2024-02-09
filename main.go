package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		stats   map[string]int
		domains []string
		total   int
		lines   int
	)

	in := bufio.NewScanner(os.Stdin)
	stats = make(map[string]int)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Printf("Wrong Input: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("Wrong Input: %v (line #%d)\n", fields[1], lines)
			return
		}

		total += visits
		if _, ok := stats[domain]; !ok {
			domains = append(domains, domain)
		}
		stats[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := stats[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
