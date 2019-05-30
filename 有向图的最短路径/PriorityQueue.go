package graph

type PriorityQueue struct {
	nodes []Vertex
	count int
}

func (pq PriorityQueue) Len() int {
	return len(pq.nodes)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.nodes[i].dist < pq.nodes[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}
