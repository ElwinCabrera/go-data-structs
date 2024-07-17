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
	InsertEnd(v any) *Node
	InsertFront(v any) *Node
	InsertBefore(v any, n *Node) *Node
	InsertAfter(v any, n *Node) *Node
	//InsertSortedAscBasedOnNodeWeight(value any, weight int) *Node
	InsertSortedDescBasedOnNodeWeight(value any, weight int) *Node
	Remove(n *Node) bool
	PopFront() *Node
	PopBack() *Node
	Find(v any) []*Node
	Len() int
	Head() *Node
	Clear()
}

func insertInListSortedDesc(ll List, weight int) {

	list_it := InitListIterator(ll).Begin()

	nodeInserted := false
	if list_it.Get() == nil {
		ll.InsertEnd(weight)
	} else if weight <= list_it.Get().Value.(int) {
		//insert front
		ll.InsertFront(weight)
	} else {
		lastNonNilNode := list_it.Get()
		for list_it.Get() != nil {
			if weight <= list_it.Get().Value.(int) {
				//insert before head.next
				ll.InsertBefore(weight, list_it.Get())
				nodeInserted = true
				break
			}
			if list_it.Get() != nil {
				lastNonNilNode = list_it.Get()
			}
			list_it.Next()
		}
		if !nodeInserted {
			//insert after head.next //baisically insert end but since we know where the list ends just do it inplace
			ll.InsertAfter(weight, lastNonNilNode)
			//ll.InsertEnd(weight)
		}
	}
}
