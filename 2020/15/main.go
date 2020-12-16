package main

import (
	"fmt"
	"time"
)

var sample = []int{0, 3, 6}
var input = []int{0, 14, 6, 20, 1, 4}

func main() {
	start := time.Now()
	play(input, 2020)
	fmt.Println(time.Since(start))
	start = time.Now()
	play(input, 30000000)
	fmt.Println(time.Since(start))

	// last=&main.num{val:257, times:1, turnA:2020, turnB:0} spoken=384
	// 177.48µs
	// last=&main.num{val:8546398, times:1, turnA:30000000, turnB:0} spoken=3611842
	// 4.641186742s
}

type num struct {
	val   int
	times int
	turnA int
	turnB int
}

func play(in []int, limit int) {
	spoken := map[int]*num{}
	var last *num
	for i := 0; i < limit; i++ {
		if i < len(in) {
			last = &num{val: in[i], times: 1, turnA: i + 1}
			spoken[in[i]] = last
		} else {
			if last.times == 1 {
				last = spoken[0]
				last.times += 1
				if last.turnB > 0 {
					last.turnA = last.turnB
				}
				last.turnB = i + 1
			} else {
				nv := last.turnB - last.turnA
				var ok bool
				last, ok = spoken[nv]
				if !ok {
					last = &num{val: nv, times: 1, turnA: i + 1}
					spoken[nv] = last
				} else {
					last.times += 1
					if last.turnB > 0 {
						last.turnA = last.turnB
					}
					last.turnB = i + 1
				}
			}
		}
	}
	fmt.Printf("last=%#v spoken=%v\n", last, len(spoken))
}
