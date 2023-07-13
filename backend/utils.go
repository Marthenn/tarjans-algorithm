package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	set "github.com/golang-collections/collections/set"
)

type pair struct { // pair of strings
	a string
	b string
}

func InPairList(u string, v string, list []pair) bool {
	pair1 := pair{u, v}
	pair2 := pair{v, u}

	for _, p := range list {
		if p == pair1 || p == pair2 {
			return true
		}
	}
	return false
}

func Min(a, b int) int {
	if a < b {
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

func FileToAdjList(name string, names *set.Set) map[string][]string {
	adjList := make(map[string][]string)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dir = dir + "/tests/" + name
	fmt.Println(dir)

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

func ListToSet(list []string) *set.Set {
	set := set.New()

	for _, v := range list {
		set.Insert(v)
	}

	return set
}

func SetToList(set *set.Set) []string {
	list := []string{}

	set.Do(func(x interface{}) {
		list = append(list, x.(string))
	})

	return list
}
