package main

import (
	set "github.com/golang-collections/collections/set"
	stack "github.com/golang-collections/collections/stack"
)

var time int // time of discovery of each vertex

func tarjanDFS(u string, low map[string]int, disc map[string]int, st *stack.Stack, st_mem map[string]bool, adjList map[string][]string) {
	disc[u] = time
	low[u] = time
	time++

	st.Push(u)
	st_mem[u] = true

	for _, v := range adjList[u] {
		if disc[v] == -1 {
			tarjanDFS(v, low, disc, st, st_mem, adjList)
			low[u] = Min(low[u], low[v])
		} else if st_mem[v] {
			low[u] = Min(low[u], disc[v])
		}
	}
}

func Tarjan(adjList map[string][]string, names *set.Set) map[string]int {
	// setup the variables for tarjan
	low := make(map[string]int)
	disc := make(map[string]int)
	st_mem := make(map[string]bool)
	st := stack.New()

	names.Do(func(x interface{}) {
		low[x.(string)] = -1
		disc[x.(string)] = -1
		st_mem[x.(string)] = false
	})

	time = 0 // reset time

	// run tarjan
	names.Do(func(x interface{}) {
		v := x.(string)
		if disc[v] == -1 {
			tarjanDFS(v, low, disc, st, st_mem, adjList)
		}
	})

	return low // the clustering the SCC
}

func bridgeDFS(u string, low map[string]int, disc map[string]int, visited map[string]bool, parent string, adjList map[string][]string, res *[]pair) {
	visited[u] = true
	disc[u] = time
	low[u] = time
	time++

	for _, v := range adjList[u] {
		if !visited[v] {
			parent = u
			bridgeDFS(v, low, disc, visited, parent, adjList, res)
			low[u] = Min(low[u], low[v])
			if low[v] > disc[u] {
				*res = append(*res, pair{u, v})
			}
		} else if v != parent {
			low[u] = Min(low[u], disc[v])
		}
	}
}

func Bridge(adjList map[string][]string, names *set.Set) []pair {
	visited := make(map[string]bool)
	low := make(map[string]int)
	disc := make(map[string]int)

	names.Do(func(x interface{}) {
		visited[x.(string)] = false
		low[x.(string)] = -1
		disc[x.(string)] = -1
	})

	time = 0

	res := make([]pair, 0)

	names.Do(func(x interface{}) {
		v := x.(string)
		if !visited[v] {
			bridgeDFS(v, low, disc, visited, "", adjList, &res)
		}
	})

	return res
}
