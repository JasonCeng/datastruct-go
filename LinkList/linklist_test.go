package LinkList

import "testing"

func TestLinkList(t *testing.T) {
	list := LinkList{}
	list.AppendNode(1)
	list.AppendNode(2)
	list.AppendNode(3)
	list.AppendNode(4)
	list.AppendNode(5)
	list.AppendNode(6)
	list.printLinkList()
	list.DeleteNthFromEnd(3)
	list.printLinkList()

	list.PreAppendNode(0)
	list.printLinkList()

	list.DeleteNode(2)
	list.printLinkList()

	list.UpdateNode(3, 4)
	list.printLinkList()
}
