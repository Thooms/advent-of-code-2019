package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func count(low, high int) int {
	count := 0
	for n := int64(low); n <= int64(high); n++ {
		foundEqualDigit := false
		digitsOnlyIncrease := true
		currentDigit := int64(n/1e5) % 10
		// iterate through the rest of the digits, from high to low
		for m := int64(1e4); m > 0; m /= 10 {
			d := int64(n/m) % 10
			if d == currentDigit {
				foundEqualDigit = true
			}
			if d < currentDigit {
				digitsOnlyIncrease = false
				break
			}
			currentDigit = d
		}
		if digitsOnlyIncrease && foundEqualDigit {
			count++
			fmt.Println("Found number", n)
		}
	}
	return count
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	text, _ := ioutil.ReadAll(stdin)
	bounds := strings.Split(strings.TrimSpace(string(text)), "-")
	low, _ := strconv.Atoi(bounds[0])
	high, _ := strconv.Atoi(bounds[1])
	fmt.Println(count(low, high))
}
