package array

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		elementsCount := i + 1
		result[i] = make([]int, elementsCount)
		lastJIndex := elementsCount - 1

		for j := 0; j < elementsCount; j++ {
			if j == lastJIndex || j == 0 {
				result[i][j] = 1
				continue
			}

			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}
