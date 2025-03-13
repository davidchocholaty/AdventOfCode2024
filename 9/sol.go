package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input file
	inputRaw, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := string(inputRaw)

	// ==================================================
	// PART 1
	// ==================================================

	// Format input
	var IDCount []string
	var space []string

	for i, char := range content {
		if i%2 == 0 {
			IDCount = append(IDCount, string(char))
		} else {
			space = append(space, string(char))
		}
	}

	var input []interface{}
	for id, n := range IDCount {
		nInt, _ := strconv.Atoi(n)
		for i := 0; i < nInt; i++ {
			input = append(input, id)
		}

		if id < len(space) {
			spaceInt, _ := strconv.Atoi(space[id])
			for i := 0; i < spaceInt; i++ {
				input = append(input, ".")
			}
		}
	}

	// 2 pointers
	pointer := len(input) - 1
	var reorderedInput []int

	for i, char := range input {
		if char != "." {
			reorderedInput = append(reorderedInput, char.(int))
		} else {
			reorderedInput = append(reorderedInput, input[pointer].(int))
			input[pointer] = "."

			// Decrement pointer until get to next number
			for input[pointer] == "." {
				pointer--
			}
		}

		if i >= pointer {
			break
		}
	}

	tot := 0
	for i, val := range reorderedInput {
		tot += val * i
	}

	fmt.Println("part 1: ", tot)

	// ==================================================
	// PART 2
	// ==================================================

	// Format input (different to part 1)
	ID := 0
	var inputPart2 [][2]interface{}

	for i, n := range content {
		nInt, _ := strconv.Atoi(string(n))
		if i%2 == 0 {
			inputPart2 = append(inputPart2, [2]interface{}{strconv.Itoa(ID), nInt})
			ID++
		} else {
			inputPart2 = append(inputPart2, [2]interface{}{".", nInt})
		}
	}

	// Swap blocks
	// Loop through numbers from the back
	for i := len(inputPart2) - 1; i >= 0; i-- {
		id := inputPart2[i][0]
		size := inputPart2[i][1].(int)

		if id != "." {
			// Go through available spaces from the front
			idxMax := -1
			for j, item := range inputPart2 {
				if item[0] == id && item[1] == size {
					idxMax = j
					break
				}
			}

			for j := 0; j < idxMax; j++ {
				char := inputPart2[j][0]
				available := inputPart2[j][1].(int)

				if char == "." {
					// Can move item to this space?
					if size <= available {
						// Decrement amount of free space at pos
						inputPart2[j][1] = available - size

						// Find index of the current item
						idx := -1
						for k, item := range inputPart2 {
							if item[0] == id && item[1] == size {
								idx = k
								break
							}
						}

						// Replace it with free space
						inputPart2[idx] = [2]interface{}{".", size}

						// Insert number in right place
						newInputPart2 := make([][2]interface{}, 0)
						newInputPart2 = append(newInputPart2, inputPart2[:j]...)
						newInputPart2 = append(newInputPart2, [2]interface{}{id, size})
						newInputPart2 = append(newInputPart2, inputPart2[j:]...)
						inputPart2 = newInputPart2

						break
					}
				}
			}
		}
	}

	tot2 := 0
	pos := 0

	for _, item := range inputPart2 {
		val := item[0]
		count := item[1].(int)

		for i := 0; i < count; i++ {
			if val != "." {
				valInt, _ := strconv.Atoi(val.(string))
				tot2 += valInt * pos
			}
			pos++
		}
	}

	fmt.Println("part 2: ", tot2)
}
