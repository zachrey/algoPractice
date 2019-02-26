package linkstack

import "errors"

// Node 节点
type Node struct {
	Val  int
	Next *Node
}

// LinkStack 链式栈
type LinkStack struct {
	head *Node
	N    int
}

// Push 进栈
func (ls *LinkStack) Push(val int) int {
	node := new(Node)
	node.Val = val
	node.Next = ls.head
	ls.head = node
	ls.N++
	return ls.N
}

// Pop 出栈
func (ls *LinkStack) Pop() (int, error) {
	if ls.head == nil {
		return 0, errors.New("The stack is empty")
	}
	rst := ls.head.Val
	ls.head = ls.head.Next
	ls.N--
	return rst, nil
}

// GetStackAll 获取所有的值
func (ls *LinkStack) GetStackAll() []int {
	rst := []int{}
	for p := ls.head; p != nil; p = p.Next {
		rst = append(rst, p.Val)
	}
	return rst
}
