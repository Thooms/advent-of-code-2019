package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day3.input")
	if err != nil {
		panic(err)
	}

	input := []string{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		input = append(input, r)
	}

	l := len(input[0])

	for i := 1; i < len(input); i++ {
		j := (3 * i) % l
		c := input[i][j]
		if c == '#' {
			fmt.Println("Found tree at", i, j)
		}
	}
}
