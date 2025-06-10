package HashTable

import "fmt"

type HashNode struct {
	Key   int
	Value interface{}
	Next  *HashNode
}

type HashTable struct {
	Array  []*HashNode
	Size   int
	Length int
}

// 初始化哈希表
func NewHashTable(length int) *HashTable {
	return &HashTable{
		Array:  make([]*HashNode, length),
		Size:   0,
		Length: length,
	}
}

// hash算法
func (h *HashTable) hash(key int) int {
	return int(key) % h.Length
}

// 增：插入哈希节点到哈希表
// 改：修改某个哈希节点的值
func (h *HashTable) Put(key int, value interface{}) {
	index := h.hash(key)
	node := h.Array[index]

	if node != nil {
		if node.Key == key {
			node.Value = value
			return
		}
		node = node.Next
	}
	// 新建节点并插入哈希表
	newNode := &HashNode{
		Key:   key,
		Value: value,
		Next:  h.Array[index], // 头插法
	}
	h.Array[index] = newNode
	h.Size++
}

// 查：查找某个哈希节点的值
func (h *HashTable) Get(key int) (interface{}, bool) {
	index := h.hash(key)
	node := h.Array[index]

	for node != nil {
		if node.Key == key {
			return node.Value, true
		}
		node = node.Next
	}
	return nil, false
}

// 删：删除某个哈希节点
func (h *HashTable) Delete(key int) bool {
	index := h.hash(key)
	node := h.Array[index]

	// 处理头结点
	if node != nil && node.Key == key {
		h.Array[index] = node.Next
		h.Size--
		return true
	}

	// 处理非头节点
	prev := node
	for prev.Next != nil {
		if prev.Next.Key == key {
			prev.Next = prev.Next.Next
			h.Size--
			return true
		}
		prev = prev.Next
	}
	return false
}

// 打印链表
func (h *HashTable) printHashTable() {
	for i, node := range h.Array {
		fmt.Printf("Bucket: %d:", i)
		for node != nil {
			fmt.Printf("Key: %v, Value: %v\n", node.Key, node.Value)
			node = node.Next
		}
		fmt.Println("")
	}
}
