package tests

import (
	"github.com/ElwinCabrera/go-containers/trees"
	"testing"
)

func testInsertAndPreInAndPostOrderValues(t *testing.T) {
	btree := trees.NewBinaryTree[int]()

	insert_values := []int{20, 10, 30, 9, 11, 21, 31}
	treeTestInsertHelper(t, btree, &insert_values)
	treeOrder1 := btree.PreOrderValues()
	treeOrder2 := btree.InOrderValues()
	treeOrder3 := btree.PostOrderValues()
	verifyTreeNodeOrder(t, []int{20, 10, 9, 11, 30, 21, 31}, &treeOrder1)
	verifyTreeNodeOrder(t, []int{9, 10, 11, 20, 21, 30, 31}, &treeOrder2)
	verifyTreeNodeOrder(t, []int{9, 11, 10, 21, 31, 30, 20}, &treeOrder3)

	btree.Clear()
	insert_values_sorted_asc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	treeTestInsertHelper(t, btree, &insert_values_sorted_asc)
	treeOrder1 = btree.PreOrderValues()
	treeOrder2 = btree.InOrderValues()
	treeOrder3 = btree.PostOrderValues()
	verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder1)
	verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder2)
	verifyTreeNodeOrder(t, sortArrayDesc(insert_values_sorted_asc), &treeOrder3)

	btree.Clear()
	insert_values_sorted_desc := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	treeTestInsertHelper(t, btree, &insert_values_sorted_desc)
	treeOrder1 = btree.PreOrderValues()
	treeOrder2 = btree.InOrderValues()
	treeOrder3 = btree.PostOrderValues()
	verifyTreeNodeOrder(t, insert_values_sorted_desc, &treeOrder1)
	sortedAsc := sortArrayAsc(insert_values_sorted_desc)
	verifyTreeNodeOrder(t, sortedAsc, &treeOrder2)
	verifyTreeNodeOrder(t, sortedAsc, &treeOrder3)
	//
	btree.Clear()
	insert_values = getArrayOfRandomUniqueValues(30)
	treeTestInsertHelper(t, btree, &insert_values)
	treeOrder1 = btree.InOrderValues()
	verifyTreeNodeOrder(t, sortArrayAsc(insert_values), &treeOrder1)

}

func testSizeMinMaxAndClear(t *testing.T) {
	//create tree and sanity check
	btree := trees.NewBinaryTree[int]()
	insert_values := getArrayOfRandomUniqueValues(20)
	treeTestInsertHelper(t, btree, &insert_values)
	verifyTree(t, btree, insert_values)

	insert_values = sortArrayAsc(insert_values)
	minVal := btree.Min()
	maxVal := btree.Max()

	if minVal != insert_values[0] {
		t.Fatalf("Min did not return correct values. got %v, expected %v", minVal, insert_values[0])
	}
	if maxVal != insert_values[len(insert_values)-1] {
		t.Fatalf("Max did not return correct values. got %v, expected %v", maxVal, insert_values[len(insert_values)-1])
	}

	btree.Clear()
	if btree.Size() != 0 || btree.Root() != nil {
		t.Fatalf("Clear did not clear all values")
	}

}

func testRemove(t *testing.T) {
	//create tree and sanity check
	btree := trees.NewBinaryTree[int]()
	insert_values := getArrayOfRandomUniqueValues(20)
	nodes := treeTestInsertHelper(t, btree, &insert_values)
	verifyTree(t, btree, insert_values)

	//Testing RemoveNode
	for _, node := range nodes {
		btree.RemoveNode(node)
		insert_values = insert_values[1:]
		verifyTree(t, btree, insert_values)
	}
	if btree.Size() != 0 {
		t.Fatalf("tree failed to remove all nodes. tree has size %d, expected 0", btree.Size())
	}

	//testing RemoveValue
	btree.Clear()
	insert_values = getArrayOfRandomUniqueValues(20)
	nodes = treeTestInsertHelper(t, btree, &insert_values)
	for _, node := range nodes {
		btree.RemoveValue(node.Value)
		insert_values = insert_values[1:]
		verifyTree(t, btree, insert_values)
	}
	if btree.Size() != 0 {
		t.Fatalf("tree failed to remove all nodes. tree has size %d, expected 0", btree.Size())
	}

	//Testing  RemoveValue with a tree that has duplicate values
	btree.Clear()
	insert_values = getArrayOfRandomUniqueValues(10)

	for _, val := range insert_values {
		insert_values = append(insert_values, val)
	}
	nodes = treeTestInsertHelper(t, btree, &insert_values)
	for len(insert_values) > 0 {
		btree.RemoveValue(insert_values[0])
		treeSize := btree.Size()
		if treeSize != len(insert_values)-2 {
			t.Fatalf("tree size incorrect after removing duplicate values.  got %d, expected %d", treeSize, len(insert_values)-2)
		}
		values := insert_values[1 : len(insert_values)/2]
		for i := (len(insert_values) / 2) + 1; i < len(insert_values); i++ {
			values = append(values, insert_values[i])
		}
		insert_values = values
		verifyTree(t, btree, insert_values)
	}

	//Testing RemoveValue with a tree that only has duplicate values
	btree.Clear()
	insert_values = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	nodes = treeTestInsertHelper(t, btree, &insert_values)
	for _, val := range insert_values {
		btree.RemoveValue(val)
		treeSize := btree.Size()
		if treeSize != 0 {
			t.Fatalf("Tree failed to remove values. tree has size %d, expected 0", treeSize)
		}
		insert_values = []int{}
		verifyTree(t, btree, insert_values)
	}

}

