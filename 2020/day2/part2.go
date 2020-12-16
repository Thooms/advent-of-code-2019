package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type entry struct {
	lowerBound int
	upperBoud  int
	c          rune
	s          string
}

func (e *entry) String() string {
	return fmt.Sprintf("{%d-%d (%s) '%s'}", e.lowerBound, e.upperBoud, string(e.c), e.s)
}

func main() {
	data, err := ioutil.ReadFile("day2.input")
	if err != nil {
		panic(err)
	}

	input := []entry{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		e := entry{}
		var c string
		if _, err := fmt.Sscanf(r, "%d-%d %s %s", &e.lowerBound, &e.upperBoud, &c, &e.s); err != nil {
			panic(err)
		}
		e.c = []rune(c)[0]
		input = append(input, e)
	}

	for i, e := range input {
		occurences := 0
		if rune(e.s[e.lowerBound-1]) == e.c {
			occurences++
		}
		if rune(e.s[e.upperBoud-1]) == e.c {
			occurences++
		}
		if occurences == 1 {
			fmt.Println(i, "->", e.String())
		}
	}
}
