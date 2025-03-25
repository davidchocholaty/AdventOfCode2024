package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var initNumbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		initNumbers = append(initNumbers, num)
	}

	// PART 1
	tot := 0
	for _, num := range initNumbers {
		for i := 0; i < 2000; i++ {
			num = ((num * 64) ^ num) % 16777216
			num = ((num / 32) ^ num) % 16777216
			num = ((num * 2048) ^ num) % 16777216
		}
		tot += num
	}

	fmt.Println("part 1: ", tot)

	// PART 2

	var secretNums [][]int
	for _, num := range initNumbers {
		vals := []int{num}
		for i := 0; i < 2000; i++ {
			num = ((num * 64) ^ num) % 16777216
			num = ((num / 32) ^ num) % 16777216
			num = ((num * 2048) ^ num) % 16777216

			numStr := strconv.Itoa(num)
			val, _ := strconv.Atoi(string(numStr[len(numStr)-1]))
			vals = append(vals, val)
		}
		secretNums = append(secretNums, vals)
	}

	sequencePrice := make(map[string][]int)
	for _, nums := range secretNums {
		var sequences []string

		differences := make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			differences[i] = nums[i+1] - nums[i]
		}

		for i := 0; i <= len(differences)-4; i++ {
			sequence := differences[i : i+4]
			seqStr := fmt.Sprintf("%v", sequence)

			found := false
			for _, s := range sequences {
				if s == seqStr {
					found = true
					break
				}
			}

			if !found {
				sequencePrice[seqStr] = append(sequencePrice[seqStr], nums[i+4])
				sequences = append(sequences, seqStr)
			}
		}
	}

	maxVal := 0
	for _, prices := range sequencePrice {
		sum := 0
		for _, price := range prices {
			sum += price
		}
		if sum > maxVal {
			maxVal = sum
		}
	}

	fmt.Println("part 2: ", maxVal)
}
