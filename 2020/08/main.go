package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

func interpret(ls []string) int {
	rxInstruction := regexp.MustCompile("(...) ([+-][0-9]+)")
	acc := 0

	processed := map[int]bool{}
	i := 0
	for {
		instruction := ls[i]
		if i == len(ls)-1 {
			fmt.Printf("reached last instruction i=%#v acc=%#v\n", i, acc)
			return acc
		}
		if processed[i] {
			fmt.Printf("loop at instruction=%#v i=%#v\n", instruction, i)
			return acc
		} else {
			processed[i] = true
		}

		if len(instruction) == 0 {
			log.Fatalf("empty instruction i=%v\n", i)
		}

		rss := rxInstruction.FindAllStringSubmatch(instruction, -1)
		op := rss[0][1]
		arg, err := strconv.ParseInt(rss[0][2], 10, 64)
		fail(err)

		// fmt.Printf("i=%v acc=%v op=%#v arg=%#v\n", i, acc, op, arg)

		switch op {
		case "nop":
			i += 1
		case "acc":
			acc += int(arg)
			i += 1
		case "jmp":
			i += int(arg)
		default:
			log.Fatalf("unexpected op=%#v", op)
		}
		// fmt.Printf("i=%v acc=%v\n", i, acc)
	}
}

func findError(ls []string) int {
	rxInstruction := regexp.MustCompile("(...) ([+-][0-9]+)")
	processed := map[int]bool{}
	switched := map[int]bool{}
	attempts := 1
	acc := 0
	i := 0
	flag := true
	for {
		instruction := ls[i]
		if i == len(ls)-1 {
			fmt.Printf("reached last instruction i=%#v acc=%#v\n", i, acc)
			return acc
		}
		if processed[i] {
			// fmt.Printf("attempt=%v loops at instruction=%#v i=%#v\n", attempts, instruction, i)
			i = 0
			acc = 0
			flag = true
			attempts += 1
			processed = map[int]bool{}
			continue
		} else {
			processed[i] = true
		}

		rss := rxInstruction.FindAllStringSubmatch(instruction, -1)
		op := rss[0][1]
		arg, err := strconv.ParseInt(rss[0][2], 10, 64)
		fail(err)

		// fmt.Printf("i=%v acc=%v op=%#v arg=%#v\n", i, acc, op, arg)
		if flag && !switched[i] {
			switch op {
			case "nop":
				op = "jmp"
				flag = false
				switched[i] = true
				// fmt.Printf("switching nop to jmp i=%v original=%#v\n", i, instruction)
			case "jmp":
				op = "nop"
				flag = false
				switched[i] = true
				// fmt.Printf("switching jmp to nop i=%v original=%#v\n", i, instruction)
			}
		}

		switch op {
		case "nop":
			i += 1
		case "acc":
			acc += int(arg)
			i += 1
		case "jmp":
			i += int(arg)
		default:
			log.Fatalf("unexpected op=%#v", op)
		}
		// fmt.Printf("i=%v acc=%v\n", i, acc)
	}
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	fail(err)

	lines := strings.Split(string(buf), "\n")
	acc := interpret(lines)
	accFixed := findError(lines)

	fmt.Printf("acc=%#v\n", acc)
	fmt.Printf("accFixed=%#v\n", accFixed)
}
