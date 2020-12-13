package main

type data1 struct {
	earliest int
	busIDs   []int
}

var sample1 = &data1{
	earliest: 939,
	busIDs: []int{
		7,
		13,
		59,
		31,
		19,
	},
}

var input1 = &data1{
	earliest: 1008713,
	busIDs: []int{
		13,
		17,
		19,
		23,
		29,
		37,
		41,
		353,
		467,
	},
}

// offset -> busID
var sample2 = map[int]int{
	0: 7,
	1: 13,
	4: 59,
	6: 31,
	7: 19,
}
