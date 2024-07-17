package tests

import (
	"github.com/ElwinCabrera/go-data-structs/trees"
	"testing"
)

func testInsertAndPreInAndPostOrderValues(t *testing.T, tree trees.Tree[int]) {
	tree.Clear()
	insert_values := []int{20, 10, 30, 9, 11, 21, 31}
	treeTestInsertHelper(t, tree, &insert_values)
	treeOrder1 := tree.PreOrderValues()
	treeOrder2 := tree.InOrderValues()
	treeOrder3 := tree.PostOrderValues()
	verifyTreeNodeOrder(t, []int{20, 10, 9, 11, 30, 21, 31}, &treeOrder1)
	verifyTreeNodeOrder(t, []int{9, 10, 11, 20, 21, 30, 31}, &treeOrder2)
	verifyTreeNodeOrder(t, []int{9, 11, 10, 21, 31, 30, 20}, &treeOrder3)

	tree.Clear()
	insert_values_sorted_asc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	treeTestInsertHelper(t, tree, &insert_values_sorted_asc)
	treeOrder1 = tree.PreOrderValues()
	treeOrder2 = tree.InOrderValues()
	treeOrder3 = tree.PostOrderValues()
	verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder1)
	verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder2)
	verifyTreeNodeOrder(t, sortArrayDesc(insert_values_sorted_asc), &treeOrder3)

	tree.Clear()
	insert_values_sorted_desc := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	treeTestInsertHelper(t, tree, &insert_values_sorted_desc)
	treeOrder1 = tree.PreOrderValues()
	treeOrder2 = tree.InOrderValues()
	treeOrder3 = tree.PostOrderValues()
	verifyTreeNodeOrder(t, insert_values_sorted_desc, &treeOrder1)
	sortedAsc := sortArrayAsc(insert_values_sorted_desc)
	verifyTreeNodeOrder(t, sortedAsc, &treeOrder2)
	verifyTreeNodeOrder(t, sortedAsc, &treeOrder3)
	//
	tree.Clear()
	insert_values = getArrayOfRandomNonNegativeUniqueValues(30)
	treeTestInsertHelper(t, tree, &insert_values)
	treeOrder1 = tree.InOrderValues()
	verifyTreeNodeOrder(t, sortArrayAsc(insert_values), &treeOrder1)

}

