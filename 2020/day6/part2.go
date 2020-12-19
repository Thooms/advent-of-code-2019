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
		input = append(input, r)
	}

	acc := 0
	for _, r := range input {
		rs := strings.Split(r, "\n")
		counts := map[rune]int{}
		for _, participant := range rs {
			for _, c := range participant {
				counts[c]++
			}
		}
		for _, i := range counts {
			if i == len(rs) {
				acc += 1
			}
		}
	}
	fmt.Println(acc)
}
