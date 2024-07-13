package tests

import (
	"github.com/ElwinCabrera/go-containers/stack-queue-set"
	"testing"
)

func testStackPush(t *testing.T, size int) (*stack_queue_set.Stack, []int) {
	stack := stack_queue_set.InitStack()
	values := getArrayOfRandomNonNegativeUniqueValues(size)
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

func testQueuePush(t *testing.T, size int) (*stack_queue_set.Queue, []int) {
	queue := stack_queue_set.InitQueue()
	values := getArrayOfRandomNonNegativeUniqueValues(size)
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

func testSetAdd(t *testing.T) {
	set := stack_queue_set.NewSet()
	size := 20
	insertedValues := getArrayOfRandomNonNegativeUniqueValues(size)
	for _, value := range insertedValues {
		set.Add(value)
	}

	if !arraysEqual(getIntArrayFromAnyArray(set.Values()), insertedValues) {
		t.Fatalf("Set values not equal to expected inserted values. got: %v, expected: %v", set.Values(), insertedValues)
	}
	//test with duplicate values
	setLen := len(set.Values())
	for _, value := range insertedValues {
		insertedValues = append(insertedValues, value)
		set.Add(value)
		if len(set.Values()) != setLen {
			t.Fatalf("Length of set not equal to expected length. got: %v, expected: %v", len(set.Values()), setLen)
		}
	}

	if len(set.Values()) != setLen {
		t.Fatalf("Length of set not equal to expected length. got: %v, expected: %v", len(set.Values()), setLen)
	}

	if arraysEqual(getIntArrayFromAnyArray(set.Values()), insertedValues) {
		t.Fatalf("Set values are equal to list with duplicate values")
	}

}

func testSetRemove(t *testing.T) {
	set := stack_queue_set.NewSet()
	size := 10
	insertedValues := getArrayOfRandomNonNegativeUniqueValues(size)
	for _, value := range insertedValues {
		set.Add(value)
	}

	if !arraysEqual(getIntArrayFromAnyArray(set.Values()), insertedValues) {
		t.Fatalf("Set values not equal to expected inserted values. got: %v, expected: %v", set.Values(), insertedValues)
	}

	var notInsertedValues []int
	for _, value := range insertedValues {
		notInsertedValues = append(notInsertedValues, value+len(insertedValues))
	}
	setSize := len(set.Values())
	for i := 0; i < size; i++ {
		set.Remove(notInsertedValues[0])
		notInsertedValues = notInsertedValues[1:]
		if len(set.Values()) != setSize {
			t.Fatalf("Remove expected length. got: %v, expected: %v", len(set.Values()), setSize)
		}
	}

	for i := 0; i < size; i++ {
		set.Remove(insertedValues[0])
		insertedValues = insertedValues[1:]
		if !arraysEqual(getIntArrayFromAnyArray(set.Values()), insertedValues) {
			t.Fatalf("Set values not equal to expected inserted values. got: %v, expected: %v", set.Values(), insertedValues)
		}
	}
	if len(set.Values()) != 0 {
		t.Fatalf("Remove expected length. got: %v, expected: %v", len(set.Values()), 0)
	}

}

func testSetContains(t *testing.T) {
	set := stack_queue_set.NewSet()
	size := 10
	insertedValues := getArrayOfRandomNonNegativeUniqueValues(size)
	for _, value := range insertedValues {
		set.Add(value)
	}

	if !arraysEqual(getIntArrayFromAnyArray(set.Values()), insertedValues) {
		t.Fatalf("Set values not equal to expected inserted values. got: %v, expected: %v", set.Values(), insertedValues)
	}

	var notInsertedValues []int
	for _, value := range insertedValues {
		notInsertedValues = append(notInsertedValues, value+len(insertedValues))
	}

	for i := 0; i < size; i++ {
		containsValue := set.Contains(notInsertedValues[0])
		notInsertedValues = notInsertedValues[1:]
		if containsValue {
			t.Fatalf("Contains returned true when it shouldn't have")
		}
	}

	for i := 0; i < size; i++ {
		containsValue := set.Contains(insertedValues[0])
		insertedValues = insertedValues[1:]
		if !containsValue {
			t.Fatalf("Contains returned false when it shouldn't have")
		}
	}

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

func TestSet(t *testing.T) {
	testSetAdd(t)
	testSetRemove(t)
	testSetContains(t)
}
