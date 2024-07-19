package tests

import (
	"fmt"
	bitstructs "github.com/ElwinCabrera/go-data-structs/bit-structs"
	"github.com/ElwinCabrera/go-data-structs/utils"
	"math/rand"
	"testing"
)

func TestBitSequence(t *testing.T) {

	//generate a random 64bit number
	OriginalRandNum := rand.Uint64()

	//set the bit-structs of that random number to our bit sequence
	bs := bitstructs.NewBitSequence(64)
	mutableRandomNum := OriginalRandNum
	for i := 0; i < 64; i++ {

		setFlag := numToBool(uint8(mutableRandomNum & 0x1))
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

	//set the bit-structs of that random number to our bit sequence
	bs.SetNextBitStart(0)
	mutableRandomNum = OriginalRandNum
	for i := 0; i < 64; i++ {
		setFlag := numToBool(uint8(mutableRandomNum & 0x1))
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

	//set the bit-structs of that random number to our bit sequence
	bs.SetNextByteStart(0)
	mutableRandomNum = OriginalRandNum
	for i := 0; i < bitstructs.BYTE_LENGTH; i++ {
		b := uint8(mutableRandomNum & 0xFF)
		bs.SetNextByte(b)
		if bs.GetByte(i) != b {
			t.Fatalf("SetByte(%d) failed", i)
		}
		mutableRandomNum >>= bitstructs.BYTE_LENGTH
	}

	bitSeqNum = bs.GetXBytes(8)
	if bitSeqNum != OriginalRandNum {
		t.Fatalf("bit sequence number mismatch: %d vs %d", bitSeqNum, OriginalRandNum)
	}

}

func TestNumToHexString(t *testing.T) {

	end := int(^uint16(0))

	for i := 0; i <= end; i++ {
		expectedHex := fmt.Sprintf("%X", i)
		hexStr := utils.NumToHexString(uint(i))
		//fmt.Println(hexStr)
		if expectedHex != hexStr {
			t.Fatalf("NumToHexString(uint): expected %s, actual %s", expectedHex, hexStr)
		}
	}

}

func TestHexToNum(t *testing.T) {

	end := int(^uint16(0))
	for expectedNum := 0; expectedNum <= end; expectedNum++ {
		hexStr := fmt.Sprintf("%X", expectedNum)
		num := utils.HexStringToInt(hexStr)

		if num != expectedNum {
			t.Fatalf("HexStringToInt(string): expected %d, actual %d", expectedNum, num)
		}
	}
}

func numToBool(n uint8) bool {
	if n == 0 {
		return false
	}
	return true
}
