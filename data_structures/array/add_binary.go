package array

import (
	"strings"
)

func addBinary(a string, b string) string {
	aLastIndex, bLastIndex := len(a)-1, len(b)-1
	if bLastIndex > aLastIndex {
		return addBinary(b, a)
	}

	bitSum := 0
	result := make([]string, aLastIndex+1)
	for i := 0; i <= aLastIndex; i++ {
		curIndex := aLastIndex - i
		bitSum += int(a[curIndex] - '0')
		if bLastIndex >= i {
			bitSum += int(b[bLastIndex-i] - '0')
		}

		switch bitSum {
		case 0:
			result[curIndex] = "0"
		case 1:
			result[curIndex] = "1"
			bitSum = 0
		case 2:
			result[curIndex] = "0"
			bitSum = 1
		case 3:
			result[curIndex] = "1"
			bitSum = 1
		}
	}

	if bitSum > 0 {
		result = append([]string{"1"}, result...)
	}

	return strings.Join(result, "")
}
