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

func main() {
	fh, err := os.Open("./input.txt")
	fail(err)

	var count int
	rx := regexp.MustCompile("(\\d+)-(\\d+) (.): (.+)")

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		rs := rx.FindAllStringSubmatch(scanner.Text(), -1)

		// min, err := strconv.Atoi(rs[0][1])
		// fail(err)
		// max, err := strconv.Atoi(rs[0][2])
		// fail(err)
		// ch := rs[0][3]
		// pass := rs[0][4]

		// x := strings.Count(pass, ch)

		// if x >= min && x <= max {
		// 	count++
		// }

		// fmt.Printf("min=%v max=%v ch=%#v pass=%#v x=%v count=%v\n", min, max, ch, pass, x, count)

		fst, err := strconv.Atoi(rs[0][1])
		fail(err)
		snd, err := strconv.Atoi(rs[0][2])
		fail(err)
		ch := rs[0][3]
		pass := rs[0][4]

		var match int
		if pass[fst-1] == ch[0] {
			match++
		}
		if pass[snd-1] == ch[0] {
			match++
		}

		if match == 1 {
			count++
		}

		fmt.Printf("fst=%v snd=%v ch=%#v pass=%#v match=%v count=%v\n", fst, snd, ch, pass, match, count)

		// fmt.Printf("%#v %#v %#v %#v\n", min, max, ch, pass)
	}
	fmt.Printf("%#v\n", count)

	err = scanner.Err()
	fail(err)
}
