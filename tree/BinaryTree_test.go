package tree

import "testing"

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	bt := NewBinaryTree("A")
	bt.Root.Left = NewNode("B")
	bt.Root.Right = NewNode("C")
	bt.Root.Left.Left = NewNode("D")
	bt.Root.Left.Right = NewNode("E")
	bt.Root.Right.Left = NewNode("F")
	bt.Root.Right.Right = NewNode("G")

	bt.InOrderTraverse()
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	bt := NewBinaryTree("A")
	bt.Root.Left = NewNode("B")
	bt.Root.Right = NewNode("C")
	bt.Root.Left.Left = NewNode("D")
	bt.Root.Left.Right = NewNode("E")
	bt.Root.Right.Left = NewNode("F")
	bt.Root.Right.Right = NewNode("G")

	bt.PreOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	bt := NewBinaryTree("A")
	bt.Root.Left = NewNode("B")
	bt.Root.Right = NewNode("C")
	bt.Root.Left.Left = NewNode("D")
	bt.Root.Left.Right = NewNode("E")
	bt.Root.Right.Left = NewNode("F")
	bt.Root.Right.Right = NewNode("G")

	bt.PostOrderTraverse()
}
