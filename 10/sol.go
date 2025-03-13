package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	i, j int
}

func main() {
	// Read input file
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

	// Create grid
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j, ch := range lines[i] {
			n, _ := strconv.Atoi(string(ch))
			grid[i][j] = n
		}
	}

	// Find starting positions (where grid value is 0)
	var startPositions []Position
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				startPositions = append(startPositions, Position{i, j})
			}
		}
	}

	tot1 := 0
	tot2 := 0

	for _, pos := range startPositions {
		paths := dfs(pos, grid)

		// Count unique destinations reached
		uniquePaths := make(map[Position]bool)
		for _, p := range paths {
			uniquePaths[p] = true
		}
		tot1 += len(uniquePaths)

		// Count number of paths walked
		tot2 += len(paths)
	}

	fmt.Println("part 1: ", tot1)
	fmt.Println("part 2: ", tot2)
}

func validMoves(pos Position, grid [][]int) []Position {
	// Return positions of valid moves from pos - move up by 1 each step
	var children []Position
	i, j := pos.i, pos.j

	// up
	if i > 0 {
		if grid[i-1][j]-grid[i][j] == 1 {
			children = append(children, Position{i - 1, j})
		}
	}
	// down
	if i < len(grid)-1 {
		if grid[i+1][j]-grid[i][j] == 1 {
			children = append(children, Position{i + 1, j})
		}
	}
	// right
	if j < len(grid[0])-1 {
		if grid[i][j+1]-grid[i][j] == 1 {
			children = append(children, Position{i, j + 1})
		}
	}
	// left
	if j > 0 {
		if grid[i][j-1]-grid[i][j] == 1 {
			children = append(children, Position{i, j - 1})
		}
	}

	return children
}

func dfs(startPos Position, grid [][]int) []Position {
	stack := []Position{startPos}
	// Keep track of all 9s that reach (can get there multiple ways so will have duplicates)
	var reached []Position

	// Track visited positions to avoid cycles
	visited := make(map[Position]bool)

	for len(stack) > 0 {
		// Pop from stack
		currPos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Skip if already visited
		if visited[currPos] {
			continue
		}
		visited[currPos] = true

		if grid[currPos.i][currPos.j] == 9 {
			reached = append(reached, currPos)
		}

		for _, pos := range validMoves(currPos, grid) {
			stack = append(stack, pos)
		}
	}

	return reached
}
