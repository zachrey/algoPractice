package linkqueue

import "errors"

type node struct {
	Val  int
	Next *node
}

// LinkQueue 链式队列
type LinkQueue struct {
	head *node
	N    int
}

// Push 入队
func (lq *LinkQueue) Push(val int) int {
	node := new(node)
	node.Next = nil
	node.Val = val
	q := lq.head
	if q == nil { // 一个元素都没有
		lq.head = node
	} else {
		for ; q.Next != nil; q = q.Next { // 找到尾节点
		}
		q.Next = node
	}
	lq.N++
	return lq.N
}

// Pop 出队
func (lq *LinkQueue) Pop() (int, int, error) {
	if lq.head == nil {
		return 0, 0, errors.New("The queue is empty")
	}
	rst := lq.head.Val
	lq.head = lq.head.Next
	lq.N--
	return rst, lq.N, nil
}

// GetAll 获取队列所有的值
func (lq *LinkQueue) GetAll() []int {
	rst := []int{}
	for p := lq.head; p != nil; p = p.Next {
		rst = append(rst, p.Val)
	}
	return rst
}
