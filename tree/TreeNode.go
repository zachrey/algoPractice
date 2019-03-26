package tree

import "fmt"

// Node 树 节点
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

// NewNode 产生一个新的节点
func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

func (node *Node) String() string {
	return fmt.Sprintf("v: %+v, left: %+v, right: %+v", node.Data, node.Left, node.Right)
}
