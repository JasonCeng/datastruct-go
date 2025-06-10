package HashTable

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	fmt.Println("======创建HashTable======")
	hashTable := NewHashTable(4)
	hashTable.Put(2, "002")
	hashTable.Put(3, "003")
	hashTable.Put(5, "005")
	hashTable.Put(7, "007")
	hashTable.printHashTable()

	fmt.Println("======获取HashTable key=3 的值======")
	value, _ := hashTable.Get(3)
	fmt.Println(value)

	fmt.Println("======删除HashTable key=7 的节点======")
	hashTable.Delete(7)
	hashTable.printHashTable()
}
