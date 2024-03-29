package main

import (
	"os"
	"path/filepath"

	graph "github.com/dominikbraun/graph"
	draw "github.com/dominikbraun/graph/draw"
	set "github.com/golang-collections/collections/set"
)

var colors_list []string = []string{"red", "blue", "green", "yellow", "orange", "purple", "pink", "brown", "black", "white"}

func OutputGraph(adjList map[string][]string, names *set.Set, low map[string]int, bridges []pair) string {
	// low set
	low_set := set.New()
	for _, v := range low {
		low_set.Insert(v)
	}

	// initialize low to color mapping
	count := 0
	colors := map[int]string{}
	low_set.Do(func(x interface{}) {
		colors[x.(int)] = colors_list[count]
		count++
	})

	// initialize the graph
	g := graph.New(graph.StringHash, graph.Directed())
	names.Do(func(x interface{}) {
		x_color := colors[low[x.(string)]]
		g.AddVertex(x.(string), graph.VertexAttribute("color", x_color)) // give the vertex a color based on its low value
	})
	for k, v := range adjList {
		for _, w := range v {
			if InPairList(k, w, bridges) {
				g.AddEdge(k, w, graph.EdgeAttribute("color", "red"))
			} else {
				g.AddEdge(k, w)
			}
		}
	}

	file, _ := os.Create("output.dot")
	_ = draw.DOT(g, file)

	// return the path to the output.dot
	path, err := filepath.Abs("output.dot")
	if err != nil {
		panic(err)
	}
	return path
}
