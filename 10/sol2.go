package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	row, col int
}

type Trail []Position

func parseInput(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var mapData [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row[i] = num
		}
		mapData = append(mapData, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mapData, nil
}

func findTrailheads(mapData [][]int) []Position {
	var trailheads []Position
	for row := 0; row < len(mapData); row++ {
		for col := 0; col < len(mapData[0]); col++ {
			if mapData[row][col] == 0 {
				trailheads = append(trailheads, Position{row, col})
			}
		}
	}
	return trailheads
}

func bfsScore(mapData [][]int, start Position) int {
	queue := []Position{start}
	visited := make(map[Position]bool)
	visited[start] = true
	reachableNines := make(map[Position]bool)

	directions := []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentHeight := mapData[current.row][current.col]

		for _, dir := range directions {
			next := Position{current.row + dir.row, current.col + dir.col}

			// Check if position is within bounds
			if next.row >= 0 && next.row < len(mapData) && next.col >= 0 && next.col < len(mapData[0]) {
				// Check if not visited and height is exactly one more than current
				if !visited[next] && mapData[next.row][next.col] == currentHeight+1 {
					visited[next] = true
					queue = append(queue, next)

					if mapData[next.row][next.col] == 9 {
						reachableNines[next] = true
					}
				}
			}
		}
	}

	return len(reachableNines)
}

func bfsRating(mapData [][]int, start Position) int {
	type QueueItem struct {
		pos  Position
		path Trail
	}

	queue := []QueueItem{{start, Trail{start}}}
	distinctTrails := make(map[string]bool)
	directions := []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		currentPos := item.pos
		path := item.path
		currentHeight := mapData[currentPos.row][currentPos.col]

		for _, dir := range directions {
			nextPos := Position{currentPos.row + dir.row, currentPos.col + dir.col}

			// Check if position is within bounds
			if nextPos.row >= 0 && nextPos.row < len(mapData) && nextPos.col >= 0 && nextPos.col < len(mapData[0]) {
				// Check if height is exactly one more than current
				if mapData[nextPos.row][nextPos.col] == currentHeight+1 {
					newPath := make(Trail, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextPos)

					if mapData[nextPos.row][nextPos.col] == 9 {
						// Convert path to string for use as map key
						pathKey := fmt.Sprintf("%v", newPath)
						distinctTrails[pathKey] = true
					} else {
						queue = append(queue, QueueItem{nextPos, newPath})
					}
				}
			}
		}
	}

	return len(distinctTrails)
}

// calculateTotalScoreAndRating calculates the sum of scores and ratings for all trailheads
func calculateTotalScoreAndRating(mapData [][]int) (int, int) {
	trailheads := findTrailheads(mapData)
	totalScore := 0
	totalRating := 0

	for _, trailhead := range trailheads {
		totalScore += bfsScore(mapData, trailhead)
		totalRating += bfsRating(mapData, trailhead)
	}

	return totalScore, totalRating
}

func main() {
	inputFile := "input.txt"
	mapData, err := parseInput(inputFile)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	totalScore, totalRating := calculateTotalScoreAndRating(mapData)
	fmt.Printf("Total score of all trailheads: %d\n", totalScore)
	fmt.Printf("Total rating of all trailheads: %d\n", totalRating)
}
