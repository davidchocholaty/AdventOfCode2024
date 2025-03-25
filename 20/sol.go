package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	// Read input
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{0, 1},  // right
		{1, 0},  // down
		{-1, 0}, // up
		{0, -1}, // left
	}

	var spaces [][2]int
	var startPos, endPos [2]int
	var track [][2]int

	for i, row := range grid {
		for j, cell := range row {
			if cell == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else if cell == 'S' {
				startPos = [2]int{i, j}
			} else if cell == 'E' {
				endPos = [2]int{i, j}
			}
		}
	}

	track = append(track, startPos)
	currPos := startPos

	for len(track) < len(spaces)+2 {
		for _, direction := range directions {
			newPos := [2]int{currPos[0] + direction[0], currPos[1] + direction[1]}
			if grid[newPos[0]][newPos[1]] == '.' && !contains(track, newPos) {
				currPos = newPos
				track = append(track, currPos)
				break
			}
		}
	}
	track = append(track, endPos)

	travelTime := make(map[[2]int]int)
	for i, pos := range track {
		travelTime[pos] = i
	}

	// Part 1
	count := 0
	walls := findWalls(grid)

	for _, wall := range walls {
		var adjacentPositions []int
		for _, direction := range directions {
			newPos := [2]int{wall[0] + direction[0], wall[1] + direction[1]}
			if idx, found := travelTime[newPos]; found {
				adjacentPositions = append(adjacentPositions, idx)
			}
		}

		if len(adjacentPositions) == 2 {
			if math.Abs(float64(adjacentPositions[0]-adjacentPositions[1]))-2 >= 100 {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	// Part 2
	count2 := 0
	for i, pos1 := range track {
		for j := i + 100; j < len(track); j++ {
			pos2 := track[j]
			dist := math.Abs(float64(pos1[0]-pos2[0])) + math.Abs(float64(pos1[1]-pos2[1]))
			if dist <= 20 {
				if math.Abs(float64(travelTime[pos1]-travelTime[pos2]))-dist >= 100 {
					count2++
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", count2)
}

func findWalls(grid [][]rune) [][2]int {
	var walls [][2]int
	for i, row := range grid {
		for j, cell := range row {
			if cell == '#' {
				walls = append(walls, [2]int{i, j})
			}
		}
	}
	return walls
}

func contains(track [][2]int, pos [2]int) bool {
	for _, p := range track {
		if p == pos {
			return true
		}
	}
	return false
}
