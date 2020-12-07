package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

type rule struct {
	container  string
	containees map[string]int64
}

func parseRule(raw string) *rule {
	rxContainer := regexp.MustCompile(`(.+) bags contain `)
	rxContainees := regexp.MustCompile(`(\d+) (.+?) bags?[,.]`)

	rss := rxContainer.FindAllStringSubmatch(raw, -1)
	if len(rss) == 0 {
		log.Fatalf("no container match raw=%#v", raw)
	}
	container := rss[0][1]

	rss = rxContainees.FindAllStringSubmatch(raw, -1)
	containees := map[string]int64{}
	if len(rss) > 0 {
		for _, rs := range rss {
			var err error
			containees[rs[2]], err = strconv.ParseInt(rs[1], 10, 0)
			if err != nil {
				log.Fatalf("failed to parse %#v as int raw=%#v", rs[1], raw)
			}
		}
	}
	return &rule{container, containees}
}

func findContainers(rules map[string]map[string]int64, needle string) map[string]bool {
	seek := map[string]bool{needle: true}
	next := map[string]bool{}
	container := map[string]bool{}

	for {
		for s := range seek {
			for b, bs := range rules {
				for a := range bs {
					if a == s {
						// fmt.Printf("[%#v] can hold [%#v]\n", b, s)
						container[b] = true
						next[b] = true
					}
				}
			}
		}
		if len(next) == 0 {
			break
		} else {
			seek = next
			next = map[string]bool{}
		}
	}

	return container
}

func findNested(rules map[string]map[string]int64, outer string) (count int64) {
	seek := map[string]int64{outer: 1}
	next := map[string]int64{}

	for {
		// fmt.Printf(">> seek=%#v next=%#v count=%#v\n", seek, next, count)
		for n, c := range seek {
			for nn, nc := range rules[n] {
				inc := c * nc
				count += inc
				if _, ok := next[nn]; ok {
					next[nn] += inc
				} else {
					next[nn] = inc
				}
			}
		}

		if len(next) == 0 {
			break
		} else {
			seek = next
			next = map[string]int64{}
		}
	}

	return
}

func main() {
	fh, err := os.Open("./input.txt")
	fail(err)

	scanner := bufio.NewScanner(fh)
	rules := map[string]map[string]int64{}
	for scanner.Scan() {
		l := scanner.Text()
		r := parseRule(l)
		rules[r.container] = r.containees
	}
	err = scanner.Err()
	fail(err)

	container := findContainers(rules, "shiny gold")
	fmt.Printf("found %v container bags\n", len(container))

	nested := findNested(rules, "shiny gold")
	fmt.Printf("found %v nested bags\n", nested)
}
