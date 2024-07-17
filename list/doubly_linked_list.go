package list

import (
	"fmt"
)

type DoublyLinkedList LinkedList

func InitDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, 0}
}

func (dll *DoublyLinkedList) InsertEnd(v any) *Node {
	head := dll.head
	new_node := &Node{Value: v}
	if head == nil {
		dll.head = new_node
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = new_node
		new_node.prev = head
	}
	dll.len++
	return new_node
}

func (dll *DoublyLinkedList) InsertFront(v any) *Node {
	new_node := &Node{Value: v}
	if dll.head == nil {
		dll.head = new_node
	} else {
		new_node.next = dll.head
		dll.head.prev = new_node
		dll.head = new_node
	}
	dll.len++
	return new_node
}

func (dll *DoublyLinkedList) InsertBefore(v any, n *Node) *Node {

	new_node := &Node{Value: v}
	if dll.head == nil || n == nil {
		if n == nil && dll.head != nil {
			return nil
		}
		dll.head = new_node
	} else {
		if n.prev != nil {
			new_node.prev = n.prev
			n.prev.next = new_node
		}
		n.prev = new_node
		new_node.next = n
		if n == dll.head {
			dll.head = new_node
		}
	}
	dll.len++
	return new_node
}

func (dll *DoublyLinkedList) InsertAfter(v any, n *Node) *Node {

	new_node := &Node{Value: v}

	if dll.head == nil || n == nil {
		if n == nil && dll.head != nil {
			return nil
		}
		dll.head = new_node
	} else {
		new_node.prev = n

		if n.next != nil {
			new_node.next = n.next
			n.next.prev = new_node
		}
		n.next = new_node
	}

	dll.len++
	return new_node
}

func (dll *DoublyLinkedList) InsertSortedDescBasedOnNodeWeight(value any, weight int) *Node {

	curr := dll.head
	nodeInserted := false
	var res *Node
	if curr == nil {
		res = dll.InsertEnd(value)
	} else if weight <= curr.Weight {
		//insert front
		res = dll.InsertFront(value)
	} else {

		for curr.next != nil {
			if weight <= curr.next.Weight {
				//insert before head.next
				res = dll.InsertBefore(value, curr.next)
				nodeInserted = true
				break
			}

			curr = curr.next
		}
		if !nodeInserted && curr.next == nil {
			//basically ll.InsertEnd(value) since we know where the list ends just do it in-place
			res = dll.InsertAfter(value, curr)

		}
	}
	return res
}

func (dll *DoublyLinkedList) PopFront() *Node {
	popped_node := dll.head
	if dll.head != nil {
		dll.head = dll.head.next
		if dll.head != nil {
			dll.head.prev = nil
		}
		popped_node.next = nil
		popped_node.prev = nil
	}

	dll.len--
	return popped_node
}

func (dll *DoublyLinkedList) PopBack() *Node {
	popped_node := dll.head
	if dll.head != nil {

		for popped_node.next != nil {
			popped_node = popped_node.next
		}
		if popped_node.prev != nil {
			popped_node.prev.next = nil
		}
		popped_node.prev = nil

		if popped_node == dll.head {
			dll.head = nil
		}

	} else {
		dll.head = nil
	}

	dll.len--
	return popped_node

}

func (dll *DoublyLinkedList) Find(v any) []*Node {
	var foundNodes []*Node
	currentNode := dll.head
	for currentNode != nil {
		if currentNode.Value == v {
			foundNodes = append(foundNodes, currentNode)
		}
		currentNode = currentNode.next
	}
	return foundNodes
}

func (dll *DoublyLinkedList) Remove(n *Node) bool {
	if dll.head == nil {
		return false // this is more of an error than anything else
	}
	if n == nil {
		return true
	}

	if n == dll.head {
		dll.PopFront()
		return true
	}

	prev := n.prev
	next := n.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	n = nil
	dll.len--
	return true

}

func (dll *DoublyLinkedList) Detach(n *Node) *Node {
	if dll.head == nil || n == nil {
		return nil // this is more of an error than anything else
	}

	if n == dll.head {
		return dll.PopFront()
	}

	prev := n.prev
	next := n.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	n.next = nil
	n.prev = nil

	dll.len--
	return n

}

func (dll *DoublyLinkedList) Clear() {
	curr := dll.head
	for curr != nil {
		next := curr.next
		curr = nil
		curr = next
	}
	dll.head = nil
	dll.len = 0
}

func (dll *DoublyLinkedList) Len() int {
	return dll.len
}

func (dll *DoublyLinkedList) Head() *Node {
	return dll.head
}

func (dll *DoublyLinkedList) String() string {
	head := dll.head

	var arr []any

	//result_str := "["
	result_str := ""
	i := 0
	for head != nil {
		//v, _ := head.value.(string)
		//result_str += v + " "
		i++
		arr = append(arr, head.Value)
		head = head.next
	}
	//fmt.Printf("loop count=%v ", i)
	//fmt.Printf("slice len=%v ", len(arr))
	fmt.Printf("len=%v ", dll.len)
	//result_str += "]"
	result_str = fmt.Sprintf("%v", arr)
	return result_str
}
