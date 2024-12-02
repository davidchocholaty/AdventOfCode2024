package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSave(report []int) {

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
	// fmt.Println(string(content))
	reports := strings.Split(string(content), "\n")

	counter := 0

	for _, line := range reports {
		report := strings.SplitN(line, " ", 8)

		position := 1
		flagInc := 0 // 1 => inceasing, -1 => decreasing, 0 => beginning state

		for position < len(report) {
			prev, err := strconv.Atoi(report[position-1])
			if err != nil {
				fmt.Printf("Error converting '%s' to integer: %v\n", report[position-1], err)
				flagInc = 0
				break
			}

			curr, err := strconv.Atoi(report[position])
			if err != nil {
				fmt.Printf("Error converting '%s' to integer: %v\n", report[position], err)
				flagInc = 0
				break
			}

			diff := prev - curr

			if (diff >= -3 && diff <= -1) && flagInc != -1 {
				flagInc = 1
			} else if (diff >= 1 && diff <= 3) && flagInc != 1 {
				flagInc = -1
			} else {
				flagInc = 0
				break
			}

			position = position + 1
		}

		if flagInc != 0 {
			counter = counter + 1
		}
	}

	fmt.Println(counter)
}
