package arrstack

import "errors"

// Stack 栈
type Stack struct {
	arr []int
	N   int
}

// Push 进栈
func (s *Stack) Push(val int) int {
	s.arr = append(s.arr, val)
	s.N++
	return s.N
}

// Pop 出栈
func (s *Stack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("The stack is empty")
	}
	rst := s.arr[len(s.arr)-1]
	s.arr = s.arr[0:(len(s.arr) - 1)]
	s.N--
	return rst, nil
}

// GetStackVal 获取栈里面所有的值
func (s *Stack) GetStackVal() []int {
	return s.arr
}
