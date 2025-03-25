package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func getShortestPath(bytesSize int, lines []string) interface{} {
	grid := make([][]rune, 71)
	for i := range grid {
		grid[i] = make([]rune, 71)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for i := 0; i < bytesSize; i++ {
		parts := strings.Split(lines[i], ",")
		col, _ := strconv.Atoi(parts[0])
		row, _ := strconv.Atoi(parts[1])
		grid[row][col] = '#'
	}

	queue := make(map[Point]float64)
	for i := 0; i < 71; i++ {
		for j := 0; j < 71; j++ {
			queue[Point{i, j}] = math.Inf(1)
		}
	}

	start := Point{0, 0}
	queue[start] = 0
	end := Point{70, 70}

	directions := []Point{
		{0, 1}, {1, 0}, {-1, 0}, {0, -1},
	}

	for len(queue) > 0 {
		var currPos Point
		minDist := math.Inf(1)
		for pos, dist := range queue {
			if dist < minDist {
				minDist = dist
				currPos = pos
			}
		}

		if minDist == math.Inf(1) {
			return "BLOCKED"
		}

		if currPos == end {
			return int(minDist)
		}

		for _, d := range directions {
			newPos := Point{currPos.X + d.X, currPos.Y + d.Y}
			if val, exists := queue[newPos]; exists {
				if grid[newPos.X][newPos.Y] != '#' {
					queue[newPos] = math.Min(val, minDist+1)
				}
			}
		}

		delete(queue, currPos)
	}

	return "BLOCKED"
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", getShortestPath(1024, lines))

	startVal := 0
	endVal := 3450
	lowestBlockedByte := endVal + 1

	for startVal <= endVal {
		b := (endVal + startVal) / 2
		res := getShortestPath(b, lines)
		if res == "BLOCKED" {
			if b < lowestBlockedByte {
				lowestBlockedByte = b
			}
			endVal = b - 1
		} else {
			startVal = b + 1
		}
	}

	fmt.Println("Part 2:", lines[lowestBlockedByte-1])
}
