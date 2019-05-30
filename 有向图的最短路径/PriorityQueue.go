package graph

import "container/heap"

// PriorityQueue 优先队列
type PriorityQueue []Vertex

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Pop 删除最后一个元素
func (pq *PriorityQueue) Pop() interface{} {
	n := pq.Len()
	item := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return item
}

// Push 在最后添加元素
func (pq *PriorityQueue) Push(v interface{}) {
	*pq = append(*pq, v.(Vertex))
}

// queueUpdate 更新某个节点
func queueUpdate(pq *PriorityQueue, v Vertex) {
	for i := range *pq {
		if (*pq)[i].id == v.id {
			(*pq)[i] = v
			heap.Fix(pq, i)
			return
		}
	}
}
