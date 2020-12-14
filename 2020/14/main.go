package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	regs := map[int]int{}
	interp1(input, regs)
	var sum int
	for _, v := range regs {
		sum += v
	}
	fmt.Printf("sum=%v\n", sum)

	regs2 := map[int64]int64{}
	interp2(input, regs2)
	var sum2 int64
	for _, v := range regs2 {
		sum2 += v
	}
	fmt.Printf("sum2=%v\n", sum2)
}

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

func interp1(is []string, regs map[int]int) {
	rxMask := regexp.MustCompile(`^mask = ([X10]+)$`)
	rxWrite := regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)

	var mask string
	for _, i := range is {
		// fmt.Printf(">> instruction %v: %#v\n", c, i)
		if rxMask.MatchString(i) {
			mask = rxMask.FindStringSubmatch(i)[1]
			// fmt.Printf(">> updated mask=%v\n", mask)
		} else if rxWrite.MatchString(i) {
			addr, err := strconv.ParseInt(rxWrite.FindStringSubmatch(i)[1], 10, 64)
			fail(err)
			val, err := strconv.ParseInt(rxWrite.FindStringSubmatch(i)[2], 10, 64)
			fail(err)
			maskedVal := applyMask(mask, int(val))
			regs[int(addr)] = maskedVal
			// fmt.Printf(">> updated regs[%v]=%v\n", int(addr), maskedVal)
		} else {
			log.Fatalf("failed to match instruction %#v\n", i)
		}
	}
}

func interp2(is []string, regs map[int64]int64) {
	rxMask := regexp.MustCompile(`^mask = ([X10]+)$`)
	rxWrite := regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)

	var mask string
	for _, i := range is {
		if rxMask.MatchString(i) {
			mask = rxMask.FindStringSubmatch(i)[1]
		} else if rxWrite.MatchString(i) {
			addr, err := strconv.ParseInt(rxWrite.FindStringSubmatch(i)[1], 10, 64)
			fail(err)
			val, err := strconv.ParseInt(rxWrite.FindStringSubmatch(i)[2], 10, 64)
			fail(err)

			addrs := applyAddrMask(mask, int(addr))
			for _, a := range addrs {
				regs[a] = val
			}
		} else {
			log.Fatalf("failed to match instruction %#v\n", i)
		}
	}
}
func applyAddrMask(mask string, a int) []int64 {
	format := fmt.Sprintf("%%0%vb", len(mask))
	binVal := fmt.Sprintf(format, a)
	madr := ""
	for i, v := range mask {
		switch string(v) {
		case "X":
			madr += "X"
		case "0":
			madr += string(binVal[i])
		case "1":
			madr += "1"
		}
	}

	// 	fmt.Printf(`applyAddrMask %v
	// mask=%v
	// addr=%v
	// madr=%v
	// `, a, mask, binVal, madr)

	as := []string{}
	for _, v := range madr {
		switch string(v) {
		case "X":
			if len(as) == 0 {
				as = []string{"0", "1"}
			} else {
				nr := make([]string, 0, len(as)*2)
				for _, na := range as {
					nr = append(nr, na+"1")
					nr = append(nr, na+"0")
				}
				as = nr
			}
		default:
			if len(as) == 0 {
				as = []string{string(v)}
			} else {
				nr := make([]string, 0, len(as))
				for _, na := range as {
					nr = append(nr, na+string(v))
				}
				as = nr
			}
		}
	}

	result := make([]int64, 0, len(as))
	for _, x := range as {
		i, err := strconv.ParseInt(x, 2, 64)
		fail(err)
		result = append(result, i)
	}

	return result
}

func applyMask(mask string, val int) int {
	format := fmt.Sprintf("%%0%vb", len(mask))
	binVal := fmt.Sprintf(format, val)
	// fmt.Printf("before mask:\n%v\n%v\n", mask, binVal)

	newVal := ""
	for i, v := range mask {
		if string(v) == "X" {
			newVal += string(binVal[i])
		} else {
			newVal += string(v)
		}
	}

	iv, err := strconv.ParseInt(newVal, 2, 64)
	fail(err)
	// fmt.Printf("after mask:\n%v\n%v\n", newVal, iv)
	return int(iv)
}
