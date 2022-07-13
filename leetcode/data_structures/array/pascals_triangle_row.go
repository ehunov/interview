package array

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				result[j] = 1
				continue
			}

			if float64(j) < (float64(i)+1.0)/2.0 {
				result[j] = result[j] + result[i-j]
			} else {
				result[j] = result[i-j]
			}
		}
	}

	return result
}
