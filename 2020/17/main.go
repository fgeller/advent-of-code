package main

import (
	"fmt"
	"time"
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
	x, y, z, w int
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
				p.active[coord{x, y, 0, 0}] = ACTIVE
			}
		}
	}
	return p
}

func (c coord) neighbors(active map[coord]struct{}, limit int) int {
	var count int
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				for w := c.w - 1; w <= c.w+1; w++ {
					if c.x == x && c.y == y && c.z == z && c.w == w {
						continue
					}
					if _, ok := active[coord{x, y, z, w}]; ok {
						count += 1
						if count >= limit {
							return count
						}
					}

				}
			}
		}
	}
	return count
}

func (c coord) possibilities() []coord {
	ps := make([]coord, 0, 81)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				for w := c.w - 1; w <= c.w+1; w++ {
					ps = append(ps, coord{x, y, z, w})
				}
			}
		}
	}
	return ps
}

func (p *pocket) next() *pocket {
	n := &pocket{active: map[coord]struct{}{}}
	// TODO seen coords
	for c := range p.active {
		for _, pc := range c.possibilities() {
			nc := pc.neighbors(p.active, 4)
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

func simulate(in []string, cycles int) {
	p := newPocket(in)
	fmt.Println("Start")
	for i := 0; i < cycles; i++ {
		p = p.next()
		fmt.Printf("After %v cycle %v cubes active.\n", i+1, len(p.active))
	}
}

func main() {
	start := time.Now()
	simulate(input, 6)
	fmt.Println(time.Since(start))
}
