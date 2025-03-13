package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

func findAntennas(grid [][]rune) map[rune][][2]int {
	antennas := make(map[rune][][2]int)
	for y, row := range grid {
		for x, char := range row {
			if ('0' <= char && char <= '9') || ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
				antennas[char] = append(antennas[char], [2]int{x, y})
			}
		}
	}
	return antennas
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func calculateAntinodes(grid [][]rune, antennas map[rune][][2]int) map[[2]int]struct{} {
	rows := len(grid)
	cols := len(grid[0])
	antinodePositions := make(map[[2]int]struct{})

	for _, positions := range antennas {
		n := len(positions)
		if n < 2 {
			continue
		}

		for _, pos := range positions {
			antinodePositions[pos] = struct{}{}
		}

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				x1, y1 := positions[i][0], positions[i][1]
				x2, y2 := positions[j][0], positions[j][1]
				dx, dy := x2-x1, y2-y1
				g := gcd(dx, dy)
				dx /= g
				dy /= g

				for y := 0; y < rows; y++ {
					for x := 0; x < cols; x++ {
						if (y-y1)*(x2-x1) == (x-x1)*(y2-y1) {
							antinodePositions[[2]int{x, y}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return antinodePositions
}

func countAntinodes(filename string) int {
	grid := readInput(filename)
	antennas := findAntennas(grid)
	antinodes := calculateAntinodes(grid, antennas)
	return len(antinodes)
}

func main() {
	inputFile := "input.txt"
	result := countAntinodes(inputFile)
	fmt.Println("Total unique locations containing an antinode:", result)
}
