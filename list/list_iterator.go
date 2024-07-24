package list

type ListIterator struct {
	list    List
	current *Node
	end     *ListIterator
}

func NewListIterator(l List) *ListIterator {
	endItr := &ListIterator{nil, nil, nil}
	return &ListIterator{l, l.Head(), endItr}
}

func (it *ListIterator) Begin() *ListIterator {
	it.current = it.list.Head()
	if it.current == nil {
		return it.end
	}
	return it
}

func (it *ListIterator) Next() *ListIterator {
	it.current = it.current.next
	if it.current == nil {
		return it.end
	}
	return it
}

func (it *ListIterator) Get() *Node {
	return it.current
}

func (it *ListIterator) End() *ListIterator {
	return it.end
}

//func (it *ListIterator) Last() *ListIterator {
//
//}
//
//func (it *ListIterator) First() *ListIterator {
//
//}
