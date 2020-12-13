package main

import (
	"fmt"
	"log"
)

type instruction struct {
	op  string
	arg int
}

type ship struct {
	facing string
	x      int
	y      int
}

type waypoint struct {
	x int
	y int
}

func (s *ship) execute(is []*instruction) {
	for _, i := range is {
		switch i.op {
		case "L":
			switch i.arg {
			case 90:
				switch s.facing {
				case "N":
					s.facing = "W"
				case "E":
					s.facing = "N"
				case "S":
					s.facing = "E"
				case "W":
					s.facing = "S"
				}
			case 180:
				switch s.facing {
				case "N":
					s.facing = "S"
				case "E":
					s.facing = "W"
				case "S":
					s.facing = "N"
				case "W":
					s.facing = "E"
				}
			case 270:
				switch s.facing {
				case "N":
					s.facing = "E"
				case "E":
					s.facing = "S"
				case "S":
					s.facing = "W"
				case "W":
					s.facing = "N"
				}
			case 360:
			default:
				log.Fatalf("unexpected L arg=%#v\n", i.arg)
			}

		case "R":
			switch i.arg {
			case 90:
				switch s.facing {
				case "N":
					s.facing = "E"
				case "E":
					s.facing = "S"
				case "S":
					s.facing = "W"
				case "W":
					s.facing = "N"
				}
			case 180:
				switch s.facing {
				case "N":
					s.facing = "S"
				case "E":
					s.facing = "W"
				case "S":
					s.facing = "N"
				case "W":
					s.facing = "E"
				}
			case 270:
				switch s.facing {
				case "N":
					s.facing = "W"
				case "E":
					s.facing = "N"
				case "S":
					s.facing = "E"
				case "W":
					s.facing = "S"
				}
			case 360:
			default:
				log.Fatalf("unexpected R arg=%#v\n", i.arg)
			}

		case "N":
			s.y -= i.arg

		case "S":
			s.y += i.arg

		case "E":
			s.x += i.arg

		case "W":
			s.x -= i.arg

		case "F":
			switch s.facing {
			case "N":
				s.y -= i.arg
			case "E":
				s.x += i.arg
			case "S":
				s.y += i.arg
			case "W":
				s.x -= i.arg
			}

		default:
			log.Fatalf("unexpected instruction=%#v\n", i)
		}
	}
}

func (s *ship) executeWithWaypoint(is []*instruction, wp *waypoint) {

	for _, i := range is {
		switch i.op {
		case "N":
			wp.y -= i.arg
		case "S":
			wp.y += i.arg
		case "E":
			wp.x += i.arg
		case "W":
			wp.x -= i.arg

		case "F":
			s.x += wp.x * i.arg
			s.y += wp.y * i.arg

		case "L":
			switch i.arg {
			case 90:
				a := wp.x
				wp.x = wp.y
				wp.y = -1 * a
			case 180:
				wp.x *= -1
				wp.y *= -1
			case 270:
				a := wp.x
				wp.x = -1 * wp.y
				wp.y = a
			}

		case "R":
			switch i.arg {
			case 90:
				a := wp.x
				wp.x = -1 * wp.y
				wp.y = a
			case 180:
				wp.x *= -1
				wp.y *= -1
			case 270:
				a := wp.x
				wp.x = wp.y
				wp.y = -1 * a
			}
		}

		fmt.Printf("i=%#v s=%#v wp=%#v\n", i, s, wp)
	}
}

func main() {
	s := &ship{facing: "E"}
	s.execute(input)
	fmt.Printf("%#v s.x+s.y=%v\n", s, s.x+s.y)

	s = &ship{}
	wp := &waypoint{x: 10, y: -1}
	s.executeWithWaypoint(input, wp)
	fmt.Printf("%#v s.x+s.y=%v\n", s, s.x+s.y)
}
