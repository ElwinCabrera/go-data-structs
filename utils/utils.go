package utils

import (
	"math"
	"unicode"
)

// little endian
func NumToHexString(num uint64) string {
	if num == 0 {
		return "0"
	}
	hex := ""
	for num > 0 {
		remainder := uint8(num % 16)
		c := '0' + remainder
		if remainder >= 10 {
			remainder -= 10
			c = 'A' + remainder
		}
		num /= 16
		hex = string(c) + hex
	}

	return hex
}

// reads the hex str as little endian
func HexStringToInt(hexStr string) uint64 {
	num := uint64(0)
	idx := len(hexStr) - 1
	for idx >= 0 {
		c := hexStr[idx]
		exponent := float64(len(hexStr) - idx - 1)
		if unicode.IsDigit(rune(c)) {
			num += uint64(c-'0') * uint64(math.Pow(16, exponent))
		} else {
			num += uint64((c-'A')+10) * uint64(math.Pow(16, exponent))
		}
		idx--
	}
	return num
}
