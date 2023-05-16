package skiplist

import "math/rand"

// node 节点
type node struct {
	key   int
	value any
	next  []*node
}

type SkipList struct {
	head      *node
	maxLevels int
	p         float32 // 索引分布概率参数
	level     int     // 当前层数
}

func newSkipListNode(key int, value interface{}, level int) *node {
	return &node{
		key:   key,
		value: value,
		next:  make([]*node, level), // 根据level新建next切片
	}
}

// 创建跳表的函数
func NewSkipList(p float32, maxLevels int) *SkipList {
	head := newSkipListNode(0, nil, maxLevels)               //新建头节点
	return &SkipList{head: head, maxLevels: maxLevels, p: p} //返回跳表结构体指针
}

func (s *SkipList) Insert(key int, value any) {
	current := s.head
	update := make([]*node, s.maxLevels) //创建更新数组用来记录每层需要更新的节点

	// 搜索待插入节点的位置，并记录每一层需要更新的节点
	for i := s.level - 1; i >= 0; i-- {
		// 找到i层上第一个键值小于要插入的key值的节点
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		update[i] = current
	}

	// 如果要插入的节点已经存在，直接更新其值即可
	if current.next[0] != nil && current.next[0].key == key {
		current.next[0].value = value
		return
	}

	// 根据索引分布概率参数随机生成新节点的层数
	newLevel := s.randomLevel()

	// 如果新建节点的层数大于当前跳表的层数，需要为更高的层次更新指针
	if newLevel > s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.head
		}
		s.level = newLevel
	}

	// 创建新的节点并依次与原来的节点相连
	newNode := newSkipListNode(key, value, newLevel)
	for i := 0; i < newLevel; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (s *SkipList) Search(key int) any {
	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
	}

	if current.next[0] != nil && current.next[0].key == key {
		return current.next[0].value
	}

	return nil
}

// 随机生成节点层数的函数
func (s *SkipList) randomLevel() int {
	level := 1
	for rand.Float32() < s.p && level < s.maxLevels {
		level++
	}
	return level
}
