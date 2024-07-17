package stack_queue_set

import (
	"fmt"
	"github.com/ElwinCabrera/go-data-structs/list"
)

type PriorityQueue struct {
	list     list.List
	orderAsc bool
}

func NewPriorityQueue(orderAsc bool) *PriorityQueue {
	return &PriorityQueue{list.InitDoublyLinkedList(), orderAsc}
}

func (queue *PriorityQueue) Push(value any, priority int) {

	if queue.orderAsc {
		queue.list.InsertSortedAscBasedOnNodeWeight(value, priority)
	} else {
		queue.list.InsertSortedDescBasedOnNodeWeight(value, priority)
	}
}

func (queue *PriorityQueue) Dequeue() any {
	if queue.IsEmpty() {
		return nil
	}
	return queue.list.PopFront().Value
}

func (queue *PriorityQueue) Size() int {
	return queue.list.Len()
}

func (queue *PriorityQueue) IsEmpty() bool {
	return queue.list.Len() == 0
}

func (queue *PriorityQueue) Clear() {
	queue.list.Clear()
}

func (queue *PriorityQueue) Front() any {
	if queue.IsEmpty() {
		return nil // more like an error but will leave as nil for now and will fix later
	}
	return queue.list.Head()
}

func (queue *PriorityQueue) Back() any {
	if queue.IsEmpty() {
		return nil // more like an error but will leave as nil for now and will fix later
	}
	it := list.InitListIterator(queue.list).Begin()

	lastVal := it.Get().Value
	for ; it.Get() != nil; it = it.Next() {
		lastVal = it.Get().Value
	}
	return lastVal
}

func (queue *PriorityQueue) Values() []any {
	var values []any
	it := list.InitListIterator(queue.list)
	for ; it.Get() != nil; it = it.Next() {
		values = append(values, it.Get().Value)
	}
	return values
}

func (queue *PriorityQueue) String() string {
	return fmt.Sprintf("%v", queue.Values())
}
