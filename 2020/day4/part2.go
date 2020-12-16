package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day4.input")
	if err != nil {
		panic(err)
	}

	input := []map[string]string{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n\n") {
		m := map[string]string{}
		for _, f := range strings.Fields(r) {
			splt := strings.Split(f, ":")
			if len(splt) != 2 {
				panic(r)
			}
			m[splt[0]] = splt[1]
		}
		input = append(input, m)
	}

	for _, p := range input {

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		if byr, ok := p["byr"]; ok {
			parsedByr, err := strconv.Atoi(byr)
			if err != nil {
				panic(err)
			}
			if parsedByr < 1920 || parsedByr > 2020 {
				continue
			}
		} else {
			continue
		}

		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		if iyr, ok := p["iyr"]; ok {
			parsedIyr, err := strconv.Atoi(iyr)
			if err != nil {
				panic(err)
			}
			if parsedIyr < 2010 || parsedIyr > 2020 {
				continue
			}
		} else {
			continue
		}

		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		if eyr, ok := p["eyr"]; ok {
			parsedEyr, err := strconv.Atoi(eyr)
			if err != nil {
				panic(err)
			}
			if parsedEyr < 2020 || parsedEyr > 2030 {
				continue
			}
		} else {
			continue
		}

		// hgt (Height) - a number followed by either cm or in:
		//    If cm, the number must be at least 150 and at most 193.
		//    If in, the number must be at least 59 and at most 76.

		if hgt, ok := p["hgt"]; ok {
			l := len(hgt)
			if hgt[l-2:] == "cm" {
				parsedHgt, err := strconv.Atoi(hgt[:l-2])
				if err != nil {
					panic(err)
				}
				if parsedHgt < 150 || parsedHgt > 193 {
					continue
				}
			} else if hgt[l-2:] == "in" {
				parsedHgt, err := strconv.Atoi(hgt[:l-2])
				if err != nil {
					panic(err)
				}
				if parsedHgt < 59 || parsedHgt > 76 {
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}

		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		allowedChars := "0123456789abcdefghijklmnopqrstuvwxyz"
		if hcl, ok := p["hcl"]; ok {
			if hcl[0] != '#' {
				continue
			}
			if len(hcl) != 7 {
				panic(hcl)
			}
			for _, c := range hcl[1:] {
				if !strings.Contains(allowedChars, string(c)) {
					panic(hcl)
				}
			}
		} else {
			continue
		}

		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		allowedValues := "amb blu brn gry grn hzl oth"
		if ecl, ok := p["ecl"]; ok {
			if !strings.Contains(allowedValues, ecl) {
				continue
			}
		} else {
			continue
		}

		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		if pid, ok := p["pid"]; ok {
			if len(pid) != 9 {
				continue
			}
			if _, err := strconv.Atoi(pid); err != nil {
				panic(err)
			}
		} else {
			continue
		}

		fmt.Println(p)
	}
}
