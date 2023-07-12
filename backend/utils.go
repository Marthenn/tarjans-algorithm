package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddEdge(adjList map[string][]string, a, b string) map[string][]string {
	// check if b is alread in a's list
	for _, v := range adjList[a] {
		if v == b {
			return adjList
		}
	}

	adjList[a] = append(adjList[a], b)
	return adjList
}

func FileToAdjList(dir string) map[string][]string {
	adjList := make(map[string][]string)

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		adjList = AddEdge(adjList, words[0], words[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return adjList
}

func MakeUndirected(adjList map[string][]string) map[string][]string {
	for k, v := range adjList {
		for _, w := range v {
			adjList = AddEdge(adjList, w, k)
		}
	}

	return adjList
}
