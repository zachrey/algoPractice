package tree

import "fmt"

// BinaryTree 二叉树
type BinaryTree struct {
	Root *Node
}

// NewBinaryTree 创建一个二叉树
func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

// InOrderTree 中序遍历
func (bt *BinaryTree) InOrderTree() {
	p := bt.Root
	st := NewArrayStack()

	for !st.IsEmpty() || nil != p {
		if nil != p {
			st.Push(p)
			p = p.Left
		} else {
			curNode := st.Pop().(*Node)
			fmt.Printf("%+v ", curNode.Data)
			p = curNode.Right
		}
	}
	fmt.Println()
}
