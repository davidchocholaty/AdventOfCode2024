package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func findAntennas(grid [][]rune) map[rune][]Point {
	antennas := make(map[rune][]Point)
	for y, row := range grid {
		for x, char := range row {
			if (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
				antennas[char] = append(antennas[char], Point{X: x, Y: y})
			}
		}
	}
	return antennas
}

func calculateAntinodes(grid [][]rune, antennas map[rune][]Point) map[Point]struct{} {
	antinodes := make(map[Point]struct{})
	rows, cols := len(grid), len(grid[0])

	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				x1, y1 := positions[i].X, positions[i].Y
				x2, y2 := positions[j].X, positions[j].Y

				dx, dy := x2-x1, y2-y1

				ax, ay := x1-dx, y1-dy
				bx, by := x2+dx, y2+dy

				if ax >= 0 && ax < cols && ay >= 0 && ay < rows {
					antinodes[Point{X: ax, Y: ay}] = struct{}{}
				}
				if bx >= 0 && bx < cols && by >= 0 && by < rows {
					antinodes[Point{X: bx, Y: by}] = struct{}{}
				}
			}
		}
	}
	return antinodes
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
	fmt.Printf("Total unique locations containing an antinode: %d\n", result)
}