func testHeapInsertAndPreInAndPostOrderValues(t *testing.T, tree trees.MinMaxTree[int]) {
	//var tree trees.MinMaxTree[int]
	//tree = trees.NewMinHeap[int]()
	//tree = trees.NewMaxHeap[int]()
	tree.Clear()

	insert_values := []int{20, 10, 30, 9, 11, 21, 31}
	treeTestInsertHelper(t, tree, &insert_values)
	treeOrder1 := tree.PreOrderValues()
	treeOrder2 := tree.InOrderValues()
	treeOrder3 := tree.PostOrderValues()
	if tree.GetHeapType() == "min" {
		verifyTreeNodeOrder(t, []int{9, 20, 10, 11, 30, 21, 31}, &treeOrder1)
		verifyTreeNodeOrder(t, []int{9, 10, 11, 20, 21, 30, 31}, &treeOrder2)
		verifyTreeNodeOrder(t, []int{11, 10, 21, 31, 30, 20, 9}, &treeOrder3)
	} else {
		//verifyTreeNodeOrder(t, []int{31, 9, 20, 10, 11, 30, 21}, &treeOrder1)
		verifyTreeNodeOrder(t, []int{9, 10, 11, 20, 21, 30, 31}, &treeOrder2)
		//verifyTreeNodeOrder(t, []int{9, 11, 10, 21, 30, 20, 31}, &treeOrder3)
	}

	tree.Clear()
	insert_values_sorted_asc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	treeTestInsertHelper(t, tree, &insert_values_sorted_asc)
	treeOrder1 = tree.PreOrderValues()
	treeOrder2 = tree.InOrderValues()
	treeOrder3 = tree.PostOrderValues()
	if tree.GetHeapType() == "min" {
		verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder1)
		verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder2)
		verifyTreeNodeOrder(t, sortArrayDesc(insert_values_sorted_asc), &treeOrder3)
	} else {
		verifyTreeNodeOrder(t, []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}, &treeOrder1)
		verifyTreeNodeOrder(t, insert_values_sorted_asc, &treeOrder2)
		verifyTreeNodeOrder(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}, &treeOrder3)
	}

	tree.Clear()
	insert_values_sorted_desc := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	treeTestInsertHelper(t, tree, &insert_values_sorted_desc)
	treeOrder1 = tree.PreOrderValues()
	treeOrder2 = tree.InOrderValues()
	treeOrder3 = tree.PostOrderValues()
	if tree.GetHeapType() == "min" {
		verifyTreeNodeOrder(t, []int{1, 10, 9, 8, 7, 6, 5, 4, 3, 2}, &treeOrder1)
		verifyTreeNodeOrder(t, sortArrayAsc(insert_values_sorted_desc), &treeOrder2)
		verifyTreeNodeOrder(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, &treeOrder3)
	} else {
		verifyTreeNodeOrder(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, &treeOrder1)
		sortedAsc := sortArrayAsc(insert_values_sorted_desc)
		verifyTreeNodeOrder(t, sortedAsc, &treeOrder2)
		verifyTreeNodeOrder(t, sortedAsc, &treeOrder3)
	}

	//
	tree.Clear()
	insert_values = getArrayOfRandomNonNegativeUniqueValues(30)
	treeTestInsertHelper(t, tree, &insert_values)
	treeOrder1 = tree.InOrderValues()
	verifyTreeNodeOrder(t, sortArrayAsc(insert_values), &treeOrder1)
}

func testSizeMinMaxAndClear(t *testing.T, tree trees.Tree[int]) {
	tree.Clear()
	//create tree and sanity check
	insert_values := getArrayOfRandomNonNegativeUniqueValues(20)
	treeTestInsertHelper(t, tree, &insert_values)
	verifyTree(t, tree, insert_values)

	insert_values = sortArrayAsc(insert_values)
	minVal := tree.Min()
	maxVal := tree.Max()

	if minVal != insert_values[0] {
		t.Fatalf("Min did not return correct values. got %v, expected %v", minVal, insert_values[0])
	}
	if maxVal != insert_values[len(insert_values)-1] {
		t.Fatalf("Max did not return correct values. got %v, expected %v", maxVal, insert_values[len(insert_values)-1])
	}

	tree.Clear()
	if tree.Size() != 0 || tree.Root() != nil {
		t.Fatalf("Clear did not clear all values")
	}

}

