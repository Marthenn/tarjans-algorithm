package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	set "github.com/golang-collections/collections/set"
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

func Min(a, b string) string {
	if LexographicCompare(a, b) {
		return a
	}
	return b
}

func AddEdge(adjList map[string][]string, names *set.Set, a, b string) map[string][]string {
	// check if b is already in a's list
	for _, v := range adjList[a] {
		if v == b {
			return adjList
		}
	}

	adjList[a] = append(adjList[a], b)
	names.Insert(a)
	names.Insert(b)

	return adjList
}

func FileToAdjList(dir string, names *set.Set) map[string][]string {
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
		adjList = AddEdge(adjList, names, words[0], words[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return adjList
}

func MakeUndirected(adjList map[string][]string, names *set.Set) map[string][]string {
	for k, v := range adjList {
		for _, w := range v {
			adjList = AddEdge(adjList, names, w, k)
		}
	}

	return adjList
}
