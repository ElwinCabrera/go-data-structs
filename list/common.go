package list

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

type List interface {
	InsertEnd(v any) *Node
	InsertFront(v any) *Node
	InsertBefore(v any, n *Node) *Node
	InsertAfter(v any, n *Node) *Node
	Remove(n *Node) bool
	PopFront() *Node
	PopBack() *Node
	Find(v any) []*Node
	Len() int
	Head() *Node
	Clear()
}
