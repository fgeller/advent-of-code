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

	req := []string{"ecl:", "pid:", "eyr:", "hcl:", "byr:", "iyr:", "hgt:"}
	var pp string
	var count int

	check := func() string {
		for _, f := range req {
			if !strings.Contains(pp, f) {
				return fmt.Sprintf("missing %#v: %#v\n", f, pp)
			}
		}
		frx := regexp.MustCompile("(...):([^ ]+) ")
		fss := frx.FindAllStringSubmatch(pp, -1)

	fields:
		for _, fs := range fss {
			fn := fs[1]
			fv := fs[2]
			switch fn {
			case "byr":
				i, err := strconv.Atoi(fv)
				if err != nil {
					return fmt.Sprintf("invalid byr=%#v err=%v\n", fv, err)
				}
				if i < 1920 || i > 2002 {

					return fmt.Sprintf("invalid byr=%#v i=[%v]\n", fv, i)
				}
				continue fields

			case "iyr":
				i, err := strconv.Atoi(fv)
				if err != nil {
					return fmt.Sprintf("invalid iyr=%#v err=%v\n", fv, err)
				}
				if i < 2010 || i > 2020 {
					return fmt.Sprintf("invalid iyr=%#v\n", fv)
				}
				continue fields

			case "eyr":
				i, err := strconv.Atoi(fv)
				if err != nil {
					return fmt.Sprintf("invalid eyr=%#v err=%v\n", fv, err)
				}
				if i < 2020 || i > 2030 {
					return fmt.Sprintf("invalid eyr=%#v\n", fv)
				}
				continue fields

			case "hgt":
				if !strings.HasSuffix(fv, "cm") && !strings.HasSuffix(fv, "in") {
					return fmt.Sprintf("invalid hgt, missing suffix: %#v\n", fv)
				}
				i, err := strconv.Atoi(fv[:len(fv)-2])
				if err != nil {
					return fmt.Sprintf("invalid hgt=%#v err=%v\n", fv, err)
				}
				if strings.HasSuffix(fv, "cm") {
					if i < 150 || i > 193 {
						return fmt.Sprintf("invalid cm hgt=%#v\n", fv)
					}
				}
				if strings.HasSuffix(fv, "in") {
					if i < 59 || i > 76 {
						return fmt.Sprintf("invalid in hgt=%#v\n", fv)
					}
				}
				continue fields

			case "hcl":
				rx := regexp.MustCompile("^#[0-9a-f]{6}$")
				if !rx.MatchString(fv) {
					return fmt.Sprintf("invalid hcl: %#v\n", fv)
				}
				continue fields

			case "ecl":
				switch fv {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					continue fields
				default:
					return fmt.Sprintf("invalid ecl: %#v\n", fv)
				}

			case "pid":
				rx := regexp.MustCompile("^[0-9]{9}$")
				if !rx.MatchString(fv) {
					return fmt.Sprintf("invalid pid: %#v\n", fv)
				}
				continue fields

			case "cid":
				continue fields

			default:
				log.Fatalf("unexpected field: %#v:%#v\n", fn, fv)
			}
		}

		return ""
	}

	checkPass := func() {
		if e := check(); e == "" {
			count++
		} else {
			fmt.Print(e)
		}
		pp = ""
	}

	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
			pp += l + " "
		} else {
			checkPass()
		}
	}
	checkPass()

	err = scanner.Err()
	fail(err)

	fmt.Printf("count=%v\n", count)
}
