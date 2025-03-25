package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	towels := strings.Split(lines[0], ", ")
	patterns := lines[2:]
	return towels, patterns
}

func canForm(pattern string, towels []string) bool {
	isValid := make(map[string]bool)
	isValid[""] = true

	for i := 1; i <= len(pattern); i++ {
		subpattern := pattern[:i]
		isValid[subpattern] = false
		for _, towel := range towels {
			if strings.HasSuffix(subpattern, towel) {
				prevSubpattern := subpattern[:len(subpattern)-len(towel)]
				if isValid[prevSubpattern] {
					isValid[subpattern] = true
					break
				}
			}
		}
	}
	return isValid[pattern]
}

func countArrangements(pattern string, towels []string) int {
	counts := make(map[string]int)
	counts[""] = 1

	for i := 1; i <= len(pattern); i++ {
		subpattern := pattern[:i]
		counts[subpattern] = 0
		for _, towel := range towels {
			if strings.HasSuffix(subpattern, towel) {
				prevSubpattern := subpattern[:len(subpattern)-len(towel)]
				counts[subpattern] += counts[prevSubpattern]
			}
		}
	}
	return counts[pattern]
}

func main() {
	towels, patterns := readInput("input.txt")

	count1 := 0
	for _, pattern := range patterns {
		if canForm(pattern, towels) {
			count1++
		}
	}
	fmt.Println("Part 1:", count1)

	count2 := 0
	for _, pattern := range patterns {
		count2 += countArrangements(pattern, towels)
	}
	fmt.Println("Part 2:", count2)
}
