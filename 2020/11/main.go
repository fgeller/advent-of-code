package main

import (
	"fmt"
	"reflect"
)

func visible(sp []string, ri, ci int) int {
	above := func() int {
		for i := ri - 1; i >= 0; i-- {
			if string(sp[i][ci]) == "#" {
				return 1
			}
			if string(sp[i][ci]) == "L" {
				return 0
			}
		}
		return 0
	}
	below := func() int {
		for i := ri + 1; i < len(sp); i++ {
			if string(sp[i][ci]) == "#" {
				return 1
			}
			if string(sp[i][ci]) == "L" {
				return 0
			}
		}
		return 0
	}
	right := func() int {
		for i := ci + 1; i < len(sp[0]); i++ {
			if string(sp[ri][i]) == "#" {
				return 1
			}
			if string(sp[ri][i]) == "L" {
				return 0
			}
		}
		return 0
	}
	left := func() int {
		for i := ci - 1; i >= 0; i-- {
			if string(sp[ri][i]) == "#" {
				return 1
			}
			if string(sp[ri][i]) == "L" {
				return 0
			}
		}
		return 0
	}
	topLeft := func() int {
		c := ci - 1
		for r := ri - 1; r >= 0; r-- {
			if c < 0 {
				break
			}
			if string(sp[r][c]) == "#" {
				return 1
			}
			if string(sp[r][c]) == "L" {
				return 0
			}
			c -= 1
		}
		return 0
	}
	topRight := func() int {
		c := ci + 1
		for r := ri - 1; r >= 0; r-- {
			if c >= len(sp[0]) {
				break
			}
			if string(sp[r][c]) == "#" {
				return 1
			}
			if string(sp[r][c]) == "L" {
				return 0
			}
			c += 1
		}
		return 0
	}
	bottomRight := func() int {
		c := ci + 1
		for r := ri + 1; r < len(sp); r++ {
			if c >= len(sp[0]) {
				break
			}
			if string(sp[r][c]) == "#" {
				return 1
			}
			if string(sp[r][c]) == "L" {
				return 0
			}
			c += 1
		}
		return 0
	}
	bottomLeft := func() int {
		c := ci - 1
		for r := ri + 1; r < len(sp); r++ {
			if c < 0 {
				break
			}
			if string(sp[r][c]) == "#" {
				return 1
			}
			if string(sp[r][c]) == "L" {
				return 0
			}
			c -= 1
		}
		return 0
	}

	a, b, r, l := above(), below(), right(), left()
	tl, tr, br, bl := topLeft(), topRight(), bottomRight(), bottomLeft()
	all := a + b + r + l + tl + tr + br + bl
	return all
}

func adjacent(sp []string, ri, ci int) int {
	count := 0
	rowCount := len(sp)
	colCount := len(sp[0])

	cr := 0
	if ri > 0 {
		cr = ri - 1
	}
	for ; cr <= ri+1 && cr < rowCount; cr++ {
		cc := 0
		if ci > 0 {
			cc = ci - 1
		}
		for ; cc <= ci+1 && cc < colCount; cc++ {
			if cc == ci && cr == ri {
				continue
			}
			if string(sp[cr][cc]) == "#" {
				count += 1
			}
		}
	}
	return count
}

func round(sp []string, limit int, counter func([]string, int, int) int) []string {
	next := []string{}
	for ri, row := range sp {
		next = append(next, "")
		for ci, c := range row {
			ao := counter(sp, ri, ci)
			ns := string(c)
			switch ns {
			case "L":
				if ao == 0 {
					ns = "#"
				}
			case "#":
				if ao >= limit {
					ns = "L"
				}
			}
			next[ri] = next[ri] + ns
		}
	}

	return next
}

func countOccupied(sp []string) int {
	count := 0
	for _, l := range sp {
		for _, s := range l {
			if string(s) == "#" {
				count += 1
			}
		}
	}
	return count
}

func printPlan(sp []string) {
	for _, l := range sp {
		fmt.Println(l)
	}
	fmt.Println()
}

func solve(sp []string, limit int, counter func([]string, int, int) int) {
	n := sp
	r := 0
	for {
		nn := round(n, limit, counter)
		r += 1
		// fmt.Printf("round=%v\n", r)
		// printPlan(nn)
		if reflect.DeepEqual(nn, n) {
			fmt.Printf("found equilibrium after %v rounds. occupied=%v\n", r, countOccupied(n))
			return
		}
		n = nn
	}
}

func main() {
	solve(input, 4, adjacent)
	solve(input, 5, visible)
}
