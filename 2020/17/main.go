package main

import (
	"fmt"
	"sort"
	"strings"
)

var input = []string{
	".......#",
	"....#...",
	"...###.#",
	"#...###.",
	"....##..",
	"##.#..#.",
	"###.#.#.",
	"....#...",
}

var sample = []string{
	".#.",
	"..#",
	"###",
}

type coord struct {
	x, y, z int
}

var ACTIVE = struct{}{}

type pocket struct {
	active map[coord]struct{}
}

func newPocket(init []string) *pocket {
	p := &pocket{active: map[coord]struct{}{}}
	for y := range init {
		for x, r := range init[y] {
			if "#" == string(r) {
				p.active[coord{x, y, 0}] = ACTIVE
			}
		}
	}
	return p
}

type minMax struct {
	xMin, xMax, yMin, yMax int
}

func (c coord) neighbors(active map[coord]struct{}, limit int) int {
	var count int
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				if c.x == x && c.y == y && c.z == z {
					continue
				}
				if _, ok := active[coord{x, y, z}]; ok {
					count += 1
					if count >= limit {
						return count
					}
				}
			}
		}
	}
	return count
}

func (c coord) possibilities() []coord {
	ps := make([]coord, 0, 27)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				ps = append(ps, coord{x, y, z})
			}
		}
	}
	// fmt.Printf("coord = %#v  %v possibilities:\n", c, len(ps))
	// for _, a := range ps {
	// 	fmt.Printf("    %#v\n", a)
	// }
	// fmt.Println()

	return ps
}

func (p *pocket) next() *pocket {
	n := &pocket{active: map[coord]struct{}{}}

	// TODO seen coords
	// iterate over possibilities
	for c := range p.active {
		for _, pc := range c.possibilities() {
			nc := pc.neighbors(p.active, 4)
			// fmt.Printf("pc=%v neighbors=%v\n", pc, pc.neighbors(p.active, 4))
			_, ok := p.active[pc]
			if ok && (nc == 2 || nc == 3) {
				n.active[pc] = ACTIVE
			}
			if !ok && nc == 3 {
				n.active[pc] = ACTIVE
			}
		}
	}
	return n
}

func (p *pocket) String() string {
	zs := []int{}
	z2Active := map[int][]coord{}
	mms := map[int]*minMax{}
	for c := range p.active {

		if _, ok := z2Active[c.z]; !ok {
			zs = append(zs, c.z)
			z2Active[c.z] = []coord{}
		}
		z2Active[c.z] = append(z2Active[c.z], c)
		mm, ok := mms[c.z]
		if !ok {
			mms[c.z] = &minMax{xMin: c.x, xMax: c.x, yMin: c.y, yMax: c.y}
		} else {
			if c.x < mm.xMin {
				mm.xMin = c.x
			}
			if c.x > mm.xMax {
				mm.xMax = c.x
			}
			if c.y < mm.yMin {
				mm.yMin = c.y
			}
			if c.y > mm.yMax {
				mm.yMax = c.y
			}
		}
	}
	_ = sort.IntSlice(zs)

	result := ""

	for _, z := range zs {
		result += fmt.Sprintf("z=%v\n", z)
		l := []string{}
		height := (mms[z].yMax - mms[z].yMin) + 1
		width := (mms[z].xMax - mms[z].xMin) + 1
		for i := 0; i < height; i++ {
			l = append(l, strings.Repeat(".", width))
		}
		// TODO borked
		fmt.Printf("%#v height=%v width=%v\n", mms[z], height, width)
		for _, c := range z2Active[z] {
			fmt.Println(">> A <<")
			fmt.Printf(">> c=%#v\n", c)
			bs := []byte(l[c.y])
			bs[c.x] = byte('#')
			l[c.y] = string(bs)
			fmt.Println(">> B <<")
		}
		fmt.Println(">> C <<")

		for _, r := range l {
			result += r + "\n"
		}
	}
	result += "\n"

	return result
}

func simulate(in []string, cycles int) {
	p := newPocket(in)
	fmt.Println("Start")
	// fmt.Printf("%s\n", p)
	for i := 0; i < cycles; i++ {
		p = p.next()
		fmt.Printf("After %v cycle %v cubes active.\n", i+1, len(p.active))
	}
}

func main() {
	simulate(input, 6)
}
