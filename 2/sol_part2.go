package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)


func validate(report []string) bool {
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

	return (flagInc != 0)
}

func removeAtIndex(slice []string, index int) []string {
    return append(append([]string{}, slice[:index]...), slice[index+1:]...)
}


func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <file-path>")
	}

	filePath := os.Args[1]

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	reports := strings.Split(string(content), "\n")

	counter := 0

	for _, line := range reports {
		report := strings.SplitN(line, " ", 8)
		
		if validate(report) {
			counter = counter + 1
		} else {
			position := 0
			
			for position < len(report) {
				if validate(removeAtIndex(report, position)) {
					counter = counter + 1
					break
				}
				position = position + 1
			}
		}
	}

	fmt.Println(counter)
}
