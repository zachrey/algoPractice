package graph

import (
	"container/list"
)

// Graph 有向图
type Graph struct {
	adj []*list.List // 邻接表
	v   int          // 顶点个数
}

// New 创建一个有向图
func New(v int) *Graph {
	var graph Graph
	graph.v = v
	graph.adj = make([]*list.List, v)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return &graph
}

func (g *Graph) addEdge(s, t, w int) {
	g.adj[s].PushBack(Edge{s, t, w})
}

// Edge 边
type Edge struct {
	sid int
	tid int
	w   int
}

// Vertex 顶点
type Vertex struct {
	id   int
	dist int
}
