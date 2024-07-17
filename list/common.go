package list

type ListComparableTypes interface {
	~int | ~float32 | ~float64 | ~string
}

type Node struct {
	Value  any
	Weight int
	next   *Node
	prev   *Node
}

type LinkedList struct {
	head *Node
	len  int
}

type List interface {
	InsertEnd(v any, weight int) *Node
	InsertFront(v any, weight int) *Node
	InsertBefore(v any, weight int, n *Node) *Node
	InsertAfter(v any, weight int, n *Node) *Node
	InsertSortedAscBasedOnNodeWeight(value any, weight int) *Node
	InsertSortedDescBasedOnNodeWeight(value any, weight int) *Node
	Remove(n *Node) bool
	PopFront() *Node
	PopBack() *Node
	Find(v any) []*Node
	Len() int
	Head() *Node
	Clear()
}
