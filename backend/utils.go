package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
)

func AddEdge(adjList map[string][]string, a, b string) map[string][]string{
	adjList[a] = append(adjList[a], b)
	return adjList
}

func FileToAdjList(dir string) map[string][]string{
	adjList := make(map[string][]string)

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		words := strings.Split(line, " ")
		adjList = AddEdge(adjList, words[0], words[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return adjList
}
