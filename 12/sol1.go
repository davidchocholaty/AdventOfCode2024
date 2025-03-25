package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid, scanner.Err()
}

func calculateAreaAndPerimeter(grid []string, visited [][]bool, start [2]int, plantType byte) (int, int) {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := [][2]int{start}
	visited[start[0]][start[1]] = true

	area := 0
	perimeter := 0

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		area++

		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				if grid[nx][ny] == plantType && !visited[nx][ny] {
					visited[nx][ny] = true
					queue = append(queue, [2]int{nx, ny})
				} else if grid[nx][ny] != plantType {
					perimeter++
				}
			} else {
				perimeter++
			}
		}
	}

	return area, perimeter
}

func calculateTotalPrice(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	totalPrice := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] {
				plantType := grid[r][c]
				area, perimeter := calculateAreaAndPerimeter(grid, visited, [2]int{r, c}, plantType)
				totalPrice += area * perimeter
			}
		}
	}

	return totalPrice
}

func main() {
	inputFile := "input.txt"
	gardenMap, err := readInputFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	totalPrice := calculateTotalPrice(gardenMap)
	fmt.Printf("PART1: %d\n", totalPrice)
}
