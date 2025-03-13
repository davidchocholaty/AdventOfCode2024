package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Graph map[string]map[string]bool

func readGraph(filename string) (Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := make(Graph)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		if len(parts) != 2 {
			continue
		}
		if graph[parts[0]] == nil {
			graph[parts[0]] = make(map[string]bool)
		}
		if graph[parts[1]] == nil {
			graph[parts[1]] = make(map[string]bool)
		}
		graph[parts[0]][parts[1]] = true
		graph[parts[1]][parts[0]] = true
	}

	return graph, scanner.Err()
}

func findConnections(graph Graph) int {
	connections := make(map[string]bool)

	for node, neighbours := range graph {
		if strings.HasPrefix(node, "t") {
			for node2 := range neighbours {
				for node3 := range graph[node2] {
					if graph[node][node3] {
						trio := []string{node, node2, node3}
						sort.Strings(trio)
						connections[strings.Join(trio, ",")] = true
					}
				}
			}
		}
	}

	return len(connections)
}

func bronKerbosch(graph Graph, R, P, X map[string]bool) []map[string]bool {
	if len(P) == 0 && len(X) == 0 {
		return []map[string]bool{R}
	}
	var cliques []map[string]bool
	Pcopy := make(map[string]bool)
	for k := range P {
		Pcopy[k] = true
	}

	for node := range Pcopy {
		neighbours := graph[node]
		newR := make(map[string]bool)
		for k := range R {
			newR[k] = true
		}
		newR[node] = true

		newP := make(map[string]bool)
		newX := make(map[string]bool)
		for k := range P {
			if neighbours[k] {
				newP[k] = true
			}
		}
		for k := range X {
			if neighbours[k] {
				newX[k] = true
			}
		}

		cliques = append(cliques, bronKerbosch(graph, newR, newP, newX)...)
		delete(P, node)
		X[node] = true
	}

	return cliques
}

func main() {
	graph, err := readGraph("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", findConnections(graph))

	R, P, X := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	for k := range graph {
		P[k] = true
	}

	allCliques := bronKerbosch(graph, R, P, X)
	for _, clique := range allCliques {
		if len(clique) > 12 {
			var sortedClique []string
			for node := range clique {
				sortedClique = append(sortedClique, node)
			}
			sort.Strings(sortedClique)
			fmt.Println(sortedClique)
		}
	}
}
