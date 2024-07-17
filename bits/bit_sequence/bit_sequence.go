package bit_sequence

import (
	"fmt"
	utils "github.com/ElwinCabrera/go-data-structs/bits"
)

type BitSequence struct {
	data           *[]uint8
	numBits        int
	bytesAllocated int
	nextBitIdx     *int
	nextByteIdx    *int
}

func NewBitSequence(numBits int) BitSequence {

	bytesAllocated := numBits / utils.BYTE_LENGTH
	if numBits%utils.BYTE_LENGTH != 0 { // if there is a remainder we need to allocate an extra byte to accommodate the extra number of bits
		bytesAllocated += 1
	}

	data := make([]uint8, bytesAllocated)
	nextBitIdx := 0
	nextByteIdx := 0
	return BitSequence{data: &data, numBits: numBits, bytesAllocated: bytesAllocated, nextBitIdx: &nextBitIdx, nextByteIdx: &nextByteIdx}
}

func (bseq BitSequence) SetBit(bitIdx int, set bool) {
	if bitIdx < 0 || bitIdx >= bseq.numBits {
		panic("set_bit idx out of bounds")
	}
	byteIdx := bitIdx / utils.BYTE_LENGTH
	bitIdxInByte := bitIdx % utils.BYTE_LENGTH

	if set {
		(*bseq.data)[byteIdx] |= 1 << bitIdxInByte
	} else {
		(*bseq.data)[byteIdx] &= ^(1 << bitIdxInByte)
	}
}

func (bseq BitSequence) SetBitsFromNum(startBitIdx int, number uint64) {

	if startBitIdx < 0 || startBitIdx >= bseq.numBits {
		panic("BitSequence.setBitsFromNum(..): idx out of bounds")
	}
	for number != 0 && startBitIdx < bseq.numBits {
		bseq.SetBit(startBitIdx, utils.NumToBool(uint8(number&0x1)))
		number >>= 1
	}
}

func (bseq BitSequence) SetNextBitStart(bitIdx int) {
	if bitIdx < 0 || bitIdx >= bseq.numBits {
		panic("BitSequence.setNextBitStart(..): idx out of bounds")
	}
	*bseq.nextBitIdx = bitIdx

}

func (bseq BitSequence) SetNextBit(set bool) {
	if *bseq.nextBitIdx == -1 {
		panic("BitSequence.setNextBit(..) never called")
	}
	bseq.SetBit(*bseq.nextBitIdx, set)
	*bseq.nextBitIdx++
}

func (bseq BitSequence) SetNextByteStart(byteIdx int) {
	if byteIdx < 0 || byteIdx >= bseq.bytesAllocated {
		panic("BitSequence.setNextByteStart(): idx out of bounds")
	}
	*bseq.nextByteIdx = byteIdx
}

func (bseq BitSequence) SetNextByte(byte uint8) {
	if *bseq.nextByteIdx == -1 {
		panic("BitSequence.setNextByte() never called\n")

	}
	(*bseq.data)[*bseq.nextByteIdx] = byte
	*bseq.nextByteIdx++
}

func (bseq BitSequence) GetBit(bitIdx int) bool {
	if bitIdx < 0 || bitIdx >= bseq.numBits {
		panic("BitSequence.getBit() idx out of bounds")
		//return false
	}
	byteIdx := bitIdx / utils.BYTE_LENGTH
	bitIdxInByte := bitIdx % utils.BYTE_LENGTH

	zeroOrOne := ((*bseq.data)[byteIdx] >> bitIdxInByte) & 0x1

	return utils.NumToBool(zeroOrOne)

}

func (bseq BitSequence) GetByte(byteIdx int) uint8 {
	if byteIdx < 0 || byteIdx >= bseq.bytesAllocated {
		panic("BitSequence.getByte() idx out of bounds")
	}
	return (*bseq.data)[byteIdx]

}

func (bseq BitSequence) GetXBytes(numBytes int) uint64 {

	if numBytes > bseq.bytesAllocated {
		numBytes = bseq.bytesAllocated
	}
	res := uint64(0)

	for numBytes != 0 {
		res <<= utils.BYTE_LENGTH
		res |= uint64(bseq.GetByte(numBytes - 1))
		numBytes--
	}
	return res
}

func (bseq BitSequence) GetNextBitStart(bitIdx int) {
	if bitIdx < 0 {
		fmt.Printf("BitSequence.getNextBitStart: idx out of bounds\n")
		bitIdx = 0
	}
	if bitIdx >= bseq.numBits {
		fmt.Printf("BitSequence.getNextBitStart: idx out of bounds\n")
		bitIdx = bseq.numBits - 1
	}
	*bseq.nextBitIdx = bitIdx
}

func (bseq BitSequence) GetNextBit() bool {
	if *bseq.nextBitIdx == -1 {
		panic("BitSequence.getNextBit() never called")
	}
	res := bseq.GetBit(*bseq.nextBitIdx)
	*bseq.nextBitIdx++
	return res
}

//func (bseq BitSequence) GetNextByte() uint8 {
//	return 0
//}

func (bseq BitSequence) GetBitSeq() []uint8 {
	return *bseq.data
}

func (bseq BitSequence) GetNumBits() int {
	return bseq.numBits
}
func (bseq BitSequence) GetBytesAllocated() int {
	return len(*bseq.data)
}
func (bseq BitSequence) GetNextBitIdx() int {
	return *bseq.nextBitIdx
}
func (bseq BitSequence) GetNextByteIdx() int {
	return *bseq.nextByteIdx
}

func (bseq BitSequence) Clear() {
	for i, _ := range *bseq.data {
		(*bseq.data)[i] = 0
	}
}