func testContainsFindAndFindFirst(t *testing.T) {
	//create tree, insert values, and do small sanity check
	btree := trees.NewBinaryTree[int]()
	insert_values := getArrayOfRandomUniqueValues(10)
	treeTestInsertHelper(t, btree, &insert_values)
	verifyTree(t, btree, insert_values)

	//test contains and find with values that are definitely not in the tree
	var non_inserted_vals []int
	for _, val := range insert_values {
		non_inserted_vals = append(non_inserted_vals, val+len(insert_values))
	}
	for i := 0; i < len(insert_values); i++ {
		shouldBeTrue := btree.Contains(insert_values[i])
		shouldBeFalse := btree.Contains(non_inserted_vals[i])
		allFound := btree.Find(insert_values[i])
		shouldReturnEmpty := btree.Find(non_inserted_vals[i])

		if !shouldBeTrue {
			t.Fatalf("Contains did NOT return true for insert_values %v, expected true", insert_values[i])
		}
		if shouldBeFalse {
			t.Fatalf("Contains did NOT return false for non_inserted_value %v, expected false", non_inserted_vals[i])
		}
		if allFound == nil || len(allFound) == 0 || len(allFound) > 1 || allFound[0].Value != insert_values[i] {
			t.Fatalf("Find did not return all expected values. got %v for find value  %v", allFound, insert_values[i])
		}
		if len(shouldReturnEmpty) > 0 {
			t.Fatalf("Found values when expected nothing to be returned. got %v for find value  %v", shouldReturnEmpty, non_inserted_vals[i])
		}
	}

	//Test Find and FindFirst with a tree having duplicate values
	btree.Clear()
	insert_values = getArrayOfRandomUniqueValues(10)

	for _, val := range insert_values {
		insert_values = append(insert_values, val)
	}
	/*nodes := */ treeTestInsertHelper(t, btree, &insert_values)
	verifyTree(t, btree, insert_values)
	for i := 0; i < len(insert_values)/2; i++ {
		allFound := btree.Find(insert_values[i])
		if len(allFound) < 2 || len(allFound) > 2 || allFound[0].Value != insert_values[i] || allFound[0].Value != allFound[1].Value {
			t.Fatalf("Find did not return all expected values. got %v for find value  %v", allFound, insert_values[i])
		}
	}

	for i := 0; i < len(insert_values)/2; i++ {
		allFound := btree.Find(insert_values[i])
		firstFound := btree.FindFirst(insert_values[i])

		if len(allFound) < 2 || len(allFound) > 2 || allFound[0].Value != insert_values[i] || allFound[0].Value != allFound[1].Value {
			t.Fatalf("Find did not return all expected values. got %v for find value  %v", allFound, insert_values[i])
		}
		if firstFound == nil || firstFound.Value != insert_values[i] {
			if firstFound != nil {
				t.Fatalf("FindFirst should have retuned %v, got %v", insert_values[i], firstFound.Value)
			} else {
				t.Fatalf("FindFirst should have retuned %v, got nil", insert_values[i])
			}
		}
		nextNodeInAllFoundIdx := 1
		if allFound[0] != firstFound {
			nextNodeInAllFoundIdx = 0
		}
		btree.RemoveNode(firstFound)
		firstFound = btree.FindFirst(insert_values[i])
		if firstFound != allFound[nextNodeInAllFoundIdx] {
			t.Fatalf("FirstFound called a second time did not find the duplicate value")
		}
	}

	//Test Find and FindFirst with a tree of all same values
	btree.Clear()
	insert_values = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	treeTestInsertHelper(t, btree, &insert_values)
	verifyTree(t, btree, insert_values)
	allFound := btree.Find(insert_values[0])
	if len(allFound) != len(insert_values) {
		t.Fatalf("Find should have gotten all %v values but got %v for find value  %v", len(insert_values), allFound, insert_values[0])
	}

}

func TestBinaryTree(t *testing.T) {
	testInsertAndPreInAndPostOrderValues(t)
	testSizeMinMaxAndClear(t)
	testRemove(t)
	testContainsFindAndFindFirst(t)
}
