package graph

import "testing"

func Test_graph(t *testing.T) {
	graph := New(5)
	graph.addEdge(0, 1, 2)
	graph.addEdge(0, 2, 2)
	graph.addEdge(0, 3, 2)
	graph.addEdge(0, 4, 2)
	graph.addEdge(0, 5, 2)

	t.Logf("%v", graph.adj[1].Len())
}
