package tests

import (
	"testing"

	"github.com/ElwinCabrera/go-data-structs/trees"
)

func treeTestInsertHelper[T trees.TreeNodeValue](t *testing.T, btree trees.Tree[T], insertValues *[]T) []*trees.TreeNode[T] {
	var nodes []*trees.TreeNode[T]
	for i, value := range *insertValues {
		newNode := btree.Insert(value)
		if i == 0 && btree.Root().Value != value {
			t.Fatalf("Insert: expected %v, got %v", value, btree.Root().Value)
		}
		nodes = append(nodes, newNode)
	}

	if len(*insertValues) != btree.Size() {
		t.Fatalf("Insert: incorrect tree size. Expected %v, got %v", btree.Size(), len(*insertValues))
	}

	return nodes

}

func verifyTree[T trees.TreeNodeValue](t *testing.T, btree trees.Tree[T], insertValues []T) {
	inOrder := btree.InOrderValues()
	if len(inOrder) != len(insertValues) {
		t.Fatalf("wrong tree size expected %v, got %v", len(insertValues), len(inOrder))
	}
	if len(inOrder) == 0 && len(insertValues) == 0 {
		return
	}
	sortedValuesAsc := sortArrayAsc(insertValues)

	for i, value := range sortedValuesAsc {
		if inOrder[i].Value != value {
			t.Fatalf("Tree verification failed. expected %v, got %v", value, inOrder[i].Value)
		}
	}

}

func verifyTreeNodeOrder[T trees.TreeNodeValue](t *testing.T, expectedOrder []T, actualTreeOrder *[]*trees.TreeNode[T]) {
	if len(expectedOrder) != len(*actualTreeOrder) {
		t.Fatalf("Insert: tree did not insert all values correctly. Expected tree size %v, got %v", len(expectedOrder), len(*actualTreeOrder))
	}

	for i, node := range *actualTreeOrder {
		if (expectedOrder)[i] != node.Value {
			t.Fatalf("Insert: values were inserted in wrong order")
		}
	}
}

func findMaxNodeValueInList(nodes []*trees.TreeNode[int]) int {
	if len(nodes) == 0 {
		return -1
	}
	maxNodeVal := nodes[0].Value
	for _, node := range nodes {
		if node.Value > maxNodeVal {
			maxNodeVal = node.Value
		}
	}
	return maxNodeVal
}

func findMinNodeValueInList(nodes []*trees.TreeNode[int]) int {
	if len(nodes) == 0 {
		return -1
	}
	minNodeVal := nodes[0].Value
	for _, node := range nodes {
		if node.Value < minNodeVal {
			minNodeVal = node.Value
		}
	}
	return minNodeVal
}
