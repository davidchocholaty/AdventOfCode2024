package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Position struct {
	X, Y int
}

type Direction struct {
	DX, DY int
}

type Node struct {
	Pos Position
	Dir Direction
}

var directions = []Direction{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

type Grid [][]rune

func readGrid(filename string) (Grid, Position, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, Position{}, err
	}
	defer file.Close()

	var grid Grid
	var start Position

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		for col, char := range line {
			if char == 'S' {
				start = Position{row, col}
			}
		}
		row++
	}

	return grid, start, scanner.Err()
}

func changeDirection(dir Direction) []Direction {
	if dir == (Direction{0, 1}) || dir == (Direction{0, -1}) {
		return []Direction{{1, 0}, {-1, 0}}
	}
	return []Direction{{0, 1}, {0, -1}}
}

func nextMoves(grid Grid, pos Position, dir Direction, score int) ([]Node, []int) {
	var moves []Node
	var scores []int

	newPos := Position{pos.X + dir.DX, pos.Y + dir.DY}
	if newPos.X >= 0 && newPos.X < len(grid) && newPos.Y >= 0 && newPos.Y < len(grid[0]) && grid[newPos.X][newPos.Y] != '#' {
		moves = append(moves, Node{newPos, dir})
		scores = append(scores, score+1)
	}

	newDirs := changeDirection(dir)
	for _, newDir := range newDirs {
		moves = append(moves, Node{pos, newDir})
		scores = append(scores, score+1000)
	}

	return moves, scores
}

func updateNeighbours(neighbours map[Node](struct {
	Score    int
	Previous []Node
}), curr, new Node, score int) {
	if _, exists := neighbours[new]; !exists {
		neighbours[new] = struct {
			Score    int
			Previous []Node
		}{score, []Node{curr}}
	} else {
		lowestScore := neighbours[new].Score
		if score < lowestScore {
			neighbours[new] = struct {
				Score    int
				Previous []Node
			}{score, []Node{curr}}
		} else if score == lowestScore {
			neighbours[new] = struct {
				Score    int
				Previous []Node
			}{score, append(neighbours[new].Previous, curr)}
		}
	}
}

func dijkstra(grid Grid, start Position) (int, map[Node](struct {
	Score    int
	Previous []Node
})) {
	queue := make(map[Node]int)
	neighbours := make(map[Node](struct {
		Score    int
		Previous []Node
	}))

	for i, row := range grid {
		for j, char := range row {
			if char == 'E' || char == '.' || char == 'S' {
				for _, dir := range directions {
					node := Node{Position{i, j}, dir}
					queue[node] = math.MaxInt32
				}
			}
		}
	}

	queue[Node{start, Direction{0, 1}}] = 0
	lowestScore := math.MaxInt32
	//var endNode Node

	for len(queue) > 0 {
		var currNode Node
		minScore := math.MaxInt32
		for node, score := range queue {
			if score < minScore {
				minScore = score
				currNode = node
			}
		}

		currPos, currDir := currNode.Pos, currNode.Dir
		currScore := queue[currNode]
		delete(queue, currNode)

		if currScore == math.MaxInt32 {
			break
		}
		if currScore > lowestScore {
			break
		}
		if grid[currPos.X][currPos.Y] == 'E' {
			// endNode = currNode

			if currScore < lowestScore {
				lowestScore = currScore
			}
		}

		moves, scores := nextMoves(grid, currPos, currDir, currScore)
		for i, newNode := range moves {
			newScore := scores[i]
			if _, exists := queue[newNode]; exists {
				if newScore < queue[newNode] {
					queue[newNode] = newScore
				}
			}
			updateNeighbours(neighbours, currNode, newNode, newScore)
		}
	}

	return lowestScore, neighbours
}

func countUniquePositions(neighbours map[Node](struct {
	Score    int
	Previous []Node
}), endNode Node) int {
	stack := []Node{endNode}
	visited := make(map[Node]bool)
	positionSet := make(map[Position]bool)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		positionSet[curr.Pos] = true
		for _, prev := range neighbours[curr].Previous {
			if !visited[prev] {
				stack = append(stack, prev)
			}
		}
	}

	return len(positionSet)
}

func main() {
	grid, start, err := readGrid("input.txt")
	if err != nil {
		fmt.Println("Error reading grid:", err)
		return
	}

	lowestScore, neighbours := dijkstra(grid, start)
	fmt.Println("PART 1:", lowestScore)

	endNode := Node{}
	for node := range neighbours {
		if grid[node.Pos.X][node.Pos.Y] == 'E' {
			endNode = node
			break
		}
	}
	fmt.Println("PART 2:", countUniquePositions(neighbours, endNode))
}
