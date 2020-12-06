package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

func checkGroup(g string) int {
	rxAnswer := regexp.MustCompile("^[a-z]$")
	uniq := ""
	for i := range g {
		ch := g[i : i+1]
		if rxAnswer.MatchString(ch) && !strings.Contains(uniq, ch) {
			uniq += ch
		}
	}
	return len(uniq)
}

func checkCommonGroup(g string) int {
	as := strings.Split(g, "|")
	common := as[0]
	for _, a := range as {
		common = strings.Map(
			func(r rune) rune {
				if strings.Contains(a, string(r)) {
					return r
				}
				return -1
			},
			common,
		)
	}
	return len(common)
}

func main() {
	fh, err := os.Open("./input.txt")
	fail(err)

	scanner := bufio.NewScanner(fh)

	groups := []string{}
	grp := ""
	addGroup := func() {
		groups = append(groups, strings.TrimRight(grp, "|"))
		grp = ""
	}
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			addGroup()
		} else {
			grp += l + "|"
		}
	}
	addGroup()
	err = scanner.Err()
	fail(err)

	var count int
	for _, g := range groups {
		count += checkGroup(g)
	}
	fmt.Printf("count=%v\n", count)

	var common int
	for _, g := range groups {
		common += checkCommonGroup(g)
	}
	fmt.Printf("common=%v\n", common)
}
