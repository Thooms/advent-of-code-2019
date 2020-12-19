package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day6.input")
	if err != nil {
		panic(err)
	}

	input := []string{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n\n") {
		input = append(input, strings.Join(strings.Split(r, "\n"), ""))
	}

	acc := 0
	for _, r := range input {
		s := map[rune]struct{}{}
		for _, c := range r {
			s[c] = struct{}{}
		}
		acc += len(s)
	}
	fmt.Println(acc)
}
