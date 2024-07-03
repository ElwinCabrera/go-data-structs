package tests

import (
	"slices"
	"testing"

	"github.com/ElwinCabrera/go-containers/list"
)

func test_InsertEnd(t *testing.T, l list.List, size int) {
	l.Clear()
	var insert_values []int
	idx := 0

	if l.Len() != 0 {
		t.Fatalf("Test_InsertEnd(..) initial list length is not zero %v %v \n", len(insert_values), l.Len())

	}
	insert_values = append(insert_values, 1)
	l.InsertEnd(insert_values[idx])
	idx++

	if !verifyListContents(l, insert_values) {
		t.Fatalf("Test_InsertEnd(..) test failed to verify list contents\n")
	}
	insert_values = append(insert_values, 2)
	l.InsertEnd(insert_values[idx])
	idx++
	if !verifyListContents(l, insert_values) {
		t.Fatalf("Test_InsertEnd(..) test failed to verify list contents\n")
	}

	l.Clear()
	_, insert_values = insertListRandomData(l, size)

	if !verifyListContents(l, insert_values) {
		t.Fatalf("InsertEnd(..) test failed to verify random list contents\n")
	}

}

func test_InsertFront(t *testing.T, l list.List, size int) {
	l.Clear()
	_, insert_values := insertListRandomData(l, size)

	l.Clear()
	for i := len(insert_values) - 1; i >= 0; i-- {
		l.InsertFront(insert_values[i])
	}
	if !verifyListContents(l, insert_values) {
		t.Fatalf("Test_InsertFront(..) test failed to verify random list contents\n")
	}
}

func test_InsertBefore(t *testing.T, l list.List, size int) {
	l.Clear()
	_, insert_values := insertListRandomData(l, size)

	if !verifyListContents(l, insert_values) {
		t.Fatalf("Test_InsertBefore(..) test failed to verify random list contents\n")
	}

	i := 0
	it := list.InitListIterator(l)
	for it := it.Begin(); it.Get() != nil; it = it.Next() {
		new_val := (i + 1) * -1
		l.InsertBefore(new_val, it.Get())
		insert_values = slices.Insert(insert_values, i, new_val)
		if !verifyListContents(l, insert_values) {
			t.Fatalf("Test_InsertBefore(..) test failed to verify random list contents\n")
		}
		i += 2
	}

}

func test_InsertAfter(t *testing.T, l list.List, size int) {
	l.Clear()
	_, insert_values := insertListRandomData(l, size)

	if !verifyListContents(l, insert_values) {
		t.Fatalf("Test_InsertAfter(..) test failed to verify random list contents\n")
	}

	i := 0
	it := list.InitListIterator(l)
	for it := it.Begin(); it.Get() != nil; it = it.Next().Next() {
		new_val := (i + 1) * -1
		l.InsertAfter(new_val, it.Get())

		if i+1 < len(insert_values) {
			insert_values = slices.Insert(insert_values, i+1, new_val)
		} else {
			insert_values = append(insert_values, new_val)
		}

		if !verifyListContents(l, insert_values) {
			t.Fatalf("Test_InsertAfter(..) test failed to verify random list contents\n")
		}
		i += 2
		//it = it.Next()

	}
}

func test_PopFront(t *testing.T, l list.List, size int) {
	l.Clear()
	_, insert_values := insertListRandomData(l, size)

	for i := 0; i < size; i++ {
		node := l.PopFront()
		if isNodeInList(l, node) || node == l.Head() {
			t.Fatalf("Test_PopFront(..) test failed to remove node from list\n")
		}
		insert_values = insert_values[1:]
		if !verifyListContents(l, insert_values) {
			t.Fatalf("Test_PopFront(..) test failed to verify random list contents\n")
		}
	}
}

func test_PopBack(t *testing.T, l list.List, size int) {
	l.Clear()
	_, insert_values := insertListRandomData(l, size)

	for i := 0; i < size; i++ {
		node := l.PopBack()
		if isNodeInList(l, node) {
			t.Fatalf("Test_PopBack(..) test failed to remove node from list\n")
		}
		insert_values = insert_values[:len(insert_values)-1]
		if !verifyListContents(l, insert_values) {
			t.Fatalf("Test_PopBack(..) test failed to verify random list contents\n")
		}
	}
}

func test_Remove(t *testing.T, l list.List, size int) {
	l.Clear()
	insertListRandomData(l, size)

	if !removeRandomListNodes(l, size) {
		t.Fatalf("Test_Remove(..) test failed\n")
	}
}

func Test_SinglyLinkedList(t *testing.T) {
	ll := list.InitSinglyLinkedList()
	size := 10
	test_InsertEnd(t, ll, size)
	test_InsertFront(t, ll, size)
	test_InsertBefore(t, ll, size)
	test_InsertAfter(t, ll, size)
	test_PopFront(t, ll, size)
	test_PopBack(t, ll, size)
	test_Remove(t, ll, size)
}

func Test_DoublyLinkedList(t *testing.T) {
	ll := list.InitDoublyLinkedList()
	size := 10
	test_InsertEnd(t, ll, size)
	test_InsertFront(t, ll, size)
	test_InsertBefore(t, ll, size)
	test_InsertAfter(t, ll, size)
	test_PopFront(t, ll, size)
	test_PopBack(t, ll, size)
	test_Remove(t, ll, size)
}
