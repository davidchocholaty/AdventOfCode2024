package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// PARSE INPUT
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	inputs := make(map[string]int)
	for i := 0; i < 90; i++ {
		parts := strings.Split(lines[i], ": ")
		val, _ := strconv.Atoi(parts[1])
		inputs[parts[0]] = val
	}

	operations := make(map[[3]string]string)
	for i := 91; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		operations[[3]string{parts[0], parts[2], parts[4]}] = parts[1]
	}

	// PART 1

	for len(operations) > 0 {
		for opKey := range operations {
			var1, var2, res := opKey[0], opKey[1], opKey[2]
			if _, ok1 := inputs[var1]; ok1 {
				if _, ok2 := inputs[var2]; ok2 {
					input1 := inputs[var1]
					input2 := inputs[var2]
					operand := operations[opKey]

					switch operand {
					case "AND":
						if input1 != 0 && input2 != 0 {
							inputs[res] = 1
						} else {
							inputs[res] = 0
						}
					case "OR":
						if input1 != 0 || input2 != 0 {
							inputs[res] = 1
						} else {
							inputs[res] = 0
						}
					case "XOR":
						inputs[res] = input1 ^ input2
					}

					delete(operations, opKey)
				}
			}
		}
	}

	var numBuilder strings.Builder
	for i := 45; i >= 0; i-- {
		var key string
		if i < 10 {
			key = fmt.Sprintf("z0%d", i)
		} else {
			key = fmt.Sprintf("z%d", i)
		}
		numBuilder.WriteString(strconv.Itoa(inputs[key]))
	}

	numStr := numBuilder.String()
	num, _ := strconv.ParseInt(numStr, 2, 64)
	fmt.Println("part 1: ", num)

	// PART 2: ripple adder

	type Operation struct {
		input1, input2, res, operand string
	}

	var operationsList []Operation
	inputOperandMap := make(map[string][]string)

	for i := 91; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		input1, operand, input2, res := parts[0], parts[1], parts[2], parts[4]
		operationsList = append(operationsList, Operation{input1, input2, res, operand})
		inputOperandMap[input1] = append(inputOperandMap[input1], operand)
		inputOperandMap[input2] = append(inputOperandMap[input2], operand)
	}

	wrongGates := make(map[string]bool)

	for _, op := range operationsList {
		input1, input2, res, operand := op.input1, op.input2, op.res, op.operand

		if res != "z00" && res != "z01" && res != "z45" {
			if res[0] == 'z' && operand != "XOR" {
				wrongGates[res] = true
			}

			if operand == "XOR" {
				if (input1[0] != 'x' && input1[0] != 'y') &&
					(input2[0] != 'x' && input2[0] != 'y') &&
					res[0] != 'z' {
					wrongGates[res] = true
				}
			}

			if operand == "AND" {
				if input1 != "x00" && input2 != "x00" && input1 != "y00" && input2 != "y00" {
					hasOR := false
					for _, op := range inputOperandMap[res] {
						if op == "OR" {
							hasOR = true
							break
						}
					}
					if !hasOR {
						wrongGates[res] = true
					}
				}
			}

			if operand == "XOR" {
				for _, op := range inputOperandMap[res] {
					if op == "OR" {
						wrongGates[res] = true
						break
					}
				}
			}
		}
	}

	var wrongGatesList []string
	for gate := range wrongGates {
		wrongGatesList = append(wrongGatesList, gate)
	}

	for i := 0; i < len(wrongGatesList)-1; i++ {
		for j := i + 1; j < len(wrongGatesList); j++ {
			if wrongGatesList[i] > wrongGatesList[j] {
				wrongGatesList[i], wrongGatesList[j] = wrongGatesList[j], wrongGatesList[i]
			}
		}
	}

	fmt.Println("part 2: ", strings.Join(wrongGatesList, ","))
}
