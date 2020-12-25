package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day9.input")
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

	ws := 25
	h := map[int]int{} // keep the latest 25 in the map

	for i, n := range input {
		fmt.Println("examining", n)
		h[n]++
		if i >= ws {
			cnt, ok := h[input[i-ws]]
			if ok {
				if cnt == 1 { // delete if 0
					delete(h, input[i-ws])
				} else {
					h[input[i-ws]]--
				}
			}

			if i < len(input)-1 {
				// lookahead, check if any two *different* numbers can sum to it
				found := false
				for j := i - ws; j < i; j++ {
					a := input[j]
					if _, ok := h[n-a]; ok && (n-a) != a {
						found = true
						fmt.Println("Found a matching pair that sums to", n, "->", a, n-a)
						break
					}
				}
				if !found {
					fmt.Println("Could not find anything that sums to", n)
					goto _found
				}

			}
		}
	}
_found:
}
