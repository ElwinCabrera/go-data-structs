package stack_queue_set

import (
	"fmt"
	"github.com/ElwinCabrera/go-data-structs/list"
)

type Stack struct {
	list            list.List
	last_val_pushed any
}

func InitStack() *Stack {
	return &Stack{list.InitDoublyLinkedList(), 0}
}

func (stack *Stack) Push(value any) {
	stack.list.InsertEnd(value, 0)
	stack.last_val_pushed = value
}

func (stack *Stack) Pop() any {
	if stack.IsEmpty() {
		return nil
	}
	return stack.list.PopBack().Value
}

func (stack *Stack) Peek() any {
	if stack.IsEmpty() {
		return nil // more like an error but will leave as nil for now and will fix later
	}
	return stack.last_val_pushed
}

func (stack *Stack) IsEmpty() bool {
	return stack.list.Len() == 0
}

func (stack *Stack) Size() int {
	return stack.list.Len()
}

func (stack *Stack) Clear() {
	stack.list.Clear()
}

func (stack *Stack) Values() []any {
	var values []any
	it := list.InitListIterator(stack.list)
	for ; it.Get() != nil; it = it.Next() {
		values = append(values, it.Get().Value)
	}
	return values
}

func (stack *Stack) String() string {
	return fmt.Sprintf("%v", stack.Values())
}
