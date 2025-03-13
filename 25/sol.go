package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func splitBlocks(lines []string) [][]string {
	var blocks [][]string
	for i := 0; i < len(lines); i += 8 {
		end := i + 7
		if end > len(lines) {
			end = len(lines)
		}
		blocks = append(blocks, lines[i:end])
	}
	return blocks
}

func countPins(block []string) [5]int {
	var pinHeights [5]int
	for _, row := range block {
		for i := 0; i < 5; i++ {
			if row[i] == '#' {
				pinHeights[i]++
			}
		}
	}
	return pinHeights
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	blocks := splitBlocks(lines)

	var keys, locks [][5]int

	for _, block := range blocks {
		pinHeights := countPins(block)
		if strings.TrimSpace(block[0]) == "....." {
			locks = append(locks, pinHeights)
		} else if strings.TrimSpace(block[len(block)-1]) == "....." {
			keys = append(keys, pinHeights)
		}
	}

	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			match := true
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 7 {
					match = false
					break
				}
			}
			if match {
				count++
			}
		}
	}

	fmt.Println(count)
}
