package tests

import (
	"github.com/ElwinCabrera/go-containers/stack-queue"
	"testing"
)

func testStackPush(t *testing.T, size int) (*stack_queue.Stack, []int) {
	stack := stack_queue.InitStack()
	values := getArrayOfRandomUniqueValues(size)
	for _, v := range values {
		stack.Push(v)
		if stack.Peek() != v {
			t.Errorf("Stack push expected %v, got %v", v, stack.Peek())
		}
	}

	if stack.Size() != len(values) {
		t.Errorf("Stack size expected %d, got %d", len(values), stack.Size())
	}
	return stack, values
}

func testQueuePush(t *testing.T, size int) (*stack_queue.Queue, []int) {
	queue := stack_queue.InitQueue()
	values := getArrayOfRandomUniqueValues(size)
	for _, v := range values {
		queue.Push(v)
		if queue.Back() != v {
			t.Errorf("Queue back expected %v, got %v", v, queue.Back())
		}
	}
	if queue.Size() != len(values) {
		t.Errorf("Queue size expected %d, got %d", len(values), queue.Size())
	}
	return queue, values
}

func TestStackPush(t *testing.T) {
	testStackPush(t, 10)
}

func TestStackPop(t *testing.T) {
	stack, values := testStackPush(t, 10)

	for i := len(values) - 1; i >= 0; i-- {
		if stack.Pop() != values[i] {
			t.Errorf("Stack pop expected %v, got %v", values[i], stack.Peek())
		}
	}
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}

func TestQueuePush(t *testing.T) {
	testQueuePush(t, 10)
}

func TestDequeue(t *testing.T) {
	queue, values := testQueuePush(t, 10)
	for _, v := range values {
		if queue.Dequeue() != v {
			t.Errorf("Dequeue expected %v, got %v", v, queue.Dequeue())
		}
	}
	if !queue.IsEmpty() {
		t.Errorf("Expected Queue to be empty")
	}

}
