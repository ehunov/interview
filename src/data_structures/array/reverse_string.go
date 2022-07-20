// https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1183/

package array

func reverseString(s []byte) []byte {
	j := len(s) - 1

	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}

	return s
}
