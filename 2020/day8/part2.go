package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type op int

const (
	nop = op(0)
	acc = op(1)
	jmp = op(2)
)

func parseOp(s string) op {
	switch s {
	case "nop":
		return nop
	case "jmp":
		return jmp
	case "acc":
		return acc
	default:
		panic(s)
	}
}

type instr struct {
	o op
	v int
}

func main() {
	data, err := ioutil.ReadFile("day8.input")
	if err != nil {
		panic(err)
	}

	input := []instr{}
	for _, r := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		s := ""
		v := 0
		if _, err := fmt.Sscanf(strings.TrimSpace(r), "%s %d", &s, &v); err != nil {
			panic(err)
		}
		input = append(input, instr{o: parseOp(s), v: v})
	}

	// bruteforce the crap out of it

	found := false
	for i := 0; !found && i < len(input); i++ {
		reg := 0                      // value of the register
		visited := map[int]struct{}{} // to check if we are in a loop
		pc := 0                       // current program counter
		for {
			// infinite loop
			if _, ok := visited[pc]; ok {
				break
			}
			// end of program, we found it !
			if pc == len(input) {
				fmt.Println("reached end of program %%")
				found = true
				break
			}

			visited[pc] = struct{}{}
			in := input[pc]
			switch in.o {
			case nop:
				if pc == i {
					fmt.Println("trying to replace nop at instr", i)
					pc += in.v // replace by a jmp
					continue
				}
				pc++
			case acc:
				reg += in.v
				pc++
			case jmp:
				if pc == i {
					fmt.Println("trying to replace jmp at instr", i)
					pc++
					continue
				}
				pc += in.v
			}
		}
		if found {
			fmt.Println("found it! replaced instruction", i, " - reg =", reg)
			break
		}
	}
}
