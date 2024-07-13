package list

import (
	"fmt"
)

type SinglyLinkedList LinkedList

func InitSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{nil, 0}
}

func (sll *SinglyLinkedList) InsertEnd(v any) *Node {
	head := sll.head
	new_node := &Node{v, nil, nil}
	if head == nil {
		sll.head = new_node
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = new_node
	}
	sll.len++
	return new_node
}

func (sll *SinglyLinkedList) InsertFront(v any) *Node {
	new_node := &Node{v, nil, nil}
	new_node.next = sll.head
	sll.head = new_node
	sll.len++
	return new_node
}

func (sll *SinglyLinkedList) InsertBefore(v any, n *Node) *Node {

	new_node := &Node{v, nil, nil}
	if sll.head == nil || n == nil {
		if n == nil && sll.head != nil {
			return nil
		}
		sll.head = new_node
	} else if n == sll.head {
		new_node.next = sll.head
		sll.head = new_node

	} else {
		curr := sll.head
		for curr.next != n && curr != nil {
			curr = curr.next
		}
		curr.next = new_node
		new_node.next = n
	}
	sll.len++
	return new_node
}

func (sll *SinglyLinkedList) InsertAfter(v any, n *Node) *Node {
	if n == nil {
		return nil
	}
	new_node := &Node{v, nil, nil}
	hold := n.next
	n.next = new_node
	new_node.next = hold
	sll.len++
	return new_node
}

func (sll *SinglyLinkedList) PopFront() *Node {
	popped_node := sll.head
	if sll.head != nil {
		sll.head = sll.head.next
	}
	sll.len--
	return popped_node
}

func (sll *SinglyLinkedList) PopBack() *Node {
	popped_node := sll.head
	if sll.head != nil && sll.head.next != nil {
		curr := sll.head
		for curr.next.next != nil {
			curr = curr.next
		}
		popped_node = curr.next
		curr.next = nil
	} else {
		sll.head = nil
	}
	sll.len--
	return popped_node

}

func (sll *SinglyLinkedList) Find(v any) []*Node {
	var foundNodes []*Node
	currentNode := sll.head
	for currentNode != nil {
		if currentNode.Value == v {
			foundNodes = append(foundNodes, currentNode)
		}
		currentNode = currentNode.next
	}
	return foundNodes
}

func (sll *SinglyLinkedList) Remove(n *Node) bool {
	if sll.head == nil {
		return false // this is more of an error than anything else
	}
	if n == sll.head {
		sll.PopFront()
		return true
	}
	curr := sll.head
	for curr.next != n && curr != nil {
		curr = curr.next
	}
	if curr == nil {
		return false
	}
	hold := curr.next.next
	curr.next = hold
	n.next = nil
	n = nil
	sll.len--
	return true

}

func (sll *SinglyLinkedList) Clear() {
	curr := sll.head
	for curr != nil {
		next := curr.next
		curr = nil
		curr = next
	}
	sll.head = nil
	sll.len = 0
}

func (sll *SinglyLinkedList) Len() int {
	return sll.len
}

func (sll *SinglyLinkedList) Head() *Node {
	return sll.head
}

func (sll *SinglyLinkedList) String() string {
	head := sll.head

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
	fmt.Printf("len=%v ", sll.len)
	//result_str += "]"
	result_str = fmt.Sprintf("%v", arr)
	return result_str
}
