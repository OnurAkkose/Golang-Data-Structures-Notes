package main

import "fmt"

// Graph struct represents a graph using an adjacency list.
type Graph struct {
	vertices map[int][]int // The vertices map stores each vertex and its corresponding list of adjacent vertices.
}

// NewGraph function creates and returns a new Graph.
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)}
}

// AddEdge method adds an edge between two vertices in the graph.
func (g *Graph) AddEdge(v1, v2 int) {
	g.vertices[v1] = append(g.vertices[v1], v2) // Adds v2 to the adjacency list of v1.
	g.vertices[v2] = append(g.vertices[v2], v1) // Adds v1 to the adjacency list of v2 to ensure it's an undirected graph.
}

// Print method prints the graph in the form of adjacency lists.
func (g *Graph) Print() {
	for key, value := range g.vertices { // Iterates over each vertex and its adjacent vertices.
		fmt.Printf("%d -> %v\n", key, value)
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	graph.Print()
}
