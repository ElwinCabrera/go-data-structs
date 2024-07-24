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
	return &PriorityQueue{list.NewDoublyLinkedList(), orderAsc}
}

func (queue *PriorityQueue) Push(value any, weight float64) {

	if queue.orderAsc {
		queue.list.InsertSortedAscBasedOnNodeWeight(value, weight)
	} else {
		queue.list.InsertSortedDescBasedOnNodeWeight(value, weight)
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
	it := list.NewListIterator(queue.list).Begin()

	lastVal := it.Get().Value
	for ; it.Get() != nil; it = it.Next() {
		lastVal = it.Get().Value
	}
	return lastVal
}

func (queue *PriorityQueue) Values() []any {
	var values []any
	it := list.NewListIterator(queue.list)
	for ; it.Get() != nil; it = it.Next() {
		values = append(values, it.Get().Value)
	}
	return values
}

func (queue *PriorityQueue) String() string {
	return fmt.Sprintf("%v", queue.Values())
}
