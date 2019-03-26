package skiplist

import "math/rand"

// Node 节点
type Node struct {
	Key     int         // 根据当前字段排序
	Data    interface{} // 数据
	ForWard []*Node     // 当前节点每层指向下一个的指针
}

// Skiplist 跳表
type Skiplist struct {
	Length int   // 数据个数
	Level  int   // 层级
	Header *Node // 头指针
}

const (
	maxLevel int = 11
)

// 获取随机的层级数
func (list *Skiplist) randomLevel() int {
	level := 1
	for ; level < list.Level && rand.Uint32()&0x1 == 1; level++ {
	}
	return level
}

// Add 添加值
func (list *Skiplist) Add(key int, data interface{}) {
	// 获取当前数据的插入层级
	level := list.randomLevel()
	// 创建新的节点
	newForward := make([]*Node, level)
	// 开始查询
	p := list.Header

	for l := level - 1; l >= 0; l-- {
		for {
			nextNode := p.ForWard[l]
			if nextNode == nil || nextNode.Key > key {
				newForward[l] = p // 第 l 层的插入点
				break
			} else if nextNode.Key == key {
				nextNode.Data = data // 如果排序下表相等，直接覆盖了？？？
				return
			} else {
				p = nextNode
			}
		}
	}

	// 各个层级的插入操作
	newNode := &Node{ForWard: newForward, Key: key, Data: data}
	for i, pNode := range newForward {
		pNode.ForWard[i], newNode.ForWard[i] = newNode, pNode.ForWard[i]
	}

	list.Length++

}

// GetListData 返回全部数据
func (list *Skiplist) GetListData() []interface{} {
	datas := make([]interface{}, list.Length)
	for i, p := 0, list.Header.ForWard[0]; p != nil; p = p.ForWard[0] {
		datas[i] = p.Data
		i++
	}
	return datas
}

// NewList 初始化一个 跳表
func NewList() *Skiplist {
	sl := new(Skiplist)
	sl.Level = 0
	sl.Header = newNodeByMaxLevel(maxLevel)
	return sl
}

func newNodeByMaxLevel(maxLevel int) *Node {
	pNodes := make([]*Node, maxLevel)
	return &Node{ForWard: pNodes}
}
