package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Position struct {
	Row int
	Col int
}

func readInput(warehouseFile, movesFile string) ([]string, []string, error) {
	// Read warehouse file
	warehouse, err := readLines(warehouseFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading warehouse file: %w", err)
	}

	// Read moves file
	movesList, err := readLines(movesFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading moves file: %w", err)
	}

	return warehouse, movesList, nil
}

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

func getDirection(symbol rune) (int, int) {
	switch symbol {
	case '<':
		return 0, -1
	case '>':
		return 0, 1
	case '^':
		return -1, 0
	case 'v':
		return 1, 0
	default:
		return 0, 0
	}
}

func findRobot(grid [][]rune) Position {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				return Position{Row: i, Col: j}
			}
		}
	}
	return Position{-1, -1}
}

func updateGridPart1(grid [][]rune, direction rune) [][]rune {
	di, dj := getDirection(direction)

	robot := findRobot(grid)

	i, j := robot.Row+di, robot.Col+dj
	canMove := false
	for grid[i][j] != '#' {
		if grid[i][j] == '.' {
			canMove = true
			break
		}
		i += di
		j += dj
	}

	if canMove {
		i, j := robot.Row+di, robot.Col+dj
		for grid[i][j] == 'O' {
			i += di
			j += dj
		}
		// Note: if there are no boxes to move, the next line gets overwritten
		// by moving the robot to this pos
		grid[i][j] = 'O'
		grid[robot.Row][robot.Col] = '.'
		grid[robot.Row+di][robot.Col+dj] = '@'
	}

	return grid
}

func getAllBoxes(pos Position, di int, grid [][]rune) (map[Position]bool, bool) {
	boxes := make(map[Position]bool)
	stack := []Position{pos}
	canMove := true

	for len(stack) > 0 {
		currPos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		i, j := currPos.Row, currPos.Col
		if grid[i][j] == '[' {
			boxes[Position{i, j}] = true
			boxes[Position{i, j + 1}] = true
			stack = append(stack, Position{i + di, j})
			stack = append(stack, Position{i + di, j + 1})
		} else if grid[i][j] == ']' {
			boxes[Position{i, j}] = true
			boxes[Position{i, j - 1}] = true
			stack = append(stack, Position{i + di, j})
			stack = append(stack, Position{i + di, j - 1})
		} else if grid[i][j] == '#' {
			canMove = false
		}
	}

	return boxes, canMove
}

func updateGridPart2(grid [][]rune, direction rune) [][]rune {
	robot := findRobot(grid)
	di, dj := getDirection(direction)

	if grid[robot.Row+di][robot.Col+dj] == '.' {
		grid[robot.Row+di][robot.Col+dj] = '@'
		grid[robot.Row][robot.Col] = '.'
		return grid
	}

	if direction == '>' || direction == '<' {
		i, j := robot.Row+di, robot.Col+dj
		canMove := false
		for grid[i][j] != '#' {
			if grid[i][j] == '.' {
				canMove = true
				break
			}
			i += di
			j += dj
		}

		if canMove {
			j := robot.Col + dj
			for grid[robot.Row][j] == '[' || grid[robot.Row][j] == ']' {
				j += dj
			}

			startPos, endPos := 0, 0
			if direction == '>' {
				startPos, endPos = robot.Col+2, j+1
			} else {
				startPos, endPos = j, robot.Col
			}

			symb := '['
			for col := startPos; col < endPos; col++ {
				grid[robot.Row][col] = symb
				if symb == '[' {
					symb = ']'
				} else {
					symb = '['
				}
			}

			grid[robot.Row][robot.Col] = '.'
			grid[robot.Row+di][robot.Col+dj] = '@'
		}
	} else if direction == '^' || direction == 'v' {
		i, j := robot.Row+di, robot.Col+dj

		boxMap, canMove := getAllBoxes(Position{i, j}, di, grid)

		if canMove {

			var boxes []Position
			for pos := range boxMap {
				boxes = append(boxes, pos)
			}

			if direction == '^' {
				sort.Slice(boxes, func(i, j int) bool {
					return boxes[i].Row < boxes[j].Row ||
						(boxes[i].Row == boxes[j].Row && boxes[i].Col < boxes[j].Col)
				})
			} else {
				sort.Slice(boxes, func(i, j int) bool {
					return boxes[i].Row > boxes[j].Row ||
						(boxes[i].Row == boxes[j].Row && boxes[i].Col > boxes[j].Col)
				})
			}

			for _, box := range boxes {
				grid[box.Row+di][box.Col] = grid[box.Row][box.Col]
				grid[box.Row][box.Col] = '.'
			}

			grid[robot.Row][robot.Col] = '.'
			grid[robot.Row+di][robot.Col+dj] = '@'
		}
	}

	return grid
}

func createGrid(warehouse []string) [][]rune {
	grid := make([][]rune, len(warehouse))
	for i, line := range warehouse {
		grid[i] = []rune(line)
	}
	return grid
}

func createLargerWarehouse(warehouse []string) [][]rune {
	largerWarehouse := make([][]rune, len(warehouse))
	for i, line := range warehouse {
		var newLine []rune
		for _, char := range line {
			switch char {
			case '#':
				newLine = append(newLine, '#', '#')
			case 'O':
				newLine = append(newLine, '[', ']')
			case '.':
				newLine = append(newLine, '.', '.')
			case '@':
				newLine = append(newLine, '@', '.')
			default:
				newLine = append(newLine, char, '.')
			}
		}
		largerWarehouse[i] = newLine
	}
	return largerWarehouse
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func main() {
	warehouse, movesList, err := readInput("warehouse.txt", "moves.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Part 1
	grid := createGrid(warehouse)
	for _, moveLine := range movesList {
		for _, move := range moveLine {
			grid = updateGridPart1(grid, move)
		}
	}

	total := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'O' {
				total += 100*i + j
			}
		}
	}
	fmt.Println("Part 1:", total)

	// Part 2
	largerGrid := createLargerWarehouse(warehouse)
	for _, moveLine := range movesList {
		for _, move := range moveLine {
			largerGrid = updateGridPart2(largerGrid, move)
		}
	}

	total = 0
	for i := range largerGrid {
		for j := range largerGrid[i] {
			if largerGrid[i][j] == '[' {
				total += 100*i + j
			}
		}
	}
	fmt.Println("Part 2:", total)
}
