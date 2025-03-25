package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evolveStonesGrouped(initialStones []int, blinks int) int {
	stoneCounts := make(map[int]int)
	for _, stone := range initialStones {
		stoneCounts[stone]++
	}

	for i := 0; i < blinks; i++ {
		newStoneCounts := make(map[int]int)
		for stone, count := range stoneCounts {
			if stone == 0 {
				newStoneCounts[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				mid := len(strconv.Itoa(stone)) / 2
				left, _ := strconv.Atoi(strconv.Itoa(stone)[:mid])
				right, _ := strconv.Atoi(strconv.Itoa(stone)[mid:])
				newStoneCounts[left] += count
				newStoneCounts[right] += count
			} else {
				newStoneCounts[stone*2024] += count
			}
		}
		stoneCounts = newStoneCounts
	}

	total := 0
	for _, count := range stoneCounts {
		total += count
	}

	return total
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputData := strings.TrimSpace(string(data))
	stones := []int{}
	for _, s := range strings.Fields(inputData) {
		stone, _ := strconv.Atoi(s)
		stones = append(stones, stone)
	}

	blinks := 75

	totalStones := evolveStonesGrouped(stones, blinks)

	fmt.Println(totalStones)
}
