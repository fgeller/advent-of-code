package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

func parseInput(fn string) []int {
	buf, err := ioutil.ReadFile(fn)
	fail(err)

	lines := strings.Split(string(buf), "\n")
	result := []int{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		i, err := strconv.ParseInt(l, 10, 64)
		fail(err)
		result = append(result, int(i))
	}
	return result
}

func findWeakness(in []int, win int) int {
outer:
	for i := win; i < len(in); i++ {
		x := in[i]
		for j := i - win; j < i; j++ {
			for k := j + 1; k < i; k++ {
				a := in[j]
				b := in[k]
				if a+b == x && a != b {
					// fmt.Printf("found sum x=%v a=%v b=%v\n", x, a, b)
					continue outer
				}
			}
		}
		// fmt.Printf("found weakness=%v i=%v\n", x, i)
		return x
	}
	log.Fatal("did not find weakness")
	return 0
}

func findWeakness2(in []int, targetSum int) int {
	fmt.Printf("looking for targetSum=%v\n", targetSum)
outer:
	for i := 0; i < len(in)-1; i++ {
		sum, min, max := in[i], in[i], in[i]
		for j := i + 1; j < len(in); j++ {
			sum += in[j]
			if min > in[j] {
				min = in[j]
			}
			if max < in[j] {
				max = in[j]
			}
			if sum == targetSum {
				// fmt.Printf("found contiguous set between in[%v]=%v and in[%v]=%v\n", i, in[i], j, in[j])
				// fmt.Printf("min=%v max=%v sum=%v\n", min, max, min+max)
				return min + max
			} else if sum > targetSum {
				// fmt.Printf("sum too large, backtracking\n")
				continue outer
			}
		}
	}

	log.Fatal("did not find weakness 2")
	return 0
}

func main() {
	input := parseInput("input.txt")
	weak := findWeakness(input, 25)
	weak2 := findWeakness2(input, weak)

	fmt.Printf("weak=%#v weak2=%#v\n", weak, weak2)
}
