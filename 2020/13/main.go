package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	part1(input1)
	start := time.Now()
	part2(sample2)
	fmt.Printf("time: %s\n", time.Since(start))

	// found t=672754131923874
	// time: 2h26m55.281149424s
}

func part1(in *data1) {
	mins := map[int]int{}
	for _, id := range in.busIDs {
		c := math.Ceil(float64(in.earliest) / float64(id))
		min := int(c) * id
		mins[id] = min
		fmt.Printf("id=%v c=%v min=%v\n", id, c, min)
	}
	fmt.Printf("mins=%#v\n", mins)

	mbi := 0
	mts := math.MaxInt64
	for bi, ts := range mins {
		if ts < mts {
			mts = ts
			mbi = bi
		}
	}
	fmt.Printf("busID=%v timestamp=%v busID*diff=%v\n", mbi, mts, mbi*(mts-in.earliest))
}

func part2(in map[int]int) {
	// d := "7,13,x,x,59,x,31,19"
	d := "13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23"
	ps := map[int64]int64{}
	for i, j := range strings.Split(d, ",") {
		if j == "x" {
			continue
		}
		n, _ := strconv.ParseInt(j, 10, 64)
		ps[int64(i)] = n
	}
	fmt.Printf("d=%v\n", d)
	fmt.Printf("ps=%#v\n", ps)
	inc := ps[ps[0]] * ps[0]
	fmt.Printf("inc=%v\n", inc)

	var pot = int64(1)
	var x int64
search:
	for {
		if x > pot {
			fmt.Printf("looking at x=%v l=%v\n", x, math.Log10(float64(x)))
			pot *= 10
		}
		x += inc
		t := x - ps[0]
		for i, m := range ps {
			if (t+i)%m != 0 {
				continue search
			}
		}

		fmt.Printf("found t=%v\n", t)
		return
	}
}
