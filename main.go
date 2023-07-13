package main

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"

	"io/ioutil"
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/golang-collections/collections/set"
)

func main() {
	adjList := make(map[string][]string)
	names := set.New()

	fmt.Println("Graph input:")
	fmt.Println("1. File")
	fmt.Println("2. Type")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Enter file path:")
		var dir string
		fmt.Scanln(&dir)
		adjList = FileToAdjList(dir, names)
	} else {
		fmt.Println("Enter number of edges:")
		var n int
		var u, v string
		fmt.Scanln(&n)

		for i := 0; i < n; i++ {
			fmt.Scanln(&u, &v)
			adjList = AddEdge(adjList, names, u, v)
		}
	}

	low := Tarjan(adjList, names)
	bridge := Bridge(adjList, names)

	g := OutputGraph(adjList, names, low, bridge)

	outputData, err := ioutil.ReadFile(g)
	if err != nil {
		log.Fatal(err)
	}

	outputGraph, err := graphviz.ParseBytes(outputData)
	if err != nil {
		log.Fatal(err)
	}

	gv := graphviz.New()
	gv.SetLayout(graphviz.CIRCO)

	outputImage, err := gv.RenderImage(outputGraph)
	if err != nil {
		log.Fatal(err)
	}

	// check if the file exists
	_, err = os.Stat("output.png")
	if err == nil {
		os.Remove("output.png")
	}

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(file, outputImage)
	if err != nil {
		log.Fatal(err)
	}

	// execute start command to open the image
	cmd := exec.Command("cmd", "start", "output.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error opening image")
		log.Fatal(err)
	}
}
