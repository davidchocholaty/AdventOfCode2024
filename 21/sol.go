package main

import (
	"fmt"
	"math"
)

// Directional Keypad
var directionalKeypad = map[[2]int]string{
	{0, 1}: "^",
	{0, 2}: "A",
	{1, 0}: "<",
	{1, 1}: "v",
	{1, 2}: ">",
}

// Numeric Keypad
var numericKeypad = map[[2]int]string{
	{0, 0}: "7",
	{0, 1}: "8",
	{0, 2}: "9",
	{1, 0}: "4",
	{1, 1}: "5",
	{1, 2}: "6",
	{2, 0}: "1",
	{2, 1}: "2",
	{2, 2}: "3",
	{3, 1}: "0",
	{3, 2}: "A",
}

var cache = map[string]int{}

func getNeighbours(pos [2]int, keypad map[[2]int]string) [][2]interface{} {
	valid := [][2]interface{}{}
	moves := map[[2]int]string{
		{-1, 0}: "^",
		{0, -1}: "<",
		{1, 0}:  "v",
		{0, 1}:  ">",
	}
	for direction, move := range moves {
		newPos := [2]int{pos[0] + direction[0], pos[1] + direction[1]}
		if _, exists := keypad[newPos]; exists {
			valid = append(valid, [2]interface{}{newPos, move})
		}
	}
	return valid
}

func getShortestPaths(sourcePos, targetPos [2]int, keypad map[[2]int]string) []string {
	var paths []string
	var stack = [][3]interface{}{{sourcePos, "", []interface{}{}}}
	minLen := math.MaxInt
	for len(stack) > 0 {
		currPos, path, visited := stack[len(stack)-1][0], stack[len(stack)-1][1].(string), stack[len(stack)-1][2].([]interface{})
		stack = stack[:len(stack)-1]

		if currPos, ok := currPos.([2]int); ok {
			if currPos == targetPos {
				paths = append(paths, path+"A")
				if len(paths[len(paths)-1]) < minLen {
					minLen = len(paths[len(paths)-1])
				}
			}

			for _, neighbour := range getNeighbours(currPos, keypad) {
				neighbourPos := neighbour[0].([2]int)
				if !contains(visited, neighbourPos) {
					stack = append(stack, [3]interface{}{neighbourPos, path + neighbour[1].(string), append(visited, currPos)})
				}
			}
		}
	}

	var shortestPaths []string
	for _, path := range paths {
		if len(path) == minLen {
			shortestPaths = append(shortestPaths, path)
		}
	}
	return shortestPaths
}

func contains(arr []interface{}, item [2]int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func inputSequence(sequence string, depth, nDirectionalKeypads int) int {
	if val, exists := cache[fmt.Sprintf("%s-%d", sequence, depth)]; exists {
		return val
	}

	if depth == nDirectionalKeypads {
		return len(sequence)
	}

	var keypad map[[2]int]string
	if depth == 0 {
		keypad = numericKeypad
	} else {
		keypad = directionalKeypad
	}

	invKeypad := make(map[string][2]int)
	for k, v := range keypad {
		invKeypad[v] = k
	}

	sourcePos := invKeypad["A"]
	totalButtonPresses := 0
	for _, char := range sequence {
		targetPos := invKeypad[string(char)]
		shortestPaths := getShortestPaths(sourcePos, targetPos, keypad)
		countPathButtonPresses := []int{}
		for _, path := range shortestPaths {
			countPathButtonPresses = append(countPathButtonPresses, inputSequence(path, depth+1, nDirectionalKeypads))
		}
		totalButtonPresses += min(countPathButtonPresses)
		sourcePos = targetPos
	}

	cache[fmt.Sprintf("%s-%d", sequence, depth)] = totalButtonPresses
	return totalButtonPresses
}

func min(nums []int) int {
	minVal := math.MaxInt
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func calculateScore(codes []string, nDirectionalKeypads int) int {
	totSum := 0
	for _, code := range codes {
		score := inputSequence(code, 0, nDirectionalKeypads)
		totSum += score * int(code[len(code)-2]-'0')
	}
	return totSum
}

func main() {
	codes := []string{"140A", "143A", "349A", "582A", "964A"}

	// PART 1
	cache = map[string]int{}
	fmt.Println("part 1:", calculateScore(codes, 3))

	// PART 2
	cache = map[string]int{}
	fmt.Println("part 2:", calculateScore(codes, 26))
}
