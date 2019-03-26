package tree

// BST 查找二叉树
type BST struct {
	*BinaryTree
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

// NewBinarySearchTress 创建一个查找二叉树
func NewBinarySearchTress(
	rootV interface{},
	compareFunc func(v, nodeV interface{}) int,
) *BST {
	if compareFunc == nil {
		return nil
	}

	return &BST{
		BinaryTree:  NewBinaryTree(rootV),
		compareFunc: compareFunc,
	}
}
