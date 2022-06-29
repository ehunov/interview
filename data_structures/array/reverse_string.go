package array

func reverseString(s []byte) []byte {
	j := len(s) - 1

	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}

	return s
}
