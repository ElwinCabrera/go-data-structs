package bitstructs

import (
	"fmt"
)

var BYTE_LENGTH int = 8

type BitSequence struct {
	data           *[]uint8
	numBits        *int
	bytesAllocated *int
	nextBitIdx     *int
	nextByteIdx    *int
}

func NewBitSequence(numBits int) BitSequence {

	bytesAllocated := numBits / BYTE_LENGTH
	if numBits%BYTE_LENGTH != 0 { // if there is a remainder we need to allocate an extra byte to accommodate the extra number of bit-structs
		bytesAllocated += 1
	}

	data := make([]uint8, bytesAllocated)
	nextBitIdx := 0
	nextByteIdx := 0
	return BitSequence{data: &data, numBits: &numBits, bytesAllocated: &bytesAllocated, nextBitIdx: &nextBitIdx, nextByteIdx: &nextByteIdx}
}

func NewBitSequenceFromByteArray(data *[]uint8, bitLen int) BitSequence {

	bitSeq := NewBitSequence(bitLen)

	bitSeq.SetNextByteStart(0)
	for i := 0; i < len(*data); i++ {
		bitSeq.SetNextByte((*data)[i])
	}

	return bitSeq
}

func NewDynamicBitSequence() BitSequence {
	nextBitIdx := 0
	nextByteIdx := 0

	return BitSequence{data: new([]byte), numBits: new(int), bytesAllocated: new(int), nextBitIdx: &nextBitIdx, nextByteIdx: &nextByteIdx}
}

func (bseq BitSequence) AppendBitFront(num byte) {
	if num < 0 || num > 1 {
		panic("num must be between 0 and 1 but got " + fmt.Sprint(num))
	}
	bseq.addBitsToSequence(1)
	bseq.shiftRightByOne()
	bseq.SetBit(0, NumToBool(num))
}

func (bseq BitSequence) AppendBitEnd(num byte) {
	if num < 0 || num > 1 {
		panic("num must be between 0 and 1 but got " + fmt.Sprint(num))
	}
	bseq.addBitsToSequence(1)
	bseq.SetBit(*bseq.numBits-1, NumToBool(num))
}

func (bseq BitSequence) shiftRightByOne() {
	for i := *bseq.numBits - 1; i >= 0; i-- {
		if i-1 >= 0 {
			prevBit := bseq.GetBit(i - 1)
			bseq.SetBit(i, prevBit)
		}
	}
	bseq.SetBit(0, false)
}

func (bseq BitSequence) shiftLeftByOne() {
	for i := 0; i < *bseq.numBits; i++ {
		if i+1 < *bseq.numBits {
			nextBit := bseq.GetBit(i + 1)
			bseq.SetBit(i, nextBit)
		}
	}
	bseq.SetBit(*bseq.numBits-1, false)
}

func (bseq BitSequence) SetBit(bitIdx int, set bool) {
	if bitIdx < 0 || bitIdx >= *bseq.numBits {
		panic("set_bit idx out of bounds")
	}
	byteIdx := bitIdx / BYTE_LENGTH
	bitIdxInByte := bitIdx % BYTE_LENGTH

	if set {
		(*bseq.data)[byteIdx] |= 1 << bitIdxInByte
	} else {
		(*bseq.data)[byteIdx] &= ^(1 << bitIdxInByte)
	}
}

func (bseq BitSequence) SetBitsFromNum(startBitIdx int, number uint64) {

	if startBitIdx < 0 || startBitIdx >= *bseq.numBits {
		panic("BitSequence.setBitsFromNum(..): idx out of bounds")
	}
	for number != 0 && startBitIdx < *bseq.numBits {
		bseq.SetBit(startBitIdx, NumToBool(uint8(number&0x1)))
		number >>= 1
		startBitIdx++
	}
}

