package main

import (
	"fmt"
	"io/ioutil"
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

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	for _, p := range input {
		// check how many fields are present
		fc := 0
		for _, fn := range fields {
			if _, ok := p[fn]; ok {
				fc++
			}
		}
		// check whether the cid is there
		_, hasCid := p["cid"]

		if fc == 8 || (fc == 7 && !hasCid) {
			fmt.Println(p)
		}
	}
}
