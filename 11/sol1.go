package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stripLeadingZeros(stone string) string {
	stone = strings.TrimLeft(stone, "0")
	if len(stone) == 0 {
		return "0"
	}
	return stone
}

func applyRules(stones []string) []string {
	newStones := []string{}
	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if len(stone)%2 == 0 {
			firstHalf := stripLeadingZeros(stone[:len(stone)/2])
			secondHalf := stripLeadingZeros(stone[len(stone)/2:])
			newStones = append(newStones, firstHalf, secondHalf)
		} else {
			num, _ := strconv.Atoi(stone)
			newStones = append(newStones, strconv.Itoa(num*2024))
		}
	}
	return newStones
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	testInput := strings.TrimSpace(string(data))
	stones := strings.Fields(testInput)

	for i := 0; i < 25; i++ {
		stones = applyRules(stones)
	}

	fmt.Println("Part 1:", len(stones))
}
