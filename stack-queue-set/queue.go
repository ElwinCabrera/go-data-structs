package stack_queue_set

import (
	"fmt"
	"github.com/ElwinCabrera/go-data-structs/list"
)

type Queue struct {
	list     list.List
	last_val any
}

func InitQueue() *Queue {
	return &Queue{list.InitDoublyLinkedList(), 0}
}

func (queue *Queue) Push(value any) {
	queue.list.InsertEnd(value, 0)
	queue.last_val = value
}

func (queue *Queue) Dequeue() any {
	if queue.IsEmpty() {
		return nil
	}
	return queue.list.PopFront().Value
}

func (queue *Queue) Size() int {
	return queue.list.Len()
}

func (queue *Queue) IsEmpty() bool {
	return queue.list.Len() == 0
}

func (queue *Queue) Clear() {
	queue.list.Clear()
}

func (queue *Queue) Front() any {
	if queue.IsEmpty() {
		return nil // more like an error but will leave as nil for now and will fix later
	}
	return queue.list.Head()
}

func (queue *Queue) Back() any {
	return queue.last_val
}

func (queue *Queue) Values() []any {
	var values []any
	it := list.InitListIterator(queue.list)
	for ; it.Get() != nil; it = it.Next() {
		values = append(values, it.Get().Value)
	}
	return values
}

func (queue *Queue) String() string {
	return fmt.Sprintf("%v", queue.Values())
}
