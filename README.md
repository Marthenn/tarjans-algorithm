# tarjans-algorithm
## DISCLAIMER
- For some reason, sometimes the program will result in some errors when running the program. When it happens, just rerun the program several times until the output is correct. I don't know why this happens, but I think it happens because of go memory management or execution process since when I take it slow usually the program run just fine.

## Description
- This is a simple implementation of Tarjan's Algorithm in C++.
- The algorithm is used to find strongly connected components in a graph.
- Other than that, it can also be used to find bridges in a graph.

## Requirements
- go version go1.20.3
- windows operating system (for image popup, if not in windows then just open the output.png file)

## Library Used
- `github.com/goccy/go-graphviz` for visualizing the graph.
- `github.com/golang-collections/collections` for using stack and set data structure.
- `github.com/dominikbraun/graph` for creating the graph for graphviz.

## How to Use
- Open the terminal in the root directory of the project.
- If you want to compile the program, run the following command `go build` (by default it should already being built when cloned).
- Run the following command to run the program `./main`.

## Algorithm
- Tarjan's algorithm is a graph theory algorithm for finding the strongly connected components of a graph.
- It runs in linear time O(|V|+|E|) and uses depth-first search.
- It is named for its discoverer, Robert Tarjan.
- To find bridge using Tarjan's Algorithm, we just need to add a `parent` variable to keep track of the parent of each node during the DFS.

## Edges in Graph During DFS
- Tree Edge: An edge that is part of the DFS tree.
- Back Edge: An edge that connects a vertex to one of its ancestors in the DFS traversal. Back edge indicates a cycle in the graph.
- Forward Edge: An edge that connects a vertex to a descendant vertex during the DFS traversal.
- Cross Edge: An edge that connects a vertex to a vertex that is neither its ancestor nor descendant in the DFS traversal. This edge type can go between two vertices in the same DFS tree or between two vertices in different DFS trees. Cross edges are also called as `cross-links`.

## References
- https://youtu.be/wUgWX0nc4NY
- https://youtu.be/thLQYBlz2DM
- https://www.geeksforgeeks.org/tarjan-algorithm-find-strongly-connected-components/
- https://www.geeksforgeeks.org/bridge-in-a-graph/
