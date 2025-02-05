package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
	// fmt.Println(string(content))
}
