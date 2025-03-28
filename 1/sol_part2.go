package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
)

func str_2_int(str string) int {
	intValue, err := strconv.Atoi(str)

	// Check for errors
	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}

	return intValue
}

func obtain_separated_lists(content []byte) ([]int, []int) {
	lines := strings.Split(string(content), "\n")
	var list1, list2 []int

	for _, line := range lines {
		parts := strings.SplitN(line, "   ", 2)
		if len(parts) > 0 {
			list1 = append(list1, str_2_int(parts[0])) // First part before "   "
		}
		if len(parts) > 1 {
			list2 = append(list2, str_2_int(parts[1])) // Second part after "   "
		}
	}

	return list1, list2
}

func getCount(value int, counts map[int]int) int {
	if count, exists := counts[value]; exists {
		return count
	}
	return 0
}

func main() {
	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <file-path>")
	}

	// Get the file path from the first command-line argument
	filePath := os.Args[1]

	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the content of the file to the console
	list1, list2 := obtain_separated_lists(content)

	counts := make(map[int]int)
	for _, value := range list2 {
		counts[value]++
	}

	score := 0

	for _, value := range list1 {
		score = score + value * getCount(value, counts)
	}

	fmt.Println(score)
}
