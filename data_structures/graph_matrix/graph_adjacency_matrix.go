package main

import "fmt"

// Graph struct represents a graph using an adjacency matrix.
type Graph struct {
	vertices int     // The number of vertices in the graph.
	matrix   [][]int // The adjacency matrix representing edges between vertices.
}

// NewGraph function creates and returns a new Graph with the specified number of vertices.
func NewGraph(vertices int) *Graph {
	// Initialize a 2D slice to represent the adjacency matrix.
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices) // Each row in the matrix is initialized with zeros.
	}
	return &Graph{
		vertices: vertices, // Store the number of vertices.
		matrix:   matrix,   // Store the initialized adjacency matrix.
	}
}

// AddEdge method adds an edge between two vertices in the graph.
func (g *Graph) AddEdge(v1, v2 int) {
	g.matrix[v1][v2] = 1 // Set the matrix entry to 1 to indicate an edge from v1 to v2.
	g.matrix[v2][v1] = 1 // Set the matrix entry to 1 to indicate an edge from v2 to v1, since it's an undirected graph.
}

// Print method prints the adjacency matrix of the graph.
func (g *Graph) Print() {
	for i := 0; i < g.vertices; i++ { // Loop over each vertex.
		fmt.Printf("%d: ", i)
		for j := 0; j < g.vertices; j++ { // Loop over each possible adjacent vertex.
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	graph.Print()
}
