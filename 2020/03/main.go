package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

func main() {
	fh, err := os.Open("./input.txt")
	fail(err)

	mp := [][]int{}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		raw := scanner.Text()
		line := make([]int, len(raw))
		for i, b := range raw {
			if b == rune('#') {
				line[i] = 1
			}
		}
		mp = append(mp, line)
	}
	err = scanner.Err()
	fail(err)

	var x, y, count int
	lineLen := len(mp[0])

	for {
		x = x + 1
		if x >= lineLen {
			x = x - lineLen
		}

		y = y + 2
		if y >= len(mp) {
			break
		}

		if mp[y][x] > 0 {
			count++
		}

		// l := fmt.Sprintf("%v", mp[y])
		// marker := " "
		// for j := 0; j < x; j++ {
		// 	marker += "  "
		// }
		// marker += "^"
		// for len(marker) < len(l) {
		// 	marker += " "
		// }
		// fmt.Printf("%v\n", l)
		// fmt.Printf("%v  x=%v y=%v mp[y][x]=%v count=%v\n", marker, x, y, mp[y][x], count)
	}

	fmt.Printf("count=%v\n", count)
	// fmt.Printf("prod=%v\n", 70*220*63*76*29)
}

// 1 1: 70
// 3 1: 220
// 5 1: 63
// 7 1: 76
// 1 2: 29
