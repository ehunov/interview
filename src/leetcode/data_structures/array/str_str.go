// https://leetcode.com/explore/learn/card/array-and-string/203/introduction-to-string/1161/

package array

func strStr(haystack string, needle string) int {
	haystackLen, needleLen := len(haystack), len(needle)

	for i := 0; i < haystackLen-needleLen+1; i++ {
		if needle == haystack[i:i+needleLen] {
			return i
		}
	}

	return -1
}
