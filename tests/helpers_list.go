package tests

import (
	"fmt"
	"github.com/ElwinCabrera/go-data-structs/list"
	"math/rand"
)

func insertListRandomData(l list.List, size int) (list.List, []int) {
	values := getArrayOfRandomNonNegativeUniqueValues(size)
	for _, v := range values {
		l.InsertEnd(v, 0)
	}
	return l, values
}

func removeRandomListNodes(l list.List, num_times int) bool {

	seen_map := make(map[int]bool)
	var nodes []*list.Node

	list_len := l.Len()
	if l.Len() == 0 {
		return true
	}

	it := list.InitListIterator(l)

	for it := it.Begin(); it.Get() != nil; it = it.Next() {
		nodes = append(nodes, it.Get())
	}

	for i := 0; i < num_times; i++ {
		ran_num := rand.Intn(list_len)
		_, seen := seen_map[ran_num]
		for seen {
			ran_num = rand.Intn(10)
			_, seen = seen_map[ran_num]
		}
		seen_map[ran_num] = true

		fmt.Printf("Removed %v ", nodes[ran_num].Value)
		ok := l.Remove(nodes[ran_num])
		if isNodeInList(l, nodes[ran_num]) || !ok {
			return false
		}
		fmt.Print(l)
		fmt.Print("\n")

	}

	return true
}

func verifyListContents(l list.List, values []int) bool {
	it := list.InitListIterator(l)
	if len(values) != l.Len() {
		return false
	}
	if l.Head() == nil && len(values) == 0 && l.Len() == 0 {
		return true
	}
	idx := 0
	for it := it.Begin(); it.Get() != nil; it = it.Next() {
		if it.Get().Value != values[idx] {
			return false
		}
		idx++
	}
	return true
}

func isNodeInList(l list.List, node *list.Node) bool {
	if node == nil {
		return false
	}
	it := list.InitListIterator(l)
	idx := 0
	for it := it.Begin(); it.Get() != nil; it = it.Next() {
		if it.Get() == node {
			return true
		}
		idx++
	}
	return false
}
