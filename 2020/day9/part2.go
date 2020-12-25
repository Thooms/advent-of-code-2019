package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// 1930745883

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

	sums := make([][]int, len(input)) // i < j; sums[i][j] = sum(input[i:j+1])

	for i := 0; i < len(input)-1; i++ {
		sums[i] = make([]int, len(input)) // init to 0
		for j := i + 1; j < len(input); j++ {
			if j == i+1 {
				sums[i][j] = input[i] + input[j]
			} else {
				sums[i][j] = sums[i][j-1] + input[j]
			}

			if sums[i][j] == 1930745883 {
				fmt.Println("Found sum!", i, j, input[i:j+1])
				// find answer
				min, max := 1000000000000000000, 0 // LOLMATHS
				for _, x := range input[i : j+1] {
					if x < min {
						min = x
					}
					if x > max {
						max = x
					}
				}
				fmt.Println(min, max, min+max)
				goto _found
			}
		}
	}

_found:
}
