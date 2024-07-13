package stack_queue_set

import (
	"fmt"
	"github.com/ElwinCabrera/go-containers/list"
)

type Set struct {
	list list.List
}

func NewSet() *Set {
	return &Set{list.InitDoublyLinkedList()}
}

func (set *Set) Add(element any) {
	if !set.Contains(element) {
		set.list.InsertEnd(element)
	}
}

func (set *Set) Remove(element any) {
	valuesFound := set.list.Find(element)
	if len(valuesFound) != 0 {
		set.list.Remove(valuesFound[0])
	}
}

//will have to create a hashmap that maps index to node
//func (set *Set) Get(idx int) any {
//
//}

func (set *Set) Contains(element any) bool {
	valuesFound := set.list.Find(element)
	if len(valuesFound) != 0 {
		return true
	}
	return false
}

func (set *Set) Size() int {
	return set.list.Len()
}

func (set *Set) Empty() bool {
	return set.Size() == 0
}

func (set *Set) Clear() {
	set.list.Clear()
}

func (set *Set) Values() []any {
	var values []any
	it := list.InitListIterator(set.list)
	for it := it.Begin(); it.Get() != nil; it = it.Next() {
		values = append(values, it.Get().Value)
	}
	return values
}

func (set *Set) String() string {
	return fmt.Sprintf("%v", set.Values())
}
