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

// InOrderTraverse 中序遍历
func (bt *BinaryTree) InOrderTraverse() {
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

// PreOrderTraverse 前序遍历
func (bt *BinaryTree) PreOrderTraverse() {
	p := bt.Root
	st := NewArrayStack()

	for !st.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.Data)
			st.Push(p)
			p = p.Left
		} else {
			p = st.Pop().(*Node).Right
		}
	}
	fmt.Println()
}

// PostOrderTraverse 后序遍历
func (bt *BinaryTree) PostOrderTraverse() {
	st1 := NewArrayStack()
	st2 := NewArrayStack()

	st1.Push(bt.Root)

	for !st1.IsEmpty() {
		p := st1.Pop().(*Node)
		st2.Push(p)
		if p.Left != nil {
			st1.Push(p.Left)
		}
		if p.Right != nil {
			st1.Push(p.Right)
		}
	}

	for !st2.IsEmpty() {
		fmt.Printf("%+v ", st2.Pop().(*Node).Data)
	}
	fmt.Println()
}
