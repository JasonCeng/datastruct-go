package LinkList

import "fmt"

// 链表节点结构
type Node struct {
	Value int
	Next  *Node
}

// 链表结构
type LinkList struct {
	Head *Node
	Tail *Node
	Len  int
}

func (l *LinkList) Init() {
	l.Head = nil
	l.Tail = nil
}

// 增：在链表末尾增加节点
func (l *LinkList) AppendNode(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		current := l.Head
		// 找到尾结点
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// 增：在链表的头增加节点
func (l *LinkList) PreAppendNode(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
		l.Len++
	}
}

// 删：按值删除链表中某个节点
func (l *LinkList) DeleteNode(value int) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Len--
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.Len--
			return
		}
		current = current.Next
	}
}

// 删：删除倒数第k个节点
// 思路：采用双指针。1.先将p1向前移动k-1步；2.p1、p2同时向前移动，直至p1到达最后一个节点。3.将p2之前的一个元素的next指针指向p2之后的元素
func (l *LinkList) DeleteNthFromEnd(k int) {
	p1 := l.Head
	p2 := l.Head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

}

// 查找：按值查找某个链表节点
func (l *LinkList) SearchNode(value int) *Node {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// 改：更新链表中原有的旧值为新值
func (l *LinkList) UpdateNode(oldValue, newValue int) bool {
	node := l.SearchNode(oldValue)
	if node != nil {
		node.Value = newValue
		return true
	}
	return false
}

// 打印链表
func (l *LinkList) printLinkList() {
	current := l.Head
	for current != nil {
		if current.Next != nil {
			fmt.Print(current.Value, "->")
		} else {
			fmt.Print(current.Value)
		}
		current = current.Next
	}
	fmt.Println()
}
