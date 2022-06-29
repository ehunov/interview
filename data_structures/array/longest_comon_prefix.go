package array

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	firstWord := strs[0]

	for symbolIndex, symbol := range firstWord {
		for _, word := range strs {
			if symbolIndex > len(word)-1 || rune(word[symbolIndex]) != symbol {
				return firstWord[0:symbolIndex]
			}
		}
	}

	return firstWord
}
