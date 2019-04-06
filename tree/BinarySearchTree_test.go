package tree

import "testing"

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}

func TestBST_Find(t *testing.T) {
	bst := NewBinarySearchTress(1, compareFunc)

	bst.Add(3)
	bst.Add(1)
	bst.Add(2)
	bst.Add(7)
	bst.Add(5)

	t.Log(bst.Find(2))
}

func TestBST_Insert(t *testing.T) {
	bst := NewBinarySearchTress(1, compareFunc)

	bst.Add(3)
	bst.Add(1)
	bst.Add(2)
	bst.Add(7)
	bst.Add(5)

	bst.InOrderTraverse()
}

func TestBST_DeleteA(t *testing.T) {
	bst := NewBinarySearchTress(1, compareFunc)

	bst.Add(3)
	bst.Add(2)
	bst.Add(7)
	bst.Add(5)

	t.Log(bst.Delete(7))

	bst.InOrderTraverse()
}

func TestBST_DeleteB(t *testing.T) {
	bst := NewBinarySearchTress(1, compareFunc)

	t.Log(bst.Delete(1))
	t.Log(bst.Root)

	bst.InOrderTraverse()
}

func TestBST_DeleteC(t *testing.T) {
	bst := NewBinarySearchTress(1, compareFunc)

	bst.Add(3)
	bst.Add(2)
	bst.Add(7)
	bst.Add(5)

	t.Log(bst.Delete(1))

	bst.InOrderTraverse()
}

func TestBST_DeleteD(t *testing.T) {
	bst := NewBinarySearchTress(1, compareFunc)

	bst.Add(3)
	bst.Add(2)
	bst.Add(5)

	t.Log(bst.Delete(1))

	bst.InOrderTraverse()
}
func TestBST_DeleteE(t *testing.T) {
	bst := NewBinarySearchTress(5, compareFunc)

	bst.Add(3)
	bst.Add(2)
	bst.Add(4)
	bst.Add(1)

	t.Log(bst.Delete(5))
	bst.InOrderTraverse()
}

func TestBST_DeleteF(t *testing.T) {
	bst := NewBinarySearchTress(5, compareFunc)

	bst.Add(3)
	bst.Add(2)
	bst.Add(4)
	bst.Add(1)

	t.Log(bst.Delete(2))
	bst.InOrderTraverse()
}

func TestBST_DeleteG(t *testing.T) {
	bst := NewBinarySearchTress(5, compareFunc)

	bst.Add(3)
	bst.Add(2)
	bst.Add(4)
	bst.Add(1)

	t.Log(bst.Delete(1))
	bst.InOrderTraverse()
}
