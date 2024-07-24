package list

type ComparableTypes interface {
	~int | ~float32 | ~float64 | ~string
}

type Node struct {
	Value  any
	Weight float64
	next   *Node
	prev   *Node
}

type LinkedList struct {
	head *Node
	len  int
}

type List interface {
	InsertEnd(v any, weight float64) *Node
	InsertFront(v any, weight float64) *Node
	InsertBefore(v any, weight float64, n *Node) *Node
	InsertAfter(v any, weight float64, n *Node) *Node
	InsertSortedAscBasedOnNodeWeight(value any, weight float64) *Node
	InsertSortedDescBasedOnNodeWeight(value any, weight float64) *Node
	Remove(n *Node) bool
	PopFront() *Node
	PopBack() *Node
	Find(v any) []*Node
	Len() int
	Head() *Node
	Clear()
	ValuesAsSlice() []any
}
