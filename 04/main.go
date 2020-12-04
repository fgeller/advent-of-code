package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fail(err error) {
	if err != nil {
		log.Fatalf("failed err=%v", err)
	}
}

func main() {
	fh, err := os.Open("./input.txt")
	fail(err)

	scanner := bufio.NewScanner(fh)
	// missing cid:
	req := []string{"ecl:", "pid:", "eyr:", "hcl:", "byr:", "iyr:", "hgt:"}

	var pp string
	var count int

	check := func() {
		for _, f := range req {
			if !strings.Contains(pp, f) {
				fmt.Printf("missing %#v: %#v\n", f, pp)
				pp = ""
				return
			}
		}

		frx := regexp.MustCompile("(...):([^ ]+) ")
		fss := frx.FindAllStringSubmatch(pp, -1)
		// fmt.Printf("%#v\n\n", fss)

	fields:
		for _, fs := range fss {
			fn := fs[1]
			fv := fs[2]
			switch fn {
			case "byr":
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("invalid byr=%#v err=%v\n", fv, err)
					pp = ""
					return
				}
				if i < 1920 || i > 2002 {
					fmt.Printf("invalid byr=%#v i=[%v]\n", fv, i)
					pp = ""
					return
				}
				continue fields

			case "iyr":
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("invalid iyr=%#v err=%v\n", fv, err)
					pp = ""
					return
				}
				if i < 2010 || i > 2020 {
					fmt.Printf("invalid iyr=%#v\n", fv)
					pp = ""
					return
				}
				continue fields

			case "eyr":
				i, err := strconv.Atoi(fv)
				if err != nil {
					fmt.Printf("invalid eyr=%#v err=%v\n", fv, err)
					pp = ""
					return
				}
				if i < 2020 || i > 2030 {
					fmt.Printf("invalid eyr=%#v\n", fv)
					pp = ""
					return
				}
				continue fields

			case "hgt":
				if !strings.HasSuffix(fv, "cm") && !strings.HasSuffix(fv, "in") {
					fmt.Printf("invalid hgt suffix: %#v\n", fv)
					pp = ""
					return
				}
				i, err := strconv.Atoi(fv[:len(fv)-2])
				if err != nil {
					fmt.Printf("invalid hgt=%#v err=%v\n", fv, err)
					pp = ""
					return
				}
				if strings.HasSuffix(fv, "cm") {
					if i < 150 || i > 193 {
						fmt.Printf("invalid cm hgt=%#v\n", fv)
						pp = ""
						return
					}
				}
				if strings.HasSuffix(fv, "in") {
					if i < 59 || i > 76 {
						fmt.Printf("invalid in hgt=%#v\n", fv)
						pp = ""
						return
					}
				}
				continue fields

			case "hcl":
				rx := regexp.MustCompile("^#[0-9a-f]{6}$")
				if !rx.MatchString(fv) {
					fmt.Printf("invalid hcl: %#v\n", fv)
					pp = ""
					return
				}
				continue fields

			case "ecl":
				switch fv {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					continue fields
				default:
					fmt.Printf("invalid ecl: %#v\n", fv)
					pp = ""
					return
				}

			case "pid":
				rx := regexp.MustCompile("^[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]$")
				if !rx.MatchString(fv) {
					fmt.Printf("invalid pid: %#v\n", fv)
					pp = ""
					return
				}
				continue fields

			case "cid":
				continue fields

			default:
				log.Fatalf("unexpected field: %#v:%#v\n", fn, fv)
			}
		}

		fmt.Printf("valid: %#v\n", pp)
		count++
		pp = ""
	}

scan:
	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
			pp += l + " "
			continue scan
		} else {
			check()
		}
	}
	check()
	err = scanner.Err()
	fail(err)

	fmt.Printf("count=%v\n", count)
}
