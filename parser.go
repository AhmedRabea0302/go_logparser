package main

import (
	"fmt"
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

func newParser() parser {
	return parser{stats: make(map[string]result)}
}

func parse(p parser, line string) (parsed result, err error) {

	fields := strings.Fields(line)
	if len(fields) < 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields[1], p.lines)
		return
	}

	return
}

func update(p parser, parsed result) parser {
	domain, visits := parsed.domain, parsed.visits

	p.total += visits
	if _, ok := p.stats[domain]; !ok {
		p.domains = append(p.domains, domain)
	}
	p.stats[domain] = result{
		domain: domain,
		visits: visits + p.stats[domain].visits,
	}

	return p
}
