package tests

import (
	"fmt"
	"github.com/ElwinCabrera/go-data-structs/trees"
	"testing"
)

func TestCreateHuffmanTreeFromFrequencyMap(t *testing.T) {
	s3 := "A_DEAD_DAD_CEDED_A_BAD_BABE"
	freqMap := make(map[rune]uint64)
	for _, c := range s3 {
		freqMap[c] += 1
	}
	ht := trees.NewHuffmanTreeFromFrequencyMap(freqMap)
	huffmanCodes := ht.GetHuffmanCodes()
	if len(huffmanCodes) != len(freqMap) {
		t.Fatalf("Huffman code is not the same length as frequency map")
	}
	for elem, b := range huffmanCodes {
		fmt.Printf("(elem: %v code: %v) ", string(elem), b)
		if elem == '_' {
			if b.GetNumBits() != 2 || b.GetBit(0) != false || b.GetBit(1) != false {
				t.Fatalf("incorrect bit sequence for %v. Got %v expected 00", elem, b)
			}
		}
		if elem == 'A' {
			if b.GetNumBits() != 2 || b.GetBit(0) != true || b.GetBit(1) != false {
				t.Fatalf("incorrect bit sequence for %v. Got %v expected 01", elem, b)
			}
		}
		if elem == 'D' {
			if b.GetNumBits() != 2 || b.GetBit(0) != false || b.GetBit(1) != true {
				t.Fatalf("incorrect bit sequence for %v. Got %v expected 10", elem, b)
			}
		}
		if elem == 'E' {
			if b.GetNumBits() != 3 || b.GetBit(0) != true || b.GetBit(1) != true || b.GetBit(2) != true {
				t.Fatalf("incorrect bit sequence for %v. Got %v expected 111", elem, b)
			}
		}
		if elem == 'C' {
			if b.GetNumBits() != 4 || b.GetBit(0) != false || b.GetBit(1) != false || b.GetBit(2) != true || b.GetBit(3) != true {
				t.Fatalf("incorrect bit sequence for %v. Got %v expected 1100", elem, b)
			}
		}
		if elem == 'B' {
			if b.GetNumBits() != 4 || b.GetBit(0) != true || b.GetBit(1) != false || b.GetBit(2) != true || b.GetBit(3) != true {
				t.Fatalf("incorrect bit sequence for %v. Got %v expected 1101", elem, b)
			}
		}

	}

}

func TestCreateHuffmanTreeFromHuffmanCodes(t *testing.T) {
	s3 := "A_DEAD_DAD_CEDED_A_BAD_BABE"
	freqMap := make(map[rune]uint64)
	for _, c := range s3 {
		freqMap[c] += 1
	}
	ht := trees.NewHuffmanTreeFromFrequencyMap(freqMap)
	huffmanCodes := ht.GetHuffmanCodes()

	if len(huffmanCodes) != len(freqMap) {
		t.Fatalf("Huffman code is not the same length as frequency map")
	}
	recreated_ht := trees.NewHuffmanTreeFromHuffmanCodes(huffmanCodes)

	verifyHuffmanTreesEqual(t, ht, recreated_ht)
	fmt.Printf("Original Tree %v\n", ht)
	fmt.Printf("Recreated Tree %v\n", recreated_ht)

}

func verifyHuffmanTreesEqual[T comparable](t *testing.T, hTree1, hTree2 *trees.HuffmanTree[T]) {
	treeNodes1 := hTree1.InOrderNodes()
	treeNodes2 := hTree2.InOrderNodes()

	if len(treeNodes1) != len(treeNodes2) {
		t.Fatalf("Huffman tree nodes are not the same length tree 1 length %v, tree 2 lenght %v", len(treeNodes1), len(treeNodes2))
	}
	for i, node := range treeNodes1 {
		//if treeNodes2[i].Weight != node.Weight {
		//	t.Fatalf("Tree weight not the same")
		//}
		if treeNodes2[i].Value != node.Value {
			t.Fatalf("Tree value not the same")
		}
		if treeNodes2[i].IgnoreValue != node.IgnoreValue {
			t.Fatalf("Tree ignore value not the same")
		}
	}
}
