package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

const (
	width  = 101
	height = 103
	maxT   = width * height
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	positions := make(map[int]map[Position]int)
	for t := 0; t < maxT; t++ {
		positions[t] = make(map[Position]int)
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		pParts := strings.Split(parts[0][2:], ",")
		vParts := strings.Split(parts[1][2:], ",")

		x, _ := strconv.Atoi(pParts[0])
		y, _ := strconv.Atoi(pParts[1])
		dx, _ := strconv.Atoi(vParts[0])
		dy, _ := strconv.Atoi(vParts[1])

		for t := 0; t < maxT; t++ {
			x = (x + dx) % width
			y = (y + dy) % height
			positions[t][Position{x, y}]++
		}
	}

	// PART 1
	quadCounts := quadrantCounts(positions[99])
	tot := 1
	for _, count := range quadCounts {
		tot *= count
	}
	fmt.Println("part 1:", tot)

	// PART 2
	middleCol := (width - 1) / 2
	maxUniquePosCount := 0
	bestT := maxT

	for t := 0; t < maxT; t++ {
		uniquePosCount := 0
		for i := -2; i < 2; i++ {
			occupiedColPositions := []Position{}
			for pos := range positions[t] {
				if pos.y == middleCol+i {
					occupiedColPositions = append(occupiedColPositions, pos)
				}
			}
			uniquePosCount += len(occupiedColPositions)
		}
		if uniquePosCount > maxUniquePosCount {
			maxUniquePosCount = uniquePosCount
			bestT = t
		}
	}

	fmt.Println("part 2:", bestT+1)
}

func quadrantCounts(positionDict map[Position]int) []int {
	widthSplit := (width - 1) / 2
	heightSplit := (height - 1) / 2

	q1, q2, q3, q4 := 0, 0, 0, 0

	for pos, count := range positionDict {
		switch {
		case pos.x < widthSplit && pos.y < heightSplit:
			q1 += count
		case pos.x > widthSplit && pos.y < heightSplit:
			q2 += count
		case pos.x < widthSplit && pos.y > heightSplit:
			q3 += count
		case pos.x > widthSplit && pos.y > heightSplit:
			q4 += count
		}
	}

	return []int{q1, q2, q3, q4}
}
