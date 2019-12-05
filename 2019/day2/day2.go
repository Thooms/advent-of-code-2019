package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func run(program []int, pc int) ([]int, error) {
	// read opcode
	opcode := program[pc]
	switch opcode {
	case 99: // halting opcode
		return program, nil
	case 1, 2:
		operand_1_pos, operand_2_pos, result_pos := program[pc+1], program[pc+2], program[pc+3]
		operand_1, operand_2 := program[operand_1_pos], program[operand_2_pos]
		result := 0
		if opcode == 1 {
			result = operand_1 + operand_2
		} else { // opcode == 2
			result = operand_1 * operand_2
		}
		program[result_pos] = result
	default:
		return nil, fmt.Errorf("unexpected opcode %d at position %d", opcode, pc)
	}
	return run(program, pc+4) // jump 4 opcodes for next step
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	text, _ := ioutil.ReadAll(stdin)
	splitted := strings.Split(strings.TrimSpace(string(text)), ",")
	initialProgram := []int{}
	for _, c := range splitted {
		opcode, _ := strconv.Atoi(c)
		initialProgram = append(initialProgram, opcode)
	}

	if os.Getenv("DAY2") == "" {
		initialProgram[1] = 12
		initialProgram[2] = 2
		result, _ := run(initialProgram, 0)
		fmt.Println(result[0])
	}

	// brute force the thingy
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			// fresh memory
			mem := make([]int, len(initialProgram))
			copy(mem, initialProgram)
			mem[1] = noun
			mem[2] = verb
			result, err := run(mem, 0)
			if err != nil {
				continue
			}
			if result[0] == 19690720 {
				fmt.Println("Noun:", noun, "Verb:", verb, "Magic:", 100*noun+verb)
			}
		}
	}

}
