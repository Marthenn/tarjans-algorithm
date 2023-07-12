package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LexographicCompare(a, b string) bool {
	// check lexographically whether a < b
	// return true if a <= b or a is shorter than b

	if len(a) < len(b) {
		return true
	}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		}
	}

	return true
}

func AddEdge(adjList map[string][]string, names map[string]int, a, b string) (map[string][]string, map[string]int) {
	// check if b is already in a's list
	for _, v := range adjList[a] {
		if v == b {
			return adjList, names
		}
	}

	adjList[a] = append(adjList[a], b)
	names[a] = 1
	names[b] = 1

	return adjList, names
}

func FileToAdjList(dir string) (map[string][]string, map[string]int) {
	adjList := make(map[string][]string)
	names := make(map[string]int)

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		adjList, names = AddEdge(adjList, names, words[0], words[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return adjList, names
}

func MakeUndirected(adjList map[string][]string, names map[string]int) (map[string][]string, map[string]int) {
	for k, v := range adjList {
		for _, w := range v {
			adjList, names = AddEdge(adjList, names, w, k)
		}
	}

	return adjList, names
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
