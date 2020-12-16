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
	acc := 1

	for _, ds := range [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		dx, dy := ds[0], ds[1]
		cnt := 0
		for i := 0; i*dy < len(input); i++ {
			j := (i * dx) % l
			if input[i*dy][j] == '#' {
				cnt++
			}
		}
		fmt.Println(dx, dy, cnt)
		acc *= cnt
	}
	fmt.Println("result:", acc)
}
