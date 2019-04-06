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

// Find 根据指定的数据查找节点
func (bst *BST) Find(val interface{}) *Node {
	p := bst.Root
	for nil != p {
		comRst := bst.compareFunc(val, p.Data)
		if comRst == 0 {
			return p
		} else if comRst < 0 {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}

// Add 添加一个数据节点
func (bst *BST) Add(val interface{}) bool {
	p := bst.Root
	for nil != p {
		comRst := bst.compareFunc(val, p.Data)
		if comRst == 0 {
			return false
		} else if comRst > 0 {
			if p.Right == nil {
				n := NewNode(val)
				p.Right = n
				break
			}
			p = p.Right
		} else {
			if p.Left == nil {
				n := NewNode(val)
				p.Left = n
				break
			}
			p = p.Left
		}
	}
	return true
}

// Delete 删除节点
func (bst *BST) Delete(val interface{}) bool {
	p := bst.Root // 指向要删除的节点，初始化为根节点
	var pp *Node  // pp 记录的是 p 的父节点
	for p != nil && bst.compareFunc(val, p.Data) != 0 {
		pp = p
		comRst := bst.compareFunc(val, p.Data)
		if comRst > 0 {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	if p == nil {
		return false
	}
	// 要删除的节点有左右两个子节点的情况
	if p.Left != nil && p.Right != nil {
		minP := p.Right
		minPP := p
		for minP.Left != nil { // 找到右子树中的最小节点, 最小节点一定是在节点的左节点中
			minPP = minP
			minP = minP.Left
		}
		p.Data = minP.Data // 将最小数据放到要删除的p节点中
		p = minP           // 将要删除的节点指向 右子数找到的最小节点
		pp = minPP         // 将要删除的节点的父节点， 指向右子树中最小节点的父节点
	}

	// 如果删除的节点只有一个节点
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if pp == nil {
		bst.Root = nil // 删除的是根节点
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
	return true
}
