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

	reg := 0
	visited := map[int]struct{}{}
	pc := 0
	for {
		if _, ok := visited[pc]; ok {
			fmt.Println("already visited instruction", pc)
			break
		}
		fmt.Println("visiting", pc, "reg =", reg)
		visited[pc] = struct{}{}
		in := input[pc]
		switch in.o {
		case nop:
			pc++
		case acc:
			reg += in.v
			pc++
		case jmp:
			pc += in.v
		}
	}
}
