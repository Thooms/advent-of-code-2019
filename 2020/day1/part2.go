package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("part1.input")
	if err != nil {
		panic(err)
	}

	input := []int{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		i, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}

	for i, i_ := range input {
		for j, j_ := range input {
			if j == i {
				continue
			}
			for k, k_ := range input {
				if k == i || k == j {
					continue
				}
				if i_+j_+k_ == 2020 {
					fmt.Println(i_, j_, k_, "->", i_*j_*k_)
					os.Exit(0)
				}
			}
		}
	}
}
