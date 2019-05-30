package graph

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
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

// Dijkstra 最短路径算法
func (g *Graph) Dijkstra(s, t int) {
	predecessor := make([]int, g.v)
	vertes := make([]Vertex, g.v)
	for i := range vertes {
		vertes[i].dist = math.MaxInt16
		vertes[i].id = i
	}

	var queue PriorityQueue // 小顶堆

	inQueue := make([]bool, g.v) // 记录每个顶点是否被遍历过

	vertes[s].dist = 0           // 起点的距离设置为0
	inQueue[s] = true            // 第一个点遍历了
	heap.Push(&queue, vertes[s]) // 第一个点进入小顶堆

	for queue.Len() != 0 {
		minVertex := heap.Pop(&queue).(Vertex)
		if minVertex.id == t {
			break
		}
		for e := g.adj[minVertex.id].Front(); e != nil; e = e.Next() {
			edge := e.Value.(Edge)
			nextVertex := vertes[edge.tid]
			if minVertex.dist+edge.w < nextVertex.dist {
				nextVertex.dist = minVertex.dist + edge.w
				predecessor[nextVertex.id] = minVertex.id
				if inQueue[nextVertex.id] {
					queueUpdate(&queue, nextVertex) // 如果下一个顶点是已经遍历过得,更新就行
				} else {
					heap.Push(&queue, nextVertex)
					inQueue[nextVertex.id] = true
				}
			}
		}
	}

	fmt.Print(s)

	print(s, t, predecessor)
}

func print(s, t int, predecessor []int) {
	if s == t {
		return
	}
	print(s, predecessor[t], predecessor)
	fmt.Printf("->%d", t)
}