func testRemove(t *testing.T, tree trees.Tree[int]) {
	tree.Clear()
	//create tree and sanity check
	//tree := trees.NewBinaryTree[int]()
	insert_values := getArrayOfRandomNonNegativeUniqueValues(20)
	nodes := treeTestInsertHelper(t, tree, &insert_values)
	verifyTree(t, tree, insert_values)

	//Testing RemoveNode
	for _, node := range nodes {
		tree.RemoveNode(node)
		insert_values = insert_values[1:]
		verifyTree(t, tree, insert_values)
	}
	if tree.Size() != 0 {
		t.Fatalf("tree failed to remove all nodes. tree has size %d, expected 0", tree.Size())
	}

	//testing RemoveValue
	tree.Clear()
	insert_values = getArrayOfRandomNonNegativeUniqueValues(20)
	nodes = treeTestInsertHelper(t, tree, &insert_values)
	for _, node := range nodes {
		tree.RemoveValue(node.Value)
		insert_values = insert_values[1:]
		verifyTree(t, tree, insert_values)
	}
	if tree.Size() != 0 {
		t.Fatalf("tree failed to remove all nodes. tree has size %d, expected 0", tree.Size())
	}

	//Testing  RemoveValue with a tree that has duplicate values
	tree.Clear()
	insert_values = getArrayOfRandomNonNegativeUniqueValues(10)

	for _, val := range insert_values {
		insert_values = append(insert_values, val)
	}
	nodes = treeTestInsertHelper(t, tree, &insert_values)
	for len(insert_values) > 0 {
		tree.RemoveValue(insert_values[0])
		treeSize := tree.Size()
		if treeSize != len(insert_values)-2 {
			t.Fatalf("tree size incorrect after removing duplicate values.  got %d, expected %d", treeSize, len(insert_values)-2)
		}
		values := insert_values[1 : len(insert_values)/2]
		for i := (len(insert_values) / 2) + 1; i < len(insert_values); i++ {
			values = append(values, insert_values[i])
		}
		insert_values = values
		verifyTree(t, tree, insert_values)
	}

	//Testing RemoveValue with a tree that only has duplicate values
	tree.Clear()
	insert_values = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	nodes = treeTestInsertHelper(t, tree, &insert_values)
	for _, val := range insert_values {
		tree.RemoveValue(val)
		treeSize := tree.Size()
		if treeSize != 0 {
			t.Fatalf("Tree failed to remove values. tree has size %d, expected 0", treeSize)
		}
		insert_values = []int{}
		verifyTree(t, tree, insert_values)
	}

}

func testHeapPop(t *testing.T, tree trees.MinMaxTree[int]) {

	tree.Clear()

	insertValues := getArrayOfRandomNonNegativeUniqueValues(10)
	treeTestInsertHelper(t, tree, &insertValues)
	treeInOrderVals := tree.InOrderValues()
	verifyTreeNodeOrder(t, sortArrayAsc(insertValues), &treeInOrderVals)

	sortedInsertedVals := sortArrayAsc(insertValues)

	for i := 0; i < len(insertValues); i++ {

		tree.Pop()
		if tree.GetHeapType() == "min" {
			sortedInsertedVals = sortedInsertedVals[1:]
		} else {
			sortedInsertedVals = sortedInsertedVals[:len(sortedInsertedVals)-1]
		}

		treeOrder := tree.InOrderValues()
		verifyTreeNodeOrder(t, sortedInsertedVals, &treeOrder)

	}

	if tree.Size() != 0 {
		t.Fatalf("HeapPop expected tree size of 0 but got %v", tree.Size())
	}

	tree.Clear()
	insertValues = getArrayOfRandomNonNegativeUniqueValues(30)
	treeTestInsertHelper(t, tree, &insertValues)
	treeInOrderVals = tree.InOrderValues()
	verifyTreeNodeOrder(t, sortArrayAsc(insertValues), &treeInOrderVals)

	if tree.GetHeapType() == "min" {
		for i, node := range treeInOrderVals {

			poppedNode := tree.Pop()
			if poppedNode == nil {
				t.Fatalf("HeapPop expected a non-nil node")
			}

			if poppedNode.Value != node.Value {
				t.Fatalf("HeapPop did not return correct value. got %v, expected %v", poppedNode.Value, node.Value)
			}

			if i+1 >= len(treeInOrderVals) {
				if tree.Root() != nil {
					t.Fatalf("Heap tree should have popped all values but tree is still not empty. Got a root of %v, expected nil", tree.Root().Value)
				}
			} else {
				if tree.Root().Value != treeInOrderVals[i+1].Value {
					t.Fatalf("Heap tree new root should be the next maximum value of %v but got %v", treeInOrderVals[i+1].Value, tree.Root().Value)
				}
			}
		}
	} else {
		for i := len(treeInOrderVals) - 1; i >= 0; i-- {
			node := treeInOrderVals[i]
			poppedNode := tree.Pop()
			if poppedNode == nil {
				t.Fatalf("HeapPop expected a non-nil node")
			}

			if poppedNode.Value != node.Value {
				t.Fatalf("HeapPop did not return correct value. got %v, expected %v", poppedNode.Value, node.Value)
			}

			if i-1 < 0 {
				if tree.Root() != nil {
					t.Fatalf("Heap tree should have popped all values but tree is still not empty. Got a root of %v, expected nil", tree.Root().Value)
				}
			} else {
				if tree.Root().Value != treeInOrderVals[i-1].Value {
					t.Fatalf("Heap tree new root should be the next maximum value of %v but got %v", treeInOrderVals[i-1].Value, tree.Root().Value)
				}
			}
		}
	}

}

