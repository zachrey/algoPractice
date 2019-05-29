package myheap

import "sort"

/**
最小堆和最大堆的实现
*/

// Interface 接口
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

// Init 初始化堆
func Init(h Interface) {
	for i := h.Len()/2 - 1; i >= 0; i-- {
		down(h, i, h.Len())
	}
}

// Push 添加
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1) // 往最后面添加元素，然后向上整理堆栈
}

// Pop 出堆
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Remove 删除指定的元素
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if i != n {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix 更新指定的元素
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) { // 如果有像下的调整，则不需要再向上调整
		up(h, i)
	}
}

func up(h Interface, i int) {
	for {
		j := (i - 1) / 2             // 父节点
		if i == j || !h.Less(j, i) { // 如果没有父节点或者子节点比父节点大（小）则结束
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j := i*2 - 1 // 左子节点
		if j >= n || j < 0 {
			break
		}

		if j1 := j + 1; j1 <= n && !h.Less(j, j1) { // 右子节点
			j = j1
		}

		if !h.Less(j, i) { // 比较父节点和子节点的大小
			break
		}

		h.Swap(i, j)
		i = j
	}
	return i > i0
}
