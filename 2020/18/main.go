package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var samples = []string{
	"1 + 2 * 3 + 4 * 5 + 6",
	"2 * 3 + (4 * 5)",
	"5 + (8 * 3 + 9 + 3 * 4 * 3)",
	"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
	"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
}

func main() {
	start := time.Now()
	var sum int64
	for _, e := range input {
		res := eval(e, true)
		// fmt.Printf("[%s] =%v\n", e, res)
		sum += res
	}
	fmt.Printf("sum=%v\n", sum)
	fmt.Println(time.Since(start))
}

func eval(expr string, advanced bool) int64 {
	ae := expr
	var ee string
	for {
		ee = evalInnerMostParens(ae, advanced)
		if ee == ae {
			break
		}
		ae = ee
	}
	expr = ee

	if advanced {
		be := expr
		var pe string
		for {
			pe = evalSums(be, advanced)
			if pe == be {
				break
			}
			be = pe
		}
		expr = pe
	}

	rxDigit := regexp.MustCompile("^[0-9]$")
	rxOp := regexp.MustCompile("^[+*]$")

	var lhs *int64
	var rhs *int64
	var op string

	seen := -1
	for i := range expr {
		// fmt.Printf("i=%v seen= %v\n", i, seen)
		if i <= seen {
			// fmt.Printf("skipping i=%v seen=%v\n", i, seen)
			continue
		}
		str := string(expr[i])
		if str == " " {
			continue
		}

		// dr := int64(-1)
		// if lhs != nil {
		// 	dr = (*lhs)
		// }
		// fmt.Printf("str=%#v lhs=%v (%v) rhs=%v op=%#v\n", str, lhs, dr, rhs, op)

		if rxOp.MatchString(str) {
			op = str
		}

		if rxDigit.MatchString(str) {
			fstr := str
		search:
			for j := i + 1; j < len(expr); j++ {
				ns := string(expr[j])
				if rxDigit.MatchString(ns) {
					fstr += ns
					seen = j
				} else {
					break search
				}
			}

			dv, err := strconv.ParseInt(fstr, 10, 64)
			fail(err)

			if lhs == nil {
				lhs = &dv
			} else if rhs == nil {
				rhs = &dv
			} else {
				log.Fatalf("unexpected, neither lhs=%v nor rhs=%v are nil\n", *lhs, *rhs)
			}

			if lhs != nil && rhs != nil {
				if op == "+" {
					nv := *lhs + *rhs
					lhs = &nv
				} else if op == "*" {
					nv := *lhs * *rhs
					lhs = &nv
				} else {
					log.Fatalf("unexpected op=%#v lhs=%v rhs=%v\n", op, *lhs, *rhs)
				}
				op = ""
				rhs = nil
			}
		}
	}

	return (*lhs)
}

func evalInnerMostParens(expr string, advanced bool) string {
	// (1 + 2)
	// 1 + (1 + 2)
	// 1 + ((1 + 2) + 2)
	fstClosing := -1
	lstOpening := -1
	for i, r := range expr {
		str := string(r)
		if str == "(" {
			lstOpening = i
		} else if str == ")" {
			fstClosing = i
			break
		}
	}

	if lstOpening == -1 {
		return expr
	}

	sub := expr[lstOpening+1 : fstClosing]
	res := eval(sub, advanced)

	end := fstClosing + 1
	if end >= len(expr) {
		end = len(expr) - 1
	}

	nexpr := expr[:lstOpening]
	nexpr += fmt.Sprintf("%v", res)
	if fstClosing+1 < len(expr) {
		nexpr += expr[fstClosing+1:]
	}
	// fmt.Printf(">> expr=[%s]\n>> sub=[%s] = %v\n>> nexpr=[%s]\n", expr, sub, res, nexpr)

	return nexpr
}

// assumes no parens
func evalSums(expr string, advanced bool) string {
	rxSum := regexp.MustCompile("([0-9]+) \\+ ([0-9]+)")
	rss := rxSum.FindStringSubmatch(expr)
	// fmt.Printf(">> expr=%v\n", expr)
	// fmt.Printf(">> rss=%#v\n", rss)
	if len(rss) < 3 {
		// fmt.Printf(">> no match\n")
		return expr
	}

	lhs, err := strconv.ParseInt(rss[1], 10, 64)
	fail(err)
	rhs, err := strconv.ParseInt(rss[2], 10, 64)
	fail(err)
	res := lhs + rhs
	// fmt.Printf(">> res=%v\n", res)

	start := strings.Index(expr, rss[0])
	end := start + len(rss[0])

	// fmt.Printf(">> sub=[%s]\n", expr[start:end])

	nexpr := expr[:start]
	nexpr += fmt.Sprintf("%v", res)
	if end < len(expr) {
		nexpr += expr[end:]
	}

	// fmt.Printf(">> nexpr=[%s]\n", nexpr)
	return nexpr
}

func fail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
