package main

import (
	"fmt"
	"io/ioutil"
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

	h := map[int]struct{}{}
	for _, i := range input {
		h[i] = struct{}{}
	}

	candidatesNum, cum := 0, 1
	for k := range h {
		if _, ok := h[2020-k]; ok {
			candidatesNum++
			cum *= (2020 - k)
		}
	}
	fmt.Println("Candidates:", candidatesNum, "; product:", cum)

}
