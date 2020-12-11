package main

import (
	"fmt"
	"sync"
)

var inputM = []int{
	0,
	1,
	2,
	3,
	4,
	7,
	8,
	9,
	10,
	11,
	14,
	17,
	18,
	19,
	20,
	23,
	24,
	25,
	28,
	31,
	32,
	33,
	34,
	35,
	38,
	39,
	42,
	45,
	46,
	47,
	48,
	49,
	52,
}

var inputL = []int{
	0,
	1,
	2,
	3,
	4,
	7,
	8,
	9,
	10,
	11,
	14,
	15,
	16,
	17,
	20,
	21,
	22,
	23,
	24,
	27,
	28,
	29,
	30,
	31,
	34,
	35,
	36,
	37,
	38,
	41,
	44,
	47,
	48,
	49,
	52,
	53,
	54,
	55,
	56,
	59,
	60,
	61,
	62,
	63,
	66,
	67,
	68,
	69,
	72,
	73,
	74,
	75,
	76,
	79,
	80,
	81,
	82,
	85,
	88,
	89,
	90,
	91,
	92,
	95,
	98,
	99,
	100,
	101,
	102,
	105,
	108,
	111,
	112,
	115,
	118,
	119,
	120,
	121,
	122,
	125,
	128,
	129,
	130,
	133,
	134,
	135,
	136,
	137,
	140,
	141,
	142,
	143,
	144,
	147,
	150,
	153,
	154,
	155,
	158,
	161,
	162,
	163,
	166,
}

var inputS = []int{
	0,
	1,
	4,
	5,
	6,
	7,
	10,
	11,
	12,
	15,
	16,
	19,
	22,
}

func first(in []int) {
	ones := 0
	threes := 0
	for i := 1; i < len(in); i++ {
		diff := in[i] - in[i-1]
		switch diff {
		case 1:
			ones += 1
			continue
		case 3:
			threes += 1
			continue
		default:
			fmt.Printf("unexpected diff: %v\n", diff)
		}
	}
	fmt.Printf("%v ones * %v threes = %v\n", ones, threes, ones*threes)
}

// func second(in []int) int64 {
// 	l := len(in)

// 	if l < 2 {
// 		log.Fatal("unexpected len=%v\n", len(in))
// 	}

// 	if l == 2 {
// 		return 1
// 	}

// 	if l == 3 {
// 		x := in[0]
// 		add := int64(0)
// 		diff := in[2] - x
// 		if diff < 4 {
// 			add += 1
// 		}
// 		return add + second(in[1:])
// 	}

// 	if l == 4 {
// 		x := in[0]
// 		add := int64(0)
// 		diff := in[2] - x
// 		if diff < 4 {
// 			add += 1
// 		}
// 		diff = in[3] - x
// 		if diff < 4 {
// 			add += 1
// 		}
// 		return add + second(in[1:])
// 	}

// 	//	if l >= 5
// 	x := in[0]
// 	add := int64(0)
// 	diff := in[2] - x
// 	if diff < 4 {
// 		add += 1
// 	}
// 	diff = in[3] - x
// 	if diff < 4 {
// 		add += 1
// 	}
// 	diff = in[4] - x
// 	if diff < 4 {
// 		add += 1
// 	}
// 	return add + second(in[1:])
// }

type node struct {
	i        int
	children []*node
}

func newNode(i int) *node {
	return &node{i: i, children: []*node{}}
}

var secondCount = int64(0)

// func second(in []int) {
// 	root := newNode(0)
// 	populateChildren(root, in)
// 	fmt.Printf("%v\n", secondCount)
// }

// func populateChildren(n *node, in []int) {
// 	for j := n.i + 1; j <= n.i+3; j++ {
// 		if j >= len(in) {
// 			break
// 		}
// 		diff := in[j] - in[n.i]
// 		if diff <= 3 {
// 			n.children = append(n.children, newNode(j))
// 		}
// 	}

// 	for _, c := range n.children {
// 		populateChildren(c, in)
// 	}
// 	if len(n.children) == 0 {
// 		secondCount += 1
// 	}
// }

func populateChildren(m *sync.Mutex, wg *sync.WaitGroup, i int, in []int) {
	cs := []int{}
	for j := i + 1; j <= i+3; j++ {
		if j >= len(in) {
			break
		}
		if in[j]-in[i] <= 3 {
			cs = append(cs, j)
		}
	}

	for _, c := range cs {
		wg.Add(1)
		go populateChildren(m, wg, c, in)
	}
	if len(cs) == 0 {
		m.Lock()
		secondCount += 1
		m.Unlock()
	}
	wg.Done()
}

func second(in []int) {
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go populateChildren(&m, &wg, 0, in)
	fmt.Printf("waiting for wg\n")
	wg.Wait()
	fmt.Printf("second=%v\n", secondCount)
}

func loopPopulate(in []int) {
	todo := []int{0}
	pts := map[int]int{0: 1}
	count := 0
	l := len(in)

	for len(todo) > 0 {
		i := todo[0]
		cs := 0
		for j := i + 1; j <= i+3 && j < l; j++ {
			if in[j]-in[i] <= 3 {
				cs += 1
				if _, ok := pts[j]; ok {
					pts[j] += 1
				} else {
					pts[j] = 1
					todo = append(todo, j)
				}
			}
		}

		pts[i] -= 1
		if pts[i] == 0 {
			todo = todo[1:]
			delete(pts, i)
		}

		if cs == 0 {
			count += 1
		}
	}
	fmt.Printf("loop count=%v\n", count)
}

func branches(in []int) []int {
	result := []int{}
	for i := range in {
		bs := 0
		x := in[i]
		if len(in) > i+1 {
			a := in[i+1]
			if a-x <= 3 {
				bs += 1
			}
		}
		if len(in) > i+2 {
			a := in[i+2]
			if a-x <= 3 {
				bs += 1
			}
		}
		if len(in) > i+3 {
			a := in[i+3]
			if a-x <= 3 {
				bs += 1
			}
		}
		result = append(result, bs)
	}

	strBranches := ""
	for _, r := range result {
		strBranches += fmt.Sprintf("%1d", r)
	}
	fmt.Printf("branches:\n%#v\n", strBranches)
	return result
}

func countLeaves(branches []int) int {
	// 3321  => *7
	// 321   => *4
	// 21    => *2
	count := 1
	i := 0
	for i < len(branches)-1 {
		switch branches[i] {
		case 1:
			i += 1
			continue
		case 2:
			count *= 2
			i += 1
		case 3:
			next := branches[i+1]
			if next == 2 {
				count *= 4
				i += 2
			} else if next == 3 {
				count *= 7
				i += 3
			}
		}
	}
	fmt.Printf("count=%v\n", count)
	return count
}

func main() {
	in := inputL
	first(in)
	countLeaves(branches(in))
}
