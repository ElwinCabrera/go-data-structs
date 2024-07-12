package tests

import (
	"testing"

	"github.com/ElwinCabrera/go-containers/trees"
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

//func getHeapExpectedInOrderValuesFromInsertedValues[T trees.TreeNodeValue](insertedValues []T, heapType string) []T {
//	if insertedValues == nil || len(insertedValues) == 0 {
//		return insertedValues
//	}
//
//	expectedInOrderValues := sortArrayAsc(insertedValues)
//
//	minVal, maxVal := insertedValues[0], insertedValues[0]
//
//	for _, value := range insertedValues {
//		if value < minVal {
//			minVal = value
//		}
//		if value > maxVal {
//			maxVal = value
//		}
//	}
//
//	removeIdxForMaxVal := 0
//	removeIdxForMinVal := 0
//	for i, value := range expectedInOrderValues {
//		if value == maxVal {
//			removeIdxForMaxVal = i
//		}
//		if value == minVal {
//			removeIdxForMinVal = i
//		}
//	}
//
//	if heapType == "min" {
//
//	} else if heapType == "max" {
//
//	}
//
//	return expectedInOrderValues
//}
