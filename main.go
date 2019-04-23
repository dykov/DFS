package main

import (
	"fmt"
)

var vertices int
var graph [][]bool
var used []bool

func main() {

	fmt.Print("Number of vertices: ")
	_, _ = fmt.Scanf("%d", &vertices)

	graph = make([][]bool, vertices)
	for i := range graph {
		graph[i] = make([]bool, vertices)
	}
	used = make([]bool, vertices)

	for i := 0; i < vertices; i++ {
		var from int
		var to int
		_, _ = fmt.Scanf("%d %d\n", &from, &to)
		graph[from-1][to-1], graph[to-1][from-1] = true, true
	}

	printGraph()

	dfs(0, 0)

}

func dfs(start, from int) {

	used[start] = true
	fmt.Printf("Vertex %d is visited from vertex %d\n", start+1, from+1)
	for i := 0; i < vertices; i++ {
		if graph[start][i] && !used[i] {
			dfs(i, start)
		}
	}

}

func printGraph() {

	for i := 0; i < vertices; i++ {
		if i == 0 {
			fmt.Printf("%2v", " ")
		}
		fmt.Printf("%5d |", i+1)
	}
	fmt.Println()
	for i := 0; i < vertices; i++ {
		fmt.Print(i+1, ":")
		for j := 0; j < vertices; j++ {
			fmt.Printf("%5t |", graph[i][j])
		}
		fmt.Println()
	}

}
