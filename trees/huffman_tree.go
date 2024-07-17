package trees

import (
	"github.com/ElwinCabrera/go-data-structs/bits/bit_sequence"
	"github.com/ElwinCabrera/go-data-structs/list"
)

type HuffmanTree[T comparable] struct {
	root         *TreeNode[T]
	frequencyMap map[T]int
	huffmanCodes map[T]bit_sequence.BitSequence
}

func NewTreeFromFrequencyMap[T comparable](frequencyMap map[T]int) *HuffmanTree[T] {
	ht := &HuffmanTree[T]{}
	ll := ht.getSortedListFromFrequencyMap()
	ht.buildTreeFromSortedList(ll)
	return ht
}

func NewTreeFromHuffmanCodes[T comparable](huffmanCodes map[T]bit_sequence.BitSequence) *HuffmanTree[T] {
	root := &TreeNode[T]{IgnoreValue: true, Weight: -1}
	ht := &HuffmanTree[T]{root: root, huffmanCodes: huffmanCodes}

	for data, bitseq := range huffmanCodes {
		ht.recreateOriginalTreeFromHuffmanCodes(ht.root, data, bitseq, uint64(bitseq.GetNumBits()-1))
	}
	return ht
}

func (ht *HuffmanTree[T]) getSortedListFromFrequencyMap() list.List {
	ll := list.InitDoublyLinkedList()
	for val, weight := range ht.frequencyMap {
		treeNode := &TreeNode[T]{Value: val, Weight: weight}
		ll.InsertSortedDescBasedOnNodeWeight(treeNode, weight)
	}
	return ll
}

func (ht *HuffmanTree[T]) buildTreeFromSortedList(ll list.List) {

	done := false
	for !done {
		listNode1 := ll.PopFront()
		listNode2 := ll.PopFront()

		var leftTree *TreeNode[T] = nil
		var rightTree *TreeNode[T] = nil
		combinedWeight := 0
		if listNode1 != nil {
			combinedWeight += listNode1.Weight
			leftTree = listNode1.Value.(*TreeNode[T])
		}
		if listNode2 != nil {
			combinedWeight += listNode2.Weight
			rightTree = listNode2.Value.(*TreeNode[T])
		}
		if ll.Head() == nil {
			done = true
		}
		treeNode := &TreeNode[T]{IgnoreValue: true, Weight: combinedWeight, left: leftTree, right: rightTree}
		ll.InsertSortedDescBasedOnNodeWeight(treeNode, combinedWeight)
	}

	ht.root = ll.PopFront().Value.(*TreeNode[T])
}

func (ht *HuffmanTree[T]) GetHuffmanCodes() map[T]bit_sequence.BitSequence {
	if ht.huffmanCodes == nil {
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

	if _, ok := ht.huffmanCodes[current.Value]; !ok {
		if !current.IgnoreValue {
			bs := bit_sequence.NewBitSequence(depth)
			bs.SetBitsFromNum(0, uint64(currentCode))
			ht.huffmanCodes[current.Value] = bs
		}
	}
	ht.generateHuffmanCodes(current.left, currentCode, depth+1, false)
	ht.generateHuffmanCodes(current.right, currentCode, depth+1, true)
}

func (ht *HuffmanTree[T]) recreateOriginalTreeFromHuffmanCodes(current *TreeNode[T], data T, bitSequence bit_sequence.BitSequence, currBitIdx uint64) {
	isBitSet := bitSequence.GetBit(int(currBitIdx))
	if currBitIdx == 0 {
		leaf := &TreeNode[T]{Value: data, Weight: int(bitSequence.GetXBytes(8))}
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

func (ht *HuffmanTree[T]) decodeHuffmanCode(bitSequence bit_sequence.BitSequence) *[]T {
	//bitSequence.GetNextBitStart(0)
	//dataLen := ht.findDataLenFromBitSequence(ht.root, 0, bitSequence)

	var data *[]T
	bitSequence.GetNextBitStart(0)
	ht.decodeHuffmanCodeHelper(ht.root, data, &bitSequence)
	return data
}

func (ht *HuffmanTree[T]) decodeHuffmanCodeHelper(current *TreeNode[T], data *[]T, bitSequence *bit_sequence.BitSequence) {
	if current == nil || bitSequence.GetNextBitIdx() > bitSequence.GetNumBits() {
		return
	}
	if !current.IgnoreValue {
		*data = append(*data, current.Value)
		current = ht.root
	}
	isBitSet := bitSequence.GetNextBit()
	if !isBitSet {
		ht.decodeHuffmanCodeHelper(current.left, data, bitSequence)
	} else {
		ht.decodeHuffmanCodeHelper(current.right, data, bitSequence)
	}

}

func (ht *HuffmanTree[T]) findDataLenFromBitSequence(current *TreeNode[T], count int, bitSequence *bit_sequence.BitSequence) int {
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
