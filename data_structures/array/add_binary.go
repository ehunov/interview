package array

import (
	"math"
	"strings"
)

func addBinary(a string, b string) string {
	aLastIndex, bLastIndex := len(a)-1, len(b)-1
	maxIndex := int(math.Max(float64(aLastIndex), float64(bLastIndex)))
	result := make([]string, maxIndex+1)
	appendBits := make([]string, maxIndex+1)
	for i := 0; i <= maxIndex; i++ {
		curIndex := maxIndex - i
		appendBit := appendBits[curIndex]
		if appendBit == "" {
			appendBit = "0"
		}
		bBit := "0"
		aBit := "0"
		if bLastIndex >= i {
			bBit = string(rune(b[bLastIndex-i]))
		}
		if aLastIndex >= i {
			aBit = string(rune(a[aLastIndex-i]))
		}

		if aBit == "1" && bBit == "1" {
			result[curIndex] = appendBit
			if curIndex != 0 {
				appendBits[curIndex-1] = "1"
			} else {
				result = append([]string{"1"}, result...)
			}
		} else if (aBit == "1" || bBit == "1") && appendBit == "1" {
			result[curIndex] = "0"
			if curIndex != 0 {
				appendBits[curIndex-1] = "1"
			} else {
				result = append([]string{"1"}, result...)
			}
		} else if aBit == "0" && bBit == "0" {
			result[curIndex] = appendBit
		} else {
			result[curIndex] = "1"
		}
	}

	return strings.Join(result, "")
}
