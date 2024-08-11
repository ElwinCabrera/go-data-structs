package trees

import (
	"fmt"
	bitstructs "github.com/ElwinCabrera/go-data-structs/bit-structs"
	stack_queue_set "github.com/ElwinCabrera/go-data-structs/stack-queue-set"
)

type HuffmanTree[T comparable] struct {
	root         *TreeNode[T]
	frequencyMap map[T]uint64

	huffmanCodes map[T]bitstructs.BitSequence
}

func NewHuffmanTreeFromFrequencyMap[T comparable](frequencyMap map[T]uint64) *HuffmanTree[T] {
	if frequencyMap == nil || len(frequencyMap) == 0 {
		rootNode := &TreeNode[T]{IgnoreValue: true, Weight: 0}
		return &HuffmanTree[T]{root: rootNode, frequencyMap: frequencyMap}
	}

	ht := &HuffmanTree[T]{frequencyMap: frequencyMap}
	pq := ht.getSortedListFromFrequencyMap()
	ht.buildTreeFromSortedList(pq)
	return ht
}

func NewHuffmanTreeFromHuffmanCodes[T comparable](huffmanCodes map[T]bitstructs.BitSequence) *HuffmanTree[T] {
	root := &TreeNode[T]{IgnoreValue: true, Weight: -1}
	ht := &HuffmanTree[T]{root: root, huffmanCodes: huffmanCodes}

	for data, bitseq := range huffmanCodes {
		ht.recreateOriginalTreeFromHuffmanCodes(ht.root, data, bitseq, uint64(bitseq.GetNumBits()-1))
	}
	return ht
}

func (ht *HuffmanTree[T]) getSortedListFromFrequencyMap() *stack_queue_set.PriorityQueue {
	//ll := list.InitDoublyLinkedList()
	priorityQ := stack_queue_set.NewPriorityQueue(true)
	for val, weight := range ht.frequencyMap {
		weightInFloat := float64(weight) // can cause issues putting as the mantissa is less than 64 bits
		treeNode := &TreeNode[T]{Value: val, Weight: weightInFloat}
		//ll.InsertSortedDescBasedOnNodeWeight(treeNode, weight)
		priorityQ.Push(treeNode, weightInFloat)
	}
	return priorityQ
}

func (ht *HuffmanTree[T]) buildTreeFromSortedList(pq *stack_queue_set.PriorityQueue) {

	done := false
	for !done {
		leftTree := pq.Dequeue().(*TreeNode[T])
		rightTree := pq.Dequeue()
		parentTreeNode := &TreeNode[T]{IgnoreValue: true, left: leftTree}

		combinedWeight := leftTree.Weight

		if rightTree != nil {
			combinedWeight += rightTree.(*TreeNode[T]).Weight
			parentTreeNode.right = rightTree.(*TreeNode[T])
		}
		if pq.IsEmpty() {
			done = true
		}
		parentTreeNode.Weight = combinedWeight

		pq.Push(parentTreeNode, combinedWeight)
	}

	ht.root = pq.Dequeue().(*TreeNode[T])
}

func (ht *HuffmanTree[T]) GetHuffmanCodes() map[T]bitstructs.BitSequence {
	if ht.huffmanCodes == nil {
		ht.huffmanCodes = make(map[T]bitstructs.BitSequence)
		ht.generateHuffmanCodes(ht.root, 0, 0, false)
	}
	return ht.huffmanCodes
}

func (ht *HuffmanTree[T]) generateHuffmanCodes(current *TreeNode[T], currentCode, depth int, set bool) {
	if current == nil {
		return
	}
	currentCode <<= 1
	if set {
		currentCode |= 1
	}

	if !current.IgnoreValue {
		if _, ok := ht.huffmanCodes[current.Value]; ok {
			panic(fmt.Sprintf("We visited same data node twice. current code: %v, node value: %v", currentCode, current.Value))
		}

		bs := bitstructs.NewBitSequence(depth)
		bs.SetBitsFromNum(0, uint64(currentCode))
		ht.huffmanCodes[current.Value] = bs
	}

	ht.generateHuffmanCodes(current.left, currentCode, depth+1, false)
	ht.generateHuffmanCodes(current.right, currentCode, depth+1, true)
}

