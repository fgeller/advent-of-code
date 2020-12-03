package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	nums := []int{}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		fail(err)
		nums = append(nums, n)
	}
	err = scanner.Err()
	fail(err)

	// fmt.Printf("found nums: %#v\n", nums)

	max := 8000000
	var comparisons int
outer:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				comparisons++
				a := nums[i]
				b := nums[j]
				c := nums[k]
				if a+b+c == 2020 {
					fmt.Printf("found a=%v b=%v c=%v a*b*c=%v\n", a, b, c, a*b*c)
					break outer
				}
			}
		}
	}

	fmt.Printf("comparisons=%v δ=%v\n", comparisons, max-comparisons)
	fmt.Printf("len(nums)=%v\n", len(nums))

	// x := reduce(nums, func(memoi, i int) int {
	// 	if memoi != 0 {
	// 		return memoi
	// 	}

	// 	nf := func(memoj, j int) int {
	// 		if memoj != 0 {
	// 			return memoj
	// 		}

	// 		nnf := func(memok, k int) int {
	// 			if memok != 0 {
	// 				return memok
	// 			}

	// 			if i+j+k == 2020 {
	// 				return i * j * k
	// 			}
	// 			return 0
	// 		}

	// 		return reduce(nums, nnf)
	// 	}

	// 	return reduce(nums, nf)
	// })
	// fmt.Printf("reduc: %#v\n", x)
}

// def reduce[A, B](as: List[A], f: (A, B) => B): B

func reduce(is []int, f func(int, int) int) (memo int) {
	for _, i := range is {
		memo = f(memo, i)
	}
	return memo
}
