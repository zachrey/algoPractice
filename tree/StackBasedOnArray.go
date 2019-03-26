package tree

import "fmt"

// ArrayStack 基于数组的栈
type ArrayStack struct {
	data []interface{}
	top  int
}

// NewArrayStack 创建一个基于数组的栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 32),
		top:  -1,
	}
}

// IsEmpty 栈是否为空
func (as *ArrayStack) IsEmpty() bool {
	if as.top == -1 {
		return true
	}
	return false
}

// Push 入栈
func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top++
	}

	if as.top > len(as.data)-1 {
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

// Pop 出栈
func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	v := as.data[as.top]
	as.top--
	return v
}

// Top 获取顶部元素
func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.data[as.top]
}

// Flush 情况栈
func (as *ArrayStack) Flush() {
	as.top = -1
}

// Print 打印
func (as *ArrayStack) Print() {
	if as.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := as.top; i >= 0; i-- {
			fmt.Println(as.data[i])
		}
	}
}
