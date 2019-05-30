package graph

import "testing"

func Test_graph(t *testing.T) {
	graph := New(6)
	graph.addEdge(0, 1, 10)
	graph.addEdge(1, 2, 15)
	graph.addEdge(2, 5, 5)
	graph.addEdge(1, 3, 2)
	graph.addEdge(3, 2, 1)
	graph.addEdge(3, 5, 12)
	graph.addEdge(0, 4, 15)
	graph.addEdge(4, 5, 10)

	graph.Dijkstra(0, 5)
}
