package list

type ListIterator struct {
	list    List
	current *Node
}

func InitListIterator(l List) *ListIterator {
	return &ListIterator{l, l.Head()}
}

func (it *ListIterator) Begin() *ListIterator {
	it.current = it.list.Head()
	return it
}

func (it *ListIterator) Next() *ListIterator {
	it.current = it.current.next
	return it
}

func (it *ListIterator) Get() *Node {
	return it.current
}
