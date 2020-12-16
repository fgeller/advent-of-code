package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("ticket scanning error rate=%v\n", ticketScanningErrorRate(inputIntvs, inputTickets))
	part2M(inputIntvs, inputTickets)
}

func part2(is map[string][]*intv, ts [][]int) {
	vts := dropInvalidTickets(is, ts)
	fn2is := map[string]map[int]bool{}
	fns := []string{}
	for fn := range is {
		fns = append(fns, fn)
		fn2is[fn] = map[int]bool{}
	}

	for fn, rs := range is {
	seekPossiblePositions:
		for i := 0; i < len(vts[0]); i++ {
			for _, t := range vts {
				v := t[i]
				if !rs[0].isValid(v) && !rs[1].isValid(v) {
					continue seekPossiblePositions
				}
			}
			fn2is[fn][i] = true
		}
	}

	mapping := map[string]int{}
	removePosition := func(in map[string]map[int]bool, p int) map[string]map[int]bool {
		result := map[string]map[int]bool{}
		for fn, is := range in {
			nis := map[int]bool{}
			for a := range is {
				if a != p {
					nis[a] = true
				}
			}
			result[fn] = nis
		}
		return result
	}

findMapping:
	for {
		if len(mapping) == len(fns) {
			fmt.Printf("found all mappings\n")
			break findMapping
		}

		for _, fn := range fns {
			if len(fn2is[fn]) == 1 {
				i := -1
				for x := range fn2is[fn] {
					i = x
				}
				fmt.Printf("found mapping fn=%v -> i=%v\n", fn, i)
				mapping[fn] = i
				fn2is = removePosition(fn2is, i)
				continue findMapping
			}
		}
	}

	product := 1
	for fn, i := range mapping {
		if strings.HasPrefix(fn, "departure") {
			fmt.Printf("fn=%v product=%v i=%v fv=%v\n", fn, product, i, myTicket[i])
			product *= myTicket[i]
		}
	}
	fmt.Printf("product=%v\n", product)
}

func part2M(is map[string][]*intv, ts [][]int) {
	vts := dropInvalidTickets(is, ts)
	fn2is := map[string]map[int]bool{}
	fns := []string{}
	for fn := range is {
		fns = append(fns, fn)
		fn2is[fn] = map[int]bool{}
	}

	for fn, rs := range is {
	seekPossiblePositions:
		for i := 0; i < len(vts[0]); i++ {
			for _, t := range vts {
				v := t[i]
				if !rs[0].isValid(v) && !rs[1].isValid(v) {
					continue seekPossiblePositions
				}
			}
			fn2is[fn][i] = true
		}
	}

	mapping := map[string]int{}

findMapping:
	for {
		if len(mapping) == len(fns) {
			fmt.Printf("found all mappings\n")
			break findMapping
		}

		for _, fn := range fns {
			if len(fn2is[fn]) == 1 {
				i := -1
				for x := range fn2is[fn] {
					i = x
				}
				fmt.Printf("found mapping fn=%v -> i=%v\n", fn, i)
				mapping[fn] = i
				for _, is := range fn2is {
					delete(is, i)
				}
				continue findMapping
			}
		}
	}

	product := 1
	for fn, i := range mapping {
		if strings.HasPrefix(fn, "departure") {
			fmt.Printf("fn=%v product=%v i=%v fv=%v\n", fn, product, i, myTicket[i])
			product *= myTicket[i]
		}
	}
	fmt.Printf("product=%v\n", product)
}

func dropInvalidTickets(is map[string][]*intv, ts [][]int) [][]int {
	valid := [][]int{}
outer:
	for _, t := range ts {
		for _, v := range t {
			if !isValid(is, v) {
				continue outer
			}
		}
		valid = append(valid, t)
	}
	fmt.Printf("found %v valid tickets out of %v\n", len(valid), len(ts))
	return valid
}

func isValid(is map[string][]*intv, v int) bool {
	for _, fi := range is {
		if v >= fi[0].low && v <= fi[0].high {
			return true
		}
		if v >= fi[1].low && v <= fi[1].high {
			return true
		}
	}
	return false
}

func ticketScanningErrorRate(is map[string][]*intv, ts [][]int) int {
	sum := 0
	for _, t := range ts {
		for _, v := range t {
			if !isValid(is, v) {
				// fmt.Printf("invalid ticket=%v v=%v\n", t, v)
				sum += v
			}
		}
	}
	return sum
}
