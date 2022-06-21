package array

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1

	for i := lastIndex; i > -1; i-- {
		if i == lastIndex {
			digits[i]++
		}
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}

			digits[i-1]++
		} else {
			break
		}
	}

	return digits
}
