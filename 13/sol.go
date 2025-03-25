package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solveSystem(a, b, target [2]float64) (float64, float64) {
	a1, a2 := a[0], a[1]
	b1, b2 := b[0], b[1]
	c1, c2 := target[0], target[1]

	det := a1*b2 - b1*a2
	if det == 0 {
		return 0, 0 // No valid solution
	}

	x := (c1*b2 - b1*c2) / det
	y := (a1*c2 - c1*a2) / det
	return x, y
}

func parseNumbers(s string) [2]float64 {
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(s, -1)
	var numbers [2]float64
	for i, match := range matches {
		if i >= 2 {
			break
		}
		numbers[i], _ = strconv.ParseFloat(match, 64)
	}
	return numbers
}

func main() {
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	indices := []int{-1}
	for i, line := range lines {
		if line == "" {
			indices = append(indices, i)
		}
	}

	tot1, tot2 := 0.0, 0.0
	for i := 0; i < len(indices); i++ {
		var descript []string
		if i == len(indices)-1 {
			descript = lines[indices[i]+1:]
		} else {
			descript = lines[indices[i]+1 : indices[i+1]]
		}

		machine := map[string][2]float64{
			"A": parseNumbers(descript[0]),
			"B": parseNumbers(descript[1]),
			"X": parseNumbers(descript[2]),
		}

		// PART 1
		a, b := solveSystem(machine["A"], machine["B"], machine["X"])
		if a == float64(int(a)) && b == float64(int(b)) && a <= 100 && b <= 100 {
			tot1 += 3*a + b
		}

		// PART 2
		xValues := machine["X"]
		xValues[0] += 10000000000000
		xValues[1] += 10000000000000
		machine["X"] = xValues
		a, b = solveSystem(machine["A"], machine["B"], machine["X"])
		if a == float64(int(a)) && b == float64(int(b)) {
			tot2 += 3*a + b
		}
	}

	fmt.Println(int(tot1))
	fmt.Println(int(tot2))
}
