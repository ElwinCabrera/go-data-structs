package tests

import (
	utils "github.com/ElwinCabrera/go-data-structs/bits"
	"github.com/ElwinCabrera/go-data-structs/bits/bit_sequence"
	"math/rand"
	"testing"
)

func TestBitSequence(t *testing.T) {

	//generate a random 64bit number
	OriginalRandNum := rand.Uint64()

	//set the bits of that random number to our bit sequence
	bs := bit_sequence.NewBitSequence(64)
	mutableRandomNum := OriginalRandNum
	for i := 0; i < 64; i++ {

		setFlag := utils.NumToBool(uint8(mutableRandomNum & 0x1))
		bs.SetBit(i, setFlag)
		if bs.GetBit(i) != setFlag {
			t.Fatalf("SetBit(%d) failed", i)
		}
		mutableRandomNum >>= 1
	}

	bitSeqNum := bs.GetXBytes(8)
	if bitSeqNum != OriginalRandNum {
		t.Fatalf("bit sequence number mismatch: %d vs %d", bitSeqNum, OriginalRandNum)
	}

	//Test SetNextBit
	bs.Clear()
	OriginalRandNum = rand.Uint64()

	//set the bits of that random number to our bit sequence
	bs.SetNextBitStart(0)
	mutableRandomNum = OriginalRandNum
	for i := 0; i < 64; i++ {
		setFlag := utils.NumToBool(uint8(mutableRandomNum & 0x1))
		bs.SetNextBit(setFlag)
		if bs.GetBit(i) != setFlag {
			t.Fatalf("SetBit(%d) failed", i)
		}
		mutableRandomNum >>= 1
	}

	bitSeqNum = bs.GetXBytes(8)
	if bitSeqNum != OriginalRandNum {
		t.Fatalf("bit sequence number mismatch: %d vs %d", bitSeqNum, OriginalRandNum)
	}

	//Testing SetNextByte
	bs.Clear()
	OriginalRandNum = rand.Uint64()

	//set the bits of that random number to our bit sequence
	bs.SetNextByteStart(0)
	mutableRandomNum = OriginalRandNum
	for i := 0; i < utils.BYTE_LENGTH; i++ {
		b := uint8(mutableRandomNum & 0xFF)
		bs.SetNextByte(b)
		if bs.GetByte(i) != b {
			t.Fatalf("SetByte(%d) failed", i)
		}
		mutableRandomNum >>= utils.BYTE_LENGTH
	}

	bitSeqNum = bs.GetXBytes(8)
	if bitSeqNum != OriginalRandNum {
		t.Fatalf("bit sequence number mismatch: %d vs %d", bitSeqNum, OriginalRandNum)
	}

}
