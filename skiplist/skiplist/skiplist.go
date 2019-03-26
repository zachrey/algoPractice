package skiplist

const (
	// MaxNumberOfLevels 最大的索引层级，应该根据 链表总数动态计算出来
	MaxNumberOfLevels int = 11
)

// Node 节点
type Node struct {
	Key     int // key值
	Value   int // value值
	Forward []*Node
}

// Skiplist 跳表整体
type Skiplist struct {
	Level  int
	Header *Node
}

// NewSkiplist 初始化一个skiplist
func NewSkiplist() *Skiplist {
	var sl *Skiplist
	// 初始化空间
	sl = &Skiplist{}
	// 设置跳表的层level，初始的层为0层（数组从0开始）
	sl.Level = 0
	// 生成header部分
	sl.Header = newNodeOfLevel(sl.Level)
	// 将header的forward数组清空
	for i := 0; i < MaxNumberOfLevels; i++ {
		sl.Header.Forward[i] = nil
	}
	return sl
}

func newNodeOfLevel(maxLevel int) *Node {
	nodeArr := make([]*Node, maxLevel)
	return &Node{Forward: nodeArr}
}
