package list


import (
	"fmt"
)


type DoublyLinkedList LinkedList

func InitDoublyLinkedList() *DoublyLinkedList{
	return &DoublyLinkedList{nil, 0}
}


func (dll *DoublyLinkedList) InsertEnd(v any) *Node{
	head := dll.head
	new_node := &Node{v, nil, nil}
	if head == nil{
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

func (dll *DoublyLinkedList) InsertFront(v any) *Node{
	new_node := &Node{v, nil, nil}
	new_node.next = dll.head
	dll.head.prev = new_node
	dll.head = new_node
	dll.len++
	return new_node
}

func (dll *DoublyLinkedList) InsertBefore(v any, n *Node) *Node{
	new_node := &Node{v, nil, nil}
	if dll.head == nil || n == nil{
		dll.head = new_node
		return new_node
	}
	new_node.prev = n.prev
	if n.prev != nil {
		n.prev.next = new_node
	}
	n.prev = new_node
	new_node.next = n
	dll.len++
	return new_node
}

func (dll *DoublyLinkedList) InsertAfter(v any, n *Node) *Node{
	new_node := &Node{v, nil, nil}
	if dll.head == nil || n == nil{
		dll.head = new_node
		return new_node
	}
	new_node.prev = n
	new_node.next = n.next
	if n.next != nil{
		n.next.prev = new_node
	}
	n.next = new_node
	dll.len++
	return new_node
}




func (dll *DoublyLinkedList) PopFront() *Node {
	popped_node := dll.head
	if dll.head == nil {
		return nil
	}
	if dll.head.next == nil {
		dll.head = nil
		return popped_node
	}
	dll.head = dll.head.next
	dll.head.prev = nil
	dll.len--
	return popped_node
}

func (dll *DoublyLinkedList) PopBack() *Node {
	if dll.head == nil || dll.head.next == nil{
		dll.head = nil
		return nil
	}
	curr := dll.head
	for curr.next.next != nil {
		curr = curr.next
	}
	popped_node := curr.next
	popped_node.prev = nil
	curr.next = nil
	dll.len--
	return popped_node

}



func (dll *DoublyLinkedList) Remove(n *Node) bool{
	if n == nil {
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
	n.next = nil
	n.prev = nil
	n = nil
	dll.len--
	return true

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


func (dll *DoublyLinkedList) String() string{
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