func testContainsFindAndFindFirst(t *testing.T, tree trees.Tree[int]) {
	tree.Clear()
	//create tree, insert values, and do small sanity check

	insert_values := getArrayOfRandomNonNegativeUniqueValues(10)
	treeTestInsertHelper(t, tree, &insert_values)
	verifyTree(t, tree, insert_values)

	//test contains and find with values that are definitely not in the tree
	var non_inserted_vals []int
	for _, val := range insert_values {
		non_inserted_vals = append(non_inserted_vals, val+len(insert_values))
	}
	for i := 0; i < len(insert_values); i++ {
		shouldBeTrue := tree.Contains(insert_values[i])
		shouldBeFalse := tree.Contains(non_inserted_vals[i])
		allFound := tree.Find(insert_values[i])
		shouldReturnEmpty := tree.Find(non_inserted_vals[i])

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
	tree.Clear()
	insert_values = getArrayOfRandomNonNegativeUniqueValues(10)

	for _, val := range insert_values {
		insert_values = append(insert_values, val)
	}
	/*nodes := */ treeTestInsertHelper(t, tree, &insert_values)
	verifyTree(t, tree, insert_values)
	for i := 0; i < len(insert_values)/2; i++ {
		allFound := tree.Find(insert_values[i])
		if len(allFound) < 2 || len(allFound) > 2 || allFound[0].Value != insert_values[i] || allFound[0].Value != allFound[1].Value {
			t.Fatalf("Find did not return all expected values. got %v for find value  %v", allFound, insert_values[i])
		}
	}

	for i := 0; i < len(insert_values)/2; i++ {
		allFound := tree.Find(insert_values[i])
		firstFound := tree.FindFirst(insert_values[i])

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
		tree.RemoveNode(firstFound)
		firstFound = tree.FindFirst(insert_values[i])
		if firstFound != allFound[nextNodeInAllFoundIdx] {
			t.Fatalf("FirstFound called a second time did not find the duplicate value")
		}
	}

	//Test Find and FindFirst with a tree of all same values
	tree.Clear()
	insert_values = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	treeTestInsertHelper(t, tree, &insert_values)
	verifyTree(t, tree, insert_values)
	allFound := tree.Find(insert_values[0])
	if len(allFound) != len(insert_values) {
		t.Fatalf("Find should have gotten all %v values but got %v for find value  %v", len(insert_values), allFound, insert_values[0])
	}

}

func runCommonTreeTests(t *testing.T, tree trees.Tree[int]) {
	testSizeMinMaxAndClear(t, tree)
	testRemove(t, tree)
	testContainsFindAndFindFirst(t, tree)
}

func TestBinaryTree(t *testing.T) {
	btree := trees.NewBinaryTree[int]()
	testInsertAndPreInAndPostOrderValues(t, btree)
	runCommonTreeTests(t, btree)

}

func TestMinAndMaxHeap(t *testing.T) {

	minHeap := trees.NewMinHeap[int]()
	testHeapInsertAndPreInAndPostOrderValues(t, minHeap)
	testHeapPop(t, minHeap)
	runCommonTreeTests(t, minHeap)

	maxHeap := trees.NewMaxHeap[int]()
	testHeapInsertAndPreInAndPostOrderValues(t, maxHeap)
	testHeapPop(t, maxHeap)
	runCommonTreeTests(t, maxHeap)
}
