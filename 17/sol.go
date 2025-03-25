package main

import (
	"fmt"
	"sort"
)

type Registers struct {
	A, B, C int
}

func getCombo(operand int, registers *Registers) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers.A
	case 5:
		return registers.B
	case 6:
		return registers.C
	default:
		return -1
	}
}

func executeInstruction(opcode, operand int, registers *Registers) (int, string) {
	combo := getCombo(operand, registers)
	switch opcode {
	case 0:
		registers.A /= (1 << combo)
	case 1:
		registers.B ^= operand
	case 2:
		registers.B = combo % 8
	case 3:
		if registers.A != 0 {
			return 0, "jump"
		}
	case 4:
		registers.B ^= registers.C
	case 5:
		return combo % 8, "output"
	case 7:
		registers.C = registers.A / (1 << combo)
	}
	return 0, ""
}

func runProgram(registers Registers, program []int) []int {
	instructionPointer := 0
	var outputs []int

	for instructionPointer < len(program) {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]
		output, status := executeInstruction(opcode, operand, &registers)

		if status == "output" {
			outputs = append(outputs, output)
		}

		if status == "jump" {
			instructionPointer = operand
		} else {
			instructionPointer += 2
		}
	}
	return outputs
}

func main() {
	program := []int{2, 4, 1, 5, 7, 5, 1, 6, 4, 3, 5, 5, 0, 3, 3, 0}
	outputs := runProgram(Registers{A: 47792830, B: 0, C: 0}, program)
	fmt.Println("PART 1:", outputs)

	// PART 2
	stack := [][2]int{{1 << (3 * 15), 15}} // 8^15

	for len(stack) > 0 {
		sort.Slice(stack, func(i, j int) bool { return stack[i][0] > stack[j][0] })
		currA, index := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		stepSize := 1 << (3 * index)

		for value := currA; value < currA+(1<<(3*(index+1)))+stepSize; value += stepSize {
			output := runProgram(Registers{A: value, B: 0, C: 0}, program)
			if len(output) >= index && equalSlices(output[index:], program[index:]) {
				if index == 0 {
					fmt.Println("PART 2:", value)
					return
				} else {
					stack = append(stack, [2]int{value, index - 1})
				}
			}
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
