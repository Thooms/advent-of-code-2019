package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func rowID(p string) int {
	l, u := 0, 128
	for _, c := range p {
		mid := int(l + (u-l)/2)
		switch c {
		case 'F':
			l, u = l, mid
		case 'B':
			l, u = mid, u
		default:
			panic(c)
		}
	}
	return l
}

func seatID(p string) int {
	l, u := 0, 8
	for _, c := range p {
		mid := int(l + (u-l)/2)
		switch c {
		case 'L':
			l, u = l, mid
		case 'R':
			l, u = mid, u
		default:
			panic(c)
		}
	}
	return l
}

func main() {
	data, err := ioutil.ReadFile("day5.input")
	if err != nil {
		panic(err)
	}

	input := []string{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		input = append(input, r)
	}

	max := 0
	for _, s := range input {
		r := rowID(s[:7])
		se := seatID(s[7:])
		id := r*8 + se
		if id > max {
			max = id
		}
	}
	fmt.Println(max)
}