func (bseq BitSequence) SetNextBitStart(bitIdx int) {
	if bitIdx < 0 || bitIdx >= *bseq.numBits {
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
	if byteIdx < 0 || byteIdx >= *bseq.bytesAllocated {
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
	if bitIdx < 0 || bitIdx >= *bseq.numBits {
		panic("BitSequence.getBit() idx out of bounds")
		//return false
	}
	byteIdx := bitIdx / BYTE_LENGTH
	bitIdxInByte := bitIdx % BYTE_LENGTH

	zeroOrOne := ((*bseq.data)[byteIdx] >> bitIdxInByte) & 0x1

	return NumToBool(zeroOrOne)

}

func (bseq BitSequence) GetByte(byteIdx int) uint8 {
	if byteIdx < 0 || byteIdx >= *bseq.bytesAllocated {
		panic("BitSequence.getByte() idx out of bounds")
	}
	return (*bseq.data)[byteIdx]

}

func (bseq BitSequence) GetXBytes(numBytes int) uint64 {

	if numBytes > *bseq.bytesAllocated {
		numBytes = *bseq.bytesAllocated
	}
	res := uint64(0)

	for numBytes != 0 {
		res <<= BYTE_LENGTH
		res |= uint64(bseq.GetByte(numBytes - 1))
		numBytes--
	}
	return res
}

func (bseq BitSequence) GetNextBitStart(bitIdx int) {
	if bitIdx < 0 || bitIdx >= *bseq.numBits {
		panic("BitSequence.getNextBitStart: idx out of bounds\n")
	}
	*bseq.nextBitIdx = bitIdx
}

func (bseq BitSequence) GetNextBit() bool {
	if *bseq.nextBitIdx == -1 {
		panic("BitSequence.getNextBitStart() never called")
	}
	if *bseq.nextBitIdx >= *bseq.numBits {
		panic("BitSequence.GetNextBit() called one too many times. Index Out of bounds")
	}
	res := bseq.GetBit(*bseq.nextBitIdx)
	*bseq.nextBitIdx++
	return res
}

func (bseq BitSequence) GetNextByte() uint8 {
	if *bseq.nextByteIdx == -1 {
		panic("BitSequence.getNextByteStart() never called")
	}
	if *bseq.nextByteIdx >= *bseq.bytesAllocated {
		panic("BitSequence.getNextByte() called one too many times. Index Out of bounds")
	}
	res := bseq.GetByte(*bseq.nextByteIdx)
	*bseq.nextByteIdx++
	return res
}

func (bseq BitSequence) ReverseBytes() {
	if *bseq.bytesAllocated < 2 {
		return
	}
	start := 0
	end := len(*bseq.data) - 1
	for start < end {
		hold := (*bseq.data)[start]
		(*bseq.data)[start] = (*bseq.data)[end]
		(*bseq.data)[end] = hold
		start++
		end--
	}
}

func (bseq BitSequence) addBitsToSequence(numBitsToAdd int) {
	*bseq.numBits += numBitsToAdd
	bytesNeeded := *bseq.numBits / BYTE_LENGTH
	if *bseq.numBits%BYTE_LENGTH != 0 {
		bytesNeeded += 1
	}

	for bytesNeeded != *bseq.bytesAllocated || bseq.data == nil || *bseq.bytesAllocated == 0 {
		*bseq.data = append(*bseq.data, 0x00)
		*bseq.bytesAllocated = len(*bseq.data)
	}
}

func (bseq BitSequence) ExpandNumOfBitsToUseRemainingBitsInLastByte() {
	numRemainingBits := 0
	if *bseq.numBits%BYTE_LENGTH != 0 {
		numRemainingBits = BYTE_LENGTH - (*bseq.numBits % BYTE_LENGTH)
	}
	*bseq.numBits += numRemainingBits
}

// the first byte in the slice is the least significant byte and the last byte in array is the most significant
func (bseq BitSequence) GetBitSeq() []uint8 {
	return *bseq.data
}

// the first byte in the slice is the most significant byte and the last byte in array is the least significant
func (bseq BitSequence) GetReversedBitSeq() []uint8 {
	var reversedBytes []byte
	for i := *bseq.bytesAllocated - 1; i >= 0; i-- {
		reversedBytes = append(reversedBytes, (*bseq.data)[i])
	}
	return reversedBytes
}

func (bseq BitSequence) GetNumBits() int {
	return *bseq.numBits
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

func (bseq BitSequence) String() string {
	result := ""

	//for i :=  0; i < bseq.bytesAllocated; i++ {
	//	result += fmt.Sprintf("%02X", (*bseq.data)[i])
	//}
	for i := *bseq.numBits - 1; i >= 0; i-- {
		result += fmt.Sprintf("%v", BoolToInt(bseq.GetBit(i)))
	}
	return result
}

func NumToBool(n uint8) bool {
	if n == 0 {
		return false
	}
	return true
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
