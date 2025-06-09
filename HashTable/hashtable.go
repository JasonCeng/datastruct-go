package HashTable

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