func (ht *HuffmanTree[T]) recreateOriginalTreeFromHuffmanCodes(current *TreeNode[T], data T, bitSequence bitstructs.BitSequence, currBitIdx uint64) {
	isBitSet := bitSequence.GetBit(int(currBitIdx))
	if currBitIdx == 0 {
		leaf := &TreeNode[T]{Value: data, Weight: float64(bitSequence.GetXBytes(8))}
		if !isBitSet && current.left == nil {
			current.left = leaf
		}
		if isBitSet && current.right == nil {
			current.right = leaf
		}
		return
	}

	if !isBitSet {
		if current.left == nil {
			current.left = &TreeNode[T]{IgnoreValue: true, Weight: -1}
		}
		ht.recreateOriginalTreeFromHuffmanCodes(current.left, data, bitSequence, currBitIdx-1)
	} else {
		if current.right == nil {
			current.right = &TreeNode[T]{IgnoreValue: true, Weight: -1}
		}
		ht.recreateOriginalTreeFromHuffmanCodes(current.right, data, bitSequence, currBitIdx-1)
	}

}

func (ht *HuffmanTree[T]) DecodeBitSequence(bitSequence *bitstructs.BitSequence) *[]T {
	//bitSequence.GetNextBitStart(0)
	//dataLen := ht.findDataLenFromBitSequence(ht.root, 0, bitSequence)

	var data *[]T = new([]T)
	bitSequence.GetNextBitStart(0)
	for bitSequence.GetNextBitIdx() < bitSequence.GetNumBits() {
		ht.decodeHuffmanCodeHelper(ht.root, data, bitSequence)
	}

	return data
}

func (ht *HuffmanTree[T]) decodeHuffmanCodeHelper(current *TreeNode[T], data *[]T, bitSequence *bitstructs.BitSequence) {
	if current == nil {
		return
	}
	if !current.IgnoreValue {
		*data = append(*data, current.Value)
		//current = ht.root	// not a bug but a limitation if done this way. While this does work if the bit sequence length gets is large enough then the stack could grow uncontrollably
		return
	}
	if bitSequence.GetNextBitIdx() >= bitSequence.GetNumBits() {
		return
	}
	isBitSet := bitSequence.GetNextBit()
	if !isBitSet {
		ht.decodeHuffmanCodeHelper(current.left, data, bitSequence)
	} else {
		ht.decodeHuffmanCodeHelper(current.right, data, bitSequence)
	}

}

func (ht *HuffmanTree[T]) findDataLenFromBitSequence(current *TreeNode[T], count int, bitSequence *bitstructs.BitSequence) int {
	if current == nil || bitSequence.GetNextBitIdx() > bitSequence.GetNumBits() {
		return count
	}
	if !current.IgnoreValue {
		count++
		current = ht.root
	}
	isBitSet := bitSequence.GetNextBit()
	if !isBitSet {
		count = ht.findDataLenFromBitSequence(current.left, count, bitSequence)
	} else {
		count = ht.findDataLenFromBitSequence(current.right, count, bitSequence)
	}
	return count
}

func (ht *HuffmanTree[T]) InOrderNodes() []*TreeNode[T] {
	var result []*TreeNode[T]
	treeInOrderValueHelper(ht.root, &result)
	return result
}

func (ht *HuffmanTree[T]) String() string {
	nodes := ht.InOrderNodes()
	res := "["
	for _, node := range nodes {
		if node.IgnoreValue {
			res += fmt.Sprintf("(Val: PlaceHolder, Weight: %v) ", node.Weight)
		} else {
			res += fmt.Sprintf("(Val: %v, Weight: %v) ", node.Value, node.Weight)
		}
	}
	res += "]"
	return res
}
