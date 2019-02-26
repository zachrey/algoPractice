package arrqueue

// ArrQueue 数组队列
type ArrQueue struct {
	arr []int
	N   int
}

// Push 入队
func (aq *ArrQueue) Push(val int) int {
	aq.arr = append(aq.arr, val)
	aq.N++
	return aq.N
}

// Pop 入队
func (aq *ArrQueue) Pop() (int, int) {
	rst := aq.arr[0]
	aq.arr = aq.arr[1:]
	aq.N--
	return rst, aq.N
}

// GetAll 获取队列所有的值
func (aq ArrQueue) GetAll() []int {
	return aq.arr
}
