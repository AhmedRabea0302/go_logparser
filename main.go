package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}
type parser struct {
	stats   map[string]result // Total visits per domain
	domains []string          // Unique domains
	total   int               // Total visits for all domains
	lines   int               // Number of parse lines
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	p := parser{
		stats: make(map[string]result),
	}

	for in.Scan() {
		p.lines++

		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Printf("Wrong Input: %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("Wrong Input: %v (line #%d)\n", fields[1], p.lines)
			return
		}

		p.total += visits
		if _, ok := p.stats[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
		p.stats[domain] = result{
			domain: domain,
			visits: visits + p.stats[domain].visits,
		}
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
